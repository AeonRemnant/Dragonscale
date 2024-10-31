package main

import (
	"github.com/aeonremnant/dragonscale/internal/api"
	"github.com/aeonremnant/dragonscale/internal/docker"
)

func main() {
	go api.StartRESTServer()
	go docker.DockerClient()

	select {}
}
