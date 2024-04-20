// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package vml_test

import (
	"encoding/xml"
	"testing"

	"github.com/qnfnypen/gooxml/schema/urn/schemas_microsoft_com/vml"
)

func TestOfcExtrusionConstructor(t *testing.T) {
	v := vml.NewOfcExtrusion()
	if v == nil {
		t.Errorf("vml.NewOfcExtrusion must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.OfcExtrusion should validate: %s", err)
	}
}

func TestOfcExtrusionMarshalUnmarshal(t *testing.T) {
	v := vml.NewOfcExtrusion()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewOfcExtrusion()
	xml.Unmarshal(buf, v2)
}
