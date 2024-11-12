package examples

import (
	"bufio"
	"context"
	"io"
	"os"

	"github.com/docker/docker/client"
)

func LoadImage(filepath string) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(
		client.FromEnv,
		client.WithHost("unix:///home/ilya/.docker/desktop/docker.sock"),
		client.WithAPIVersionNegotiation(),
	)
	if err != nil {
		panic(err)
	}

	f, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(f)
	load, err := cli.ImageLoad(ctx, reader, false)
	if err != nil {
		panic(err)
	}
	
	cli.ImageTag(ctx, "f8a2146330cc", "gonewserver:latest")

	defer load.Body.Close()
	io.Copy(os.Stdout, load.Body)
}