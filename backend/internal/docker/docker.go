package docker

import (
	"context"
	"fmt"

	containertypes "github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func ListRunningContainers() ([]string, error) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, fmt.Errorf("Failed to create Docker client: %w", err)
	}
	defer cli.Close()

	containers, err := cli.ContainerList(ctx, containertypes.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to list containers: %w", err)
	}

	var containerIDs []string
	for _, container := range containers {
		containerIDs = append(containerIDs, container.ID)
	}

	return containerIDs, nil

}
