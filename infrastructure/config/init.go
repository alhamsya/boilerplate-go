package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/go-gcfg/gcfg"
	"gopkg.in/ini.v1"
)

//ReadConfig read config file .ini
func (mainConfig *MainConfig) ReadConfig(module string) {
	environ := os.Getenv("ENV")
	pathBase := fmt.Sprintf("etc/%s", "service")
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("[INIT] fail get working directory: %v", err)
	}

	if environ == "" {
		environ = "development"
		pathBase = fmt.Sprintf("files/etc/%s", "service")
	}

	filePath := fmt.Sprintf("%s/%s/%s.%s.ini", dir, pathBase, module, environ)

	cfg, err := ini.LoadSources(ini.LoadOptions{
		SpaceBeforeInlineComment: true,
	}, filePath)
	if err != nil {
		log.Fatalf("[INIT] fail load file .ini: %v", err)
	}

	ini.PrettyFormat = true
	err = cfg.SaveTo(filePath)
	if err != nil {
		log.Fatalf("[INIT] fail load file .ini: %v", err)
	}

	configStr, err := ioutil.ReadFile(filepath.Clean(filePath))
	if err != nil {
		log.Fatalf("[INIT] ReadFile: %v", err)
	}
	err = gcfg.FatalOnly(gcfg.ReadStringInto(mainConfig, string(configStr)))
	if err != nil {
		log.Fatalf("[INIT] fail FatalOnly: %v", err)
	}
}