// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package pml_test

import (
	"encoding/xml"
	"testing"

	"github.com/qnfnypen/gooxml/schema/soo/pml"
)

func TestCT_SlideMasterConstructor(t *testing.T) {
	v := pml.NewCT_SlideMaster()
	if v == nil {
		t.Errorf("pml.NewCT_SlideMaster must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_SlideMaster should validate: %s", err)
	}
}

func TestCT_SlideMasterMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_SlideMaster()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_SlideMaster()
	xml.Unmarshal(buf, v2)
}
