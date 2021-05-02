package db

import (
	"fmt"

	"github.com/KarasWinds/The-trading-platform-of-Pokemon-Card-Game/pkg/setting"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDBEngine(databaseSetting *setting.DatabaseSettingS) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
			databaseSetting.UserName,
			databaseSetting.Password,
			databaseSetting.Host,
			databaseSetting.DBName,
			databaseSetting.Charset,
			databaseSetting.ParseTime,
		)}), &gorm.Config{})

	if err != nil {
		return nil, err
	}
	return db, nil
}
