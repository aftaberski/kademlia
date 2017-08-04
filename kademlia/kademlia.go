package kademlia

import (
	"fmt"
	"net"
	"net/rpc"
	"time"
)

type Kademlia struct {
	routes *RoutingTable
	// NetworkID is an arbitrary string unique to each Kademlia instance
	// to prevent different instances from merging together
	NetworkID string
}

// NewKademlia creates a new *Kademlia instance
func NewKademlia(contact *Contact, networkID string) (ret *Kademlia) {
	ret = new(Kademlia)
	ret.routes = NewRoutingTable(*contact)
	ret.NetworkID = networkID
	return
}

type RPCHeader struct {
	Sender    *Contact
	NetworkID string
}

// HandleRPC is called for every incoming RPC
func (k *Kademlia) HandleRPC(request, response *RPCHeader) error {
	if request.NetworkID != k.NetworkID {
		return fmt.Errorf("Expected NetworkID %s, but instead got %s", k.NetworkID, request.NetworkID)
	}
	if request.Sender != nil {
		k.routes.Update(*request.Sender)
	}
	response.Sender = &k.routes.contact
	return nil
}

type KademliaCore struct {
	kad *Kademlia
}

func RPCClient(contact Contact) (*rpc.Client, error) {
	connection, err := net.DialTimeout("tcp", contact.Address, 5*time.Second)
	if err != nil {
		return nil, err
	}

	return rpc.NewClient(connection), nil
}

func (k *Kademlia) Serve() (err error) {
	rpc.Register(&KademliaCore{k})

	l, err := net.Listen("tcp", k.routes.contact.Address)
	if err == nil {
		go rpc.Accept(l)
	}
	return
}
