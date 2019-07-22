package logger

import (
	"fmt"
	"reflect"
	"strconv"
	"sync"
	"testing"
)

// flag -race required
func TestEntry_WithField(t *testing.T) {
	routinesNumber := 10
	deep := 3

	var wg sync.WaitGroup
	wg.Add(routinesNumber)

	for i := 0; i < routinesNumber; i++ {
		go func(i int) {
			meta := meta{
				UserID:    strconv.Itoa(i),
				RequestID: strconv.Itoa(i),
			}

			e := entry{meta: meta}

			for j := 0; j < deep; j++ {
				key := fmt.Sprintf("%d_%d", i, j)
				fields := Fields{key: "log"}

				e2 := e.WithFields(fields).(*entry)

				if len(e2.content) != j+1 {
					t.Error()
				}

				if !reflect.DeepEqual(meta, e2.meta) {
					t.Error("Meta has been changed")
				}

				// Checks that's contains parent content
				for key, expected := range e.content {
					if actual := e2.content[key]; actual != expected {
						t.Errorf("Content not exsists or not equals. Expected '%s', actual '%s'", expected, actual)
					}
				}

				// Checks that's meta exists and didn't been changed
				if !reflect.DeepEqual(meta, e2.meta) {
					t.Error("Meta has been changed")
				}

				e = *e2
			}

			wg.Done()
		}(i)
	}

	wg.Wait()
}

func TestEntry_WithFields_ContentImmutable(t *testing.T) {
	immutable := &entry{
		content: []field{
			{"1", "1"},
		},
	}

	e := immutable.WithFields(Fields{"2": "2"}).(*entry)

	if len(immutable.content) != 1 {
		t.Error("Parent content modified")
	}

	if len(e.content) != 2 {
		t.Error("Not enough elements in content")
	}
}
