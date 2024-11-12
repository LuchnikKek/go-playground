package examples

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
)

func ListAllImages() {
	// DOCKER_HOST="unix:///home/ilya/.docker/desktop/docker.sock" go run main.go
	ctx := context.Background()

	cli, err := client.NewClientWithOpts(
		client.FromEnv,
		client.WithHost("unix:///home/ilya/.docker/desktop/docker.sock"),
		client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	filtersList := filters.NewArgs()
	// filtersList.Add("label", "")
	// filtersList.Add("dangling", "true")
	images, err := cli.ImageList(ctx, image.ListOptions{
		All: true,
		Filters: filtersList,
	})
	if err != nil {
		panic(err)
	}

	for _, image := range images {
		if len(image.RepoTags) == 0 {
			fmt.Printf("Image %s\n", image.ID)
		} else {
			fmt.Printf("Image from repo %s with size %d, %v\n", image.RepoTags[0], image.Size, image.ID)
			// fmt.Printf("%#v\n", image)
		}
		// fmt.Println(image.ID)
	}
	
}