package server

import (
	"time"
)

type ServerConfig struct {
	// this cluster id is used for server transport to tell
	// different global cluster
	ClusterID            string   `json:"cluster_id"`
	EtcdClusterAddresses string   `json:"etcd_cluster_addresses"`
	BroadcastInterface   string   `json:"broadcast_interface"`
	BroadcastAddr        string   `json:"broadcast_addr"`
	RedisAPIPort         int      `json:"redis_api_port"`
	HttpAPIPort          int      `json:"http_api_port"`
	DataDir              string   `json:"data_dir"`
	DataRsyncModule      string   `json:"data_rsync_module"`
	LocalRaftAddr        string   `json:"local_raft_addr"`
	Tags                 []string `json:"tags"`

	ElectionTick int                   `json:"election_tick"`
	TickDuration time.Duration         `json:"tick_duration"`
	Namespaces   []NamespaceNodeConfig `json:"namespaces"`
}

type NamespaceNodeConfig struct {
	Name           string `json:"name"`
	LocalReplicaID uint64 `json:"local_replica_id"`
}

type ConfigFile struct {
	ServerConf ServerConfig `json:"server_conf"`
}
