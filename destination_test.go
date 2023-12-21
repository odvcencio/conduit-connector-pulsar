package pulsar

import (
	"context"
	"fmt"
	"testing"

	"github.com/apache/pulsar-client-go/pulsar"
	sdk "github.com/conduitio/conduit-connector-sdk"
	"github.com/google/uuid"
	"github.com/matryer/is"
)

var cfg = map[string]string{"servers": "localhost:6650",
	"topic": "persistent://public/default/test-topic"}

func TestTeardown_NoOpen(t *testing.T) {
	is := is.New(t)
	con := NewDestination()
	err := con.Teardown(context.Background())
	is.NoErr(err)
}

func TestTeardown_Open(t *testing.T) {
	is := is.New(t)
	ctx := context.Background()
	con := NewDestination()

	err := con.Configure(ctx, cfg)
	is.NoErr(err)
	err = con.Open(ctx)
	is.NoErr(err)
	err = con.Teardown(ctx)
	is.NoErr(err)
}

func TestDestinationIntegrationWrite(t *testing.T) {
	is := is.New(t)
	ctx := context.Background()

	amountOfRecords := 6
	pulsarMessages := make([]*pulsar.ProducerMessage, amountOfRecords)
	sdkRecords := make([]sdk.Record, amountOfRecords)

	// generate a few records here
	for i := 0; i < amountOfRecords; i++ {
		// make the pulsar messages
		pulsarMessages[i] = &pulsar.ProducerMessage{
			Payload: []byte(fmt.Sprintf("test-payload-%d", i)),
		}
		// make the sdk records
		sdkRecords[i] = sdk.Util.Source.NewRecordCreate(
			[]byte(uuid.NewString()),
			nil,
			nil,
			sdk.RawData(pulsarMessages[i].Payload),
		)
	}

	// create a consumer to verify messages were produced to the destination
	cClient, err := pulsar.NewClient(pulsar.ClientOptions{
		URL: "pulsar://" + cfg["servers"],
	})
	is.NoErr(err)

	defer cClient.Close()

	consumer, err := cClient.Subscribe(pulsar.ConsumerOptions{
		Topic:            cfg["topic"],
		SubscriptionName: "test-sub",
		Type:             pulsar.Shared,
	})
	is.NoErr(err)

	defer consumer.Close()

	testDest := NewDestination()
	defer func() {
		err := testDest.Teardown(ctx)
		is.NoErr(err)
	}()

	err = testDest.Configure(ctx, cfg)
	is.NoErr(err)

	err = testDest.Open(ctx)
	is.NoErr(err)

	count, err := testDest.Write(ctx, sdkRecords)
	is.NoErr(err)
	is.Equal(count, len(sdkRecords))
}
