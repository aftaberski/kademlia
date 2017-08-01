package kademlia

import (
	"encoding/hex"
	"fmt"
	"log"
	"math/rand"
)

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
	buffer := make([]byte, IDLength)
	_, err := rand.Read(buffer)
	if err != nil {
		log.Print(err)
	}

	for i, b := range buffer {
		ret[i] = b
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

func (node NodeID) PrefixLen(other NodeID) int {
	distance := node.Xor(other)
	for i := 0; i < IDLength; i++ {
		for j := 0; j < 8; j++ {
			if (distance[i]>>uint8(7-j))&0x1 != 0 {
				return i*8 + j
			}
		}
	}
	return -1
}
