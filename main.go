package main

import (
	"github.com/codegangsta/cli"
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
	log.Printf("Ready to accept commands.\n")

	go d.ValidatePersistentNetworks()
	err = h.ServeUnix("sriov", 0)
	if err != nil {
		log.Fatal("Run app error: %s", err.Error())
		os.Exit(1)
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
