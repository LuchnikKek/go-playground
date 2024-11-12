package examples

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
)

func ListContainers() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(
		client.FromEnv, 
		client.WithHost("unix:///home/ilya/.docker/desktop/docker.sock"),
		client.WithAPIVersionNegotiation(),
	)
	if err != nil {
		panic(err)
	}
	
	filtersList := filters.NewArgs()
	filtersList.Add("name", "hw") // Name contains "hw"

	containers, err := cli.ContainerList(ctx, container.ListOptions{
		All: true, //all, not only running
		Filters: filtersList,
	})
	if err != nil {
		panic(err)
	}
	
	for _, cont := range containers {
		fmt.Printf("%T{id: %s, name:%s, image:%s} is %s\n", cont, cont.ID[:10], cont.Names[0][1:], cont.Image, cont.State)
	}
}