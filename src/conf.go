package main

import (
	"log"

	"github.com/BurntSushi/toml"
)

var (
	conf = tConfig{}
)

type tConfig struct {
	IPRetrieval []string `toml:"ip_retrieval"`
	TorCheck    string   `toml:"tor_check"`
	More        string   `toml:"more"`
	UA          string   `toml:"ua"`
}

func readConfig() {
	var err error
	_, err = toml.Decode(configString, &conf)
	if err != nil {
		log.Fatalf("Error unmarshal %q, %q", configString, err)
	}
}
