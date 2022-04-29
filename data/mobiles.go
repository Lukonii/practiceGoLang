package data

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-playground/validator/v10"
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

// validation for platform
func (m *Mobile) Validate() error {
	validate := validator.New()
	return validate.Struct(m)
}

// GetMobiles returns a list of mobiles
func GetMobiles() Mobiles {
	return mobileList
}
func AddMobile(m *Mobile) {
	m.ID = getNextMobID()
	mobileList = append(mobileList, m)
}
func getNextMobID() int {
	mobListLen := mobileList[len(mobileList)-1]
	return mobListLen.ID + 1
}
func UpdateMobile(id int, m *Mobile) error {
	_, pos, err := findMobile(id)
	if err != nil {
		return err
	}

	m.ID = id
	mobileList[pos] = m

	return nil
}

var ErrMobileNotFound = fmt.Errorf("Mobile not found")

func findMobile(id int) (*Mobile, int, error) {
	for i, m := range mobileList {
		if m.ID == id {
			return m, i, nil
		}
	}

	return nil, -1, ErrMobileNotFound
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
	&Mobile{
		ID:          4,
		Platform:    "Android",
		OsVersion:   "7.1",
		AppName:     "AppOne",
		AppVersion:  "1.2.2",
		CountryCode: "SL",
	},
	&Mobile{
		ID:          5,
		Platform:    "IOS",
		OsVersion:   "10.1",
		AppName:     "AppOne",
		AppVersion:  "1.2.1",
		CountryCode: "CN",
	},
}
