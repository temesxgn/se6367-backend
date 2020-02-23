package main

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/temesxgn/se6367-backend/config"
	logging "github.com/temesxgn/se6367-backend/logger"
	"github.com/temesxgn/se6367-backend/metrics"
	"github.com/temesxgn/se6367-backend/server"
)

var logger *logrus.Entry

func init() {
	logger = logging.CreateLogger(metrics.BackendServer)
}

func main() {
	e := server.New()
	logger.Fatal(e.Start(fmt.Sprintf(":%v", config.GetServerPort())))
}
