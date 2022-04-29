package handlers

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/Lukonii/practiceGoLang/data"
	"github.com/gorilla/mux"
)

type Networks struct {
	l *log.Logger
}

// NewNetworks creates a networks handler with the given logger
func NewNetworks(l *log.Logger) *Networks {
	return &Networks{l}
}

func (n *Networks) GetNetworks(rw http.ResponseWriter, r *http.Request) {
	n.l.Println("Handle GET Networks")

	ln := data.GetNetworks()

	err := ln.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}
func (n *Networks) AddNetwork(rw http.ResponseWriter, r *http.Request) {
	n.l.Println("Handle POST Network")

	net := r.Context().Value(KeyNetwork{}).(data.Network)
	data.AddNetwork(&net)
}
func (n Networks) UpdateNetworks(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
		return
	}

	n.l.Println("Handle PUT Network")
	net := r.Context().Value(KeyNetwork{}).(data.Network)

	err = data.UpdateNetwork(id, &net)
	if err == data.ErrNetworkNotFound {
		http.Error(rw, "Network not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Network not found", http.StatusInternalServerError)
		return
	}
}

type KeyNetwork struct{}

func (n Networks) MiddlewareValidateNetwork(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		net := data.Network{}

		err := net.FromJSON(r.Body)
		if err != nil {
			n.l.Println("[ERROR] deserializinig network", err)
			http.Error(rw, "Error reading network", http.StatusBadRequest)
			return
		}
		// validate the network
		/*err = net.Validate()
		if err != nil {
			n.l.Println("[ERROR] validating network", err)
			http.Error(
				rw,
				fmt.Sprintf("Error validating network: %s", err),
				http.StatusBadRequest,
			)
			return
		}*/
		// add network to the context
		ctx := context.WithValue(r.Context(), KeyNetwork{}, net)
		r = r.WithContext(ctx)

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		if r != nil {
			next.ServeHTTP(rw, r)
		}
	})
}
