package jetstream

import (
	"log"

	"github.com/nats-io/nats.go"
)

const (
	streamName     = "ORDERS"
	streamSubjects = "ORDERS.*"
)

func JetStreamInit() (nats.JetStreamContext, error) {
	// Connect to NATS
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		return nil, err
	}

	// Create JetStream Context
	// The JetStreamContext allows JetStream messaging and stream management
	js, err := nc.JetStream()
	if err != nil {
		return nil, err
	}

	// Create a stream if it does not exist
	err = CreateStream(js)
	if err != nil {
		return nil, err
	}

	return js, nil
}

func CreateStream(jetStream nats.JetStreamContext) error {
	stream, err := jetStream.StreamInfo(streamName)
	if err != nil {
		log.Println(err)
	}

	// stream not found, create it
	if stream == nil {
		log.Printf("Creating stream: %s\n", streamName)

		_, err = jetStream.AddStream(&nats.StreamConfig{
			Name:     streamName,
			Subjects: []string{streamSubjects},
		})
		if err != nil {
			return err
		}
	}
	return nil
}
