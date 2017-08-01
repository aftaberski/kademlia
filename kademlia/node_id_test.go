package kademlia

import (
	"fmt"
	"testing"
)

func TestNewRandomNodeIDLength(t *testing.T) {
	id := NewRandomNodeID()

	if len(id) != IDLength {
		t.Error(fmt.Sprintf("Expected %d", IDLength))
	}
}

func TestEqual(t *testing.T) {
	id := NewRandomNodeID()

	if !id.Equals(id) {
		t.Error("IDs should be equal to themselves")
	}
}

func TestLess(t *testing.T) {
	id := NewRandomNodeID()
	lesserID := id

	// decrement lesserID by 1 in the least significant, non-zero byte
	for i := IDLength - 1; i >= 0; i-- {
		if lesserID[i] != 0 {
			lesserID[i]--
			break
		}
	}

	if !lesserID.Less(id) {
		t.Error("Smaller ID should be less than original ID")
	}
}

func TestXor(t *testing.T) {
	expectedXor := "15d9a75528691e5ebc0a415b99ed3b98f88110bd"
	id1String := "66472dba5cf4e1cbad155ad05beb14cb19d7c65a"
	id2String := "739e8aef749dff95111f1b8bc2062f53e156d6e7"

	id1 := NewNodeID(id1String)
	id2 := NewNodeID(id2String)

	if id1.Xor(id2).String() != expectedXor {
		t.Error(fmt.Sprintf("XOR of %s and %s should be %s", id1, id2, expectedXor))
	}
}
