package keyconfig

import "encoding/base64"

// encodeBase64 encodes the given byte as a base64 string.
func encodeBase64(data []byte) string {
	sEnc := base64.StdEncoding.EncodeToString([]byte(data))
	return sEnc
}

// decodeBase64 decodes the given base64 string into a byte array.
func decodeBase64(data string) ([]byte, error) {
	sDec, err := base64.StdEncoding.DecodeString(data)
	return sDec, err
}
