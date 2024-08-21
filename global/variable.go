package global

import (
	"github.com/hespecial/gin-mall/config"
	"github.com/olivere/elastic"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Config   *config.Config
	DB       *gorm.DB
	EsClient *elastic.Client
	Log      *zap.Logger
)
