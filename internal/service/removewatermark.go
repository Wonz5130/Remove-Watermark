package service

import (
	"context"
	"os"

	pb "remove-watermark/api/removewatermark"
	"remove-watermark/internal/biz"
	"remove-watermark/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
)

type RemoveWatermarkService struct {
	pb.UnimplementedRemoveWatermarkSrvServer
	removewatermark *biz.RemoveWatermarkUsecase
	log             *log.Helper
	config          *conf.Data
}

func NewRemoveWatermarkService(
	removewatermark *biz.RemoveWatermarkUsecase,
	logger log.Logger, config *conf.Data,
) *RemoveWatermarkService {
	return &RemoveWatermarkService{
		removewatermark: removewatermark,
		log:             log.NewHelper(logger),
		config:          config,
	}
}

// 读取 pdf 文件，转为 byte
func ReadPdf(path string) ([]byte, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return content, nil
}

func (s *RemoveWatermarkService) RemoveCopyright(ctx context.Context, req *pb.RemoveWatermarkReq) (*pb.RemoveWatermarkRes, error) {
	fileUrl := req.Fileurl
	pdfUrl, err := s.removewatermark.RemoveWatermark(ctx, fileUrl)
	if err != nil {
		return nil, err
	}
	return &pb.RemoveWatermarkRes{
		Pdfurl: pdfUrl,
	}, nil
}
