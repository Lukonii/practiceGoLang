package data

type Ad struct {
	ID        int    `json:"id"`
	NetworkId int    `json:"networkId"`
	AdType    string `json:"adType"`
	Score     int    `json:"score"`
	Link      string `json:"link"`
}

type Ads []*Ad

func GetAds() Ads {
	return adList
}

var adList = []*Ad{
	&Ad{
		ID:        1,
		NetworkId: 1,
		AdType:    "banner",
		Score:     30,
		Link:      "https://linktoad.com",
	},
	&Ad{
		ID:        2,
		NetworkId: 1,
		AdType:    "interstitial",
		Score:     30,
		Link:      "https://linktoad.com",
	},
	&Ad{
		ID:        3,
		NetworkId: 1,
		AdType:    "reward",
		Score:     30,
		Link:      "https://linktoad.com",
	},
}
