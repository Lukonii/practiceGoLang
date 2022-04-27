package data

import (
	"encoding/json"
	"io"
)

type Network struct {
	ID              int      `json:"id" validate:"required"`
	Name            string   `json:"name"`
	Platform        string   `json:"platform"`
	OldestOsVersion string   `json:"oldestOsVersion"`
	CountryList     []string `json:"countryList"`
}

type Networks []*Network

// ToJSON serializes the contents of the collection to JSON
// NewEncoder provides better performance than json.Unmarshal as it does not
// have to buffer the output into an in memory slice of bytes
// this reduces allocations and the overheads of the service
func (n *Networks) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(n)
}
func (n *Network) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(n)
}

func GetNetworks() Networks {
	return networkList
}

var AdMobCountries = [6]string{"CN", "IT", "RS", "SL", "A", "D"}

var FacebookCountries = [5]string{"IT", "RS", "SL", "A", "D"}

var networkList = []*Network{
	&Network{
		ID:              1,
		Name:            "AdMob",
		Platform:        "Android",
		OldestOsVersion: "9.0",
		CountryList:     AdMobCountries[:],
	},
	&Network{
		ID:              2,
		Name:            "Facebook",
		Platform:        "Android",
		OldestOsVersion: "9.0",
		CountryList:     FacebookCountries[:],
	},
}
