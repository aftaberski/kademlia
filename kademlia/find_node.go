package kademlia

import (
	"container/heap"
	"sort"
)

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

func (k *Kademlia) IterativeFindNode(node NodeID, delta int, contacts chan Contacts) (ret Contacts) {
	done := make(chan Contacts)
	ret = make(Contacts, BucketSize)
	frontier := make(Contacts, BucketSize)

	// So we don't accidentally add the same contact twice
	seen := make(map[string]bool)

	// Initialize ret, frontier, and seen lists with local nodes
	for _, node := range k.routes.FindClosest(node, delta) {
		ret = append(ret, node)
		heap.Push(&frontier, node)
		seen[node.ID.String()] = true
	}

	// Start querying which puts responses on the done channel
	// NOTE: Not accounting for unresponsive nodes! Right now, if a node never
	// responds we'll wait here indefinitely
	// TODO: Add a timeout for each RPC
	pending := 0
	for i := 0; pending < delta && frontier.Len() > 0; i++ {
		pending++
		contact := heap.Pop(&frontier).(Contact)
		go k.sendQuery(contact, node, done)
	}

	// Iteratively look for closer nodes
	for pending > 0 {
		nodes := <-done
		pending--
		for _, n := range nodes {
			// If we haven't seen this node before then add it
			if _, ok := seen[n.ID.String()]; !ok {
				ret = append(ret, n)
				heap.Push(&frontier, n)
				seen[n.ID.String()] = true
			}
		}

		// And send the query
		for i := 0; pending < delta && frontier.Len() > 0; i++ {
			pending++
			contact := heap.Pop(&frontier).(Contact)
			go k.sendQuery(contact, node, done)
		}

		sort.Sort(ret)

		if ret.Len() > BucketSize {
			ret = ret[:BucketSize]
		}
	}
	return
}
