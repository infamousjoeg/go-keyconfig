// Package keyconfig implements a simple keychain wrapper for storing and
// retrieving configs.
//
package keyconfig

import "github.com/99designs/keyring"

// open is a function that returns a keyring.Keyring.
func open() (keyring.Keyring, error) {
	// Open keyring
	kr, err := keyring.Open(keyring.Config{
		AllowedBackends: []keyring.BackendType{
			"secret-service",
			"keychain",
			"kwallet",
			"wincred",
		},
		ServiceName:              "keyconfig",
		KeychainName:             "keyconfig",
		KeychainTrustApplication: true,
		WinCredPrefix:            "keyconfig",
		KWalletAppID:             "keyconfig",
		KWalletFolder:            "keyconfig",
		LibSecretCollectionName:  "keyconfig",
	})
	if err != nil {
		return nil, err
	}

	return kr, nil
}

// SetConfig sets the config to keychain from an interface.
// Note: The configID is used as the key for the config.
func SetConfig(configID string, config interface{}) error {
	kr, err := open()
	if err != nil {
		return err
	}

	// Encode the config to JSON string
	encodedJSON, err := encodeJSON(config)
	if err != nil {
		return err
	}
	// Encode the config to base64 string
	encodedConfig := encodeBase64(encodedJSON)
	// Populate the keyring item
	item := keyring.Item{
		Key:  configID,
		Data: []byte(encodedConfig),
	}
	// Set the config to keychain
	err = set(kr, configID, item)
	if err != nil {
		return err
	}

	return nil
}

// GetConfig gets the config from keychain and decodes it into an interface.
// Note: The configID is used as the key for the config.
func GetConfig(configID string, config interface{}) error {
	kr, err := open()
	if err != nil {
		return err
	}

	// Get the config from keychain
	encodedConfig, err := get(kr, configID)
	if err != nil {
		return err
	}
	// Decode the config from base64 string
	dJSON, _ := decodeBase64(string(encodedConfig))
	// Decode the config from JSON string
	err = decodeJSON(dJSON, config)
	if err != nil {
		return err
	}

	return nil
}

// DeleteConfig deletes the config from keychain.
// Note: The configID is used as the key for the config.
func DeleteConfig(configID string) error {
	kr, err := open()
	if err != nil {
		return err
	}

	// Delete the config from keychain
	err = delete(kr, configID)
	if err != nil {
		return err
	}

	return nil
}
