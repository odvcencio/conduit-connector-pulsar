package pulsar

// Config contains shared config parameters, common to the source and
// destination. If you don't need shared parameters you can entirely remove this
// file.
type Config struct {
	// Servers is a list of Pulsar broker servers
	Servers []string `json:"servers" validate:"required"`
	// Topic is the Pulsar topic
	Topic string `json:"topic" validate:"required"`

	// TLSTrustCertsFilePath is the path to the CA cert file
	TLSTrustCertsFilePath string `json:"trustCertsFilePath"`
	// TLSCertPath is the path to the client cert file
	TLSCertPath string `json:"certPath"`
	// TLSPrivateKeyPath is the path to the private key file
	TLSPrivateKeyPath string `json:"privateKeyPath"`
	// AllowInsecure decides if the client allows untrusted cert from broker or not
	AllowInsecure bool `json:"allowInsecure"`
}
