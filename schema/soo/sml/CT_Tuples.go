// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package sml

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"github.com/qnfnypen/gooxml"
)

type CT_Tuples struct {
	// Member Name Count
	CAttr *uint32
	// Tuple
	Tpl []*CT_Tuple
}

func NewCT_Tuples() *CT_Tuples {
	ret := &CT_Tuples{}
	return ret
}

func (m *CT_Tuples) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.CAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "c"},
			Value: fmt.Sprintf("%v", *m.CAttr)})
	}
	e.EncodeToken(start)
	setpl := xml.StartElement{Name: xml.Name{Local: "ma:tpl"}}
	for _, c := range m.Tpl {
		e.EncodeElement(c, setpl)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Tuples) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "c" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.CAttr = &pt
			continue
		}
	}
lCT_Tuples:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "tpl"}:
				tmp := NewCT_Tuple()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Tpl = append(m.Tpl, tmp)
			default:
				gooxml.Log("skipping unsupported element on CT_Tuples %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Tuples
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Tuples and its children
func (m *CT_Tuples) Validate() error {
	return m.ValidateWithPath("CT_Tuples")
}

// ValidateWithPath validates the CT_Tuples and its children, prefixing error messages with path
func (m *CT_Tuples) ValidateWithPath(path string) error {
	for i, v := range m.Tpl {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Tpl[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
