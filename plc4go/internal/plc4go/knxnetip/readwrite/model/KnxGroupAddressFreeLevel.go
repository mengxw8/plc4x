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
	"github.com/pkg/errors"
	"io"
)

// Code generated by build-utils. DO NOT EDIT.

// The data-structure of this message
type KnxGroupAddressFreeLevel struct {
	SubGroup uint16
	Parent   *KnxGroupAddress
}

// The corresponding interface
type IKnxGroupAddressFreeLevel interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *KnxGroupAddressFreeLevel) NumLevels() uint8 {
	return 1
}

func (m *KnxGroupAddressFreeLevel) InitializeParent(parent *KnxGroupAddress) {
}

func NewKnxGroupAddressFreeLevel(subGroup uint16) *KnxGroupAddress {
	child := &KnxGroupAddressFreeLevel{
		SubGroup: subGroup,
		Parent:   NewKnxGroupAddress(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastKnxGroupAddressFreeLevel(structType interface{}) *KnxGroupAddressFreeLevel {
	castFunc := func(typ interface{}) *KnxGroupAddressFreeLevel {
		if casted, ok := typ.(KnxGroupAddressFreeLevel); ok {
			return &casted
		}
		if casted, ok := typ.(*KnxGroupAddressFreeLevel); ok {
			return casted
		}
		if casted, ok := typ.(KnxGroupAddress); ok {
			return CastKnxGroupAddressFreeLevel(casted.Child)
		}
		if casted, ok := typ.(*KnxGroupAddress); ok {
			return CastKnxGroupAddressFreeLevel(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *KnxGroupAddressFreeLevel) GetTypeName() string {
	return "KnxGroupAddressFreeLevel"
}

func (m *KnxGroupAddressFreeLevel) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *KnxGroupAddressFreeLevel) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Simple field (subGroup)
	lengthInBits += 16

	return lengthInBits
}

func (m *KnxGroupAddressFreeLevel) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func KnxGroupAddressFreeLevelParse(io utils.ReadBuffer) (*KnxGroupAddress, error) {

	// Simple Field (subGroup)
	subGroup, _subGroupErr := io.ReadUint16(16)
	if _subGroupErr != nil {
		return nil, errors.Wrap(_subGroupErr, "Error parsing 'subGroup' field")
	}

	// Create a partially initialized instance
	_child := &KnxGroupAddressFreeLevel{
		SubGroup: subGroup,
		Parent:   &KnxGroupAddress{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *KnxGroupAddressFreeLevel) Serialize(io utils.WriteBuffer) error {
	ser := func() error {

		// Simple Field (subGroup)
		subGroup := uint16(m.SubGroup)
		_subGroupErr := io.WriteUint16(16, (subGroup))
		if _subGroupErr != nil {
			return errors.Wrap(_subGroupErr, "Error serializing 'subGroup' field")
		}

		return nil
	}
	return m.Parent.SerializeParent(io, m, ser)
}

func (m *KnxGroupAddressFreeLevel) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "subGroup":
				var data uint16
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.SubGroup = data
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

func (m *KnxGroupAddressFreeLevel) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := e.EncodeElement(m.SubGroup, xml.StartElement{Name: xml.Name{Local: "subGroup"}}); err != nil {
		return err
	}
	return nil
}

func (m KnxGroupAddressFreeLevel) String() string {
	return string(m.Box("", 120))
}

func (m KnxGroupAddressFreeLevel) Box(name string, width int) utils.AsciiBox {
	boxName := "KnxGroupAddressFreeLevel"
	if name != "" {
		boxName += "/" + name
	}
	childBoxer := func() []utils.AsciiBox {
		boxes := make([]utils.AsciiBox, 0)
		// Simple field (case simple)
		// uint16 can be boxed as anything with the least amount of space
		boxes = append(boxes, utils.BoxAnything("SubGroup", m.SubGroup, -1))
		return boxes
	}
	return m.Parent.BoxParent(boxName, width, childBoxer)
}
