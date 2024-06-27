package trace

import (
	"fmt"
	"io"
	"os"

	"github.com/threadedstream/trace"
)

// ParsedTraceV2 is a parsed representation of >= Go122 traces
type ParsedTraceV2 struct {
	events  []Event
	summary *trace.Summary
}

func ParseTraceV2(tr io.Reader) (*ParsedTraceV2, error) {
	r, err := NewReader(tr)
	if err != nil {
		return nil, fmt.Errorf("failed to create trace reader: %w", err)
	}
	s := trace.NewSummarizer()
	t := new(ParsedTraceV2)
	for {
		ev, err := r.ReadEvent()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, fmt.Errorf("failed to read event: %w", err)
		}
		t.events = append(t.events, ev)
		s.Event(&t.events[len(t.events)-1])
	}
	t.summary = s.Finalize()
	return t, nil
}

// IsTraceV2 returns true if filename holds a v2 trace.
func IsTraceV2(filename string) bool {
	file, err := os.Open(filename)
	if err != nil {
		return false
	}
	defer file.Close()

	ver, _, err := trace.ReadVersion(file)
	if err != nil {
		return false
	}
	return ver >= 1022
}
