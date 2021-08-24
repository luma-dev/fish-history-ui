package ptr

import (
	"github.com/luma-dev/fish-history-ui/projects/cli/gen/models"
)

func NewTimestamp(v models.Timestamp) *models.Timestamp {
	return &v
}

func NewTimestampNullable(v models.TimestampNullable) *models.TimestampNullable {
	return &v
}

func NewTimeUnit(v models.TimeUnit) *models.TimeUnit {
	return &v
}
