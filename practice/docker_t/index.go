package docker_t

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"io"
)

func main() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		fmt.Println(err)
		return
	}
	//PrintlnDockerList(ctx, cli)
	GetContainerNetworkStats(ctx, cli, "42868_8034850f_2698d6c20affd188754ca34f17f43918.0")
}

func GetDockerList(ctx context.Context, cli *client.Client) ([]container.Summary, error) {
	containers, err := cli.ContainerList(ctx, container.ListOptions{})
	if err != nil {
		return nil, err
	}
	return containers, nil
}

func PrintlnDockerList(ctx context.Context, cli *client.Client) {
	containers, err := GetDockerList(ctx, cli)
	if err != nil {
		fmt.Println(err)
		return
	}
	for i, containerInfo := range containers {
		fmt.Printf("-----------------------------%d-----------------------------\n", i)
		fmt.Println("ID:", containerInfo.ID, "\n",
			"Names:", containerInfo.Names, "\n",
			"Image:", containerInfo.Image, "\n",
			"ImageID:", containerInfo.ImageID, "\n",
			"ImageManifestDescriptor:", containerInfo.ImageManifestDescriptor, "\n",
			"Command:", containerInfo.Command, "\n",
			"Created:", containerInfo.Created, "\n",
			"Ports:", containerInfo.Ports, "\n",
			"SizeRw:", containerInfo.SizeRw, "\n",
			"SizeRootFs:", containerInfo.SizeRootFs, "\n",
			"Labels:", containerInfo.Labels, "\n",
			"State:", containerInfo.State, "\n",
			"Status:", containerInfo.Status, "\n",
			"HostConfig:", containerInfo.HostConfig, "\n",
			"NetworkSettings:", containerInfo.NetworkSettings, "\n",
			"Mounts:", containerInfo.Mounts,
		)
		fmt.Printf("-----------------------------%d-----------------------------\n", i)
	}
}

type NetStats struct {
	RxBytes uint64 // 接收（下行）
	TxBytes uint64 // 发送（上行）
}

// GetContainerNetworkStats 获取指定容器的网络上下行流量
func GetContainerNetworkStats(ctx context.Context, cli *client.Client, containerID string) (*NetStats, error) {
	stats, err := cli.ContainerStats(ctx, containerID, false)
	if err != nil {
		return nil, fmt.Errorf("get stats error: %w", err)
	}
	defer stats.Body.Close()

	// 读取所有内容为 []byte
	data, err := io.ReadAll(stats.Body)
	if err != nil {
		return nil, fmt.Errorf("reading stats body: %w", err)
	}

	fmt.Println("---------------------", string(data))

	return &NetStats{}, nil
}
