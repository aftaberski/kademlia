package kademlia

import "container/list"

const BucketSize = 20

type Contact struct {
	id NodeID
}

// RoutingTable uses container/list for creating
// a doubly-linked list
type RoutingTable struct {
	node    NodeID
	buckets [IDLength * 8]*list.List
}

func NewRoutingTable(node NodeID) (ret RoutingTable) {
	for i := 0; i < IDLength*8; i++ {
		// Create linked list for each bit in the ID
		ret.buckets[i] = list.New()
	}
	ret.node = node
	return
}

func inBucket(bucket *list.List, nodeID NodeID) (ret *Contact) {
	curr := bucket.Front()
	for curr != nil {
		if interface{}(curr).(*Contact).id == nodeID {
			return interface{}(curr).(*Contact)
		}
		curr = curr.Next()
	}
	return nil
}

func (table *RoutingTable) Update(contact *Contact) {
	prefixLen := contact.id.Xor(table.node).PrefixLen()
	bucket := table.buckets[prefixLen]
	// Locate the contact inside of the bucket if already exists
	element := inBucket(bucket, table.node)
	if element == nil {
		if bucket.Len() <= BucketSize {
			bucket.PushFront(contact)
		}
		// TODO: Handle insertion when the list is full by evicting old elements if
		// they don't respond to a ping.
	} else {
		bucket.MoveToFront(interface{}(element).(*list.Element))
	}
}
