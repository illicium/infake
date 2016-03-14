package timerange

import (
	"errors"
	"time"
)

type TimeRange struct {
	From time.Time
	To   time.Time
	Step time.Duration
}

type TimeRangeConfig struct {
	From string
	To   string
	Step string
}

func New(c TimeRangeConfig) (*TimeRange, error) {
	var from time.Time
	var to time.Time
	var step time.Duration
	var err error

	if c.From != "" {
		from, err = time.Parse(time.RFC3339Nano, c.From)

		if err != nil {
			return nil, err
		}
	}

	if c.To != "" {
		to, err = time.Parse(time.RFC3339Nano, c.To)

		if err != nil {
			return nil, err
		}
	}

	if c.Step != "" {
		step, err = time.ParseDuration(c.Step)

		if err != nil {
			return nil, err
		}
	}

	return &TimeRange{from, to, step}, nil
}

func (tr *TimeRange) Values() ([]time.Time, error) {
	var from time.Time
	var to time.Time
	var step time.Duration

	now := time.Now()

	if !tr.From.IsZero() {
		from = tr.From
	} else {
		from = now
	}

	if !tr.To.IsZero() {
		to = tr.To
	} else {
		to = now
	}

	if tr.Step != 0 {
		step = tr.Step
	} else {
		step = to.Sub(from)
	}

	if from.Equal(to) && step == 0 {
		return []time.Time{from}, nil
	}

	if from.Before(to) && step < 0 {
		return nil, errors.New("infinite sequence: From < To, Step < 0")
	}

	if to.Before(from) && step > 0 {
		return nil, errors.New("infinite sequence: To < From, Step > 0")
	}

	var n uint

	if step > 0 { // from < to
		n = uint(to.Sub(from)/step + 1)
	} else {
		n = uint(from.Sub(to)/-step + 1)
	}

	vals := make([]time.Time, 0, n)

	for t := from; t.Equal(to) || (step > 0 && t.Before(to)) || (step < 0 && t.After(to)); t = t.Add(step) {
		vals = append(vals, t)
	}

	return vals, nil
}
