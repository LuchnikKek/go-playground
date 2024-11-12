package examples

import (
	"context"
	"io"
	"os"

	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
)

func PullImage() {
	ctx := context.Background()

	cli, err := client.NewClientWithOpts(
		client.FromEnv,
		client.WithAPIVersionNegotiation(),
		client.WithHost("unix:///home/ilya/.docker/desktop/docker.sock"))
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	out, err := cli.ImagePull(ctx, "hello-world", image.PullOptions{})
	if err != nil {
		panic(err)
	}

	defer out.Close()
	io.Copy(os.Stdout, out)
}