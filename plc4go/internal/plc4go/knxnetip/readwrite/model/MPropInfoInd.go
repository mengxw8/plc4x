//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//

package model

import (
	"encoding/xml"
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"io"
)

// Code generated by build-utils. DO NOT EDIT.

// The data-structure of this message
type MPropInfoInd struct {
	Parent *CEMI
}

// The corresponding interface
type IMPropInfoInd interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *MPropInfoInd) MessageCode() uint8 {
	return 0xF7
}

func (m *MPropInfoInd) InitializeParent(parent *CEMI) {
}

func NewMPropInfoInd() *CEMI {
	child := &MPropInfoInd{
		Parent: NewCEMI(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastMPropInfoInd(structType interface{}) *MPropInfoInd {
	castFunc := func(typ interface{}) *MPropInfoInd {
		if casted, ok := typ.(MPropInfoInd); ok {
			return &casted
		}
		if casted, ok := typ.(*MPropInfoInd); ok {
			return casted
		}
		if casted, ok := typ.(CEMI); ok {
			return CastMPropInfoInd(casted.Child)
		}
		if casted, ok := typ.(*CEMI); ok {
			return CastMPropInfoInd(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *MPropInfoInd) GetTypeName() string {
	return "MPropInfoInd"
}

func (m *MPropInfoInd) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *MPropInfoInd) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	return lengthInBits
}

func (m *MPropInfoInd) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func MPropInfoIndParse(io utils.ReadBuffer) (*CEMI, error) {

	// Create a partially initialized instance
	_child := &MPropInfoInd{
		Parent: &CEMI{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *MPropInfoInd) Serialize(io utils.WriteBuffer) error {
	ser := func() error {

		return nil
	}
	return m.Parent.SerializeParent(io, m, ser)
}

func (m *MPropInfoInd) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var token xml.Token
	var err error
	foundContent := false
	token = start
	for {
		switch token.(type) {
		case xml.StartElement:
			foundContent = true
			tok := token.(xml.StartElement)
			switch tok.Name.Local {
			}
		}
		token, err = d.Token()
		if err != nil {
			if err == io.EOF && foundContent {
				return nil
			}
			return err
		}
	}
}

func (m *MPropInfoInd) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return nil
}

func (m MPropInfoInd) String() string {
	return string(m.Box("", 120))
}

func (m MPropInfoInd) Box(name string, width int) utils.AsciiBox {
	boxName := "MPropInfoInd"
	if name != "" {
		boxName += "/" + name
	}
	childBoxer := func() []utils.AsciiBox {
		boxes := make([]utils.AsciiBox, 0)
		return boxes
	}
	return m.Parent.BoxParent(boxName, width, childBoxer)
}
