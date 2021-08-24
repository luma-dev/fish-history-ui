package parser

import (
	"testing"

	"github.com/go-ptr/go-ptr/ptr"
	"github.com/google/go-cmp/cmp"

	"github.com/luma-dev/fish-history-ui/projects/cli/gen/models"
	iptr "github.com/luma-dev/fish-history-ui/projects/cli/internal/ptr"
)

func TestParseHistories(t *testing.T) {
	got := ParseHistories([]byte(`
- cmd: echo hello 1
  when: 1600000000
- cmd: echo hello 2\necho and bye
  when: 1600000004
- cmd: cat hello you
  when: 1600000006
  paths:
    - ./hello
    - ./you
- cmd: echo hello 3
  when: 1600000008
- cmd: cat hello
  when: 1600000010
  paths:
    - ./hello
`))
	want := []models.History{
		{
			Cmd:  ptr.NewString("echo hello 1"),
			When: iptr.NewTimestamp(1600000000),
		},
		{
			Cmd:  ptr.NewString("echo hello 2\necho and bye"),
			When: iptr.NewTimestamp(1600000004),
		},
		{
			Cmd:  ptr.NewString("cat hello you"),
			When: iptr.NewTimestamp(1600000006),
			Paths: []string{
				"./hello",
				"./you",
			},
		},
		{
			Cmd:  ptr.NewString("echo hello 3"),
			When: iptr.NewTimestamp(1600000008),
		},
		{
			Cmd:  ptr.NewString("cat hello"),
			When: iptr.NewTimestamp(1600000010),
			Paths: []string{
				"./hello",
			},
		},
	}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("Unmarshal() mismatch (-want +got):\n%s", diff)
	}
}
