// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wml

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"time"

	"github.com/qnfnypen/gooxml"
)

type CT_RPrChange struct {
	AuthorAttr string
	DateAttr   *time.Time
	// Annotation Identifier
	IdAttr int64
	RPr    *CT_RPrOriginal
}

func NewCT_RPrChange() *CT_RPrChange {
	ret := &CT_RPrChange{}
	ret.RPr = NewCT_RPrOriginal()
	return ret
}

func (m *CT_RPrChange) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:author"},
		Value: fmt.Sprintf("%v", m.AuthorAttr)})
	if m.DateAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:date"},
			Value: fmt.Sprintf("%v", *m.DateAttr)})
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:id"},
		Value: fmt.Sprintf("%v", m.IdAttr)})
	e.EncodeToken(start)
	serPr := xml.StartElement{Name: xml.Name{Local: "w:rPr"}}
	e.EncodeElement(m.RPr, serPr)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_RPrChange) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.RPr = NewCT_RPrOriginal()
	for _, attr := range start.Attr {
		if attr.Name.Local == "author" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.AuthorAttr = parsed
			continue
		}
		if attr.Name.Local == "date" {
			parsed, err := ParseStdlibTime(attr.Value)
			if err != nil {
				return err
			}
			m.DateAttr = &parsed
			continue
		}
		if attr.Name.Local == "id" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 64)
			if err != nil {
				return err
			}
			m.IdAttr = parsed
			continue
		}
	}
lCT_RPrChange:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "rPr"}:
				if err := d.DecodeElement(m.RPr, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on CT_RPrChange %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_RPrChange
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_RPrChange and its children
func (m *CT_RPrChange) Validate() error {
	return m.ValidateWithPath("CT_RPrChange")
}

// ValidateWithPath validates the CT_RPrChange and its children, prefixing error messages with path
func (m *CT_RPrChange) ValidateWithPath(path string) error {
	if err := m.RPr.ValidateWithPath(path + "/RPr"); err != nil {
		return err
	}
	return nil
}
