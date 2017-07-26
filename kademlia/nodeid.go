package kademlia

import (
	"encoding/hex"
	"fmt"
	"math/rand"
)

// IDLength is the length of the NodeID
const IDLength = 20

// NodeID is a byte slice of IDLength
type NodeID [IDLength]byte

func (node NodeID) String() string {
	fmt.Println(node[0:IDLength])
	return hex.EncodeToString(node[0:IDLength])
}

// NewNodeID uses named return arguments to avoid
// having to declare one of our own
func NewNodeID(s string) (ret NodeID) {
	decoded, _ := hex.DecodeString(s)
	for i := 0; i < IDLength; i++ {
		ret[i] = decoded[i]
	}
	return
}

// NewRandomNodeID creates a new NodeID
func NewRandomNodeID() (ret NodeID) {
	for i := 0; i < IDLength; i++ {
		ret[i] = uint8(rand.Intn(256))
	}
	return
}

// Equals method checks if two nodes are the same
func (node NodeID) Equals(other NodeID) bool {
	for i := 0; i < IDLength; i++ {
		if node[i] != other[i] {
			return false
		}
	}
	return true
}

// Less method sorts NodeIDs assuming most significant
// byte comes first (big-endian)
func (node NodeID) Less(other interface{}) bool {
	for i := 0; i < IDLength; i++ {
		fmt.Println(node[i])
		fmt.Println(other.(NodeID)[i])
		if node[i] != other.(NodeID)[i] {
			return node[i] < other.(NodeID)[i]
		}
	}
	return false
}

// Xor exclusive ors two nodes to find the distance metric
func (node NodeID) Xor(other interface{}) (ret NodeID) {
	for i := 0; i < IDLength; i++ {
		ret[i] = node[i] ^ other.(NodeID)[i]
	}
	return
}
