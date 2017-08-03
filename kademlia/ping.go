package kademlia

import "log"

type PingRequest struct {
	RPCHeader
}

type PingResponse struct {
	RPCHeader
}

// Ping verifies a node is active
func (kc *KademliaCore) Ping(request *PingRequest, response *PingResponse) (err error) {
	if err = kc.kad.HandleRPC(&request.RPCHeader, &response.RPCHeader); err == nil {
		log.Printf("Ping from %s\n", request.RPCHeader)
	}
	return
}
