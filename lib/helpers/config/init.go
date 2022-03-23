package config

import (
	"fmt"
	"github.com/go-gcfg/gcfg"
	"gopkg.in/ini.v1"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

//ReadConfig read config file .ini
func (serviceConfig *ServiceConfig) ReadConfig(module string) {
	environ := os.Getenv("ENV")
	pathBase := fmt.Sprintf("etc/%s", "configure")
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("[INIT] [CONFIG] fail get working directory: %v", err)
	}

	if environ == "" {
		environ = "development"
		pathBase = fmt.Sprintf("files/etc/%s", "configure")
	}

	filePath := fmt.Sprintf("%s/%s/%s.%s.ini", dir, pathBase, module, environ)

	cfg, err := ini.LoadSources(ini.LoadOptions{
		AllowShadows:             true,
		SpaceBeforeInlineComment: true,
	}, filePath)
	if err != nil {
		log.Fatalf("[INIT] [CONFIG] fail LoadSources .ini: %v", err)
	}

	ini.PrettyFormat = true
	err = cfg.SaveTo(filePath)
	if err != nil {
		log.Fatalf("[INIT] [CONFIG] fail SaveTo .ini: %v", err)
	}

	configStr, err := ioutil.ReadFile(filepath.Clean(filePath))
	if err != nil {
		log.Fatalf("[INIT] [CONFIG] fail ReadFile: %v", err)
	}
	err = gcfg.FatalOnly(gcfg.ReadStringInto(serviceConfig, string(configStr)))
	if err != nil {
		log.Fatalf("[INIT] [CONFIG] fail FatalOnly: %v", err)
	}

	log.Printf("[IGNORE] [CONFIG] [%s] initialize", strings.ToUpper(module))
}
