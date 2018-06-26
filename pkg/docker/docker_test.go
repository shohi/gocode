package docker

import (
	"io"
	"log"
	"os"
	"testing"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
)

func TestCreateContainerUseDockerAPI(t *testing.T) {
	ctx := context.Background()
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	reader, err := cli.ImagePull(ctx, "docker.io/library/alpine", types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}
	io.Copy(os.Stdout, reader)

	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: "alpine",
		Cmd:   []string{"echo", "hello world"},
		Tty:   true,
	}, nil, nil, "")
	if err != nil {
		panic(err)
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	status, err := cli.ContainerWait(ctx, resp.ID)
	t.Logf("status ==> %v", status)
	if err != nil {
		panic(err)
	}

	out, err := cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true})
	if err != nil {
		panic(err)
	}

	io.Copy(os.Stdout, out)
}

func TestListAllImages(t *testing.T) {
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	images, err := cli.ImageList(context.Background(), types.ImageListOptions{})
	if err != nil {
		panic(err)
	}

	for _, image := range images {
		log.Println(image.ID)
	}
}

func TestListAllContainers(t *testing.T) {
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	// List all containers
	listAllOpts := types.ContainerListOptions{All: true}
	containers, err := cli.ContainerList(context.Background(), listAllOpts)

	// listRunningOpts := types.ContainerListOptions{All: false}
	// containers, err := cli.ContainerList(context.Background(), listRunningOpts)

	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		log.Println(container.ID)
	}
}
