package controllers

import (
	"log"

	"github.com/BurntSushi/toml"
)

//API struct com o servidor e banco de dados
type Api struct {
	Server   string
	Database string
}

func (c *Api) Read() {
	if _, err := toml.DecodeFile("config.toml", &c); err != nil {
		log.Fatal(err)
	}

}
