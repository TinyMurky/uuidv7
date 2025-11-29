package uuidv7_test

import (
	"reflect"
	"testing"
	"time"

	"github.com/TinyMurky/uuidv7"
)

func TestParse(t *testing.T) {
	testTime := time.Date(2006, time.January, 1, 0, 0, 0, 0, time.Local)

	for range 10 {
		testTime = testTime.Add(time.Second)
		uuid, err := uuidv7.FromTime(testTime)

		if err != nil {
			t.Errorf("expect no err, got FromTime err: %s", err.Error())
		}

		uuidParsed, err := uuidv7.Parse(uuid.String())

		if err != nil {
			t.Errorf("expect no err, got Parse err: %s", err.Error())
		}

		if !reflect.DeepEqual(uuid, uuidParsed) {
			t.Errorf("uuidParsed expect %v, got %v", uuid, uuidParsed)
		}
	}
}
