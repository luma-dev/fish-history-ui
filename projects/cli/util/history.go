package util

import (
	"github.com/luma-dev/fish-history-ui/projects/cli/gen/models"
	iptr "github.com/luma-dev/fish-history-ui/projects/cli/internal/ptr"
)

func ShiftHistories(h []models.History, t int64) (r []models.History) {
	for _, e := range h {
		r = append(r, ShiftHistory(e, t))
	}
	return
}

func ShiftHistory(h models.History, t int64) (r models.History) {
	r = h
	r.When = iptr.NewTimestamp(models.Timestamp(int64(*h.When) + t))
	return
}
