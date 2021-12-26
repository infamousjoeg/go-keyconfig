package keyconfig

import (
	"fmt"
	"log"

	"github.com/keybase/go-keychain"
)

// listConfigs is a string array function that returns all secrets in keychain
// with the label `keyconfig`.
func listConfigs() []string {
	// Note: OSX use the term "account" to refer to the config id.
	configIDs, err := keychain.GetGenericPasswordAccounts("keyconfig")
	if err != nil {
		log.Fatalln(err)
	}

	return configIDs
}

// configExists is a boolean function to verify a config is present in keychain
func configExists(configID string) bool {
	allconfigIDs := listConfigs()

	// Search all the available configIDs for this one
	for _, id := range allconfigIDs {
		if id == configID {
			return true
		}
	}
	return false
}

// set is a non-return function that adds the secret and secret value to
// keychain.
func set(configID string, config []byte) error {
	// Add new generic password query to keychain
	query := keychain.NewGenericPassword(
		"keyconfig", configID, "keyconfig", config, "",
	)
	query.SetSynchronizable(keychain.SynchronizableNo)
	query.SetAccessible(keychain.AccessibleAfterFirstUnlock)

	err := keychain.AddItem(query)
	if err != nil {
		return err
	}
	// Duplicate item error (TODO: Update existing item)
	if err == keychain.ErrorDuplicateItem {
		return keychain.ErrorDuplicateItem
	}
	// Verify the config was set in keychain successfully
	if !configExists(configID) {
		err := fmt.Errorf("config not found in keychain")
		return err
	}

	return nil
}

// get is a non-return function that retrieves a config and returns it
// as a struct.
func get(configID string) ([]byte, error) {
	// Build query for config retrieval from Keychain
	query := keychain.NewItem()
	query.SetSecClass(keychain.SecClassGenericPassword)
	query.SetService("keyconfig")
	query.SetAccount(configID)
	query.SetMatchLimit(keychain.MatchLimitOne)
	query.SetReturnData(true)
	results, err := keychain.QueryItem(query)
	if err != nil {
		return nil, err
	}
	if len(results) != 1 {
		return nil, keychain.ErrorItemNotFound
	}
	encodedConfig := results[0].Data

	return encodedConfig, nil
}

// delete is a non-return function that deletes a config from keychain.
func delete(configID string) error {
	// Build query for config deletion from Keychain
	query := keychain.NewItem()
	query.SetSecClass(keychain.SecClassGenericPassword)
	query.SetService("keyconfig")
	query.SetAccount(configID)
	query.SetMatchLimit(keychain.MatchLimitOne)
	query.SetReturnData(true)
	results, err := keychain.QueryItem(query)
	if err != nil {
		return err
	}
	if len(results) != 1 {
		return keychain.ErrorItemNotFound
	}

	err = keychain.DeleteItem(query)
	if err != nil {
		return err
	}
	// Verify the config was deleted from keychain successfully
	if configExists(configID) {
		err := fmt.Errorf("config not deleted from keychain")
		return err
	}

	return nil
}
