package uuidv7_test

import (
	"testing"
	"time"

	"github.com/TinyMurky/uuidv7"
)

// TestUUIDMonotonicityInMicroSecond test uuid in micro second is monotonically increasing
func TestUUIDMonotonicityInMicroSecond(t *testing.T) {
	testTime := time.Date(2006, time.January, 1, 0, 0, 0, 0, time.Local)

	preUUID, err := uuidv7.FromTime(testTime)

	if err != nil {
		t.Errorf("expect no err, got FromTime err: %s", err.Error())
	}

	for range 10000 {
		testTime = testTime.Add(time.Microsecond)
		currentUUID, err := uuidv7.FromTime(testTime)

		if err != nil {
			t.Errorf("expect no err, got FromTime err: %s", err.Error())
		}

		if preUUID.String() >= currentUUID.String() {
			t.Errorf("expect preUUID \"%s\" less than currentUUID \"%s\", got larger or equel instead", preUUID, currentUUID)
		}
	}
}

// TestZeroUUIDv7 test ZeroUUIDv7
func TestZeroUUIDv7(t *testing.T) {
	zero := uuidv7.ZeroUUIDv7()

	if !zero.IsZero() {
		t.Errorf("expect IsZero is true, got false, uuid: %v", zero)
	}

	expectZeroUUIDStr := "00000000-0000-0000-0000-000000000000"
	if zero.String() != expectZeroUUIDStr {
		t.Errorf("expect ZeroUUIDv7 string is %q, got %q", expectZeroUUIDStr, zero.String())
	}
}

func BenchmarkNewUUIDv7(b *testing.B) {
	for b.Loop() {
		uuidv7.New()
	}
}

func BenchmarkNewUUIDv7Parallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			uuidv7.New()
		}
	})
}
