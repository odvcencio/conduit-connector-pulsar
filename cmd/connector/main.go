package main

import (
	sdk "github.com/conduitio/conduit-connector-sdk"

	pulsar "github.com/odvcencio/conduit-connector-pulsar"
)

func main() {
	sdk.Serve(pulsar.Connector)
}
