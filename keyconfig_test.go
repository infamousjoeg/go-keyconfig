package keyconfig_test

import (
	"testing"

	"github.com/infamousjoeg/go-keyconfig"
)

type TestConfig struct {
	BaseURL  string `json:"BaseURL"`
	Username string `json:"Username"`
	Password string `json:"Password"`
	AuthType string `json:"AuthType"`
}

func TestSetConfig(t *testing.T) {
	config := TestConfig{
		BaseURL:  "https://example.com",
		Username: "username",
		Password: "password",
		AuthType: "basic",
	}

	err := keyconfig.SetConfig("test", config)
	if err != nil {
		t.Errorf("Error setting config: %s", err)
	}
}

func TestGetConfig(t *testing.T) {
	var config TestConfig

	err := keyconfig.GetConfig("test", &config)
	if err != nil {
		t.Errorf("Error getting config: %s", err)
	}

	if config.BaseURL != "https://example.com" {
		t.Errorf("BaseURL is incorrect: %s", config.BaseURL)
	}

	if config.Username != "username" {
		t.Errorf("Username is incorrect: %s", config.Username)
	}

	if config.Password != "password" {
		t.Errorf("Password is incorrect: %s", config.Password)
	}

	if config.AuthType != "basic" {
		t.Errorf("AuthType is incorrect: %s", config.AuthType)
	}
}

func TestDeleteConfig(t *testing.T) {
	err := keyconfig.DeleteConfig("test")
	if err != nil {
		t.Errorf("Error deleting config: %s", err)
	}
}
