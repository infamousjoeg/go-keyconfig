package keyconfig

import "encoding/json"

// encodeJSON encodes the given struct as a JSON string.
func encodeJSON(data interface{}) ([]byte, error) {
	return json.Marshal(data)
}

// decodeJSON decodes the given JSON string into a struct.
func decodeJSON(data []byte, config interface{}) error {
	return json.Unmarshal(data, &config)
}
