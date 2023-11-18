package main

import (
	"flag"
	"path/filepath"

	"github.com/nextmicro/logger"
	"github.com/nextmicro/next"
	"github.com/nextmicro/next/config"
	"github.com/nextmicro/next/transport/grpc"
	"github.com/nextmicro/next/transport/http"
	"next.data.layout/internal/conf"

	_ "go.uber.org/automaxprocs"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// flagconf is the config flag.
	flagconf string
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
}

func newApp(logger logger.Logger, gs *grpc.Server, hs *http.Server) (*next.Next, error) {
	return next.New(
		next.Logger(logger),
		next.Server(
			gs,
			hs,
		),
	)
}

func main() {
	flag.Parse()
	c, err := config.Init(filepath.Join(flagconf, config.BizConfFile()))
	if err != nil {
		panic(err)
	}
	defer c.Close()

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	app, cleanup, err := wireApp(bc.Data, logger.DefaultLogger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}
