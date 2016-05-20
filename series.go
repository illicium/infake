package infake

import (
	"log"
	"math/rand"
	"text/template"
	"time"

	"revision.aeip.apigee.net/dia/infake/timerange"
)

type Series struct {
	Id        string
	TimeRange *timerange.TimeRange
	Name      StringTemplate
	Tags      map[string]StringTemplate
	Fields    map[string]Value
	Variables []Variable

	times []time.Time
}

type SeriesConfig struct {
	Id        string
	TimeRange timerange.TimeRangeConfig
	Name      string
	Tags      map[string]string
	Fields    map[string]ValueConfig
	Variables []Variable
}

func NewSeries(c SeriesConfig) (*Series, error) {
	var err error

	s := &Series{
		Id:     c.Id,
		Tags:   make(map[string]StringTemplate),
		Fields: make(map[string]Value),
	}

	// TimeRange

	s.TimeRange, err = timerange.New(c.TimeRange)

	if err != nil {
		return nil, err
	}

	// Name

	var nameTpl *template.Template

	nameTpl, err = template.New("Name").Parse(c.Name)

	if err != nil {
		return nil, err
	}

	s.Name = StringTemplate{nameTpl}

	// Tags

	for k, v := range c.Tags {
		tagTpl, err := template.New(k).Parse(v)

		if err != nil {
			return nil, err
		}

		s.Tags[k] = StringTemplate{tagTpl}
	}

	// Fields

	for k, valCfg := range c.Fields {
		val, err := NewValue(valCfg)

		if err != nil {
			return nil, err
		}

		s.Fields[k] = val
	}

	// times

	s.times, err = s.TimeRange.Values()

	if err != nil {
		return nil, err
	}

	return s, nil
}

func (s *Series) getFields(rnd *rand.Rand) map[string]interface{} {
	f := make(map[string]interface{})

	for k, v := range s.Fields {
		f[k] = v.Get(rnd)
	}

	return f
}

func (s *Series) genPoint(rnd *rand.Rand, boundVars map[string]interface{}, vars []Variable, t time.Time, c chan<- Point) error {
	if len(vars) > 0 {
		expanded, err := vars[0].Expand()

		if err != nil {
			return err
		}

		for _, ev := range expanded {
			boundVars[ev.Name] = ev.Value

			if err := s.genPoint(rnd, boundVars, vars[1:], t, c); err != nil {
				return err
			}
		}
	} else {
		name, err := s.Name.Execute(boundVars)

		if err != nil {
			return err
		}

		tags := make(map[string]string)

		for k, v := range s.Tags {
			tag, err := v.Execute(boundVars)

			if err != nil {
				return err
			}

			tags[k] = tag
		}

		fields := s.getFields(rnd)

		p, err := NewPoint(name, tags, fields, t)

		if err != nil {
			return err
		}

		c <- p
	}

	return nil
}

func (s *Series) Generate(rndSrc rand.Source) (<-chan Point, error) {
	c := make(chan Point)

	go func() {
		defer close(c)

		rnd := rand.New(rndSrc)

		log.Printf("Generating series: %q\n", s.Id)

		for _, t := range s.times {
			boundVars := make(map[string]interface{})

			if err := s.genPoint(rnd, boundVars, s.Variables, t, c); err != nil {
				log.Print(err)
			}
		}
	}()

	return c, nil
}
