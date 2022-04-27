package handlers

import (
	"context"
	"log"
	"net/http"

	"github.com/Lukonii/practiceGoLang/data"
)

type Mobiles struct {
	l *log.Logger
}

// NewMobiles creates a Mobiles handler with the given logger
func NewMobiles(l *log.Logger) *Mobiles {
	return &Mobiles{l}
}

func (n *Mobiles) GetMobiles(rw http.ResponseWriter, r *http.Request) {
	n.l.Println("Handle GET Mobiles")

	ln := data.GetMobiles()

	err := ln.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

type KeyMobile struct{}

func (n Mobiles) MiddlewareValidateMobile(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		mob := data.Mobile{}

		err := mob.FromJSON(r.Body)

		if err != nil {
			n.l.Println("[ERROR] deserializinig mobile", err)
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
