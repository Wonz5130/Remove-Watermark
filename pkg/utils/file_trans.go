package utils

import (
	"context"
	"fmt"
	"image"
	"image/png"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"sort"
	"strings"
	"sync"

	pkgError "github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/sync/errgroup"
	"gopkg.in/resty.v1"
)

const tmpDir = "/tmp"

func RemoveAllByPathList(p ...string) error {
	for i := range p {
		go os.RemoveAll(p[i])
	}
	return nil
}

// filePath: 本地绝对路径
func Download(fileUrl string) (filePath string, err error) {
	client := resty.New()
	filePath = path.Join(tmpDir, uuid.NewV4().String()+path.Ext(fileUrl))
	_, err = client.R().SetOutput(filePath).Get(fileUrl)
	return filePath, err
}

// word to pdf, 仅支持.doc, .docx
func DocToPdf(sourceFilePath string) (outFilePath string, err error) {
	outDir := path.Join(tmpDir, uuid.NewV4().String())
	os.Mkdir(outDir, os.ModePerm)
	outFilePath = path.Join(outDir, path.Base(sourceFilePath))
	// exec.Command 前两个参数加上 "/bin/sh", "-c", 解决报错 no such file or directory
	cmd := exec.Command("/bin/sh", "-c", fmt.Sprintf(`libreoffice --headless --invisible --convert-to pdf %s --outdir %s`, sourceFilePath, outDir))
	if err := cmd.Run(); err != nil {
		return "", err
	}
	return outFilePath, err
}

func PdfToPngs(sourceFilePath string) (outDir string, err error) {
	pngsDir := path.Join(tmpDir, uuid.NewV4().String())
	os.Mkdir(pngsDir, os.ModePerm)
	cmd := exec.Command("/bin/sh", "-c", fmt.Sprintf(`pdftoppm -png -r 150 %s %s`, sourceFilePath, path.Join(pngsDir, path.Base(sourceFilePath))))
	if err := cmd.Run(); err != nil {
		return "", err
	}
	return pngsDir, nil
}

func PngsToPdf(sourceDir string) (filePath string, err error) {
	filePath = path.Join(tmpDir, uuid.NewV4().String()+".pdf")
	files, _ := ioutil.ReadDir(sourceDir)
	filePathList := make([]string, len(files))
	for i := range files {
		filePathList[i] = path.Join(sourceDir, files[i].Name())
	}
	sort.Strings(filePathList)
	cmdImgStr := strings.Join(filePathList, " ")
	cmd := exec.Command("/bin/sh", "-c", fmt.Sprintf(`convert %s %s`, cmdImgStr, filePath))
	if err := cmd.Run(); err != nil {
		return "", err
	}
	return filePath, nil
}

func CoverImgs(sourceDir string) (outDir string, err error) {
	newPngsDir := path.Join(tmpDir, uuid.NewV4().String())
	os.Mkdir(newPngsDir, os.ModePerm)
	files, _ := ioutil.ReadDir(sourceDir)
	var wg sync.WaitGroup
	wg.Add(len(files))
	g, _ := errgroup.WithContext(context.Background())
	for _, f := range files {
		f := f
		// 并发
		g.Go(func() (err error) {
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in ", r)
					err = pkgError.Cause(fmt.Errorf("recovered CoverImgs in %s", r))
				}
			}()
			err = CutPng(path.Join(sourceDir, f.Name()), newPngsDir)
			if err != nil {
				return err
			}
			return nil
		})
	}
	if err := g.Wait(); err != nil {
		return "", pkgError.Wrap(err, "GetSubsectionGoals")
	}
	return newPngsDir, nil
}

func readImage(filePath string) (image.Image, error) {
	fd, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer fd.Close()
	img, _, err := image.Decode(fd)
	if err != nil {
		return nil, err
	}
	return img, nil
}

func cropImage(img image.Image, crop image.Rectangle) (image.Image, error) {
	type subImager interface {
		SubImage(r image.Rectangle) image.Image
	}
	simg, ok := img.(subImager)
	if !ok {
		return nil, fmt.Errorf("image does not support cropping")
	}

	return simg.SubImage(crop), nil
}

func writeImage(img image.Image, name string) error {
	fd, err := os.Create(name)
	if err != nil {
		return err
	}
	defer fd.Close()

	return png.Encode(fd, img)
}

func CutPng(filePath, pngDir string) error {
	img, err := readImage(filePath)
	if err != nil {
		return pkgError.Wrap(err, `readImage`)
	}
	imgW := img.Bounds().Dx()
	imgH := img.Bounds().Dy()

	cropLeftTopX := imgW * 60 / 796
	cropLeftTopY := imgH * 60 / 1121
	cropRightBottomX := imgW - imgW*60/796
	cropRightBottomY := imgH - imgH*60/1121

	img, err = cropImage(img, image.Rect(cropLeftTopX, cropLeftTopY, cropRightBottomX, cropRightBottomY))
	if err != nil {
		return err
	}
	newFilePath := path.Join(pngDir, path.Base(filePath))
	return writeImage(img, newFilePath)
}

// func CoverByWhitePng(filePath string, pngDir string) error {
// 	img_file, err := os.Open(filePath)
// 	if err != nil {
// 		return pkgError.Wrap(err, `OpenPngError`)
// 	}
// 	defer img_file.Close()
// 	img, err := png.Decode(img_file)
// 	if err != nil {
// 		return pkgError.Wrap(err, `DecodePngError`)
// 	}
// 	width := img.Bounds().Dx()
// 	headMaskImg := image.NewRGBA(image.Rect(0, 0, width, 100))
// 	// endMask := image.NewRGBA(image.Rect(0, 0, width, 100))
// 	headMaskImgColor := color.RGBA{0, 0, 0, 255}
// 	draw.Draw(headMaskImg, headMaskImg.Bounds(), &image.Uniform{headMaskImgColor}, image.Point{}, draw.Src)

// 	//把水印写在右下角，并向0坐标偏移10个像素
// 	offset := image.Pt(0, 0)
// 	b := img.Bounds()
// 	//根据b画布的大小新建一个新图像
// 	m := image.NewRGBA(b)

// 	//image.ZP代表Point结构体，目标的源点，即(0,0)
// 	//draw.Src源图像透过遮罩后，替换掉目标图像
// 	//draw.Over源图像透过遮罩后，覆盖在目标图像上（类似图层）
// 	draw.Draw(m, b, img, image.Point{}, draw.Src)
// 	draw.Draw(m, headMaskImg.Bounds().Add(offset), headMaskImg, image.Point{}, draw.Over)

// 	newFilePath := path.Join(pngDir, path.Base(filePath))
// 	imgw, err := os.Create(newFilePath)
// 	if err != nil {
// 		return pkgError.Wrap(err, `os.Create`)
// 	}
// 	png.Encode(imgw, m)
// 	defer imgw.Close()
// 	return nil
// }
