package examples

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
)

func CreateNetwork() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(
		client.FromEnv, 
		client.WithHost("unix:///home/ilya/.docker/desktop/docker.sock"),
		client.WithAPIVersionNegotiation(),
	)
	if err != nil {
		panic(err)
	}

	imageName := "postgres"

	out, err := cli.ImagePull(ctx, imageName, image.PullOptions{})
	if err != nil {
		panic(err)
	}
	defer out.Close()
	io.Copy(os.Stdout, out)

	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Hostname: "postgresC",
		Domainname: "postgresD",
		Image: imageName,
		Env: []string{
			"POSTGRES_PASSWORD=123qwe",
		},
	}, nil, nil, nil, "postgresDB")
	if err != nil {
		panic(err)
	}

	if err := cli.ContainerStart(ctx, resp.ID, container.StartOptions{}); err != nil {panic(err)}

	fmt.Println("Created container with ID: ", resp.ID)
	networkCreate, err := cli.NetworkCreate(ctx, "mynet", network.CreateOptions{Attachable: true})
	if err != nil {
		panic(err)
	}
	fmt.Println("Created network with networkId: ", networkCreate.ID)

	err = cli.NetworkConnect(ctx, networkCreate.ID, resp.ID, &network.EndpointSettings{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected successfully")
}
