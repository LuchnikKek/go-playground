package examples

import (
	"context"
	"io"
	"os"

	"github.com/docker/docker/client"
)

func StatsContainer() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(
		client.FromEnv, 
		client.WithHost("unix:///home/ilya/.docker/desktop/docker.sock"),
		client.WithAPIVersionNegotiation(),
	)
	if err != nil {
		panic(err)
	}
	
	statsResp, err := cli.ContainerStats(ctx, "2e17d636450d", true)
	if err != nil {
		panic(err)
	}
	defer statsResp.Body.Close()

	io.Copy(os.Stdout, statsResp.Body)
}
