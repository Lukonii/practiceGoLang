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
	&Ad{
		ID:        11,
		NetworkId: 3,
		AdType:    "reward",
		Score:     20,
		Link:      "https://linktoad.com",
	},
	&Ad{
		ID:        12,
		NetworkId: 3,
		AdType:    "reward",
		Score:     30,
		Link:      "https://linktoad.com",
	},
	&Ad{
		ID:        13,
		NetworkId: 3,
		AdType:    "reward",
		Score:     20,
		Link:      "https://linktoad.com",
	},
	&Ad{
		ID:        14,
		NetworkId: 3,
		AdType:    "reward",
		Score:     10,
		Link:      "https://linktoad.com",
	},
	&Ad{
		ID:        15,
		NetworkId: 3,
		AdType:    "banner",
		Score:     20,
		Link:      "https://linktoad.com",
	},
	&Ad{
		ID:        16,
		NetworkId: 3,
		AdType:    "banner",
		Score:     30,
		Link:      "https://linktoad.com",
	},
	&Ad{
		ID:        17,
		NetworkId: 3,
		AdType:    "banner",
		Score:     20,
		Link:      "https://linktoad.com",
	},
	&Ad{
		ID:        18,
		NetworkId: 3,
		AdType:    "banner",
		Score:     10,
		Link:      "https://linktoad.com",
	},
}
