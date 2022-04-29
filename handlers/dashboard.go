package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/Lukonii/practiceGoLang/data"
	"github.com/Lukonii/practiceGoLang/src"
)

type Dashboard struct {
	l *log.Logger
}

func NewDashboard(l *log.Logger) *Dashboard {
	return &Dashboard{l}
}
func (d *Dashboard) Wellcome(rw http.ResponseWriter, r *http.Request) {
	json.NewEncoder(rw).Encode("This app has no frontend!")
	json.NewEncoder(rw).Encode("Read more: https://github.com/Lukonii/practiceGoLang/blob/main/README.md")
}
func (d *Dashboard) GetDashboard(rw http.ResponseWriter, r *http.Request) {
	d.l.Println("Handle GET Dashboard")
	ads := data.GetAds()
	net := data.GetNetworks()

	// We will join ads and networks by id and if there is at least one ad for specific type well display it
	// To keep this part simple we will ignore other conditions like platrofm and versions
	// it is calculated in GetNeworskForGivenMobile part in calculations
	json.NewEncoder(rw).Encode("Banner:")
	PrintCountriesPerNet(rw, ads, net, 1)

	json.NewEncoder(rw).Encode("")
	json.NewEncoder(rw).Encode("Interstitial:")
	PrintCountriesPerNet(rw, ads, net, 2)

	json.NewEncoder(rw).Encode("")
	json.NewEncoder(rw).Encode("Reward:")
	PrintCountriesPerNet(rw, ads, net, 3)

}
func PrintCountriesPerNet(rw http.ResponseWriter, ads data.Ads, net data.Networks, adtype int) {
	for i := 0; i < len(net); i++ {
		adst := src.FilterAdsByType(ads, adtype)
		if src.IsNetworkUsingAds(net[i], adst) {
			s := strings.Join(net[i].CountryList, ", ")
			json.NewEncoder(rw).Encode("Countries: " + s + " Net: " + net[i].Name)
		}
	}
}
