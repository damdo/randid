package randid

import (
	"testing"
)

func TestID(t *testing.T) {
	id := ID()

	if len(id) != 32 {
		t.Fatalf("Id returned is incorrect: %s", id)
	}
}

func TestSizedID(t *testing.T) {
	id := SizedID(23)

	if len(id) != 23 {
		t.Fatalf("Id returned is incorrect: %s", id)
	}
}

func TestHugeSizedID(t *testing.T) {
	id := SizedID(122)

	if len(id) != 122 {
		t.Fatalf("Id returned is incorrect: %s", id)
	}
}

func TestChangeDefLenID(t *testing.T) {
	DefaultLen = 12
	id := ID()

	if len(id) != 12 {
		t.Fatalf("Id returned is incorrect: %s", id)
	}
}
