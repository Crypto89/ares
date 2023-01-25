package main

import (
	"net/http"
	"os"

	"github.com/crypto89/ares/api"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func main() {
	apiKey := os.Getenv("OGAME_API_KEY")

	v1, err := api.NewAPI(apiKey)
	if err != nil {
		logrus.WithError(err).Fatal("failed to create API instance")
	}

	r := mux.NewRouter()
	r.PathPrefix("/api").Handler(v1.Handler)

    logrus.Infof("Starting listering on %s", "127.0.0.1:34343")
	if err := http.ListenAndServe("127.0.0.1:34343", r); err != nil {
		logrus.WithError(err).Fatal("failed to create lister")
	}
}
