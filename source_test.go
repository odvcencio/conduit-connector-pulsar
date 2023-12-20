package pulsar_test

import (
	"context"
	"testing"

	pulsar "github.com/odvcencio/conduit-connector-pulsar"
	"github.com/matryer/is"
)

func TestTeardownSource_NoOpen(t *testing.T) {
	is := is.New(t)
	con := pulsar.NewSource()
	err := con.Teardown(context.Background())
	is.NoErr(err)
}
