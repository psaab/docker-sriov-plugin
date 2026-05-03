package driver

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func getRightClient() (*client.Client, error) {
	return client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
}

// IsNetworkIdValid reports whether the given network id exists in dockerd.
// The bool result is only meaningful when err == nil; on any error
// (e.g. dockerd unreachable) callers must NOT treat the network as stale,
// since deleting persisted config on a transient failure would be destructive.
func IsNetworkIdValid(id string) (bool, error) {
	cli, err := getRightClient()
	if err != nil {
		return false, err
	}
	networks, err := cli.NetworkList(context.Background(), types.NetworkListOptions{})
	if err != nil {
		return false, err
	}

	for _, network := range networks {
		if network.ID == id {
			return true, nil
		}
	}
	return false, nil
}
