package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

var (
	app    = NewApp()
	server = new(srv)
)

func init() {
	app.Action = cli.ShowAppHelp
	app.Commands = []cli.Command{
		startCommand,
	}

	app.Flags = []cli.Flag{
		HTTPHostFlag,
		HTTPPortFlag,
		GRPCHostFlag,
		GRPCPortFlag,
		ServiceNameFlag,
		JaegerAddressFlag,
		ConsulAddressFlag,
		AuthEndpointFlag,
		MasterDataEndpointFlag,
		NewFeedEndpointFlag,
		ProductEndpointFlag,
		SearchEndpointFlag,
		StorageEndpointFlag,
		TokenKeyFlag,
		AnalyticsEndpointFlag,
		NotificationEndpointFlag,
	}

}

// Start ...
func Start(ctx *cli.Context) error {
	if err := server.loadConfig(ctx); err != nil {
		return err
	}

	// if err := server.loadLogger(); err != nil {
	// 	return err
	// }

	if err := server.loadAuthenticator(); err != nil {
		return err
	}
	if err := server.loadGRPCClient(); err != nil {
		return err
	}
	if err := server.loadGRPCServer(); err != nil {
		return err
	}

	go server.startHTTPServer()
	if err := server.startGRPCServer(); err != nil {
		return err
	}
	return nil
}

func main() {
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
