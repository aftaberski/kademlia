package main

import (
	"fmt"

	"./kademlia"
)

func main() {
	node1 := kademlia.NewRandomNodeID()
	node2 := kademlia.NewRandomNodeID()
	// fmt.Println("Node 1: ", node1)
	// fmt.Println("Node 2: ", node2)
	// fmt.Println("Equal? ", node1.Equals(node2))
	// fmt.Println("Less? ", node1.Less(node2))
	fmt.Println("Xor", node1.Xor(node2))
	fmt.Println("PrefixLen Xor:", node1.Xor(node2).PrefixLen(node2))
	// fmt.Println("PrefixLen Node 1:", node1.PrefixLen())
	// fmt.Println("PrefixLen Node 2:", node2.PrefixLen())
}
