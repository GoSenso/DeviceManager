package config

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/hashicorp/consul/api"
)

const AppName = "DeviceManager"

type Config struct {
	CassandraAdress string //TODO:
}

func LoadConfig() *Config {

	consulCfg := api.DefaultConfig()
	consulCfg.Address = "http://consul.gosense.com"
	consulCl, err := api.NewClient(consulCfg)
	if err != nil {
		log.Fatal(fmt.Errorf("failed to connect wiht consul: %w", err))
	}

	kv := consulCl.KV()
	pair, _, err := kv.Get(AppName, nil)
	if err != nil {
		log.Fatalf("Error during downloading key value: %v", err)
	}

	if pair == nil {
		log.Printf("Key was not found.\n")
	} else {
		fmt.Printf("Key Value DeviceManager': %s\n", pair.Value)
	}

	var appConfig Config
	err = json.Unmarshal(pair.Value, &appConfig)
	if err != nil {
		log.Fatalf("Error durign deserialization: %v", err)
	}
	return &appConfig
}
