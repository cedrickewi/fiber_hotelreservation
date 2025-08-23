package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/cedrickewi/hotel-reservation/internals/services"
	"go.uber.org/zap"
)

type application struct {
	store  *services.Storage
	config config
	logger *zap.SugaredLogger
}

type config struct {
	port        int
	addr        string
	frontendURL string
	apiURL      string
}

type healthCheckResponse struct {
	Msg string `json:"message"`
	Code int	`json:"code"`
}

func (app *application) run(mux http.Handler, port int) error {
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	app.logger.Infof("Starting server on port %d", port)

	err := srv.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}
