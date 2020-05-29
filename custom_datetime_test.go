// Copyright 2020 Frank Göldner. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

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
	CustomDateTimeFormat = "02•01•2006 15:04"
}

func TestCustomDateTime_MarshalUnmarshalXML(t *testing.T) {
	type Element struct {
		CreatedAt CustomDateTime `xml:""`
	}
	dt := CustomDateTime(time.Date(2009, 11, 10, 13, 37, 0, 0, time.UTC))
	// or just use:
	// dt, _ := NewCustomDateTime("10•11•2009 13:37")
	e := Element{dt}

	want := `<Element><CreatedAt>10•11•2009 13:37</CreatedAt></Element>`
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

func TestCustomDateTime_MarshalUnmarshalXMLAttr(t *testing.T) {
	type Element struct {
		Created CustomDateTime `xml:"created_at,attr"`
	}
	dt, _ := NewCustomDateTime("10•11•2009 13:37")
	e := Element{dt}

	want := `<Element created_at="10•11•2009 13:37"></Element>`
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

func TestCustomDateTime_MarshalUnmarshalText(t *testing.T) {
	type Element struct {
		CreatedAt CustomDateTime `xml:",chardata"`
	}
	dt, _ := NewCustomDateTime("10•11•2009 13:37")
	e := Element{dt}

	want := `<Element>10•11•2009 13:37</Element>`
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

func TestCustomDateTime_MarshalUnmarshalTextCDATA(t *testing.T) {
	type Element struct {
		CreatedAt CustomDateTime `xml:",cdata"`
	}
	dt, _ := NewCustomDateTime("10•11•2009 13:37")
	e := Element{dt}

	want := `<Element><![CDATA[10•11•2009 13:37]]></Element>`
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

func TestCustomDateTime_MarshalUnmarshalJSON(t *testing.T) {
	type Element struct {
		CreatedAt CustomDateTime `json:"created_at"`
	}
	dt, _ := NewCustomDateTime("10•11•2009 13:37")
	e := Element{dt}

	want := `{"created_at":"10•11•2009 13:37"}`
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

func TestCustomDateTime_MarshalUnmarshalYAML(t *testing.T) {
	type Element struct {
		CreatedAt CustomDateTime `yaml:"created_at"`
	}
	dt, _ := NewCustomDateTime("10•11•2009 13:37")
	e := Element{dt}

	want := "created_at: 10•11•2009 13:37\n"
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
