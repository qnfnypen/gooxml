// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package diagram_test

import (
	"encoding/xml"
	"testing"

	"github.com/qnfnypen/gooxml/schema/soo/dml/diagram"
)

func TestCT_TextPropsConstructor(t *testing.T) {
	v := diagram.NewCT_TextProps()
	if v == nil {
		t.Errorf("diagram.NewCT_TextProps must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed diagram.CT_TextProps should validate: %s", err)
	}
}

func TestCT_TextPropsMarshalUnmarshal(t *testing.T) {
	v := diagram.NewCT_TextProps()
	buf, _ := xml.Marshal(v)
	v2 := diagram.NewCT_TextProps()
	xml.Unmarshal(buf, v2)
}
