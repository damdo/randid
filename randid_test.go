package randid

import (
	"testing"
)

func TestID(t *testing.T) {
	id, err := ID()
	if err != nil {
		t.Fatalf("Ar error occurred during TestID %s\n", err)
	}

	if len(id) != 32 {
		t.Fatalf("Id returned is incorrect: %s", id)
	}
}

func TestSizedID(t *testing.T) {
	id, err := SizedID(23)
	if err != nil {
		t.Fatalf("Ar error occurred during TestSizedID %s\n", err)
	}

	if len(id) != 23 {
		t.Fatalf("Id returned is incorrect: %s", id)
	}
}

func TestHugeSizedID(t *testing.T) {
	id, err := SizedID(122)
	if err != nil {
		t.Fatalf("Ar error occurred during TestHugeSizedID %s\n", err)
	}

	if len(id) != 122 {
		t.Fatalf("Id returned is incorrect: %s", id)
	}
}

func TestMoreHugeSizedID(t *testing.T) {
	id, err := SizedID(1000000)
	if err != nil {
		t.Fatalf("Ar error occurred during TestMoreSizedID %s\n", err)
	}

	if len(id) != 1000000 {
		t.Fatalf("Id returned is incorrect: %s", id)
	}
}

func TestChangeDefLenID(t *testing.T) {
	DefaultLen = 12
	id, err := ID()
	if err != nil {
		t.Fatalf("Ar error occurred during TestChangeDefID %s\n", err)
	}

	if len(id) != 12 {
		t.Fatalf("Id returned is incorrect: %s", id)
	}
}
