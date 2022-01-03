package main

import (
	"fmt"
	"strconv"

	"github.com/urfave/cli"
)

// Config ...
type Config struct {
	HTTP ConnAddress
	GRPC ConnAddress

	JaegerAddress string
	ConsulAddress string

	ServiceName      string
	AuthEndpoint     string
	UserEndpoint     string
	ProductEndpoint  string
	CategoryEndpoint string
	SearchEndpoint   string

	CookieName      string
	TokenKey        string
	TokenExpiration int
}

// Mongo ...
type Mongo struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

func (m *Mongo) getConnectString() string {
	conn := "mongodb://"
	if m.Username != "" {
		conn += fmt.Sprintf("%s:%s@", m.Username, m.Password)
	}
	conn += fmt.Sprintf("%s:%s/%s", m.Host, m.Port, m.Database)
	return conn
}

// ConnAddress ...
type ConnAddress struct {
	Host string
	Port string
}

func (c *ConnAddress) String() string {
	return fmt.Sprintf("%s:%s", c.Host, c.Port)
}

func (c *ConnAddress) GetPortInt() int {
	i2, _ := strconv.ParseInt(c.Port, 10, 64)
	return int(i2)
}

func migrateFlags(action func(ctx *cli.Context) error) func(*cli.Context) error {
	return func(ctx *cli.Context) error {
		for _, name := range ctx.FlagNames() {
			ctx.GlobalSet(name, ctx.String(name))
		}
		return action(ctx)
	}
}
