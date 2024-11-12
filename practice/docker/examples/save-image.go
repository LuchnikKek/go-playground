package examples

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/docker/docker/client"
)

func SaveImage(imageid string, filepath string) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(
		client.FromEnv, 
		client.WithHost("unix:///home/ilya/.docker/desktop/docker.sock"),
		client.WithAPIVersionNegotiation(),
	)
	if err != nil {
		panic(err)
	}

	images := make([]string,1)
	images[0] = imageid
	reader, err := cli.ImageSave(ctx, images)
	if err != nil {
		panic(err)
	}
	defer reader.Close()

	file, err := os.Create(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	
	writtenBytes, err := io.Copy(file, reader)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Bytes written %d\n", writtenBytes)
}