package pulsar

// Config contains shared config parameters, common to the source and
// destination. If you don't need shared parameters you can entirely remove this
// file.
type Config struct {
	// Servers is a list of Pulsar broker servers
	Servers []string `json:"servers" validate:"required"`
	// Topic is the Pulsar topic
	Topic string `json:"topic" validate:"required"`
}
