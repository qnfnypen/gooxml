// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package chartDrawing_test

import (
	"encoding/xml"
	"testing"

	"github.com/qnfnypen/gooxml/schema/soo/dml/chartDrawing"
)

func TestEG_AnchorConstructor(t *testing.T) {
	v := chartDrawing.NewEG_Anchor()
	if v == nil {
		t.Errorf("chartDrawing.NewEG_Anchor must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chartDrawing.EG_Anchor should validate: %s", err)
	}
}

func TestEG_AnchorMarshalUnmarshal(t *testing.T) {
	v := chartDrawing.NewEG_Anchor()
	buf, _ := xml.Marshal(v)
	v2 := chartDrawing.NewEG_Anchor()
	xml.Unmarshal(buf, v2)
}
