package biz

import (
	"context"
	"fmt"
	"path"

	"remove-watermark/internal/conf"
	"remove-watermark/pkg/utils"

	"github.com/go-kratos/kratos/v2/log"
	pkgError "github.com/pkg/errors"
	// third "remove-watermark/third_party"
)

type RemoveWatermarkRepo interface {
	// GetWatermarks(ctx context.Context, name string, offset int32, limit int32, status int32) (count int32, data []*model.Watermark, newLimit int32, err error)
}

type RemoveWatermarkUsecase struct {
	repo   RemoveWatermarkRepo
	log    *log.Helper
	config *conf.Data
}

func NewRemoveWatermarkUsecase(
	repo RemoveWatermarkRepo,
	logger log.Logger, config *conf.Data) *RemoveWatermarkUsecase {
	return &RemoveWatermarkUsecase{
		repo:   repo,
		log:    log.NewHelper(logger),
		config: config,
	}
}

func (uc *RemoveWatermarkUsecase) RemoveWatermark(ctx context.Context, fileUrl string) (pdfUrl string, err error) {
	var filePath string
	var pdfFilePath string
	var pngsDir string
	var newPngsDir string
	var outPdfFilePath string
	// defer：函数执行完再进行删除所有文件
	defer func() {
		utils.RemoveAllByPathList(filePath, pdfFilePath, pngsDir, newPngsDir, outPdfFilePath)
	}()
	filePath, err = utils.Download(fileUrl)
	ext := path.Ext(filePath)
	if !(ext == ".doc" || ext == ".docx" || ext == ".pdf") {
		return "", fmt.Errorf("不支持的后缀: %s", ext)
	}
	pdfFilePath = filePath
	// 如果是word，就转为pdf
	if ext == ".doc" || ext == ".docx" {
		if pdfFilePath, err = utils.DocToPdf(filePath); err != nil {
			return "", pkgError.Wrap(err, `DocToPdf`)
		}
	}
	pngsDir, err = utils.PdfToPngs(pdfFilePath)
	if err != nil {
		return "", pkgError.Wrap(err, `PdfToPngs`)
	}
	newPngsDir, err = utils.CoverImgs(pngsDir)
	if err != nil {
		return "", pkgError.Wrap(err, `CoverImgs`)
	}
	outPdfFilePath, err = utils.PngsToPdf(newPngsDir)
	if err != nil {
		return "", pkgError.Wrap(err, `PngsToPdf`)
	}
	// pdfurl, err := third.UploadYcoss(outPdfFilePath, path.Join("pdf", uuid.NewV4().String()+".pdf"), uc.config.YcOss.AccessToken)
	// if err != nil {
	// 	return "", pkgError.Wrap(err, `UploadYcoss`)
	// }
	// return pdfurl, nil
	return outPdfFilePath, nil
}
