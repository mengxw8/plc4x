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
type ModbusPDUReadDeviceIdentificationRequest struct {
	Parent *ModbusPDU
}

// The corresponding interface
type IModbusPDUReadDeviceIdentificationRequest interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *ModbusPDUReadDeviceIdentificationRequest) ErrorFlag() bool {
	return false
}

func (m *ModbusPDUReadDeviceIdentificationRequest) FunctionFlag() uint8 {
	return 0x2B
}

func (m *ModbusPDUReadDeviceIdentificationRequest) Response() bool {
	return false
}

func (m *ModbusPDUReadDeviceIdentificationRequest) InitializeParent(parent *ModbusPDU) {
}

func NewModbusPDUReadDeviceIdentificationRequest() *ModbusPDU {
	child := &ModbusPDUReadDeviceIdentificationRequest{
		Parent: NewModbusPDU(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastModbusPDUReadDeviceIdentificationRequest(structType interface{}) *ModbusPDUReadDeviceIdentificationRequest {
	castFunc := func(typ interface{}) *ModbusPDUReadDeviceIdentificationRequest {
		if casted, ok := typ.(ModbusPDUReadDeviceIdentificationRequest); ok {
			return &casted
		}
		if casted, ok := typ.(*ModbusPDUReadDeviceIdentificationRequest); ok {
			return casted
		}
		if casted, ok := typ.(ModbusPDU); ok {
			return CastModbusPDUReadDeviceIdentificationRequest(casted.Child)
		}
		if casted, ok := typ.(*ModbusPDU); ok {
			return CastModbusPDUReadDeviceIdentificationRequest(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *ModbusPDUReadDeviceIdentificationRequest) GetTypeName() string {
	return "ModbusPDUReadDeviceIdentificationRequest"
}

func (m *ModbusPDUReadDeviceIdentificationRequest) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *ModbusPDUReadDeviceIdentificationRequest) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	return lengthInBits
}

func (m *ModbusPDUReadDeviceIdentificationRequest) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func ModbusPDUReadDeviceIdentificationRequestParse(io utils.ReadBuffer) (*ModbusPDU, error) {

	// Create a partially initialized instance
	_child := &ModbusPDUReadDeviceIdentificationRequest{
		Parent: &ModbusPDU{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *ModbusPDUReadDeviceIdentificationRequest) Serialize(io utils.WriteBuffer) error {
	ser := func() error {

		return nil
	}
	return m.Parent.SerializeParent(io, m, ser)
}

func (m *ModbusPDUReadDeviceIdentificationRequest) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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

func (m *ModbusPDUReadDeviceIdentificationRequest) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return nil
}

func (m ModbusPDUReadDeviceIdentificationRequest) String() string {
	return string(m.Box("", 120))
}

func (m ModbusPDUReadDeviceIdentificationRequest) Box(name string, width int) utils.AsciiBox {
	boxName := "ModbusPDUReadDeviceIdentificationRequest"
	if name != "" {
		boxName += "/" + name
	}
	childBoxer := func() []utils.AsciiBox {
		boxes := make([]utils.AsciiBox, 0)
		return boxes
	}
	return m.Parent.BoxParent(boxName, width, childBoxer)
}
