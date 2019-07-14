package model

type PublicKeys struct {
	Keys []PublicKey `json:"keys"`
}

type PublicKey struct {
	Kid string `json:"kid"`
	E   string `json:"e"`
	Kty string `json:"kty"`
	Alg string `json:"alg"`
	N   string `json:"n"`
	Use string `json:"use"`
}
