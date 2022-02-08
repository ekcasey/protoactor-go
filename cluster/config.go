package cluster

import (
	"time"

	"github.com/AsynkronIT/protoactor-go/actor"

	"github.com/AsynkronIT/protoactor-go/remote"
)

type Config struct {
	Name                  string
	Address               string
	ClusterProvider       ClusterProvider
	RemoteConfig          remote.Config
	TimeoutTime           time.Duration
	GossipInterval        time.Duration
	GossipRequestTimeout  time.Duration
	GossipFanOut          int
	GossipMaxSend         int
	MemberStrategyBuilder func(kind string) MemberStrategy
	Kinds                 map[string]*actor.Props
}

func Configure(clusterName string, clusterProvider ClusterProvider, remoteConfig remote.Config, kinds ...*Kind) *Config {
	config := &Config{
		Name:                  clusterName,
		ClusterProvider:       clusterProvider,
		TimeoutTime:           time.Second * 5,
		GossipInterval:        time.Millisecond * 300,
		GossipRequestTimeout:  time.Millisecond * 500,
		GossipFanOut:          3,
		GossipMaxSend:         50,
		MemberStrategyBuilder: newDefaultMemberStrategy,
		RemoteConfig:          remoteConfig,
		Kinds:                 make(map[string]*actor.Props),
	}

	for _, kind := range kinds {
		config.Kinds[kind.Kind] = kind.Props
	}

	return config
}

func (c *Config) WithTimeout(t time.Duration) *Config {
	c.TimeoutTime = t
	return c
}

type Kind struct {
	Kind  string
	Props *actor.Props
}

func NewKind(kind string, props *actor.Props) *Kind {
	return &Kind{
		Kind:  kind,
		Props: props,
	}
}
