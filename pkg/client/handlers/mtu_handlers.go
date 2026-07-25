// ==============================================================================
// MasterDnsVPN
// Author: MasterkinG32
// Github: https://github.com/masterking32
// Year: 2026
// ==============================================================================
package handlers

import (
	"net"

	Enums "github.com/hiddifydeveloper/MasterDnsVPN/pkg/enums"
	VpnProto "github.com/hiddifydeveloper/MasterDnsVPN/pkg/vpnproto"
)

func init() {
	RegisterHandler(Enums.PACKET_MTU_UP_RES, handleMTUResponse)
	RegisterHandler(Enums.PACKET_MTU_DOWN_RES, handleMTUResponse)
}

func handleMTUResponse(c ClientContext, packet VpnProto.Packet, addr *net.UDPAddr) error {
	return c.HandleMTUResponse(packet)
}
