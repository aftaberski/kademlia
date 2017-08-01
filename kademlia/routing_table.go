package kademlia

type RoutingTable struct {
	contact  Contact
	kbuckets [IDBytesLength]*KBucket
}

func NewRoutingTable(contact Contact) *RoutingTable {
	rt := &RoutingTable{
		contact:  contact,
		kbuckets: [IDBytesLength]*KBucket{},
	}

	for i := 0; i < IDBytesLength; i++ {
		rt.kbuckets[i] = NewKBucket()
	}

	return rt
}

// Update adds contact to the appropriate kbucket
func (rt *RoutingTable) Update(contact Contact) {
	prefixLength := contact.ID.PrefixLen(rt.contact.ID)
	if prefixLength == -1 {
		return
	}

	bucket := rt.kbuckets[prefixLength]
	bucket.Update(contact)
}

// FindClosest returns slice of contacts that are closest to target NodeID
func (rt *RoutingTable) FindClosest(node NodeID, delta int) Contacts {
	prefixLength := node.PrefixLen(rt.contact.ID)

	contacts := Contacts{}
	if prefixLength == -1 {
		contacts = append(contacts, rt.contact)
		prefixLength = IDBytesLength - 1 // 159
	}

	// Add all contacts from buckets < prefixLength
	// Return contacts if # of contacts > BucketSize
	for i := prefixLength; i >= 0; i-- {
		// Get element from front of bucket
		element := rt.kbuckets[i].Front()
		// Iterate through linked list
		for element != nil {
			if len(contacts) < BucketSize {
				contacts = append(contacts, element.Value.(Contact))
			} else {
				return contacts
			}
			element = element.Next()
		}
	}

	// Assuming there is still more room for contacts
	// So start adding contacts from buckets of index
	// greater than prefixLength
	for i := prefixLength + 1; i < IDBytesLength; i++ {
		element := rt.kbuckets[i].Front()
		for element != nil {
			if len(contacts) < BucketSize {
				contacts = append(contacts, element.Value.(Contact))
				element = element.Next()
			} else {
				return contacts
			}
		}
	}

	if len(contacts) > delta {
		return contacts[:delta]
	}

	return contacts
}
