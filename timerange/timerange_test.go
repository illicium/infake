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

func TestTimeRangeContains(t *testing.T) {
	tr := TimeRange{
		From: time.Date(2016, time.January, 1, 0, 0, 0, 0, time.UTC),
		To:   time.Date(2016, time.January, 1, 23, 59, 59, 0, time.UTC),
	}

	assert.True(t, tr.Contains(tr.From))
	assert.False(t, tr.Contains(tr.To))

	assert.True(t, tr.Contains(time.Date(2016, time.January, 1, 5, 0, 0, 0, time.UTC)))
	assert.False(t, tr.Contains(time.Date(2015, time.December, 31, 0, 0, 0, 0, time.UTC)))
	assert.False(t, tr.Contains(time.Date(2016, time.January, 1, 23, 59, 59, 123, time.UTC)))
}

func TestTimeRangeContainsInclusive(t *testing.T) {
	tr := TimeRange{
		From: time.Date(2016, time.January, 1, 0, 0, 0, 0, time.UTC),
		To:   time.Date(2016, time.January, 1, 23, 59, 59, 0, time.UTC),
	}

	assert.True(t, tr.ContainsInclusive(tr.From))
	assert.True(t, tr.ContainsInclusive(tr.To))

	assert.True(t, tr.ContainsInclusive(time.Date(2016, time.January, 1, 5, 0, 0, 0, time.UTC)))
	assert.False(t, tr.ContainsInclusive(time.Date(2015, time.December, 31, 0, 0, 0, 0, time.UTC)))
	assert.False(t, tr.ContainsInclusive(time.Date(2016, time.January, 1, 23, 59, 59, 123, time.UTC)))
}
