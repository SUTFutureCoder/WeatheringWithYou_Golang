package util

import (
	"github.com/go-ini/ini"
	"os"
	"fmt"
	"log"
)

func GetConfig(confFile, section string) (map[string]string, error) {
	dir, _ := os.Getwd()
	confDir := dir + "/conf/" + confFile + ".ini"
	cfg, err := ini.Load(confDir)
	if err != nil {
		log.Fatal(fmt.Sprintf("conf load error: %s", err))
		return nil, err
	}
	cfg.BlockMode = false
	hash := cfg.Section(section).KeysHash()
	return hash, err
}
