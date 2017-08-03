package kademlia

type Contact struct {
	ID      NodeID
	Address string
}

func NewContact(node NodeID, address string) Contact {
	return Contact{
		ID:      node,
		Address: address,
	}
}

type Contacts []Contact

// Implement the sort.Interface which requires
// Len, Less, and Swap methods
func (h Contacts) Len() int {
	return len(h)
}

func (h Contacts) Less(i, j int) bool {
	return h[i].ID.Less(h[j].ID)
}

func (h Contacts) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *Contacts) Push(x interface{}) {
	*h = append(*h, x.(Contact))
}

func (contacts *Contacts) Pop() interface{} {
	oldContacts := *contacts
	oldLength := len(oldContacts)
	element := oldContacts[oldLength-1]
	*contacts = oldContacts[0 : oldLength-1]
	return element
}
