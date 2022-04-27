package data

import (
	"encoding/json"
	"io"
)

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

// ToJSON serializes the contents of the collection to JSON
// NewEncoder provides better performance than json.Unmarshal as it does not
// have to buffer the output into an in memory slice of bytes
// this reduces allocations and the overheads of the service
func (m *Mobiles) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(m)
}

func (m *Mobile) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(m)
}

// GetMobiles returns a list of mobiles
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
