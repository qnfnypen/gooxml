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

	"github.com/qnfnypen/gooxml"
	"github.com/qnfnypen/gooxml/schema/soo/ofc/math"
)

type CT_CustomXmlRun struct {
	// Custom XML Markup Namespace
	UriAttr *string
	// Element name
	ElementAttr string
	// Custom XML Element Properties
	CustomXmlPr *CT_CustomXmlPr
	EG_PContent []*EG_PContent
}

func NewCT_CustomXmlRun() *CT_CustomXmlRun {
	ret := &CT_CustomXmlRun{}
	return ret
}

func (m *CT_CustomXmlRun) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.UriAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:uri"},
			Value: fmt.Sprintf("%v", *m.UriAttr)})
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:element"},
		Value: fmt.Sprintf("%v", m.ElementAttr)})
	e.EncodeToken(start)
	if m.CustomXmlPr != nil {
		secustomXmlPr := xml.StartElement{Name: xml.Name{Local: "w:customXmlPr"}}
		e.EncodeElement(m.CustomXmlPr, secustomXmlPr)
	}
	if m.EG_PContent != nil {
		for _, c := range m.EG_PContent {
			c.MarshalXML(e, xml.StartElement{})
		}
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_CustomXmlRun) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "uri" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.UriAttr = &parsed
			continue
		}
		if attr.Name.Local == "element" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ElementAttr = parsed
			continue
		}
	}
lCT_CustomXmlRun:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "customXmlPr"}:
				m.CustomXmlPr = NewCT_CustomXmlPr()
				if err := d.DecodeElement(m.CustomXmlPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "fldSimple"}:
				tmppcontent := NewEG_PContent()
				tmp := NewCT_SimpleField()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				tmppcontent.FldSimple = append(tmppcontent.FldSimple, tmp)
				m.EG_PContent = append(m.EG_PContent, tmppcontent)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "hyperlink"}:
				tmppcontent := NewEG_PContent()
				tmppcontent.Hyperlink = NewCT_Hyperlink()
				if err := d.DecodeElement(tmppcontent.Hyperlink, &el); err != nil {
					return err
				}
				m.EG_PContent = append(m.EG_PContent, tmppcontent)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "subDoc"}:
				tmppcontent := NewEG_PContent()
				tmppcontent.SubDoc = NewCT_Rel()
				if err := d.DecodeElement(tmppcontent.SubDoc, &el); err != nil {
					return err
				}
				m.EG_PContent = append(m.EG_PContent, tmppcontent)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "customXml"}:
				tmppcontent := NewEG_PContent()
				tmpcontentruncontent := NewEG_ContentRunContent()
				tmpcontentruncontent.CustomXml = NewCT_CustomXmlRun()
				if err := d.DecodeElement(tmpcontentruncontent.CustomXml, &el); err != nil {
					return err
				}
				m.EG_PContent = append(m.EG_PContent, tmppcontent)
				tmppcontent.EG_ContentRunContent = append(tmppcontent.EG_ContentRunContent, tmpcontentruncontent)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "smartTag"}:
				tmppcontent := NewEG_PContent()
				tmpcontentruncontent := NewEG_ContentRunContent()
				tmpcontentruncontent.SmartTag = NewCT_SmartTagRun()
				if err := d.DecodeElement(tmpcontentruncontent.SmartTag, &el); err != nil {
					return err
				}
				m.EG_PContent = append(m.EG_PContent, tmppcontent)
				tmppcontent.EG_ContentRunContent = append(tmppcontent.EG_ContentRunContent, tmpcontentruncontent)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "sdt"}:
				tmppcontent := NewEG_PContent()
				tmpcontentruncontent := NewEG_ContentRunContent()
				tmpcontentruncontent.Sdt = NewCT_SdtRun()
				if err := d.DecodeElement(tmpcontentruncontent.Sdt, &el); err != nil {
					return err
				}
				m.EG_PContent = append(m.EG_PContent, tmppcontent)
				tmppcontent.EG_ContentRunContent = append(tmppcontent.EG_ContentRunContent, tmpcontentruncontent)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "dir"}:
				tmppcontent := NewEG_PContent()
				tmpcontentruncontent := NewEG_ContentRunContent()
				tmpcontentruncontent.Dir = NewCT_DirContentRun()
				if err := d.DecodeElement(tmpcontentruncontent.Dir, &el); err != nil {
					return err
				}
				m.EG_PContent = append(m.EG_PContent, tmppcontent)
				tmppcontent.EG_ContentRunContent = append(tmppcontent.EG_ContentRunContent, tmpcontentruncontent)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "bdo"}:
				tmppcontent := NewEG_PContent()
				tmpcontentruncontent := NewEG_ContentRunContent()
				tmpcontentruncontent.Bdo = NewCT_BdoContentRun()
				if err := d.DecodeElement(tmpcontentruncontent.Bdo, &el); err != nil {
					return err
				}
				m.EG_PContent = append(m.EG_PContent, tmppcontent)
				tmppcontent.EG_ContentRunContent = append(tmppcontent.EG_ContentRunContent, tmpcontentruncontent)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "r"}:
				tmppcontent := NewEG_PContent()
				tmpcontentruncontent := NewEG_ContentRunContent()
				tmpcontentruncontent.R = NewCT_R()
				if err := d.DecodeElement(tmpcontentruncontent.R, &el); err != nil {
					return err
				}
				m.EG_PContent = append(m.EG_PContent, tmppcontent)
				tmppcontent.EG_ContentRunContent = append(tmppcontent.EG_ContentRunContent, tmpcontentruncontent)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "proofErr"}:
				tmppcontent := NewEG_PContent()
				tmpcontentruncontent := NewEG_ContentRunContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprunlevelelts.ProofErr = NewCT_ProofErr()
				if err := d.DecodeElement(tmprunlevelelts.ProofErr, &el); err != nil {
					return err
				}
				m.EG_PContent = append(m.EG_PContent, tmppcontent)
				tmppcontent.EG_ContentRunContent = append(tmppcontent.EG_ContentRunContent, tmpcontentruncontent)
				tmpcontentruncontent.EG_RunLevelElts = append(tmpcontentruncontent.EG_RunLevelElts, tmprunlevelelts)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "permStart"}:
				tmppcontent := NewEG_PContent()
				tmpcontentruncontent := NewEG_ContentRunContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprunlevelelts.PermStart = NewCT_PermStart()
				if err := d.DecodeElement(tmprunlevelelts.PermStart, &el); err != nil {
					return err
				}
				m.EG_PContent = append(m.EG_PContent, tmppcontent)
				tmppcontent.EG_ContentRunContent = append(tmppcontent.EG_ContentRunContent, tmpcontentruncontent)
				tmpcontentruncontent.EG_RunLevelElts = append(tmpcontentruncontent.EG_RunLevelElts, tmprunlevelelts)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "permEnd"}:
				tmppcontent := NewEG_PContent()
				tmpcontentruncontent := NewEG_ContentRunContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprunlevelelts.PermEnd = NewCT_Perm()
				if err := d.DecodeElement(tmprunlevelelts.PermEnd, &el); err != nil {
					return err
				}
				m.EG_PContent = append(m.EG_PContent, tmppcontent)
				tmppcontent.EG_ContentRunContent = append(tmppcontent.EG_ContentRunContent, tmpcontentruncontent)
				tmpcontentruncontent.EG_RunLevelElts = append(tmpcontentruncontent.EG_RunLevelElts, tmprunlevelelts)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "ins"}:
				tmppcontent := NewEG_PContent()
				tmpcontentruncontent := NewEG_ContentRunContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprunlevelelts.Ins = NewCT_RunTrackChange()
				if err := d.DecodeElement(tmprunlevelelts.Ins, &el); err != nil {
					return err
				}
				m.EG_PContent = append(m.EG_PContent, tmppcontent)
				tmppcontent.EG_ContentRunContent = append(tmppcontent.EG_ContentRunContent, tmpcontentruncontent)
				tmpcontentruncontent.EG_RunLevelElts = append(tmpcontentruncontent.EG_RunLevelElts, tmprunlevelelts)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "del"}:
				tmppcontent := NewEG_PContent()
				tmpcontentruncontent := NewEG_ContentRunContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprunlevelelts.Del = NewCT_RunTrackChange()
				if err := d.DecodeElement(tmprunlevelelts.Del, &el); err != nil {
					return err
				}
				m.EG_PContent = append(m.EG_PContent, tmppcontent)
				tmppcontent.EG_ContentRunContent = append(tmppcontent.EG_ContentRunContent, tmpcontentruncontent)
				tmpcontentruncontent.EG_RunLevelElts = append(tmpcontentruncontent.EG_RunLevelElts, tmprunlevelelts)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "moveFrom"}:
				tmppcontent := NewEG_PContent()
				tmpcontentruncontent := NewEG_ContentRunContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprunlevelelts.MoveFrom = NewCT_RunTrackChange()
				if err := d.DecodeElement(tmprunlevelelts.MoveFrom, &el); err != nil {
					return err
				}
				m.EG_PContent = append(m.EG_PContent, tmppcontent)
				tmppcontent.EG_ContentRunContent = append(tmppcontent.EG_ContentRunContent, tmpcontentruncontent)
				tmpcontentruncontent.EG_RunLevelElts = append(tmpcontentruncontent.EG_RunLevelElts, tmprunlevelelts)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "moveTo"}:
				tmppcontent := NewEG_PContent()
				tmpcontentruncontent := NewEG_ContentRunContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprunlevelelts.MoveTo = NewCT_RunTrackChange()
				if err := d.DecodeElement(tmprunlevelelts.MoveTo, &el); err != nil {
					return err
				}
				m.EG_PContent = append(m.EG_PContent, tmppcontent)
				tmppcontent.EG_ContentRunContent = append(tmppcontent.EG_ContentRunContent, tmpcontentruncontent)
				tmpcontentruncontent.EG_RunLevelElts = append(tmpcontentruncontent.EG_RunLevelElts, tmprunlevelelts)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "bookmarkStart"}:
				tmppcontent := NewEG_PContent()
				tmpcontentruncontent := NewEG_ContentRunContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.BookmarkStart = NewCT_Bookmark()
				if err := d.DecodeElement(tmprangemarkupelements.BookmarkStart, &el); err != nil {
					return err
				}
				m.EG_PContent = append(m.EG_PContent, tmppcontent)
				tmppcontent.EG_ContentRunContent = append(tmppcontent.EG_ContentRunContent, tmpcontentruncontent)
				tmpcontentruncontent.EG_RunLevelElts = append(tmpcontentruncontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "bookmarkEnd"}:
				tmppcontent := NewEG_PContent()
				tmpcontentruncontent := NewEG_ContentRunContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.BookmarkEnd = NewCT_MarkupRange()
				if err := d.DecodeElement(tmprangemarkupelements.BookmarkEnd, &el); err != nil {
					return err
				}
				m.EG_PContent = append(m.EG_PContent, tmppcontent)
				tmppcontent.EG_ContentRunContent = append(tmppcontent.EG_ContentRunContent, tmpcontentruncontent)
				tmpcontentruncontent.EG_RunLevelElts = append(tmpcontentruncontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "moveFromRangeStart"}:
				tmppcontent := NewEG_PContent()
				tmpcontentruncontent := NewEG_ContentRunContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.MoveFromRangeStart = NewCT_MoveBookmark()
				if err := d.DecodeElement(tmprangemarkupelements.MoveFromRangeStart, &el); err != nil {
					return err
				}
				m.EG_PContent = append(m.EG_PContent, tmppcontent)
				tmppcontent.EG_ContentRunContent = append(tmppcontent.EG_ContentRunContent, tmpcontentruncontent)
				tmpcontentruncontent.EG_RunLevelElts = append(tmpcontentruncontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "moveFromRangeEnd"}:
				tmppcontent := NewEG_PContent()
				tmpcontentruncontent := NewEG_ContentRunContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.MoveFromRangeEnd = NewCT_MarkupRange()
				if err := d.DecodeElement(tmprangemarkupelements.MoveFromRangeEnd, &el); err != nil {
					return err
				}
				m.EG_PContent = append(m.EG_PContent, tmppcontent)
				tmppcontent.EG_ContentRunContent = append(tmppcontent.EG_ContentRunContent, tmpcontentruncontent)
				tmpcontentruncontent.EG_RunLevelElts = append(tmpcontentruncontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "moveToRangeStart"}:
				tmppcontent := NewEG_PContent()
				tmpcontentruncontent := NewEG_ContentRunContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.MoveToRangeStart = NewCT_MoveBookmark()
				if err := d.DecodeElement(tmprangemarkupelements.MoveToRangeStart, &el); err != nil {
					return err
				}
				m.EG_PContent = append(m.EG_PContent, tmppcontent)
				tmppcontent.EG_ContentRunContent = append(tmppcontent.EG_ContentRunContent, tmpcontentruncontent)
				tmpcontentruncontent.EG_RunLevelElts = append(tmpcontentruncontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "moveToRangeEnd"}:
				tmppcontent := NewEG_PContent()
				tmpcontentruncontent := NewEG_ContentRunContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.MoveToRangeEnd = NewCT_MarkupRange()
				if err := d.DecodeElement(tmprangemarkupelements.MoveToRangeEnd, &el); err != nil {
					return err
				}
				m.EG_PContent = append(m.EG_PContent, tmppcontent)
				tmppcontent.EG_ContentRunContent = append(tmppcontent.EG_ContentRunContent, tmpcontentruncontent)
				tmpcontentruncontent.EG_RunLevelElts = append(tmpcontentruncontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "commentRangeStart"}:
				tmppcontent := NewEG_PContent()
				tmpcontentruncontent := NewEG_ContentRunContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CommentRangeStart = NewCT_MarkupRange()
				if err := d.DecodeElement(tmprangemarkupelements.CommentRangeStart, &el); err != nil {
					return err
				}
				m.EG_PContent = append(m.EG_PContent, tmppcontent)
				tmppcontent.EG_ContentRunContent = append(tmppcontent.EG_ContentRunContent, tmpcontentruncontent)
				tmpcontentruncontent.EG_RunLevelElts = append(tmpcontentruncontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "commentRangeEnd"}:
				tmppcontent := NewEG_PContent()
				tmpcontentruncontent := NewEG_ContentRunContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CommentRangeEnd = NewCT_MarkupRange()
				if err := d.DecodeElement(tmprangemarkupelements.CommentRangeEnd, &el); err != nil {
					return err
				}
				m.EG_PContent = append(m.EG_PContent, tmppcontent)
				tmppcontent.EG_ContentRunContent = append(tmppcontent.EG_ContentRunContent, tmpcontentruncontent)
				tmpcontentruncontent.EG_RunLevelElts = append(tmpcontentruncontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "customXmlInsRangeStart"}:
				tmppcontent := NewEG_PContent()
				tmpcontentruncontent := NewEG_ContentRunContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CustomXmlInsRangeStart = NewCT_TrackChange()
				if err := d.DecodeElement(tmprangemarkupelements.CustomXmlInsRangeStart, &el); err != nil {
					return err
				}
				m.EG_PContent = append(m.EG_PContent, tmppcontent)
				tmppcontent.EG_ContentRunContent = append(tmppcontent.EG_ContentRunContent, tmpcontentruncontent)
				tmpcontentruncontent.EG_RunLevelElts = append(tmpcontentruncontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "customXmlInsRangeEnd"}:
				tmppcontent := NewEG_PContent()
				tmpcontentruncontent := NewEG_ContentRunContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CustomXmlInsRangeEnd = NewCT_Markup()
				if err := d.DecodeElement(tmprangemarkupelements.CustomXmlInsRangeEnd, &el); err != nil {
					return err
				}
				m.EG_PContent = append(m.EG_PContent, tmppcontent)
				tmppcontent.EG_ContentRunContent = append(tmppcontent.EG_ContentRunContent, tmpcontentruncontent)
				tmpcontentruncontent.EG_RunLevelElts = append(tmpcontentruncontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "customXmlDelRangeStart"}:
				tmppcontent := NewEG_PContent()
				tmpcontentruncontent := NewEG_ContentRunContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CustomXmlDelRangeStart = NewCT_TrackChange()
				if err := d.DecodeElement(tmprangemarkupelements.CustomXmlDelRangeStart, &el); err != nil {
					return err
				}
				m.EG_PContent = append(m.EG_PContent, tmppcontent)
				tmppcontent.EG_ContentRunContent = append(tmppcontent.EG_ContentRunContent, tmpcontentruncontent)
				tmpcontentruncontent.EG_RunLevelElts = append(tmpcontentruncontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "customXmlDelRangeEnd"}:
				tmppcontent := NewEG_PContent()
				tmpcontentruncontent := NewEG_ContentRunContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CustomXmlDelRangeEnd = NewCT_Markup()
				if err := d.DecodeElement(tmprangemarkupelements.CustomXmlDelRangeEnd, &el); err != nil {
					return err
				}
				m.EG_PContent = append(m.EG_PContent, tmppcontent)
				tmppcontent.EG_ContentRunContent = append(tmppcontent.EG_ContentRunContent, tmpcontentruncontent)
				tmpcontentruncontent.EG_RunLevelElts = append(tmpcontentruncontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "customXmlMoveFromRangeStart"}:
				tmppcontent := NewEG_PContent()
				tmpcontentruncontent := NewEG_ContentRunContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CustomXmlMoveFromRangeStart = NewCT_TrackChange()
				if err := d.DecodeElement(tmprangemarkupelements.CustomXmlMoveFromRangeStart, &el); err != nil {
					return err
				}
				m.EG_PContent = append(m.EG_PContent, tmppcontent)
				tmppcontent.EG_ContentRunContent = append(tmppcontent.EG_ContentRunContent, tmpcontentruncontent)
				tmpcontentruncontent.EG_RunLevelElts = append(tmpcontentruncontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "customXmlMoveFromRangeEnd"}:
				tmppcontent := NewEG_PContent()
				tmpcontentruncontent := NewEG_ContentRunContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CustomXmlMoveFromRangeEnd = NewCT_Markup()
				if err := d.DecodeElement(tmprangemarkupelements.CustomXmlMoveFromRangeEnd, &el); err != nil {
					return err
				}
				m.EG_PContent = append(m.EG_PContent, tmppcontent)
				tmppcontent.EG_ContentRunContent = append(tmppcontent.EG_ContentRunContent, tmpcontentruncontent)
				tmpcontentruncontent.EG_RunLevelElts = append(tmpcontentruncontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "customXmlMoveToRangeStart"}:
				tmppcontent := NewEG_PContent()
				tmpcontentruncontent := NewEG_ContentRunContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CustomXmlMoveToRangeStart = NewCT_TrackChange()
				if err := d.DecodeElement(tmprangemarkupelements.CustomXmlMoveToRangeStart, &el); err != nil {
					return err
				}
				m.EG_PContent = append(m.EG_PContent, tmppcontent)
				tmppcontent.EG_ContentRunContent = append(tmppcontent.EG_ContentRunContent, tmpcontentruncontent)
				tmpcontentruncontent.EG_RunLevelElts = append(tmpcontentruncontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "customXmlMoveToRangeEnd"}:
				tmppcontent := NewEG_PContent()
				tmpcontentruncontent := NewEG_ContentRunContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CustomXmlMoveToRangeEnd = NewCT_Markup()
				if err := d.DecodeElement(tmprangemarkupelements.CustomXmlMoveToRangeEnd, &el); err != nil {
					return err
				}
				m.EG_PContent = append(m.EG_PContent, tmppcontent)
				tmppcontent.EG_ContentRunContent = append(tmppcontent.EG_ContentRunContent, tmpcontentruncontent)
				tmpcontentruncontent.EG_RunLevelElts = append(tmpcontentruncontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "oMathPara"}:
				tmppcontent := NewEG_PContent()
				tmpcontentruncontent := NewEG_ContentRunContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmpmathcontent := NewEG_MathContent()
				tmpmathcontent.OMathPara = math.NewOMathPara()
				if err := d.DecodeElement(tmpmathcontent.OMathPara, &el); err != nil {
					return err
				}
				m.EG_PContent = append(m.EG_PContent, tmppcontent)
				tmppcontent.EG_ContentRunContent = append(tmppcontent.EG_ContentRunContent, tmpcontentruncontent)
				tmpcontentruncontent.EG_RunLevelElts = append(tmpcontentruncontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_MathContent = append(tmprunlevelelts.EG_MathContent, tmpmathcontent)
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "oMath"}:
				tmppcontent := NewEG_PContent()
				tmpcontentruncontent := NewEG_ContentRunContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmpmathcontent := NewEG_MathContent()
				tmpmathcontent.OMath = math.NewOMath()
				if err := d.DecodeElement(tmpmathcontent.OMath, &el); err != nil {
					return err
				}
				m.EG_PContent = append(m.EG_PContent, tmppcontent)
				tmppcontent.EG_ContentRunContent = append(tmppcontent.EG_ContentRunContent, tmpcontentruncontent)
				tmpcontentruncontent.EG_RunLevelElts = append(tmpcontentruncontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_MathContent = append(tmprunlevelelts.EG_MathContent, tmpmathcontent)
			default:
				gooxml.Log("skipping unsupported element on CT_CustomXmlRun %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_CustomXmlRun
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_CustomXmlRun and its children
func (m *CT_CustomXmlRun) Validate() error {
	return m.ValidateWithPath("CT_CustomXmlRun")
}

// ValidateWithPath validates the CT_CustomXmlRun and its children, prefixing error messages with path
func (m *CT_CustomXmlRun) ValidateWithPath(path string) error {
	if m.CustomXmlPr != nil {
		if err := m.CustomXmlPr.ValidateWithPath(path + "/CustomXmlPr"); err != nil {
			return err
		}
	}
	for i, v := range m.EG_PContent {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/EG_PContent[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
