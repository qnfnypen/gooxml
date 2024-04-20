// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package extended_properties_test

import (
	"encoding/xml"
	"testing"

	"github.com/qnfnypen/gooxml/schema/soo/ofc/extended_properties"
)

func TestCT_PropertiesConstructor(t *testing.T) {
	v := extended_properties.NewCT_Properties()
	if v == nil {
		t.Errorf("extended_properties.NewCT_Properties must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed extended_properties.CT_Properties should validate: %s", err)
	}
}

func TestCT_PropertiesMarshalUnmarshal(t *testing.T) {
	v := extended_properties.NewCT_Properties()
	buf, _ := xml.Marshal(v)
	v2 := extended_properties.NewCT_Properties()
	xml.Unmarshal(buf, v2)
}
