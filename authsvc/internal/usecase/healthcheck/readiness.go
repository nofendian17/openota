package healthcheck

import (
	"context"
	"github.com/nofendian17/openota/authsvc/internal/entity"
)

const (
	statusUP   = "UP"   // statusUP represents the status of a component as "UP"
	statusDown = "DOWN" // statusDown represents the status of a component as "DOWN"
)

// Readiness checks the readiness of the application by verifying different components like database and Redis connections.
func (u *useCase) Readiness(ctx context.Context) (*entity.ReadinessResponse, error) {
	ready := entity.ReadinessResponse{
		Status: statusUP,
		Checks: []entity.Check{},
	}

	var checks []entity.Check

	// Check database connection
	checkDB := entity.Check{
		Name:   u.db.GormDB.Name(),
		Status: statusUP,
	}
	err := u.db.SqlDB.PingContext(ctx)
	if err != nil {
		checkDB.Status = statusDown
		checkDB.Error = err.Error()
		ready.Status = statusDown
	}
	checks = append(checks, checkDB)

	// Check Redis connection
	checkRedis := entity.Check{
		Name:   "Redis",
		Status: statusUP,
	}
	err = u.cacheClient.Ping(ctx)
	if err != nil {
		checkRedis.Status = statusDown
		checkRedis.Error = err.Error()
		ready.Status = statusDown
	}
	checks = append(checks, checkRedis)

	ready.Checks = append(ready.Checks, checks...)

	return &ready, nil
}
