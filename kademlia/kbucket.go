package kademlia

import "container/list"

// KBucket is a doubly linked list
type KBucket struct {
	*list.List
}

// NewKBucket creates a new KBucket
func NewKBucket() *KBucket {
	return &KBucket{
		list.New(),
	}
}

// Update KBucket with new contact
func (kb *KBucket) Update(contact Contact) {
	foundContact := kb.findContact(contact)
	if foundContact != nil {
		// If contact is already in bucket, move it to back of list
		kb.MoveToBack(foundContact)
	} else if kb.isFull() {
		// TODO: PING node at front of list
		// Remove it if it is unresponsive
		// Otherwise ignore contact
	} else {
		// Otherwise just add contact to the list
		kb.PushBack(contact)
	}
}

func (kb KBucket) findContact(contact Contact) *list.Element {
	return kb.findByID(contact.ID)
}

// findById finds element with nodeID in list
func (kb KBucket) findByID(nodeID NodeID) *list.Element {
	// Iterate through linked list
	for element := kb.Front(); element != nil; element = element.Next() {
		if nodeID == element.Value.(Contact).ID {
			return element
		}
	}

	return nil
}

func (kb KBucket) isFull() bool {
	return kb.Len() >= BucketSize
}
