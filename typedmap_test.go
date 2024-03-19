package typedmap_test

import (
	"testing"

	"github.com/thetechpanda/typedmap"
)

func TestNew(t *testing.T) {
	if typedmap.New[string, int]().Has(`k`) {
		t.Errorf("typedmap.New[string, int]().Has(`k`) expected false, got true")
	}

	if typedmap.NewWithMap[string, int](nil).Has(`k`) {
		t.Errorf("typedmap.NewWithMap[string, int](nil).Has(`k`) expected false, got true")
	}

	if _, ok := typedmap.NewSyncMap[string, int]().Load(`k`); ok {
		t.Errorf("typedmap.NewSyncMap[string, int]().Load(`k`) expected false, got true")
	}

	if _, ok := typedmap.NewSyncMapCompatible[string, int]().Load(`k`); ok {
		t.Errorf("typedmap.NewSyncMapCompatible[string, int]().Load(`k`) expected false, got true")
	}
}
