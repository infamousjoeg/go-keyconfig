package main

import (
	"fmt"

	"github.com/infamousjoeg/go-keyconfig"
)

type exampleConfig struct {
	BaseURL  string `json:"BaseURL"`
	Username string `json:"Username"`
	Password string `json:"Password"`
	AuthType string `json:"AuthType"`
}

func exampleSetConfig() {
	config := exampleConfig{
		BaseURL:  "https://example.com",
		Username: "username",
		Password: "password",
		AuthType: "basic",
	}

	err := keyconfig.SetConfig("example", config)
	if err != nil {
		panic(err)
	}
}

func exampleGetConfig() {
	var config exampleConfig

	err := keyconfig.GetConfig("example", &config)
	if err != nil {
		panic(err)
	}

	if config.BaseURL != "https://example.com" {
		panic(config.BaseURL)
	}

	if config.Username != "username" {
		panic(config.Username)
	}

	if config.Password != "password" {
		panic(config.Password)
	}

	if config.AuthType != "basic" {
		panic(config.AuthType)
	}

	fmt.Printf("%+v\n", config.BaseURL)
	fmt.Printf("%+v\n", config.Username)
	fmt.Printf("%+v\n", config.Password)
	fmt.Printf("%+v\n", config.AuthType)
}

func exampleDeleteConfig() {
	err := keyconfig.DeleteConfig("example")
	if err != nil {
		panic(err)
	}
}

func main() {
	exampleSetConfig()
	exampleGetConfig()
	exampleDeleteConfig()
}
