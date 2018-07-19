///Copyright (c) 2018 Ambr project
///Written by KimiKan
///Distributed under the MIT software license, see the accompanying
///file COPYING or http://www.opensource.org/licenses/mit-license.php.

package protocol

import (
	"github.com/golang/protobuf/proto"
)

type Transporter interface {
	//publish a message to network
	Publish(m proto.Message) error
	//created unit received
	OnUnitCreated(sender string, m *CreatedUnit) error
	//vote. while conflict
	OnVoteRequest(sender string, m *VoteRequest) error
	OnVoteResponse(sender string, m *VoteResponse) error
	//heartbeat to keep peer alived
	//if not provided it's okay
	//libary maintained
	OnHeartbeatRequest(sender string, m *HeartbeatRequest) error
	OnHeartbeatResponse(sender string, m *HeartbeatResponse) error
	//for replication used
	//request for lost nodes
	//it should be carefully designed
	OnReplicationRequest(sender string, m *ReplicationRequest) error
	OnReplicationResponse(sender string, m *ReplicationResponse) error
}
