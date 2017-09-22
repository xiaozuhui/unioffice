// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package presentationml

import (
	"encoding/xml"
	"log"
)

type CT_SlideViewProperties struct {
	CSldViewPr *CT_CommonSlideViewProperties
	ExtLst     *CT_ExtensionList
}

func NewCT_SlideViewProperties() *CT_SlideViewProperties {
	ret := &CT_SlideViewProperties{}
	ret.CSldViewPr = NewCT_CommonSlideViewProperties()
	return ret
}

func (m *CT_SlideViewProperties) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	secSldViewPr := xml.StartElement{Name: xml.Name{Local: "p:cSldViewPr"}}
	e.EncodeElement(m.CSldViewPr, secSldViewPr)
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "p:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_SlideViewProperties) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CSldViewPr = NewCT_CommonSlideViewProperties()
lCT_SlideViewProperties:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "cSldViewPr"}:
				if err := d.DecodeElement(m.CSldViewPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "extLst"}:
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_SlideViewProperties %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_SlideViewProperties
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_SlideViewProperties and its children
func (m *CT_SlideViewProperties) Validate() error {
	return m.ValidateWithPath("CT_SlideViewProperties")
}

// ValidateWithPath validates the CT_SlideViewProperties and its children, prefixing error messages with path
func (m *CT_SlideViewProperties) ValidateWithPath(path string) error {
	if err := m.CSldViewPr.ValidateWithPath(path + "/CSldViewPr"); err != nil {
		return err
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}