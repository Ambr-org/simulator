///Copyright (c) 2018 Ambr project
///Written by KimiKan
///Distributed under the MIT software license, see the accompanying
///file COPYING or http://www.opensource.org/licenses/mit-license.php.

package protocol

type Transporter interface {
	//publish a message to network
	Publish(m interface{}) error
	//created unit received
	OnUnitCreated(m *CreatedUnit) error
	//vote. while conflict
	OnVoteRequest(m *VoteRequest) error
	OnVoteResponse(m *VoteResponse) error
	//heartbeat to keep peer alived
	//if not provided it's okay
	//libary maintained
	OnHeartbeatRequest(m *HeartbeatRequest) error
	OnHeartbeatResponse(m *HeartbeatResponse) error
	//for replication used
	//request for lost nodes
	//it should be carefully designed
	OnReplicationRequest(m *ReplicationRequest) error
	OnReplicationResponse(m *ReplicationResponse) error
}
