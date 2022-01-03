package main

import (
	"github.com/urfave/cli"
)

var (
	// HTTPHostFlag flag
	HTTPHostFlag = cli.StringFlag{
		Name:   "http.host",
		Usage:  "http host listen",
		EnvVar: "HTTP_HOST",
		Value:  "localhost",
	}
	// HTTPPortFlag flag
	HTTPPortFlag = cli.StringFlag{
		Name:   "http.port",
		Usage:  "http port listen",
		EnvVar: "HTTP_PORT",
		Value:  "10010",
	}
	// GRPCHostFlag flag
	GRPCHostFlag = cli.StringFlag{
		Name:   "grpc.host",
		Usage:  "grpc host listen",
		EnvVar: "GRPC_HOST",
		Value:  "localhost",
	}
	// GRPCPortFlag flag
	GRPCPortFlag = cli.StringFlag{
		Name:   "grpc.port",
		Usage:  "grpc port listen",
		EnvVar: "GRPC_PORT",
		Value:  "9002",
	}
	// ServiceNameFlag flag
	ServiceNameFlag = cli.StringFlag{
		Name:   "gateway",
		Usage:  "Service name",
		EnvVar: "GATEWAY_SERVICE_NAME",
		Value:  "Gateway",
	}
	// JaegerAddressFlag flag
	JaegerAddressFlag = cli.StringFlag{
		Name:   "jaeger_address",
		Usage:  "Jaeger Address used to connect",
		EnvVar: "USER_JAEGER_ADDRESS",
		Value:  "localhost:6831",
	}
	// ConsulAddressFlag flag
	ConsulAddressFlag = cli.StringFlag{
		Name:   "consul_address",
		Usage:  "Consul Address used to connect",
		EnvVar: "USER_CONSUL_ADDRESS",
		Value:  "localhost:8500",
	}
	// AuthEndpointFlag flag
	AuthEndpointFlag = cli.StringFlag{
		Name:   "auth_service.endpoint",
		Usage:  "auth service endpoint",
		EnvVar: "AUTH_SERVICE_ENDPOINT",
		Value:  "localhost:8283",
	}
	// SearchEndpointFlag flag
	ProductEndpointFlag = cli.StringFlag{
		Name:   "product_service.endpoint",
		Usage:  "product service endpoint",
		EnvVar: "product_service_ENDPOINT",
		Value:  "localhost:8282",
	}

	// SearchEndpointFlag flag
	SearchEndpointFlag = cli.StringFlag{
		Name:   "search_service.endpoint",
		Usage:  "search service endpoint",
		EnvVar: "SEARCH_SERVICE_ENDPOINT",
		Value:  "localhost:8585",
	}

	// MasterDataEndpointFlag flag
	MasterDataEndpointFlag = cli.StringFlag{
		Name:   "master_data_service.endpoint",
		Usage:  "master data service endpoint",
		EnvVar: "MASTER_DATA_SERVICE_ENDPOINT",
		Value:  "localhost:10011",
	}
	// NewFeedEndpointFlag flag
	NewFeedEndpointFlag = cli.StringFlag{
		Name:   "new_feed_service.endpoint",
		Usage:  "new feed service endpoint",
		EnvVar: "NEW_FEED_SERVICE_ENDPOINT",
		Value:  "localhost:10013",
	}
	// StorageEndpointFlag flag
	StorageEndpointFlag = cli.StringFlag{
		Name:   "storage_service.endpoint",
		Usage:  "storage service endpoint",
		EnvVar: "STORAGE_SERVICE_ENDPOINT",
		Value:  "localhost:10015",
	}
	// AnalyticsEndpointFlag flag
	AnalyticsEndpointFlag = cli.StringFlag{
		Name:   "analytics_service.endpoint",
		Usage:  "analytics service endpoint",
		EnvVar: "ANALYTICS_SERVICE_ENDPOINT",
		Value:  "localhost:10016",
	}
	// NotificationEndpointFlag flag
	NotificationEndpointFlag = cli.StringFlag{
		Name:   "notification_service.endpoint",
		Usage:  "notification service endpoint",
		EnvVar: "NOTIFICATION_SERVICE_ENDPOINT",
		Value:  "localhost:10016",
	}

	// TokenKeyFlag flag
	TokenKeyFlag = cli.StringFlag{
		Name:   "token.key",
		Usage:  "token key",
		EnvVar: "TOKEN_KEY",
		Value:  "ezexjSqGSzefFUFRxlxgMLNlPjrRWsiA",
	}
)

var (
	startCommand = cli.Command{
		Action:      migrateFlags(Start),
		Name:        "start",
		Usage:       "Bootstrap and start worker server",
		ArgsUsage:   "<genesisPath>",
		Flags:       []cli.Flag{},
		Category:    "Crawler Worker",
		Description: `Used to start crawler worker, clone data from omada cloud`,
	}
)

// NewApp creates an app with sane defaults.
func NewApp() *cli.App {
	app := cli.NewApp()
	app.Action = cli.ShowAppHelp
	app.Name = "QandA"
	app.Author = "Le Duy Dat"
	app.Email = "duyledat197@gmail.com"
	app.Usage = "question and answer"
	return app
}
