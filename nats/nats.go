package nats

import "github.com/nats-io/stan.go"

func NatsConnect(clusterID, clientID string) (stan.Conn, error) {
	return stan.Connect(clusterID, clientID, stan.NatsURL("0.0.0.0:4222"))
}
