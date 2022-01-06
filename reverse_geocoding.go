package gobaidumap

import (
	"errors"
	"fmt"
	"github.com/guonaihong/gout"
	"strconv"
	"time"
)

type ReverseGeocodingAddressComponent struct {
	Country      string `json:"country"`
	Province     string `json:"province"`
	City         string `json:"city"`
	District     string `json:"district"`
	Street       string `json:"street"`
	StreetNumber string `json:"street_number"`
}

type ReverseGeocodingRes struct {
	Result struct {
		Status           int                              `json:"status"`
		AddressComponent ReverseGeocodingAddressComponent `json:"addressComponent"`
	} `json:"result"`
}

var uri_reverse_geocoding = "/reverse_geocoding/v3/?ak=%s&output=json&coordtype=bd09ll&location=%f,%f"

func (m *BaiduMap) ReverseGeocoding(lat, lon float64, timeout time.Duration) (*ReverseGeocodingAddressComponent, error) {

	var res ReverseGeocodingRes

	url := BASE_URL + fmt.Sprintf(uri_reverse_geocoding, m.AK, lat, lon)
	err := gout.GET(url).SetTimeout(timeout).BindJSON(&res).Do()

	if err != nil {
		return nil, err
	}

	if res.Result.Status != 0 {
		return nil, errors.New("err:" + strconv.Itoa(res.Result.Status))
	}

	return &res.Result.AddressComponent, nil

}
