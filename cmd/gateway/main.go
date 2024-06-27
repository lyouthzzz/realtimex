package main

import (
	"flag"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/lyouthzzz/realtimex/internal/gateway"
	"github.com/lyouthzzz/realtimex/internal/gateway/conf"
	_ "github.com/lyouthzzz/realtimex/pkg/log"
)

var configPath string

func main() {
	flag.StringVar(&configPath, "conf", "configs/gateway.yaml", "config path, eg: -conf default.yaml")
	//flag.Parse()

	c := config.New(
		config.WithSource(
			file.NewSource(configPath),
		),
	)

	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	app := kratos.New(
		kratos.Name("realtimex-gateway"),
		kratos.Version("v1.0.0"),
		kratos.Server(
			gateway.NewServer(),
		),
	)

	if err := app.Run(); err != nil {
		panic(err)
	}
}
