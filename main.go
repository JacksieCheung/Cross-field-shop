package main

import (
	"Cross-field-shop/model"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"time"

	"Cross-field-shop/config"
	"Cross-field-shop/log"
	"Cross-field-shop/router"
	"Cross-field-shop/router/middleware"
	"Cross-field-shop/service/history"

	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	cfg = pflag.StringP("config", "c", "", "apiserver config file path.")
)

func main() {
	pflag.Parse()

	// init config
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}

	// logger sync
	defer log.SyncLogger()

	// init db
	model.DB.Init()
	defer model.DB.Close()

	// init redis
	model.RedisDB.Init()
	defer model.RedisDB.Close()

	go history.GoCreateHistory()

	// Set gin mode.
	gin.SetMode(viper.GetString("runmode"))

	// Create the Gin engine.
	g := gin.New()

	// Routes.
	router.Load(
		// Cores.
		g,

		// MiddleWares.
		middleware.Logging(),
		middleware.RequestId(),
	)

	// Ping the server to make sure the router is working.
	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("The router has no response, or it might took too long to start up.",
				zap.String("reason", err.Error()))
		}
		log.Info("The router has been deployed successfully.")
	}()

	log.Info(
		fmt.Sprintf("Start to listening the incoming requests on http address: %s", viper.GetString("addr")))
	log.Info(http.ListenAndServe(viper.GetString("addr"), g).Error())
}

// pingServer pings the http server to make sure the router is working.
func pingServer() error {
	for i := 0; i < viper.GetInt("max_ping_count"); i++ {
		// Ping the server by sending a GET request to `/health`.
		resp, err := http.Get(viper.GetString("url") + "/sd/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}

		// Sleep for a second to continue the next ping.
		log.Info("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}
	return errors.New("Cannot connect to the router.")
}
