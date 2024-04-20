// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package math

import (
	"encoding/xml"
	"fmt"

	"github.com/qnfnypen/gooxml"
)

type CT_M struct {
	MPr *CT_MPr
	Mr  []*CT_MR
}

func NewCT_M() *CT_M {
	ret := &CT_M{}
	return ret
}

func (m *CT_M) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.MPr != nil {
		semPr := xml.StartElement{Name: xml.Name{Local: "m:mPr"}}
		e.EncodeElement(m.MPr, semPr)
	}
	semr := xml.StartElement{Name: xml.Name{Local: "m:mr"}}
	for _, c := range m.Mr {
		e.EncodeElement(c, semr)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_M) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_M:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "mPr"}:
				m.MPr = NewCT_MPr()
				if err := d.DecodeElement(m.MPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "mr"}:
				tmp := NewCT_MR()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Mr = append(m.Mr, tmp)
			default:
				gooxml.Log("skipping unsupported element on CT_M %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_M
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_M and its children
func (m *CT_M) Validate() error {
	return m.ValidateWithPath("CT_M")
}

// ValidateWithPath validates the CT_M and its children, prefixing error messages with path
func (m *CT_M) ValidateWithPath(path string) error {
	if m.MPr != nil {
		if err := m.MPr.ValidateWithPath(path + "/MPr"); err != nil {
			return err
		}
	}
	for i, v := range m.Mr {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Mr[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
