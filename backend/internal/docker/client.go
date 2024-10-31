package docker

import (
	"fmt"
	"log"

	"github.com/docker/docker/client"
)

func DockerClient() *client.Client {
	fmt.Println("Initialized Docker Client")
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Fatalf("Error creating Docker client: %v", err)
	}
	return cli
}
