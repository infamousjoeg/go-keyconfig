package main

import (
	"fmt"

	"github.com/infamousjoeg/go-keyconfig"
)

type ExampleConfig struct {
	BaseURL  string `json:"BaseURL"`
	Username string `json:"Username"`
	Password string `json:"Password"`
	AuthType string `json:"AuthType"`
}

func ExampleSetConfig() {
	config := ExampleConfig{
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

func ExampleGetConfig() {
	var config ExampleConfig

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

func ExampleDeleteConfig() {
	err := keyconfig.DeleteConfig("example")
	if err != nil {
		panic(err)
	}
}

func main() {
	ExampleSetConfig()
	ExampleGetConfig()
	ExampleDeleteConfig()
}
