package redis

import (
	"github.com/redis/go-redis/v9"
	"gitlab.com/golibs-starter/golib/actuator"
)

type Informer struct {
	client *redis.Client
	props  *Properties
}

func NewInformer(client *redis.Client, props *Properties) actuator.Informer {
	return &Informer{client: client, props: props}
}

func (d Informer) Key() string {
	return "redis"
}

func (d Informer) Value() interface{} {
	return map[string]interface{}{
		"pool": map[string]interface{}{
			"max_conns":   d.props.PoolSize,
			"total_conns": d.client.PoolStats().TotalConns,
			"idle_conns":  d.client.PoolStats().IdleConns,
			"stale_conns": d.client.PoolStats().StaleConns,

			// TODO use counter instead
			"hits":     d.client.PoolStats().Hits,
			"misses":   d.client.PoolStats().Misses,
			"timeouts": d.client.PoolStats().Timeouts,
		},
		"counter": map[string]interface{}{
			"hits":     d.client.PoolStats().Hits,
			"misses":   d.client.PoolStats().Misses,
			"timeouts": d.client.PoolStats().Timeouts,
		},
	}
}
