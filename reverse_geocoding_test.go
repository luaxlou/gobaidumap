package gobaidumap

import (
	"encoding/json"
	"log"
	"testing"
	"time"
)

func TestBaiduMap_ReverseGeocoding(t *testing.T) {
	m := New("--AK--")

	res, err := m.ReverseGeocoding(27.990458, 86.934165,time.Second *5)

	if err != nil {
		log.Println(err.Error())
	} else {

		bs, _ := json.Marshal(&res)
		log.Println(string(bs))
	}
}
