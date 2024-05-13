package rest

import (
	"context"
	"errors"
	"fmt"
	"github.com/doxanocap/pkg/router"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"push-api/internal/interfaces"
	"push-api/internal/models"
	"push-api/server/rest/middlewares"
	"syscall"
	"time"
)

type REST struct {
	log     *zap.Logger
	config  *models.Config
	router  *gin.Engine
	server  *http.Server
	service interfaces.IService

	middlewares *middlewares.Middlewares
}

func InitREST(
	config *models.Config,
	service interfaces.IService,
	log *zap.Logger) *REST {
	r := &REST{
		log:     log,
		config:  config,
		service: service,
		router:  router.InitGinRouter(config.ENV),

		middlewares: middlewares.InitMiddlewares(service, log.Named("[MIDDLEWARE]")),
	}
	r.InitRoutes()
	return r
}

func (r *REST) Run() {
	r.server = &http.Server{
		Addr:           ":" + r.config.ServerPORT,
		Handler:        r.router,
		MaxHeaderBytes: 1 << 20, // 1mb
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	go func() {
		r.log.Info(fmt.Sprintf("REST server running at: %s", r.config.ServerPORT))
		if err := r.server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			r.log.Error(fmt.Sprintf("r.ListenAndServer: %v", err))
		}
	}()

	go func() {
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
		<-ch

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := r.server.Shutdown(ctx); err != nil {
			r.log.Error(fmt.Sprintf("r.server.Stop: %s", err))
		}
		r.log.Info("REST graceful shut down...")
	}()
}
