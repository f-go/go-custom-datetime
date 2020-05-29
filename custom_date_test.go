package custom_datetime

import (
	"encoding/json"
	"encoding/xml"
	"reflect"
	"testing"
	"time"

	"gopkg.in/yaml.v2"
)

// Define the custom date fomrat
func init() {
	CustomDateFormat = "02•01•2006"
}

func TestCustomDate_MarshalUnmarshalXML(t *testing.T) {
	type Element struct {
		CreatedAt CustomDate `xml:""`
	}
	e := Element{CustomDate(time.Date(2009, 11, 10, 0, 0, 0, 0, time.UTC))}

	want := `<Element><CreatedAt>10•11•2009</CreatedAt></Element>`
	got, err := xml.Marshal(e)
	if err != nil {
		t.Error(err)
	}
	if string(got) != want {
		t.Errorf("got: %s; want: %s", got, want)
	}

	var e2 Element
	if err := xml.Unmarshal([]byte(want), &e2); err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(e, e2) {
		t.Errorf("got: %v; want: %v", e2, e)
	}
}

func TestCustomDate_MarshalUnmarshalXMLAttr(t *testing.T) {
	type Element struct {
		Created CustomDate `xml:"created_at,attr"`
	}
	e := Element{CustomDate(time.Date(2009, 11, 10, 0, 0, 0, 0, time.UTC))}

	want := `<Element created_at="10•11•2009"></Element>`
	got, err := xml.Marshal(e)
	if err != nil {
		t.Error(err)
	}
	if string(got) != want {
		t.Errorf("got: %s; want: %s", got, want)
	}

	var e2 Element
	if err := xml.Unmarshal([]byte(want), &e2); err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(e, e2) {
		t.Errorf("got: %v; want: %v", e2, e)
	}
}

func TestCustomDate_MarshalUnmarshalText(t *testing.T) {
	type Element struct {
		CreatedAt CustomDate `xml:",chardata"`
	}
	e := Element{CustomDate(time.Date(2009, 11, 10, 0, 0, 0, 0, time.UTC))}

	want := `<Element>10•11•2009</Element>`
	got, err := xml.Marshal(e)
	if err != nil {
		t.Error(err)
	}
	if string(got) != want {
		t.Errorf("got: %s; want: %s", got, want)
	}

	var e2 Element
	if err := xml.Unmarshal([]byte(want), &e2); err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(e, e2) {
		t.Errorf("got: %v; want: %v", e2, e)
	}
}

func TestCustomDate_MarshalUnmarshalTextCDATA(t *testing.T) {
	type Element struct {
		CreatedAt CustomDate `xml:",cdata"`
	}
	e := Element{CustomDate(time.Date(2009, 11, 10, 0, 0, 0, 0, time.UTC))}

	want := `<Element><![CDATA[10•11•2009]]></Element>`
	got, err := xml.Marshal(e)
	if err != nil {
		t.Error(err)
	}
	if string(got) != want {
		t.Errorf("got: %s; want: %s", got, want)
	}

	var e2 Element
	if err := xml.Unmarshal([]byte(want), &e2); err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(e, e2) {
		t.Errorf("got: %v; want: %v", e2, e)
	}
}

func TestCustomDate_MarshalUnmarshalJSON(t *testing.T) {
	type Element struct {
		CreatedAt CustomDate `json:"created_at"`
	}
	e := Element{CustomDate(time.Date(2009, 11, 10, 0, 0, 0, 0, time.UTC))}

	want := `{"created_at":"10•11•2009"}`
	got, err := json.Marshal(e)
	if err != nil {
		t.Error(err)
	}
	if string(got) != want {
		t.Errorf("got: %s; want: %s", got, want)
	}

	var e2 Element
	if err := json.Unmarshal([]byte(want), &e2); err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(e, e2) {
		t.Errorf("got: %v; want: %v", e2, e)
	}
}

func TestCustomDate_MarshalUnmarshalYAML(t *testing.T) {
	type Element struct {
		CreatedAt CustomDate `yaml:"created_at"`
	}
	e := Element{CustomDate(time.Date(2009, 11, 10, 0, 0, 0, 0, time.UTC))}

	want := "created_at: 10•11•2009\n"
	got, err := yaml.Marshal(e)
	if err != nil {
		t.Error(err)
	}
	if string(got) != want {
		t.Errorf("got: %s; want: %s", got, want)
	}

	var e2 Element
	if err := yaml.Unmarshal([]byte(want), &e2); err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(e, e2) {
		t.Errorf("got: %v; want: %v", e2, e)
	}
}
