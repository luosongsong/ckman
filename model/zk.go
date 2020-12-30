package model

type ZkStatusRsp struct {
	Host                string  `json:"host"`
	Version             string  `json:"version"`
	ServerState         string  `json:"server_state"`
	PeerState           string  `json:"peer_state"`
	AvgLatency          float64 `json:"avg_latency"`
	ApproximateDataSize float64 `json:"approximate_data_size"`
	ZnodeCount          float64 `json:"znode_count"`
}
