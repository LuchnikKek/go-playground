package examples

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func CommitContainer() {
	// По коммиту можно сделать docker run -it IMAGE_HASH|IMAGE_NAME sh
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
	
	createResp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: imageName,
		Cmd: []string{"touch", "/hello-world"},
	}, nil, nil, nil, "hw-file-container")
	if err != nil {
		panic(err)
	}
	
	if err := cli.ContainerStart(ctx, createResp.ID, container.StartOptions{}); err != nil {panic(err)}

	statusCh, errCh := cli.ContainerWait(ctx, createResp.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			panic(err)
		}
	case status := <-statusCh:
		fmt.Printf("Status is %d\n", status.StatusCode)
	}

	commitResp, err := cli.ContainerCommit(ctx, createResp.ID, container.CommitOptions{
		Reference: "hw-commit", // Image Name
		Comment: "Commited moment",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(commitResp.ID)
}