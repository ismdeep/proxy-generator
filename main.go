package main

import (
	"fmt"
	"github.com/ismdeep/proxy-generator/config"
	"github.com/ismdeep/proxy-generator/docker/compose"
	"gopkg.in/yaml.v3"
)

func main() {
	var instance compose.DockerComposeYAML
	instance.Version = "3.9"
	var proxies compose.Service
	proxies.Image = "hub.deepin.com/public/nginx-proxy:latest"
	proxies.Restart = "always"
	proxies.Volumes = []string{"./config.yaml:/config.yaml"}
	for _, v := range config.ROOT.Proxies {
		proxies.Ports = append(proxies.Ports, fmt.Sprintf("%v:%v", v.Port, v.Port))
	}
	instance.Services = make(map[string]compose.Service)
	instance.Services["proxies"] = proxies

	content, err := yaml.Marshal(instance)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(content))
}
