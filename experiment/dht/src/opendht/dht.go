package opendht

import (
	"errors"

	"berty.tech/experiment/dht"
	_ "berty.tech/experiment/dht/src/opendht/opendht"
)

var _ dht.DHT = (*DHT)(nil)

type DHT struct {
	node                   DhtRunner
	default_bootstrap_host string
	default_bootstrap_port string
}

func New() *DHT {
	return &DHT{
		node:                   NewDhtRunner(),
		default_bootstrap_host: "bootstrap.ring.cx",
		default_bootstrap_port: "4222",
	}
}

func (d *DHT) Run(port int) error {
	go d.node.Run(port)
	return nil
}

func (*DHT) RunSigned(port int, identity string) error {
	return errors.New("not implemented")
}

func (d *DHT) Bootstrap(host string, port string) error {
	d.node.Bootstrap(host, port)
	return nil
}

func (*DHT) Put(key string, value interface{}) error {
	return errors.New("not implemented")
}

func (*DHT) PutSigned(key string, value interface{}, identity string) error {
	return errors.New("not implemented")
}

func (*DHT) Get(key string) (interface{}, error) {
	return nil, errors.New("not implemented")
}

func (*DHT) GetSigned(key string, identity string) (interface{}, error) {
	return nil, errors.New("not implemented")
}

func (*DHT) Remove(key string) error {
	return errors.New("not implemented")
}

func (*DHT) RemoveSigned(key string, identity string) error {
	return errors.New("not implemented")
}
