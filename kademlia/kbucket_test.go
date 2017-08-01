package kademlia

import (
	"fmt"
	"testing"
)

func TestFindContact(t *testing.T) {
	kb := NewKBucket()

	firstContact := randomContact()
	foundContact := kb.findContact(firstContact)
	if foundContact != nil {
		t.Error("KBucket is empty, findContact should return nil")
	}

	kb.PushBack(firstContact)

	fmt.Println("finding contact")

	foundContact = kb.findContact(firstContact)
	fmt.Println("found contact")
	if foundContact == nil || foundContact.Value.(Contact) != firstContact {
		t.Error("findContact should return non-nil Contact equal to firstContact")
	}

	kb.PushFront(randomContact())

	foundContact = kb.findContact(firstContact)
	if foundContact == nil || foundContact.Value.(Contact) != firstContact {
		t.Error("findContact should return non-nil Contact equal to firstContact")
	}
}

// Helpful functions
func randomContact() (contact Contact) {
	contactID := NewRandomNodeID()
	contact = NewContact(contactID, "")
	return
}
