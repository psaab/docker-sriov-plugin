package main

import (
	"github.com/urfave/cli"
	"github.com/docker/go-plugins-helpers/network"
	"log"
	"os"

	"github.com/psaab/docker-sriov-plugin/driver"
)

const (
	version = "0.1.2"
)

// Run initializes the driver
func Run(ctx *cli.Context) error {
	d, err := driver.StartDriver()
	if err != nil {
		panic(err)
	}
	h := network.NewHandler(d)

	log.Printf("Mellanox sriov plugin started version=%v\n", version)

	// Reconcile persisted state with dockerd in the background. This MUST
	// run concurrently with ServeUnix: at boot dockerd's container-restore
	// is blocked waiting for our plugin socket, so calling NetworkList
	// inline would deadlock. d.Mutex serializes any DeleteNetwork the
	// reconciliation issues against incoming dockerd requests.
	go d.ValidatePersistentNetworksLoop()

	log.Printf("Ready to accept commands.\n")
	err = h.ServeUnix("sriov", 0)
	if err != nil {
		log.Fatalf("Run app error: %s", err.Error())
	}
	return err
}

func main() {

/*
	var flagDebug = cli.BoolFlag{
		Name:  "debug, d",
		Usage: "enable debugging",
	}
*/
	app := cli.NewApp()
	app.Name = "sriov"
	app.Usage = "Docker Networking using SRIOV/Passthrough netdevices"
	app.Version = version
	app.Flags = []cli.Flag{
	//	flagDebug,
	}
	app.Action = Run
	app.Run(os.Args)
}
