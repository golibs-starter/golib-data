package redis

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	"time"
)

func NewClient(props *Properties) (*redis.Client, error) {
	if props == nil {
		return nil, errors.New("redis config is required")
	}
	if len(props.Host) == 0 {
		return nil, errors.New("redis host is required")
	}
	if props.Port <= 0 {
		return nil, errors.New("redis port is required")
	}
	config := &redis.Options{
		Addr:            fmt.Sprintf("%s:%d", props.Host, props.Port),
		Username:        props.Username,
		Password:        props.Password,
		DB:              props.Database,
		PoolSize:        props.PoolSize,
		MinIdleConns:    props.MinIdleConns,
		ConnMaxIdleTime: props.ConnMaxIdleTime,
		ConnMaxLifetime: props.ConnMaxLifetime,
	}
	if props.EnableTLS {
		config.TLSConfig = &tls.Config{
			MinVersion: tls.VersionTLS12,
		}
	}
	client := redis.NewClient(config)
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()
	_, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, errors.WithMessage(err, "error when open redis connection")
	}
	return client, nil
}
