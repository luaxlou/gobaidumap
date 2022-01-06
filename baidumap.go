package gobaidumap

const BASE_URL = "http://api.map.baidu.com"

type BaiduMap struct {
	AK string
}

func New(ak string) *BaiduMap {

	return &BaiduMap{
		AK:ak,
	}
}
