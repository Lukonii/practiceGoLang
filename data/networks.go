package data

type Network struct {
	ID              int      `json:"id"`
	Name            string   `json:"name"`
	Platform        string   `json:"platform"`
	OldestOsVersion string   `json:"oldestOsVersion"`
	CountryList     []string `json:"countryList"`
}

type Networks []*Network

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
