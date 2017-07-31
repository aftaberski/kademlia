package kademlia

import "container/list"

const BucketSize = 20

// RoutingTable uses container/list for creating
// a doubly-linked list
type RoutingTable struct {
	node    NodeID
	buckets [IDLength * 8]*list.List
}

func NewRoutingTable(node NodeID) (ret RoutingTable) {
	for i := 0; i < IDLength*8; i++ {
		ret.buckets[i] = list.New()
	}
	ret.node = node
	return
}
