package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
)

type config struct {
	Proxies []struct {
		Name string `yaml:"name"`
		Addr string `yaml:"addr"`
		Port int    `yaml:"port"`
	} `yaml:"proxies"`
}

var WorkDir string

var ROOT config

func init() {
	WorkDir, _ = os.Getwd()
	raw, err := ioutil.ReadFile(fmt.Sprintf("%v/proxies.yaml", WorkDir))
	if err != nil {
		panic(err)
	}

	if err := yaml.Unmarshal(raw, &ROOT); err != nil {
		panic(err)
	}
}
