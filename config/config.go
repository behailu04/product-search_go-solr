package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"runtime"
	"strings"
)

type ConfigurationModel struct {
	Port string `json:"port"`
	Solr struct {
		Addr string `json:"addr"`
		Core string `json:"core"`
	} `json:"solr"`
}

var (
	Configuration = ConfigurationModel{}
)

func init() {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)
	basepath = strings.Replace(basepath, "config", "", -1)

	file := basepath + "config.json"

	raw, err := ioutil.ReadFile(file)

	if err != nil {
		panic(fmt.Sprintf("Failed to load auth configuration file: %s", err.Error()))
	}
	err = json.Unmarshal(raw, &Configuration)

	if err != nil {
		panic(fmt.Sprintf("Failed to parse auth configuration file: %s", err.Error()))
	}
}
