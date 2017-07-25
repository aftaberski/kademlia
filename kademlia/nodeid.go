package dht

import (
	"encoding/hex"
	"math/rand"
)

const IdLength = 20

type NodeID [IdLength]byte

func (node NodeID) String() string {
	return hex.EncodeToString(node[0:IdLength])
}

// NewNodeID uses named return arguments to avoid
// having to declare one of our own
func NewNodeID(s string) (ret NodeID) {
	decoded, _ := hex.DecodeString(s)
	for i := 0; i < IdLength; i++ {
		ret[i] = decoded[i]
	}
	return
}

func NewRandomNodeID() (ret NodeID) {
	for i := 0; i < IdLength; i++ {
		ret[i] = uint8(rand.Intn(256))
	}
	return
}
