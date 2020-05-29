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
	CustomTimeFormat = "15:04"
}

func TestCustomTime_MarshalUnmarshalXML(t *testing.T) {
	type Element struct {
		CreatedAt CustomTime `xml:""`
	}
	e := Element{CustomTime(time.Date(0, 1, 1, 13, 37, 0, 0, time.UTC))}

	want := `<Element><CreatedAt>13:37</CreatedAt></Element>`
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

func TestCustomTime_MarshalUnmarshalXMLAttr(t *testing.T) {
	type Element struct {
		Created CustomTime `xml:"created_at,attr"`
	}
	e := Element{CustomTime(time.Date(0, 1, 1, 13, 37, 0, 0, time.UTC))}

	want := `<Element created_at="13:37"></Element>`
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

func TestCustomTime_MarshalUnmarshalText(t *testing.T) {
	type Element struct {
		CreatedAt CustomTime `xml:",chardata"`
	}
	e := Element{CustomTime(time.Date(0, 1, 1, 13, 37, 0, 0, time.UTC))}

	want := `<Element>13:37</Element>`
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

func TestCustomTime_MarshalUnmarshalTextCDATA(t *testing.T) {
	type Element struct {
		CreatedAt CustomTime `xml:",cdata"`
	}
	e := Element{CustomTime(time.Date(0, 1, 1, 13, 37, 0, 0, time.UTC))}

	want := `<Element><![CDATA[13:37]]></Element>`
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

func TestCustomTime_MarshalUnmarshalJSON(t *testing.T) {
	type Element struct {
		CreatedAt CustomTime `json:"created_at"`
	}
	e := Element{CustomTime(time.Date(0, 1, 1, 13, 37, 0, 0, time.UTC))}

	want := `{"created_at":"13:37"}`
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

func TestCustomTime_MarshalUnmarshalYAML(t *testing.T) {
	type Element struct {
		CreatedAt CustomTime `yaml:"created_at"`
	}
	e := Element{CustomTime(time.Date(0, 1, 1, 13, 37, 0, 0, time.UTC))}

	want := "created_at: \"13:37\"\n"
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
