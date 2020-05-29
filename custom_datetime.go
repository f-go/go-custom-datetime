package custom_datetime

import (
	"encoding/json"
	"encoding/xml"
	"strings"
	"time"
)

// Customisable date type.
type CustomDateTime time.Time

// Format of customised date values.
var CustomDateTimeFormat = time.RFC3339


func (d CustomDateTime) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(time.Time(d).Format(CustomDateTimeFormat), start)
}

func (d *CustomDateTime) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) error {
	var value string
	if err := dec.DecodeElement(&value, &start); err != nil {
		return err
	}

	date, err := time.Parse(CustomDateTimeFormat, value)
	if err != nil {
		return err
	}

	*d = CustomDateTime(date)
	return nil
}

func (d CustomDateTime) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	return xml.Attr{name, time.Time(d).Format(CustomDateTimeFormat)}, nil
}

func (d *CustomDateTime) UnmarshalXMLAttr(attr xml.Attr) error {
	date, err := time.Parse(CustomDateTimeFormat, attr.Value)
	if err != nil {
		return err
	}
	*d = CustomDateTime(date)

	return nil
}

func (d CustomDateTime) MarshalText() ([]byte, error) {
	return []byte(time.Time(d).Format(CustomDateTimeFormat)), nil
}

func (d *CustomDateTime) UnmarshalText(text []byte) error {
	date, err := time.Parse(CustomDateTimeFormat, string(text))
	if err != nil {
		return err
	}
	*d = CustomDateTime(date)

	return nil
}

func (d CustomDateTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(d).Format(CustomDateTimeFormat))
}

func (d *CustomDateTime) UnmarshalJSON(b []byte) error {
	var s string
	s = strings.Trim(string(b), "\"")
	date, err := time.Parse(CustomDateTimeFormat, s)
	if err != nil {
		return err
	}
	*d = CustomDateTime(date)

	return nil
}

func (d CustomDateTime) MarshalYAML() (interface{}, error) {
	return time.Time(d).Format(CustomDateTimeFormat), nil
}

func (d *CustomDateTime) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return nil
	}

	date, err := time.Parse(CustomDateTimeFormat, strings.TrimSpace(s))
	if err != nil {
		return err
	}
	*d = CustomDateTime(date)

	return nil
}
