package infake

import (
	"reflect"
	"testing"
)

func TestExpandWithValue(t *testing.T) {
	v := Variable{
		Name:  "var",
		Value: "hello world",
	}

	want := []Variable{v}

	got, err := v.Expand()

	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("(%+v).Expand() == %+v; want %+v", v, got, want)
	}
}

func TestExpandValuesEmpty(t *testing.T) {
	v := Variable{
		Name:   "var",
		Values: []VarValue{},
	}

	_, err := v.expandValues()

	if err == nil {
		t.Fatal("expected error")
	}
}

func TestExpandValuesOne(t *testing.T) {
	v := Variable{
		Name:   "var",
		Values: []VarValue{"foo"},
	}

	v2 := Variable{
		Name:  "var",
		Value: "foo",
	}

	want := []Variable{v2}

	got, err := v.expandValues()

	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("(%+v).expandValues() == %+v; want %+v", v, got, want)
	}
}

func TestExpandValuesTwo(t *testing.T) {
	v := Variable{
		Name:   "var",
		Values: []VarValue{"foo", "bar"},
	}

	got, err := v.expandValues()

	if err != nil {
		t.Fatal(err)
	}

	if len(got) != 2 {
		t.Fatalf("(%+v).expandValues(): expected 2 items, got %v", v, len(got))
	}

	if got[0].Value != "foo" {
		t.Fatalf(`(%+v).expandValues(): expected first item to have Value "foo", got %v`, v, got[0].Value)
	}

	if got[1].Value != "bar" {
		t.Fatalf(`(%+v).expandValues(): expected first item to have Value "bar", got %v`, v, got[1].Value)
	}
}

func TestExpandNumericBadStepOne(t *testing.T) {
	v := Variable{
		Name: "var",
		From: 123,
		To:   0,
		Step: 1,
	}

	_, err := v.expandNumeric()

	if err == nil {
		t.Fatal("expected error")
	}
}

func TestExpandNumericBadStepTwo(t *testing.T) {
	v := Variable{
		Name: "var",
		From: 0,
		To:   123,
		Step: -1,
	}

	_, err := v.expandNumeric()

	if err == nil {
		t.Fatal("expected error")
	}
}

func testExpandNumeric(t *testing.T, v Variable, want []float64) {
	got, err := v.expandNumeric()

	if err != nil {
		t.Fatal(err)
	}

	if len(got) != len(want) {
		t.Fatalf("(%+v).expandNumeric(): expected %v items, got %v", v, len(want), len(got))
	}

	for i, val := range want {
		if got[i].Value != val {
			t.Fatalf("(%+v).expandNumeric(): #%v: expected %v, got %v", v, i, val, got[i].Value)
		}
	}
}

func TestExpandNumericPositive(t *testing.T) {
	v := Variable{
		Name: "var",
		From: 0,
		To:   2,
		Step: 0.5,
	}

	testExpandNumeric(t, v, []float64{0, 0.5, 1, 1.5, 2})
}

func TestExpandNumericNegative(t *testing.T) {
	v := Variable{
		Name: "var",
		From: 5,
		To:   -2,
		Step: -2.5,
	}

	testExpandNumeric(t, v, []float64{5, 2.5, 0})
}

func TestExpandNumericZeroStepOne(t *testing.T) {
	v := Variable{
		Name: "var",
		From: 1,
		To:   3,
	}

	testExpandNumeric(t, v, []float64{1, 3})
}

func TestExpandNumericZeroStepTwo(t *testing.T) {
	v := Variable{
		Name: "var",
		From: 3,
		To:   1,
	}

	testExpandNumeric(t, v, []float64{3, 1})
}

func TestExpandNumericZeroStepEqual(t *testing.T) {
	v := Variable{
		Name: "var",
		From: 123,
		To:   123,
	}

	testExpandNumeric(t, v, []float64{123})
}
