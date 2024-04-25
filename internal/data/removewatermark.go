package data

import (
	"remove-watermark/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type RemoveWatermarkRepo struct {
	data *Data
	log  *log.Helper
}

func NewRemoveWatermarkData(data *Data, logger log.Logger) biz.RemoveWatermarkRepo {
	return &RemoveWatermarkRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// func (r *RemoveWatermarkRepo) GetWatermarks(ctx context.Context, name string, offset int32, limit int32, status int32) (int32, []*model.Watermark, int32, error) {
// 	var count int64
// 	watermarks := []*model.Watermark{}

// 	tx := r.data.db.Model(model.Watermark{}).Where("deleted_at is null")
// 	if len(name) > 0 {
// 		tx = tx.Where("name = ?", name)
// 	}
// 	if status > 0 {
// 		tx = tx.Where("status = ?", status)
// 	}

// 	if err := tx.Order("updated_at desc").Offset(int(offset)).Limit(int(limit)).Find(&watermarks).Offset(-1).Limit(-1).Count(&count).Error; err != nil {
// 		return 0, nil, 0, errors.Wrap(err, "GetWatermarks err")
// 	}
// 	return int32(count), watermarks, int32(limit), nil
// }
