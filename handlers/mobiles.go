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
func (m *Mobiles) GetBestNetwork(rw http.ResponseWriter, r *http.Request) {
	m.l.Println("Handle GET best network")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
		return
	}
	m.l.Println("GET best network for given mobile id: ", id)

	nets := data.GetNetworks()
	mobs := data.GetMobiles()

	availableNetworks := data.Networks{}

	var i = 0
	for i = 0; i < len(mobs); i++ {
		if mobs[i].ID == (id) {
			m.l.Println("Given mobile: ", mobs[i])
			break
		}
	}
	m.l.Println("i: ", i)
	if i > len(mobs) {
		http.Error(rw, "Mobile not found for given ID", http.StatusBadRequest)
		return
	}
	// find list of available networks
	for j := 0; j < len(nets); j++ {
		if !(mobs[i].Platform == nets[j].Platform) { // check if platform is the same
			continue
		}

		mobVersion := src.VersionOrdinal(mobs[i].OsVersion) // ex: 10.1
		validVer := false
		for v := 0; v < len(nets[j].SuppVersions); v += 2 { // ckeck if Os version is suported
			minVersion := src.VersionOrdinal(nets[j].SuppVersions[v])   // ex: 10.0.1
			maxVersion := src.VersionOrdinal(nets[j].SuppVersions[v+1]) // ex: 10.2.0
			if (minVersion <= mobVersion) && (mobVersion <= maxVersion) {
				validVer = true
				break
			}
		}
		if !validVer {
			continue
		}

		for k := 0; k < len(nets[j].CountryList); k++ {
			if mobs[i].CountryCode == nets[j].CountryList[k] { // check for allowed country
				availableNetworks = append(availableNetworks, nets[j])
				break
			}
		}
	}
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

type KeyMobile struct{}

func (m Mobiles) MiddlewareValidateMobile(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		mob := data.Mobile{}
		m.l.Println("rrrr", r.Body)
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
