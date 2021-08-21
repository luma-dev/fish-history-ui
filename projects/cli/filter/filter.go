package filter

import (
	"math"
	"time"

	"github.com/go-ptr/go-ptr/ptr"
	"github.com/luma-dev/fishis/projects/cli/gen/models"
	iptr "github.com/luma-dev/fishis/projects/cli/internal/ptr"
)

func ChunkHistory(h []models.History, unit models.TimeUnit, limit int64, filter *models.TimeRangeFilter) models.HistoryBlocks {
	var (
		from   *models.Timestamp
		to     *models.Timestamp
		blocks = models.HistoryBlocks{
			Unit:   &unit,
			Blocks: []*models.HistoryBlock{},
		}
		block *models.HistoryBlock
		last  int64 = -1
		rest        = limit
		count int64 = 0
		check       = func() {
			if count > 0 {
				n := UnitNumber{
					Unit:   unit,
					Number: last,
				}
				block.Count = ptr.NewInt64(count)
				block.From = iptr.NewTimestamp(n.ToTimestamp())
				n.Number = n.Number + 1
				block.To = iptr.NewTimestamp(n.ToTimestamp())
				blocks.Blocks = append(blocks.Blocks, block)
				count = 0
			}
			block = &models.HistoryBlock{
				Histories: []*models.History{},
			}
		}
	)
	if filter != nil {
		if filter.From != nil {
			from = iptr.NewTimestamp(models.Timestamp(*filter.From))
		}
		if filter.To != nil {
			to = iptr.NewTimestamp(models.Timestamp(*filter.To))
		}
	}
	for _, e := range h {
		if (from != nil && *e.When < *from) || (to != nil && *to < *e.When) {
			continue
		}
		e2 := e
		n := NewUnitNumber(*e.When, unit)
		if n.Number != last {
			check()
			last = n.Number
			rest = limit
		}
		count = count + 1
		if limit < 0 || rest > 0 {
			block.Histories = append(block.Histories, &e2)
			if rest > 0 {
				rest = rest - 1
			}
		}
	}
	check()
	return blocks
}

type UnitNumber struct {
	Unit   models.TimeUnit
	Number int64
}

func NewUnitNumber(t models.Timestamp, unit models.TimeUnit) UnitNumber {
	un := UnitNumber{
		Unit: unit,
	}
	switch unit {
	case models.TimeUnitAll:
		un.Number = 0
	case models.TimeUnitYear:
		un.Number = int64(time.Unix(int64(t), 0).Year())
	case models.TimeUnitMonth:
		d := time.Unix(int64(t), 0)
		un.Number = int64(d.Year())*12 + int64(d.Month()) - 1
	case models.TimeUnitWeek:
		un.Number = (int64(t) + 60*60*24*4) / (60 * 60 * 24 * 7)
	case models.TimeUnitDay:
		un.Number = int64(t) / (60 * 60 * 24)
	case models.TimeUnitHour:
		un.Number = int64(t) / (60 * 60)
	case models.TimeUnitMinute:
		un.Number = int64(t) / 60
	default:
		un.Number = int64(t)
	}
	return un
}

func (un UnitNumber) ToTimestamp() models.Timestamp {
	switch un.Unit {
	case models.TimeUnitAll:
		if un.Number == 0 {
			return 0
		}
		return math.MaxInt32
	case models.TimeUnitYear:
		return models.Timestamp(time.Date(int(un.Number), time.January, 1, 0, 0, 0, 0, time.UTC).Unix())
	case models.TimeUnitMonth:
		return models.Timestamp(time.Date(int(un.Number/12), time.Month(un.Number%12+1), 1, 0, 0, 0, 0, time.UTC).Unix())
	case models.TimeUnitWeek:
		return models.Timestamp(un.Number*60*60*24*7 - 60*60*24*4)
	case models.TimeUnitDay:
		return models.Timestamp(un.Number * (60 * 60 * 24))
	case models.TimeUnitHour:
		return models.Timestamp(un.Number * (60 * 60))
	case models.TimeUnitMinute:
		return models.Timestamp(un.Number * 60)
	default:
		return models.Timestamp(un.Number)
	}
}
