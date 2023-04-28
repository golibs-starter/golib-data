package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"gitlab.com/golibs-starter/golib/actuator"
	"gitlab.com/golibs-starter/golib/web/log"
)

type HealthChecker struct {
	client *redis.Client
}

func NewHealthChecker(client *redis.Client) actuator.HealthChecker {
	return &HealthChecker{client: client}
}

func (h HealthChecker) Component() string {
	return "redis"
}

func (h HealthChecker) Check(ctx context.Context) actuator.StatusDetails {
	statusDetails := actuator.StatusDetails{
		Status: actuator.StatusUp,
	}
	_, err := h.client.Ping(ctx).Result()
	if err != nil {
		log.Error(ctx, "Redis health check failed, err [%s]", err.Error())
		statusDetails.Status = actuator.StatusDown
		statusDetails.Reason = "Redis health check failed"
	}
	return statusDetails
}
