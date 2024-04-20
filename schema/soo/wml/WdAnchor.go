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
	"strconv"

	"github.com/qnfnypen/gooxml"
	"github.com/qnfnypen/gooxml/schema/soo/dml"
)

type WdAnchor struct {
	WdCT_Anchor
}

func NewWdAnchor() *WdAnchor {
	ret := &WdAnchor{}
	ret.WdCT_Anchor = *NewWdCT_Anchor()
	return ret
}

func (m *WdAnchor) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:a"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:pic"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/picture"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:r"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/relationships"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:w"}, Value: "http://schemas.openxmlformats.org/wordprocessingml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:wp"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xml"}, Value: "http://www.w3.org/XML/1998/namespace"})
	start.Name.Local = "wp:anchor"
	return m.WdCT_Anchor.MarshalXML(e, start)
}

func (m *WdAnchor) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.WdCT_Anchor = *NewWdCT_Anchor()
	for _, attr := range start.Attr {
		if attr.Name.Local == "distT" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.DistTAttr = &pt
			continue
		}
		if attr.Name.Local == "distL" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.DistLAttr = &pt
			continue
		}
		if attr.Name.Local == "simplePos" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.SimplePosAttr = &parsed
			continue
		}
		if attr.Name.Local == "behindDoc" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.BehindDocAttr = parsed
			continue
		}
		if attr.Name.Local == "layoutInCell" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.LayoutInCellAttr = parsed
			continue
		}
		if attr.Name.Local == "hidden" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.HiddenAttr = &parsed
			continue
		}
		if attr.Name.Local == "distB" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.DistBAttr = &pt
			continue
		}
		if attr.Name.Local == "distR" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.DistRAttr = &pt
			continue
		}
		if attr.Name.Local == "relativeHeight" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.RelativeHeightAttr = uint32(parsed)
			continue
		}
		if attr.Name.Local == "locked" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.LockedAttr = parsed
			continue
		}
		if attr.Name.Local == "allowOverlap" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.AllowOverlapAttr = parsed
			continue
		}
	}
lWdAnchor:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", Local: "simplePos"}:
				if err := d.DecodeElement(m.SimplePos, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", Local: "positionH"}:
				if err := d.DecodeElement(m.PositionH, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", Local: "positionV"}:
				if err := d.DecodeElement(m.PositionV, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", Local: "extent"}:
				if err := d.DecodeElement(m.Extent, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", Local: "effectExtent"}:
				m.EffectExtent = NewWdCT_EffectExtent()
				if err := d.DecodeElement(m.EffectExtent, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", Local: "wrapNone"}:
				m.Choice = NewWdEG_WrapTypeChoice()
				if err := d.DecodeElement(&m.Choice.WrapNone, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", Local: "wrapSquare"}:
				m.Choice = NewWdEG_WrapTypeChoice()
				if err := d.DecodeElement(&m.Choice.WrapSquare, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", Local: "wrapTight"}:
				m.Choice = NewWdEG_WrapTypeChoice()
				if err := d.DecodeElement(&m.Choice.WrapTight, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", Local: "wrapThrough"}:
				m.Choice = NewWdEG_WrapTypeChoice()
				if err := d.DecodeElement(&m.Choice.WrapThrough, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", Local: "wrapTopAndBottom"}:
				m.Choice = NewWdEG_WrapTypeChoice()
				if err := d.DecodeElement(&m.Choice.WrapTopAndBottom, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", Local: "docPr"}:
				if err := d.DecodeElement(m.DocPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", Local: "cNvGraphicFramePr"}:
				m.CNvGraphicFramePr = dml.NewCT_NonVisualGraphicFrameProperties()
				if err := d.DecodeElement(m.CNvGraphicFramePr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "graphic"}:
				if err := d.DecodeElement(m.Graphic, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on WdAnchor %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lWdAnchor
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the WdAnchor and its children
func (m *WdAnchor) Validate() error {
	return m.ValidateWithPath("WdAnchor")
}

// ValidateWithPath validates the WdAnchor and its children, prefixing error messages with path
func (m *WdAnchor) ValidateWithPath(path string) error {
	if err := m.WdCT_Anchor.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
