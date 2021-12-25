package keyconfig

func SetConfig(configID string, config interface{}) {
	// Encode the config as a JSON string
	eJSON, _ := encodeJSON(config)
	// Encode the JSON string as a base64 string
	encodedConfig := encodeBase64(eJSON)
	// Add the config to keychain
	set(configID, []byte(encodedConfig))
}

func GetConfig(configID string, config interface{}) (interface{}, error) {
	// Get the config from keychain
	encodedConfig, err := get(configID)
	if err != nil {
		return nil, err
	}
	// Decode the config from base64 string
	dJSON, _ := decodeBase64(string(encodedConfig))
	// Decode the config from JSON string
	err = decodeJSON(dJSON, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
