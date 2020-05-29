// Copyright 2020 Frank Göldner. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package custom_datetime

import (
	"encoding/json"
	"encoding/xml"
	"strings"
	"time"
)

// Customisable time date type.
type CustomTime time.Time

// Format of customised time values.
var CustomTimeFormat = "15:04:05"


func (d CustomTime) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(time.Time(d).Format(CustomTimeFormat), start)
}

func (d *CustomTime) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) error {
	var value string
	if err := dec.DecodeElement(&value, &start); err != nil {
		return err
	}

	date, err := time.Parse(CustomTimeFormat, value)
	if err != nil {
		return err
	}

	*d = CustomTime(date)
	return nil
}

func (d CustomTime) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	return xml.Attr{name, time.Time(d).Format(CustomTimeFormat)}, nil
}

func (d *CustomTime) UnmarshalXMLAttr(attr xml.Attr) error {
	date, err := time.Parse(CustomTimeFormat, attr.Value)
	if err != nil {
		return err
	}
	*d = CustomTime(date)

	return nil
}

func (d CustomTime) MarshalText() ([]byte, error) {
	return []byte(time.Time(d).Format(CustomTimeFormat)), nil
}

func (d *CustomTime) UnmarshalText(text []byte) error {
	date, err := time.Parse(CustomTimeFormat, string(text))
	if err != nil {
		return err
	}
	*d = CustomTime(date)

	return nil
}

func (d CustomTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(d).Format(CustomTimeFormat))
}

func (d *CustomTime) UnmarshalJSON(b []byte) error {
	var s string
	s = strings.Trim(string(b), "\"")
	date, err := time.Parse(CustomTimeFormat, s)
	if err != nil {
		return err
	}
	*d = CustomTime(date)

	return nil
}

func (d CustomTime) MarshalYAML() (interface{}, error) {
	return time.Time(d).Format(CustomTimeFormat), nil
}

func (d *CustomTime) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return nil
	}

	date, err := time.Parse(CustomTimeFormat, strings.TrimSpace(s))
	if err != nil {
		return err
	}
	*d = CustomTime(date)

	return nil
}
