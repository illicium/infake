package timerange

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func testTimeRangeValues(t *testing.T, tr TimeRange, want []time.Time) {
	got, err := tr.Values()

	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, want, got)
}

func TestTimeRangeValues(t *testing.T) {
	tr := TimeRange{
		From: time.Date(2016, time.January, 1, 0, 0, 0, 0, time.UTC),
		To:   time.Date(2016, time.January, 1, 0, 0, 3, 0, time.UTC),
		Step: time.Second,
	}

	want := []time.Time{
		time.Date(2016, time.January, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2016, time.January, 1, 0, 0, 1, 0, time.UTC),
		time.Date(2016, time.January, 1, 0, 0, 2, 0, time.UTC),
		time.Date(2016, time.January, 1, 0, 0, 3, 0, time.UTC),
	}

	testTimeRangeValues(t, tr, want)
}

func TestTimeRangeValuesBackwards(t *testing.T) {
	tr := TimeRange{
		From: time.Date(2016, time.January, 1, 0, 0, 3, 0, time.UTC),
		To:   time.Date(2016, time.January, 1, 0, 0, 0, 0, time.UTC),
		Step: -1 * time.Second,
	}

	want := []time.Time{
		time.Date(2016, time.January, 1, 0, 0, 3, 0, time.UTC),
		time.Date(2016, time.January, 1, 0, 0, 2, 0, time.UTC),
		time.Date(2016, time.January, 1, 0, 0, 1, 0, time.UTC),
		time.Date(2016, time.January, 1, 0, 0, 0, 0, time.UTC),
	}

	testTimeRangeValues(t, tr, want)
}

func TestTimeRangeValuesBadStepOne(t *testing.T) {
	tr := TimeRange{
		From: time.Date(2016, time.January, 1, 0, 0, 0, 0, time.UTC),
		To:   time.Date(2016, time.January, 1, 23, 59, 59, 0, time.UTC),
		Step: -1 * time.Second,
	}

	_, err := tr.Values()

	if err == nil {
		t.Fatal("expected error")
	}
}

func TestTimeRangeValuesBadStepTwo(t *testing.T) {
	tr := TimeRange{
		From: time.Date(2016, time.January, 1, 23, 59, 59, 0, time.UTC),
		To:   time.Date(2016, time.January, 1, 0, 0, 0, 0, time.UTC),
		Step: 1 * time.Second,
	}

	_, err := tr.Values()

	if err == nil {
		t.Fatal("expected error")
	}
}

func TestTimeRangeValuesZeroStepOne(t *testing.T) {
	tr := TimeRange{
		From: time.Date(2016, time.January, 1, 0, 0, 0, 0, time.UTC),
		To:   time.Date(2016, time.January, 1, 23, 59, 59, 0, time.UTC),
	}

	want := []time.Time{
		time.Date(2016, time.January, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2016, time.January, 1, 23, 59, 59, 0, time.UTC),
	}

	testTimeRangeValues(t, tr, want)
}

func TestTimeRangeValuesZeroStepTwo(t *testing.T) {
	tr := TimeRange{
		From: time.Date(2016, time.January, 1, 23, 59, 59, 0, time.UTC),
		To:   time.Date(2016, time.January, 1, 0, 0, 0, 0, time.UTC),
	}

	want := []time.Time{
		time.Date(2016, time.January, 1, 23, 59, 59, 0, time.UTC),
		time.Date(2016, time.January, 1, 0, 0, 0, 0, time.UTC),
	}

	testTimeRangeValues(t, tr, want)
}

func TestTimeRangeValuesZeroStepEqual(t *testing.T) {
	tr := TimeRange{
		From: time.Date(2016, time.January, 1, 12, 34, 56, 0, time.UTC),
		To:   time.Date(2016, time.January, 1, 12, 34, 56, 0, time.UTC),
	}

	want := []time.Time{
		time.Date(2016, time.January, 1, 12, 34, 56, 0, time.UTC),
	}

	testTimeRangeValues(t, tr, want)
}
