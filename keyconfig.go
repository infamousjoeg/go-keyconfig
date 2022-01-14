// Package keyconfig implements a simple keychain wrapper for storing and
// retrieving configs.
//
package keyconfig

// SetConfig sets the config to keychain from an interface.
// Note: The configID is used as the key for the config.
func SetConfig(configID string, config interface{}) error {
	// Encode the config to JSON string
	encodedJSON, err := encodeJSON(config)
	if err != nil {
		return err
	}
	// Encode the config to base64 string
	encodedConfig := encodeBase64(encodedJSON)
	// Set the config to keychain
	err = set(configID, []byte(encodedConfig))
	if err != nil {
		return err
	}

	return nil
}

// GetConfig gets the config from keychain and decodes it into an interface.
// Note: The configID is used as the key for the config.
func GetConfig(configID string, config interface{}) error {
	// Get the config from keychain
	encodedConfig, err := get(configID)
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
	// Delete the config from keychain
	err := delete(configID)
	if err != nil {
		return err
	}

	return nil
}
