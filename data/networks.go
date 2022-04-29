package data

import (
	"encoding/json"
	"fmt"
	"io"
)

type Network struct {
	ID           int      `json:"id" validate:"required"`
	Name         string   `json:"name"`
	Platform     string   `json:"platform"`
	SuppVersions []string `json:"suppVersions"`
	CountryList  []string `json:"countryList"`
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
func AddNetwork(n *Network) {
	n.ID = getNextNetID()
	networkList = append(networkList, n)
}
func getNextNetID() int {
	netListLen := networkList[len(networkList)-1]
	return netListLen.ID + 1
}
func UpdateNetwork(id int, n *Network) error {
	_, pos, err := findNetwork(id)
	if err != nil {
		return err
	}

	n.ID = id
	networkList[pos] = n

	return nil
}

var ErrNetworkNotFound = fmt.Errorf("Network not found")

func findNetwork(id int) (*Network, int, error) {
	for i, n := range networkList {
		if n.ID == id {
			return n, i, nil
		}
	}

	return nil, -1, ErrNetworkNotFound
}

var AdMobCountries = [6]string{"CN", "IT", "RS", "SL", "A", "D"}

var FacebookCountries = [5]string{"IT", "RS", "SL", "A", "D"}

var networkList = []*Network{
	&Network{
		ID:           1,
		Name:         "AdMob",
		Platform:     "Android",
		SuppVersions: []string{"8.0.0", "8.8.7", "10.0.0", "15.5.0"}, //between 1-2, 3-4
		CountryList:  AdMobCountries[:],
	},
	&Network{
		ID:           2,
		Name:         "AdMob",
		Platform:     "IOS",
		SuppVersions: []string{"10.0.0", "14.3.0"},
		CountryList:  AdMobCountries[:],
	},
	&Network{
		ID:           3,
		Name:         "Facebook",
		Platform:     "Android",
		SuppVersions: []string{"8.0.0", "15.5.0"},
		CountryList:  FacebookCountries[:],
	},
	&Network{
		ID:           4,
		Name:         "Facebook",
		Platform:     "IOS",
		SuppVersions: []string{"10.0.0", "14.3.0"},
		CountryList:  FacebookCountries[:],
	},
}
