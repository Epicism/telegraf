package docker

import (
	"context"
	"crypto/tls"
	"net/http"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/swarm"
	"github.com/docker/docker/api/types/system"
	"github.com/docker/docker/client"
)

var (
	defaultHeaders = map[string]string{"User-Agent": "engine-api-cli-1.0"}
)

type Client interface {
	Info(ctx context.Context) (system.Info, error)
	ContainerList(ctx context.Context, options container.ListOptions) ([]types.Container, error)
	ContainerStats(ctx context.Context, containerID string, stream bool) (container.StatsResponseReader, error)
	ContainerInspect(ctx context.Context, containerID string) (types.ContainerJSON, error)
	ServiceList(ctx context.Context, options types.ServiceListOptions) ([]swarm.Service, error)
	TaskList(ctx context.Context, options types.TaskListOptions) ([]swarm.Task, error)
	NodeList(ctx context.Context, options types.NodeListOptions) ([]swarm.Node, error)
	DiskUsage(ctx context.Context, options types.DiskUsageOptions) (types.DiskUsage, error)
	ClientVersion() string
	Close() error
}

func NewEnvClient() (Client, error) {
	dockerClient, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return nil, err
	}
	return &SocketClient{dockerClient}, nil
}

func NewClient(host string, tlsConfig *tls.Config) (Client, error) {
	transport := &http.Transport{
		TLSClientConfig: tlsConfig,
	}
	httpClient := &http.Client{Transport: transport}

	dockerClient, err := client.NewClientWithOpts(
		client.WithHTTPHeaders(defaultHeaders),
		client.WithHTTPClient(httpClient),
		client.WithAPIVersionNegotiation(),
		client.WithHost(host))
	if err != nil {
		return nil, err
	}

	return &SocketClient{dockerClient}, nil
}

type SocketClient struct {
	client *client.Client
}

func (c *SocketClient) Info(ctx context.Context) (system.Info, error) {
	return c.client.Info(ctx)
}
func (c *SocketClient) ContainerList(ctx context.Context, options container.ListOptions) ([]types.Container, error) {
	return c.client.ContainerList(ctx, options)
}
func (c *SocketClient) ContainerStats(ctx context.Context, containerID string, stream bool) (container.StatsResponseReader, error) {
	return c.client.ContainerStats(ctx, containerID, stream)
}
func (c *SocketClient) ContainerInspect(ctx context.Context, containerID string) (types.ContainerJSON, error) {
	return c.client.ContainerInspect(ctx, containerID)
}
func (c *SocketClient) ServiceList(ctx context.Context, options types.ServiceListOptions) ([]swarm.Service, error) {
	return c.client.ServiceList(ctx, options)
}
func (c *SocketClient) TaskList(ctx context.Context, options types.TaskListOptions) ([]swarm.Task, error) {
	return c.client.TaskList(ctx, options)
}
func (c *SocketClient) NodeList(ctx context.Context, options types.NodeListOptions) ([]swarm.Node, error) {
	return c.client.NodeList(ctx, options)
}
func (c *SocketClient) DiskUsage(ctx context.Context, options types.DiskUsageOptions) (types.DiskUsage, error) {
	return c.client.DiskUsage(ctx, options)
}

func (c *SocketClient) ClientVersion() string {
	return c.client.ClientVersion()
}

func (c *SocketClient) Close() error {
	return c.client.Close()
}
