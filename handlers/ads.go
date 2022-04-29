package handlers

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/Lukonii/practiceGoLang/data"
	"github.com/gorilla/mux"
)

type Ads struct {
	l *log.Logger
}

func NewAd(l *log.Logger) *Ads {
	return &Ads{l}
}
func (a *Ads) GetAds(rw http.ResponseWriter, r *http.Request) {
	a.l.Println("Handle GET Ads")

	listAds := data.GetAds()

	err := listAds.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}
func (a *Ads) AddAd(rw http.ResponseWriter, r *http.Request) {
	a.l.Println("Handle POST Ad")

	ad := r.Context().Value(KeyAd{}).(data.Ad)
	data.AddAd(&ad)
}
func (a Ads) UpdateAds(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
		return
	}

	a.l.Println("Handle PUT Ad", id)
	ad := r.Context().Value(KeyAd{}).(data.Ad)

	err = data.UpdateAd(id, &ad)
	if err == data.ErrAdNotFound {
		http.Error(rw, "Ad not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Ad not found", http.StatusInternalServerError)
		return
	}
}

type KeyAd struct{}

func (a Ads) MiddlewareValidateAd(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		ad := data.Ad{}

		err := ad.FromJSON(r.Body)
		if err != nil {
			a.l.Println("[ERROR] deserializinig ad", err)
			http.Error(rw, "Error reading ad", http.StatusBadRequest)
			return
		}
		// add ad to the context
		ctx := context.WithValue(r.Context(), KeyAd{}, ad)
		r = r.WithContext(ctx)

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		if r != nil {
			next.ServeHTTP(rw, r)
		}
	})
}
