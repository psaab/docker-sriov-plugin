package driver

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func getRightClient() (*client.Client, error) {
	return client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
}

// ListDockerNetworkIDs returns the set of network IDs currently known to
// dockerd. The caller-supplied context bounds the round-trip so a single
// hang can't wedge validation forever. On any error the caller must NOT
// treat absent ids as stale -- deleting persisted config on a transient
// dockerd failure would be destructive.
func ListDockerNetworkIDs(ctx context.Context) (map[string]struct{}, error) {
	cli, err := getRightClient()
	if err != nil {
		return nil, err
	}
	defer cli.Close()

	networks, err := cli.NetworkList(ctx, types.NetworkListOptions{})
	if err != nil {
		return nil, err
	}
	ids := make(map[string]struct{}, len(networks))
	for _, n := range networks {
		ids[n.ID] = struct{}{}
	}
	return ids, nil
}
