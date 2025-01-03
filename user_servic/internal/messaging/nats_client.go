package messaging

import (
	"log"

	"github.com/nats-io/nats.go"
)

type NatsClient struct {
	conn *nats.Conn
}

func NewNatsClient(url string) (*NatsClient, error) {
	conn, err := nats.Connect(url)
	if err != nil {
		return nil, err
	}

	return &NatsClient{conn: conn}, nil
}

func (n *NatsClient) Publisher(subject string, message []byte) error {
	if err := n.conn.Publish(subject, message); err != nil {
		return err
	}

	log.Printf("Published message to subject: %s", subject)

	return nil
}

func (n *NatsClient) Close() {
	n.conn.Close()
}
