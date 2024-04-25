package model

type Watermark struct {
	Name         string `gorm:"column:name;type:varchar(100);default:'';comment:水印名称" json:"name"`
	OperatorId   string `gorm:"column:operator_id;type:varchar(100);default:'';comment:操作人id" json:"operatorId"`
	OperatorName string `gorm:"column:operator_name;type:varchar(100);default:'';comment:操作人名称" json:"operatorName"`
	Status       int    `gorm:"column:status;type:integer;default:0;comment:状态 1:启用2:停用" json:"status"`
	Path720      string `gorm:"column:path720;type:varchar(200);default:'';comment:720水印路径" json:"path720"`
	Path1080     string `gorm:"column:path1080;type:varchar(200);default:'';comment:1080水印路径" json:"path1080"`
	Url720       string `_:"_:url720;type:varchar(200);default:'';comment:'720水印数据'" json:"url720"`
	Url1080      string `_:"_:url1080;type:varchar(200);default:'';comment:'1080水印数据'" json:"url1080"`
}

func (Watermark) TableName() string {
	return "watermark"
}
