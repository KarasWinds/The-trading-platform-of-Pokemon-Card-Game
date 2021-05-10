package init

import (
	"flag"
	"log"
	"strings"
	"testing"

	"github.com/KarasWinds/The-trading-platform-of-Pokemon-Card-Game/global"
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
	testing.Init()
	flag.Parse()
	return nil
}

func setupSetting() error {
	setting, err := setting.NewSetting(strings.Split(config, ",")...)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}

	return nil
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = db.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}

	return nil
}
