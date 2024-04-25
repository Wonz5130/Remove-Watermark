package data

import (
	"remove-watermark/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	// "gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewRemoveWatermarkData,
)

// Data .
type Data struct {
	// db *gorm.DB
	// handler *data.DataHandler
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	log := log.NewHelper(logger)
	cleanup := func() {
		log.Info("closing the data resources")
	}
	return &Data{
		// db: global.DB,
	}, cleanup, nil
}
