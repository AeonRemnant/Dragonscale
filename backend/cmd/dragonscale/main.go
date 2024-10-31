package main

import (
	"github.com/aeonremnant/dragonscale/api"
	"github.com/aeonremnant/dragonscale/docker"
)

func main() {
	go api.StartRESTServer()
	go docker.DockerClient()

	select {}
}
