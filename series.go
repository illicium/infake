package infake

import (
	"log"
	"math/rand"
	"text/template"
)

type Series struct {
	Id        string
	Name      string
	Tags      map[string]string
	Fields    map[string]string
	Variables []Variable
}

type SeriesTemplates struct {
	Name StringTemplate
	Tags map[string]StringTemplate
}

func (s *Series) Templates() (*SeriesTemplates, error) {
	nameTpl, err := template.New("Name").Parse(s.Name)

	if err != nil {
		return nil, err
	}

	tagTpls := make(map[string]StringTemplate)

	for k, v := range s.Tags {
		tagTpl, err := template.New(k).Parse(v)

		if err != nil {
			return nil, err
		}

		tagTpls[k] = StringTemplate{tagTpl}
	}

	return &SeriesTemplates{Name: StringTemplate{nameTpl}, Tags: tagTpls}, nil
}

func (s *Series) genFields(rnd *rand.Rand) (map[string]interface{}, error) {
	f := make(map[string]interface{})

	for k, v := range s.Fields {
		switch v {
		case "int":
			f[k] = rnd.Int31n(2000) - 1000
		case "uint":
			f[k] = rnd.Uint32()
		case "float":
			f[k] = rnd.Float64()
		default:
			f[k] = v
		}
	}

	return f, nil
}

func (s *Series) genPoint(rnd *rand.Rand, tpls *SeriesTemplates, boundVars map[string]interface{}, vars []Variable, c chan<- Point) error {
	if len(vars) > 0 {
		expanded, err := vars[0].Expand()

		if err != nil {
			return err
		}

		for _, ev := range expanded {
			boundVars[ev.Name] = ev.Value

			if err := s.genPoint(rnd, tpls, boundVars, vars[1:], c); err != nil {
				return err
			}
		}
	} else {
		name, err := tpls.Name.Execute(boundVars)

		if err != nil {
			return err
		}

		tags := make(map[string]string)

		for k, v := range tpls.Tags {
			tag, err := v.Execute(boundVars)

			if err != nil {
				return err
			}

			tags[k] = tag
		}

		fields, err := s.genFields(rnd)

		if err != nil {
			return err
		}

		p, err := NewPoint(name, tags, fields)

		if err != nil {
			return err
		}

		c <- p
	}

	return nil
}

func (s *Series) Generate(rnd *rand.Rand) (<-chan Point, error) {
	c := make(chan Point)

	go func() {
		defer close(c)

		log.Printf("Generating series: %q\n", s.Id)

		tpls, err := s.Templates()

		if err != nil {
			log.Print(err)
			return
		}

		boundVars := make(map[string]interface{})

		if err := s.genPoint(rnd, tpls, boundVars, s.Variables, c); err != nil {
			log.Print(err)
		}
	}()

	return c, nil
}

func (s *Series) Expand() ([]Series, error) {
	series := []Series{}

	return series, nil
}
