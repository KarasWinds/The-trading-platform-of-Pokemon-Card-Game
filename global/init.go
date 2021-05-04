package global

import (
	"flag"
	"log"
	"strings"

	"github.com/KarasWinds/The-trading-platform-of-Pokemon-Card-Game/pkg/db"
	"github.com/KarasWinds/The-trading-platform-of-Pokemon-Card-Game/pkg/setting"
)

var (
	config string
)

func init() {
	err := setupFlag()
	if err != nil {
		log.Fatalf("init.setupFlag err: %v", err)
	}
	err = setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}
}

func setupFlag() error {
	flag.StringVar(&config, "config", "../../configs/", "指定要使用的設定檔路徑")
	flag.Parse()

	return nil
}

func setupSetting() error {
	setting, err := setting.NewSetting(strings.Split(config, ",")...)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Server", &ServerSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Database", &DatabaseSetting)
	if err != nil {
		return err
	}

	return nil
}

func setupDBEngine() error {
	var err error
	DBEngine, err = db.NewDBEngine(DatabaseSetting)
	if err != nil {
		return err
	}

	return nil
}
