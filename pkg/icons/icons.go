package icons

import (
	"io/ioutil"
)

const basePath = "/usr/share/icons/"

type Icons struct {
	Low    []byte
	Medium []byte
	High   []byte
}

func Read(theme string) *Icons {
	path := basePath + theme + "/24x24/panel/"
	return &Icons{
		Low:    getIconData(path + "volume-level-low.svg"),
		Medium: getIconData(path + "volume-level-medium.svg"),
		High:   getIconData(path + "volume-level-high.svg"),
	}
}

func getIconData(path string) []byte {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return data
}
