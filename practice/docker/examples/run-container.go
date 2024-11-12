package examples

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
)

func RunContainer() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(
		client.FromEnv, 
		client.WithHost("unix:///home/ilya/.docker/desktop/docker.sock"),
		client.WithAPIVersionNegotiation(),
	)
	if err != nil {
		panic(err)
	}
	
	imageName := "alpine"
	reader, err := cli.ImagePull(ctx, imageName, image.PullOptions{})
	if err != nil {
		panic(err)
	}
	
	defer reader.Close()
	io.Copy(os.Stdout, reader)

	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: imageName,
		Cmd: []string{"echo", "Hello, World!"},
		Tty: false,
	}, nil, nil, nil, "hw-container")
	if err != nil {
		panic(err)
	}
	
	if err := cli.ContainerStart(ctx, resp.ID, container.StartOptions{}); err != nil {panic(err)}

	statusCh, errCh := cli.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			panic(err)
		}
	case status := <-statusCh:
		fmt.Printf("Status is %d\n", status.StatusCode)
	}

	out, err := cli.ContainerLogs(ctx, resp.ID, container.LogsOptions{ShowStdout: true})
	if err != nil {
		panic(err)
	}
	stdcopy.StdCopy(os.Stdout, os.Stderr, out)

	
}