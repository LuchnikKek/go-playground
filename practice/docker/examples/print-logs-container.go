package examples

import (
	"context"
	"io"
	"os"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func PrintLogs(containerStr string) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(
		client.FromEnv,
		client.WithHost("unix:///home/ilya/.docker/desktop/docker.sock"),
		client.WithAPIVersionNegotiation(),
	)
	if err != nil {
		panic(err)
	}
	
	options := container.LogsOptions{ShowStdout: true}
	out, err := cli.ContainerLogs(ctx, containerStr, options)
	if err != nil {
		panic(err)
	}
	
	io.Copy(os.Stdout, out)
}