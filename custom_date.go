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

// Customisable date type.
type CustomDate time.Time

// Format of customised date values.
var CustomDateFormat = "2006-01-02"


func (d CustomDate) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(time.Time(d).Format(CustomDateFormat), start)
}

func (d *CustomDate) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) error {
	var value string
	if err := dec.DecodeElement(&value, &start); err != nil {
		return err
	}

	date, err := time.Parse(CustomDateFormat, value)
	if err != nil {
		return err
	}

	*d = CustomDate(date)
	return nil
}

func (d CustomDate) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	return xml.Attr{name, time.Time(d).Format(CustomDateFormat)}, nil
}

func (d *CustomDate) UnmarshalXMLAttr(attr xml.Attr) error {
	date, err := time.Parse(CustomDateFormat, attr.Value)
	if err != nil {
		return err
	}
	*d = CustomDate(date)

	return nil
}

func (d CustomDate) MarshalText() ([]byte, error) {
	return []byte(time.Time(d).Format(CustomDateFormat)), nil
}

func (d *CustomDate) UnmarshalText(text []byte) error {
	date, err := time.Parse(CustomDateFormat, string(text))
	if err != nil {
		return err
	}
	*d = CustomDate(date)

	return nil
}

func (d CustomDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(d).Format(CustomDateFormat))
}

func (d *CustomDate) UnmarshalJSON(b []byte) error {
	var s string
	s = strings.Trim(string(b), "\"")
	date, err := time.Parse(CustomDateFormat, s)
	if err != nil {
		return err
	}
	*d = CustomDate(date)

	return nil
}

func (d CustomDate) MarshalYAML() (interface{}, error) {
	return time.Time(d).Format(CustomDateFormat), nil
}

func (d *CustomDate) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return nil
	}

	date, err := time.Parse(CustomDateFormat, strings.TrimSpace(s))
	if err != nil {
		return err
	}
	*d = CustomDate(date)

	return nil
}
