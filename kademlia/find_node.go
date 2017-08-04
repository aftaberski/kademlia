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

func (k *Kademlia) sendQuery(contact Contact, node NodeID) {
	client, err := RPCClient(contact)
	if err != nil {
		return
	}

	req := k.NewFindNodeRequest(node)
	res := FindNodeResponse{}

	err = client.Call("KademliaCore.FindNode", &req, &res)
	// TODO: finish this
}

func (kc *KademliaCore) FindNode(req FindNodeRequest, res *FindNodeResponse) error {
	err := kc.kad.HandleRPC(&req.RPCHeader, &res.RPCHeader)
	if err != nil {
		return err
	}
	res.Contacts = kc.kad.routes.FindClosest(req.NodeID, BucketSize)

	return nil
}
