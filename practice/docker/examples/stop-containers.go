package examples

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func StopContainers() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(
		client.FromEnv, 
		client.WithHost("unix:///home/ilya/.docker/desktop/docker.sock"),
		client.WithAPIVersionNegotiation(),
	)
	if err != nil {
		panic(err)
	}
	
	containers, err := cli.ContainerList(ctx, container.ListOptions{})
	if err != nil {
		panic(err)
	}
	
	for _, cont := range containers {
		fmt.Printf("Stopping %T{id: %s, name:%s}... ", cont, cont.ID[:10], cont.Names[0][1:])

		timeout := 2 // опционален
		if err := cli.ContainerStop(ctx, cont.ID, container.StopOptions{Timeout: &timeout}); err != nil {
			panic(err)
		}
		fmt.Println("Success")
	}
}