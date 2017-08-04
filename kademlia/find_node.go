package kademlia

type FindNodeRequest struct {
	RPCHeader
	NodeID
}

func (k *Kademlia) NewFindNodeRequest(target NodeID) FindNodeRequest {
	header := RPCHeader{
		&k.routes.contact,
		k.NetworkID,
	}
	return FindNodeRequest{
		header,
		target,
	}
}

type FindNodeResponse struct {
	RPCHeader
	Contacts
}

// sendQuery sends FindNode RPC to the specified contact
// Then sends the result (list of Contacts) to the done channel
// so we can call sendQuery in a goroutine
func (k *Kademlia) sendQuery(contact Contact, node NodeID, done chan Contacts) {
	client, err := RPCClient(contact)
	if err != nil {
		done <- nil
		return
	}

	req := k.NewFindNodeRequest(node)
	res := FindNodeResponse{}

	err = client.Call("KademliaCore.FindNode", &req, &res)
	if err != nil {
		done <- nil
		return
	}

	done <- res.Contacts
}

func (kc *KademliaCore) FindNode(req FindNodeRequest, res *FindNodeResponse) error {
	err := kc.kad.HandleRPC(&req.RPCHeader, &res.RPCHeader)
	if err != nil {
		return err
	}
	res.Contacts = kc.kad.routes.FindClosest(req.NodeID, BucketSize)

	return nil
}

func (k *Kademlia) IterativeFindNode(node NodeID, delta int, contacts chan Contacts) {
	done := make(chan Contacts)
	ret := make(Contacts, BucketSize)
	frontier := make(Contacts, BucketSize)
	seen := make(map[string]bool)

	// Finish this :)

}
