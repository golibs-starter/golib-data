package datasource

import (
	"context"
	"database/sql"
	"gitlab.id.vin/vincart/golib/actuator"
	"gitlab.id.vin/vincart/golib/web/log"
)

type HealthChecker struct {
	connection *sql.DB
}

func NewHealthChecker(connection *sql.DB) actuator.HealthChecker {
	return &HealthChecker{connection: connection}
}

func (h HealthChecker) Component() string {
	return "datasource"
}

func (h HealthChecker) Check(ctx context.Context) actuator.StatusDetails {
	statusDetails := actuator.StatusDetails{
		Status: actuator.StatusUp,
	}
	if err := h.connection.Ping(); err != nil {
		log.Error(ctx, "Datasource health check failed, err [%s]", err.Error())
		statusDetails.Status = actuator.StatusDown
		statusDetails.Reason = "Datasource health check failed"
	}
	return statusDetails
}
