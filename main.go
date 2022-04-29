package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Lukonii/practiceGoLang/handlers"
	"github.com/gorilla/mux"
)

func main() {
	l := log.New(os.Stdout, "ad-mediation-api ", log.LstdFlags)

	// create the handlers
	nh := handlers.NewNetworks(l)
	mh := handlers.NewMobiles(l)
	dh := handlers.NewDashboard(l)
	ah := handlers.NewAd(l)

	// create a new serve mux router and register the handlers
	router := mux.NewRouter()

	getRouter := router.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/network", nh.GetNetworks)
	getRouter.HandleFunc("/mobile", mh.GetMobiles)
	getRouter.HandleFunc("/mobile-{id:[0-9]+}", mh.GetMobileNetworks)
	getRouter.HandleFunc("/mobile-{id:[0-9]+}-adtype-{adt:[1-3]}", mh.GetSuggestedAds)
	getRouter.HandleFunc("/dashboard", dh.GetDashboard)
	getRouter.HandleFunc("/ad", ah.GetAds)

	putRouterNet := router.PathPrefix("/network").Methods(http.MethodPut).Subrouter()
	putRouterNet.Use(nh.MiddlewareValidateNetwork)
	putRouterNet.HandleFunc("/{id:[0-9]+}", nh.UpdateNetworks)
	putRouterAd := router.PathPrefix("/ad").Methods(http.MethodPut).Subrouter()
	putRouterAd.Use(ah.MiddlewareValidateAd)
	putRouterAd.HandleFunc("/{id:[0-9]+}", ah.UpdateAds)
	putRouterMob := router.PathPrefix("/mobile").Methods(http.MethodPut).Subrouter()
	putRouterMob.Use(mh.MiddlewareValidateMobile)
	putRouterMob.HandleFunc("/{id:[0-9]+}", mh.UpdateMobiles)

	postRouterNet := router.Methods(http.MethodPost).Subrouter()
	postRouterNet.Use(nh.MiddlewareValidateNetwork)
	postRouterNet.HandleFunc("/network", nh.AddNetwork)
	postRouterAd := router.Methods(http.MethodPost).Subrouter()
	postRouterAd.Use(ah.MiddlewareValidateAd)
	postRouterAd.HandleFunc("/ad", ah.AddAd)
	postRouterMob := router.Methods(http.MethodPost).Subrouter()
	postRouterMob.Use(mh.MiddlewareValidateMobile)
	postRouterMob.HandleFunc("/mobile", mh.AddMobile)

	// create a new server
	s := http.Server{
		Addr:         ":9090",          // configure the bind address
		Handler:      router,           // set the default handler
		ErrorLog:     l,                // set the logger for the server
		ReadTimeout:  5 * time.Second,  // max time to read request from the client
		WriteTimeout: 10 * time.Second, // max time to write response to the client
		IdleTimeout:  120 * time.Second,
	}

	// start the server
	go func() {
		l.Println("Starting server on port 9090")

		err := s.ListenAndServe()
		if err != nil {
			l.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block until a signal is received.
	sig := <-c
	log.Println("Got signal:", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)
}
