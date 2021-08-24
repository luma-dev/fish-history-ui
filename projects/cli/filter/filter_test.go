package filter

import (
	"testing"
	"time"

	"github.com/go-ptr/go-ptr/ptr"
	"github.com/google/go-cmp/cmp"
	"github.com/luma-dev/fish-history-ui/projects/cli/gen/models"
	iptr "github.com/luma-dev/fish-history-ui/projects/cli/internal/ptr"
)

func TestChunkHistory(t *testing.T) {
	type Case struct {
		got  models.HistoryBlocks
		want models.HistoryBlocks
	}
	// 2010/1/1 00:00
	h1 := func() *models.History {
		return &models.History{
			Cmd:  ptr.NewString("echo 1"),
			When: iptr.NewTimestamp(1262304000),
		}
	}
	// 2010/1/1 01:00
	h2 := func() *models.History {
		return &models.History{
			Cmd:  ptr.NewString("echo 2"),
			When: iptr.NewTimestamp(1262307600),
		}
	}
	// 2010/1/1 02:00
	h3 := func() *models.History {
		return &models.History{
			Cmd:  ptr.NewString("echo 3"),
			When: iptr.NewTimestamp(1262311200),
		}
	}
	// 2010/1/3 01:00
	h4 := func() *models.History {
		return &models.History{
			Cmd:  ptr.NewString("echo 4"),
			When: iptr.NewTimestamp(1262480400),
		}
	}
	// 2010/1/3 01:30
	h5 := func() *models.History {
		return &models.History{
			Cmd:  ptr.NewString("echo 5"),
			When: iptr.NewTimestamp(1262482200),
		}
	}
	cases := []Case{
		{
			got: ChunkHistory([]models.History{
				*h1(),
				*h2(),
				*h3(),
				*h4(),
				*h5(),
			}, models.TimeUnitDay, -1, nil),
			want: models.HistoryBlocks{
				Unit: iptr.NewTimeUnit(models.TimeUnitDay),
				Blocks: []*models.HistoryBlock{
					{
						From:  iptr.NewTimestamp(1262304000),
						To:    iptr.NewTimestamp(1262390400),
						Count: ptr.NewInt64(3),
						Histories: []*models.History{
							h1(),
							h2(),
							h3(),
						},
					},
					{
						From:  iptr.NewTimestamp(1262476800),
						To:    iptr.NewTimestamp(1262563200),
						Count: ptr.NewInt64(2),
						Histories: []*models.History{
							h4(),
							h5(),
						},
					},
				},
			},
		},
		{
			got: ChunkHistory([]models.History{
				*h1(),
				*h2(),
				*h3(),
				*h4(),
				*h5(),
			}, models.TimeUnitDay, -1, &models.TimeRangeFilter{
				From: iptr.NewTimestampNullable(1262476800),
			}),
			want: models.HistoryBlocks{
				Unit: iptr.NewTimeUnit(models.TimeUnitDay),
				Blocks: []*models.HistoryBlock{
					{
						From:  iptr.NewTimestamp(1262476800),
						To:    iptr.NewTimestamp(1262563200),
						Count: ptr.NewInt64(2),
						Histories: []*models.History{
							h4(),
							h5(),
						},
					},
				},
			},
		},
		{
			got: ChunkHistory([]models.History{
				*h1(),
				*h2(),
				*h3(),
				*h4(),
				*h5(),
			}, models.TimeUnitDay, 2, nil),
			want: models.HistoryBlocks{
				Unit: iptr.NewTimeUnit(models.TimeUnitDay),
				Blocks: []*models.HistoryBlock{
					{
						From:  iptr.NewTimestamp(1262304000),
						To:    iptr.NewTimestamp(1262390400),
						Count: ptr.NewInt64(3),
						Histories: []*models.History{
							h1(),
							h2(),
						},
					},
					{
						From:  iptr.NewTimestamp(1262476800),
						To:    iptr.NewTimestamp(1262563200),
						Count: ptr.NewInt64(2),
						Histories: []*models.History{
							h4(),
							h5(),
						},
					},
				},
			},
		},
		{
			got: ChunkHistory([]models.History{
				*h1(),
				*h2(),
				*h3(),
				*h4(),
				*h5(),
			}, models.TimeUnitDay, 0, nil),
			want: models.HistoryBlocks{
				Unit: iptr.NewTimeUnit(models.TimeUnitDay),
				Blocks: []*models.HistoryBlock{
					{
						From:      iptr.NewTimestamp(1262304000),
						To:        iptr.NewTimestamp(1262390400),
						Count:     ptr.NewInt64(3),
						Histories: []*models.History{},
					},
					{
						From:      iptr.NewTimestamp(1262476800),
						To:        iptr.NewTimestamp(1262563200),
						Count:     ptr.NewInt64(2),
						Histories: []*models.History{},
					},
				},
			},
		},
		{
			got: ChunkHistory([]models.History{
				*h1(),
				*h2(),
				*h3(),
				*h4(),
				*h5(),
			}, models.TimeUnitAll, 0, nil),
			want: models.HistoryBlocks{
				Unit: iptr.NewTimeUnit(models.TimeUnitAll),
				Blocks: []*models.HistoryBlock{
					{
						From:      iptr.NewTimestamp(0),
						To:        iptr.NewTimestamp(2147483647),
						Count:     ptr.NewInt64(5),
						Histories: []*models.History{},
					},
				},
			},
		},
	}
	for i, c := range cases {
		if diff := cmp.Diff(c.want, c.got); diff != "" {
			t.Errorf("%d: Unmarshal() mismatch (-want +got):\n%s", i, diff)
		}
	}
}

func TestNewUnitNumber(t *testing.T) {
	type Case struct {
		date1 time.Time
		date2 time.Time
		unit  models.TimeUnit
		eq    bool
	}
	cases := []Case{
		{
			date1: time.Date(2011, time.February, 1, 0, 0, 0, 0, time.UTC),
			date2: time.Date(2014, time.January, 2, 0, 0, 0, 0, time.UTC),
			unit:  models.TimeUnitAll,
			eq:    true,
		},
		{
			date1: time.Date(2011, time.February, 1, 0, 0, 0, 0, time.UTC),
			date2: time.Date(2014, time.January, 2, 0, 0, 0, 0, time.UTC),
			unit:  models.TimeUnitYear,
			eq:    false,
		},
		{
			date1: time.Date(2014, time.February, 1, 0, 0, 0, 0, time.UTC),
			date2: time.Date(2014, time.January, 2, 0, 0, 0, 0, time.UTC),
			unit:  models.TimeUnitYear,
			eq:    true,
		},
		{
			date1: time.Date(2011, time.February, 1, 0, 0, 0, 0, time.UTC),
			date2: time.Date(2014, time.January, 2, 0, 0, 0, 0, time.UTC),
			unit:  models.TimeUnitMonth,
			eq:    false,
		},
		{
			date1: time.Date(2014, time.February, 1, 0, 0, 0, 0, time.UTC),
			date2: time.Date(2014, time.January, 2, 0, 0, 0, 0, time.UTC),
			unit:  models.TimeUnitMonth,
			eq:    false,
		},
		{
			date1: time.Date(2014, time.January, 31, 0, 0, 0, 0, time.UTC),
			date2: time.Date(2014, time.January, 2, 0, 0, 0, 0, time.UTC),
			unit:  models.TimeUnitMonth,
			eq:    true,
		},
		{
			date1: time.Date(2013, time.January, 31, 0, 0, 0, 0, time.UTC),
			date2: time.Date(2014, time.January, 2, 0, 0, 0, 0, time.UTC),
			unit:  models.TimeUnitMonth,
			eq:    false,
		},
		{
			date1: time.Date(2014, time.January, 31, 0, 0, 0, 0, time.UTC),
			date2: time.Date(2014, time.January, 2, 0, 0, 0, 0, time.UTC),
			unit:  models.TimeUnitWeek,
			eq:    false,
		},
		{
			date1: time.Date(2014, time.January, 11, 23, 59, 59, 0, time.UTC),
			date2: time.Date(2014, time.January, 5, 0, 0, 0, 0, time.UTC),
			unit:  models.TimeUnitWeek,
			eq:    true,
		},
		{
			date1: time.Date(2014, time.January, 12, 0, 0, 0, 0, time.UTC),
			date2: time.Date(2014, time.January, 11, 23, 59, 59, 0, time.UTC),
			unit:  models.TimeUnitWeek,
			eq:    false,
		},
		{
			date1: time.Date(2014, time.January, 9, 0, 0, 0, 0, time.UTC), date2: time.Date(2014, time.January, 3, 0, 0, 0, 0, time.UTC),
			unit: models.TimeUnitDay,
			eq:   false,
		},
		{
			date1: time.Date(2014, time.January, 2, 0, 0, 0, 0, time.UTC),
			date2: time.Date(2014, time.January, 3, 10, 0, 0, 0, time.UTC),
			unit:  models.TimeUnitDay,
			eq:    false,
		},
		{
			date1: time.Date(2014, time.January, 3, 0, 0, 0, 0, time.UTC),
			date2: time.Date(2014, time.January, 3, 10, 0, 0, 0, time.UTC),
			unit:  models.TimeUnitDay,
			eq:    true,
		},
		{
			date1: time.Date(2014, time.January, 3, 10, 30, 0, 0, time.UTC),
			date2: time.Date(2014, time.January, 3, 9, 0, 0, 0, time.UTC),
			unit:  models.TimeUnitHour,
			eq:    false,
		},
		{
			date1: time.Date(2014, time.January, 3, 10, 30, 0, 0, time.UTC),
			date2: time.Date(2014, time.January, 3, 10, 0, 0, 0, time.UTC),
			unit:  models.TimeUnitHour,
			eq:    true,
		},
		{
			date1: time.Date(2014, time.January, 3, 10, 10, 30, 0, time.UTC),
			date2: time.Date(2014, time.January, 3, 10, 30, 0, 0, time.UTC),
			unit:  models.TimeUnitMinute,
			eq:    false,
		},
		{
			date1: time.Date(2014, time.January, 3, 10, 30, 30, 0, time.UTC),
			date2: time.Date(2014, time.January, 3, 10, 30, 0, 0, time.UTC),
			unit:  models.TimeUnitMinute,
			eq:    true,
		},
		{
			date1: time.Date(2014, time.January, 3, 10, 30, 27, 0, time.UTC),
			date2: time.Date(2014, time.January, 3, 10, 30, 30, 0, time.UTC),
			unit:  models.TimeUnitSecond,
			eq:    false,
		},
		{
			date1: time.Date(2014, time.January, 3, 10, 30, 30, 0, time.UTC),
			date2: time.Date(2014, time.January, 3, 10, 30, 30, 0, time.UTC),
			unit:  models.TimeUnitSecond,
			eq:    true,
		},
	}
	for i, c := range cases {
		n1 := NewUnitNumber(
			models.Timestamp(c.date1.Unix()),
			c.unit,
		).Number
		n2 := NewUnitNumber(
			models.Timestamp(c.date2.Unix()),
			c.unit,
		).Number
		if (n1 == n2) != c.eq {
			verb := "not equeal"
			if c.eq {
				verb = "equeal"
			}
			t.Errorf("#%d: NewUnitNumber(%v) = %v and NewUnitNumber(%v) = %v should %v in unit %v.", i, c.date1, n1, c.date2, n2, verb, c.unit)
		}
	}
}
