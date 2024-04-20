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

	"github.com/qnfnypen/gooxml"
)

type Metadata struct {
	CT_Metadata
}

func NewMetadata() *Metadata {
	ret := &Metadata{}
	ret.CT_Metadata = *NewCT_Metadata()
	return ret
}

func (m *Metadata) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "http://schemas.openxmlformats.org/spreadsheetml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:ma"}, Value: "http://schemas.openxmlformats.org/spreadsheetml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:r"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/relationships"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:s"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/sharedTypes"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xdr"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xml"}, Value: "http://www.w3.org/XML/1998/namespace"})
	start.Name.Local = "ma:metadata"
	return m.CT_Metadata.MarshalXML(e, start)
}

func (m *Metadata) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_Metadata = *NewCT_Metadata()
lMetadata:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "metadataTypes"}:
				m.MetadataTypes = NewCT_MetadataTypes()
				if err := d.DecodeElement(m.MetadataTypes, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "metadataStrings"}:
				m.MetadataStrings = NewCT_MetadataStrings()
				if err := d.DecodeElement(m.MetadataStrings, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "mdxMetadata"}:
				m.MdxMetadata = NewCT_MdxMetadata()
				if err := d.DecodeElement(m.MdxMetadata, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "futureMetadata"}:
				tmp := NewCT_FutureMetadata()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.FutureMetadata = append(m.FutureMetadata, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "cellMetadata"}:
				m.CellMetadata = NewCT_MetadataBlocks()
				if err := d.DecodeElement(m.CellMetadata, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "valueMetadata"}:
				m.ValueMetadata = NewCT_MetadataBlocks()
				if err := d.DecodeElement(m.ValueMetadata, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "extLst"}:
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on Metadata %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lMetadata
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the Metadata and its children
func (m *Metadata) Validate() error {
	return m.ValidateWithPath("Metadata")
}

// ValidateWithPath validates the Metadata and its children, prefixing error messages with path
func (m *Metadata) ValidateWithPath(path string) error {
	if err := m.CT_Metadata.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
