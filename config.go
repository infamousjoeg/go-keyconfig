package keyconfig

import (
	"fmt"

	"github.com/99designs/keyring"
)

// configExists is a boolean function to verify a config is present in keychain
func configExists(kr keyring.Keyring, configID string) (bool, error) {
	// Check if config exists in keyring
	_, err := kr.Get(configID)
	if err != nil {
		return false, err
	}

	return true, nil
}

// set is a non-return function that adds the secret and secret value to
// keychain.
func set(kr keyring.Keyring, configID string, config keyring.Item) error {
	// Check if config exists in keyring
	exists, _ := configExists(kr, configID)
	// If config exists, delete it
	if exists {
		err := kr.Remove(configID)
		if err != nil {
			return err
		}
	}

	// Add config to keyring
	err := kr.Set(config)
	if err != nil {
		return err
	}

	// Verify the config was set in keychain successfully
	_, err = configExists(kr, configID)
	if err != nil {
		err := fmt.Errorf("config not found in keychain")
		return err
	}

	return nil
}

// get is a non-return function that retrieves a config and returns it
// as a struct.
func get(kr keyring.Keyring, configID string) ([]byte, error) {
	// Get config from keyring
	results, err := kr.Get(configID)
	if err != nil {
		return nil, err
	}

	return results.Data, nil
}

// delete is a non-return function that deletes a config from keychain.
func delete(kr keyring.Keyring, configID string) error {
	// Delete config from keyring
	err := kr.Remove(configID)
	if err != nil {
		return err
	}

	// Verify the config was deleted from keychain successfully
	_, err = configExists(kr, configID)
	if err == nil {
		err := fmt.Errorf("config not deleted from keychain")
		return err
	}

	return nil
}
