package public

import (
	"github.com/379763376/gcommon/lib"
	"github.com/jinzhu/gorm"
)

var (
	GormPool *gorm.DB
)

func InitMysql() error {
	dbpool, err := lib.GetGormPool("default")
	if err != nil {
		return err
	}
	GormPool = dbpool
	return nil
}