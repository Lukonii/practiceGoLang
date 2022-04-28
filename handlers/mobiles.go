package handlers

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/Lukonii/practiceGoLang/data"
	"github.com/Lukonii/practiceGoLang/src"
	"github.com/gorilla/mux"
)

type Mobiles struct {
	l *log.Logger
}

// NewMobiles creates a Mobiles handler with the given logger
func NewMobiles(l *log.Logger) *Mobiles {
	return &Mobiles{l}
}

func (m *Mobiles) GetMobiles(rw http.ResponseWriter, r *http.Request) {
	m.l.Println("Handle GET Mobiles")

	ln := data.GetMobiles()

	err := ln.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}
func (m *Mobiles) GetMobileNetworks(rw http.ResponseWriter, r *http.Request) {
	m.l.Println("Handle GET best network")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
		return
	}
	src.GetNeworskForGivenMobile(id, rw)
	/*
		for j := 0; j < len(availableNetworks); j++ {
			m.l.Println("Good network: ", availableNetworks[j])
		}
	*/
	/*vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
		return
	}
	m.l.Println("GET best network for given mobile id: ", id)
	mob := r.Context().Value(KeyMobile{}).(data.Mobile) //cast interface to data
	m.l.Println("Mobile: ", mob)
	*/
}
func (m *Mobiles) GetSuggestedAds(rw http.ResponseWriter, r *http.Request) {
	m.l.Println("Sugg ads")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
		return
	}
	adt, err := strconv.Atoi(vars["adt"])
	if err != nil {
		http.Error(rw, "Unable to convert ad type", http.StatusBadRequest)
		return
	}
	src.FindBestAdForGivenMobileAndType(src.GetNeworskForGivenMobile(id, rw), adt, rw)

}

type KeyMobile struct{}

func (m Mobiles) MiddlewareValidateMobile(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		mob := data.Mobile{}
		err := mob.FromJSON(r.Body)
		if err != nil {
			m.l.Println("[ERROR] deserializinig mobile", err)
			http.Error(rw, "Error reading mobile", http.StatusBadRequest)
			return
		}
		// validate the mobile
		/*err = net.Validate()
		if err != nil {
			n.l.Println("[ERROR] validating mobile", err)
			http.Error(
				rw,
				fmt.Sprintf("Error validating mobile: %s", err),
				http.StatusBadRequest,
			)
			return
		}*/
		// add mobile to the context
		ctx := context.WithValue(r.Context(), KeyMobile{}, mob)
		r = r.WithContext(ctx)

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(rw, r)
	})
}
