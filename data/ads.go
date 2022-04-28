package data

import (
	"encoding/json"
	"io"
)

type Ad struct {
	ID        int    `json:"id"`
	NetworkId int    `json:"networkId"`
	AdType    string `json:"adType"`
	Score     int    `json:"score"`
	Link      string `json:"link"`
}

var AdTypes = [3]string{"banner", "interstitial", "reward"}

type Ads []*Ad

func GetAds() Ads {
	return adList
}
func (a *Ads) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(a)
}
func (a *Ad) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(a)
}
func (ads *Ads) FilterAdsByType(adt int) Ads {
	var filtered = Ads{}

	adtype := ""
	switch adt {
	case 1:
		adtype = AdTypes[0] //banner
	case 2:
		adtype = AdTypes[1] //interstitial
	case 3:
		adtype = AdTypes[2] //reward
	}

	for i := 0; i < len(adList); i++ {
		if adList[i].AdType == adtype {
			filtered = append(filtered, adList[i])
		}
	}
	return filtered
}

var adList = []*Ad{
	&Ad{
		ID:        1,
		NetworkId: 1,
		AdType:    "banner",
		Score:     15,
		Link:      "https://linktoad.com",
	},
	&Ad{
		ID:        2,
		NetworkId: 1,
		AdType:    "interstitial",
		Score:     15,
		Link:      "https://linktoad.com",
	},
	&Ad{
		ID:        3,
		NetworkId: 1,
		AdType:    "reward",
		Score:     5,
		Link:      "https://linktoad.com",
	},
	&Ad{
		ID:        4,
		NetworkId: 1,
		AdType:    "reward",
		Score:     20,
		Link:      "https://linktoad.com",
	},
	&Ad{
		ID:        5,
		NetworkId: 1,
		AdType:    "reward",
		Score:     10,
		Link:      "https://linktoad.com",
	},
	&Ad{
		ID:        6,
		NetworkId: 1,
		AdType:    "reward",
		Score:     30,
		Link:      "https://linktoad.com",
	},
	&Ad{
		ID:        7,
		NetworkId: 1,
		AdType:    "reward",
		Score:     30,
		Link:      "https://linktoad.com",
	},
	&Ad{
		ID:        8,
		NetworkId: 1,
		AdType:    "reward",
		Score:     10,
		Link:      "https://linktoad.com",
	},
	&Ad{
		ID:        9,
		NetworkId: 1,
		AdType:    "reward",
		Score:     10,
		Link:      "https://linktoad.com",
	},
	&Ad{
		ID:        10,
		NetworkId: 1,
		AdType:    "reward",
		Score:     5,
		Link:      "https://linktoad.com",
	},
}
