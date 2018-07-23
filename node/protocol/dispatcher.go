package protocol

type Dispatcher interface {
	//created unit received
	OnSendUnitArrived(sender string, m *SendUnit) error
	OnRecvUnitArrived(sender string, m *RecvUnit) error

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
