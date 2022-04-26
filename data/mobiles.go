package data

type Mobile struct {
	ID          int    `json:"id"`
	Platform    string `json:"platform" validate:"required"`
	OsVersion   string `json:"osVersion"`
	AppName     string `json:"appName"`
	AppVersion  string `json:"appVersion"`
	CountryCode string `json:"countryCode"`
}

// Products is a collection of Product
type Mobiles []*Mobile

// GetProducts returns a list of products
func GetMobiles() Mobiles {
	return mobileList
}

var mobileList = []*Mobile{
	&Mobile{
		ID:          1,
		Platform:    "Android",
		OsVersion:   "8.1",
		AppName:     "AppOne",
		AppVersion:  "1.2.1",
		CountryCode: "CN",
	},
	&Mobile{
		ID:          2,
		Platform:    "Android",
		OsVersion:   "9.1",
		AppName:     "AppOne",
		AppVersion:  "1.2.2",
		CountryCode: "IT",
	},
	&Mobile{
		ID:          3,
		Platform:    "Android",
		OsVersion:   "9.1",
		AppName:     "AppOne",
		AppVersion:  "1.2.2",
		CountryCode: "SL",
	},
}
