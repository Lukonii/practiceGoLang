package src

import (
	"encoding/json"
	"net/http"

	"github.com/Lukonii/practiceGoLang/data"
)

func GetNeworskForGivenMobile(id int, rw http.ResponseWriter) data.Networks {
	nets := data.GetNetworks()
	mobs := data.GetMobiles()

	availableNetworks := data.Networks{}

	var i = 0
	for i = 0; i < len(mobs); i++ {
		if mobs[i].ID == (id) {
			json.NewEncoder(rw).Encode("Selected mobile: ")
			json.NewEncoder(rw).Encode(mobs[i])
			break
		}
	}
	if i > len(mobs) {
		http.Error(rw, "Mobile not found for given ID", http.StatusBadRequest)
		return availableNetworks
	}
	// find list of available networks
	for j := 0; j < len(nets); j++ {
		if !(mobs[i].Platform == nets[j].Platform) { // check if platform is the same
			continue
		}

		mobVersion := VersionOrdinal(mobs[i].OsVersion) // ex: 10.1
		validVer := false
		for v := 0; v < len(nets[j].SuppVersions); v += 2 { // ckeck if Os version is suported
			minVersion := VersionOrdinal(nets[j].SuppVersions[v])   // ex: 10.0.1
			maxVersion := VersionOrdinal(nets[j].SuppVersions[v+1]) // ex: 10.2.0
			if (minVersion <= mobVersion) && (mobVersion <= maxVersion) {
				validVer = true
				break
			}
		}
		if !validVer {
			continue
		}

		for k := 0; k < len(nets[j].CountryList); k++ {
			if mobs[i].CountryCode == nets[j].CountryList[k] { // check for allowed country
				availableNetworks = append(availableNetworks, nets[j])
				break
			}
		}
	}
	json.NewEncoder(rw).Encode("Available networks: ")
	availableNetworks.ToJSON(rw)
	return availableNetworks
}
func FindBestAdForGivenMobileAndType(networks data.Networks, adtype int, rw http.ResponseWriter) {
	ads := data.GetAds()
	ads = FilterAdsByType(ads, adtype)
	json.NewEncoder(rw).Encode("Filtered by type: ")
	ads.ToJSON(rw)
	ads = FilterAdsByAvailableNetworks(ads, networks)
	json.NewEncoder(rw).Encode("Filtered by net: ")
	ads.ToJSON(rw)
	SortAdsByScore(ads)
	json.NewEncoder(rw).Encode("Best ads to show: ")
	ads.ToJSON(rw)
}
func FilterAdsByType(ads data.Ads, adt int) data.Ads {
	var filtered = data.Ads{}

	adtype := ""
	switch adt {
	case 1:
		adtype = data.AdTypes[0] //banner
	case 2:
		adtype = data.AdTypes[1] //interstitial
	case 3:
		adtype = data.AdTypes[2] //reward
	}

	for i := 0; i < len(ads); i++ {
		if ads[i].AdType == adtype {
			filtered = append(filtered, ads[i])
		}
	}
	return filtered
}
func FilterAdsByAvailableNetworks(ads data.Ads, net data.Networks) data.Ads {
	var filtered = data.Ads{}
	for i := 0; i < len(ads); i++ {
		for j := 0; j < len(net); j++ {
			if ads[i].NetworkId == net[j].ID {
				filtered = append(filtered, ads[i])
			}
		}
	}
	return filtered
}
func SortAdsByScore(ads data.Ads) {
	// quick sort implementation for ads score
	// we assume that score is calculated and inserted in db for every ad
	QuickSort(ads, 0, len(ads)-1)
}
func QuickSort(arr data.Ads, start, end int) {
	if (end - start) < 1 {
		return
	}

	pivot := arr[end]
	splitIndex := start

	for i := start; i < end; i++ {
		if arr[i].Score > pivot.Score {
			temp := arr[splitIndex]

			arr[splitIndex] = arr[i]
			arr[i] = temp

			splitIndex++
		}
	}

	arr[end] = arr[splitIndex]
	arr[splitIndex] = pivot

	QuickSort(arr, start, splitIndex-1)
	QuickSort(arr, splitIndex+1, end)
}

// converts to version
func VersionOrdinal(version string) string {
	const maxByte = 1<<8 - 1
	vo := make([]byte, 0, len(version)+8)
	j := -1
	for i := 0; i < len(version); i++ {
		b := version[i]
		if '0' > b || b > '9' {
			vo = append(vo, b)
			j = -1
			continue
		}
		if j == -1 {
			vo = append(vo, 0x00)
			j = len(vo) - 1
		}
		if vo[j] == 1 && vo[j+1] == '0' {
			vo[j+1] = b
			continue
		}
		if vo[j]+1 > maxByte {
			panic("VersionOrdinal: invalid version")
		}
		vo = append(vo, b)
		vo[j]++
	}
	return string(vo)
}
func IsNetworkUsingAds(net *data.Network, ads data.Ads) bool {
	for i := 0; i < len(ads); i++ {
		if net.ID == ads[i].NetworkId {
			return true
		}
	}
	return false
}
