package randid

import (
	"testing"

	"github.com/labstack/gommon/log"
)

func TestID(t *testing.T) {
	id, err := ID()
	if err != nil {
		t.Errorf("Ar error occurred during TestID %s\n", err)
	}

	if len(id) != 32 {
		t.Errorf("Id returned is incorrect: %s", id)
	}
}

func TestSizedID(t *testing.T) {
	id, err := SizedID(23)
	if err != nil {
		t.Errorf("Ar error occurred during TestSizedID %s\n", err)
	}

	if len(id) != 23 {
		t.Errorf("Id returned is incorrect: %s", id)
	}
}

func TestHugeSizedID(t *testing.T) {
	id, err := SizedID(122)
	if err != nil {
		t.Errorf("Ar error occurred during TestHugeSizedID %s\n", err)
	}

	if len(id) != 122 {
		t.Errorf("Id returned is incorrect: %s", id)
	}
}

func TestMoreHugeSizedID(t *testing.T) {
	id, err := SizedID(1000000)
	if err != nil {
		t.Errorf("Ar error occurred during TestMoreSizedID %s\n", err)
	}

	if len(id) != 1000000 {
		t.Errorf("Id returned is incorrect: %s", id)
	}
}

func TestChangeDefLenID(t *testing.T) {
	DefaultLen = 12
	id, err := ID()
	if err != nil {
		t.Errorf("Ar error occurred during TestChangeDefID %s\n", err)
	}

	if len(id) != 12 {
		t.Errorf("Id returned is incorrect: %s", id)
	}
}

func TestHardLimitLength(t *testing.T) {
	id, err := SizedID(1000001)
	if err != nil {
		log.Errorf("Error occurred on TestHardLimitLength: %s\n", err)
	}
	if len(id) != 1000000 {
		t.Errorf("Id size doesn't respect hardlimit")
	}
}
