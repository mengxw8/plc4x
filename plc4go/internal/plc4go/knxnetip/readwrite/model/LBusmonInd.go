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
	"strconv"
)

// Code generated by build-utils. DO NOT EDIT.

// The data-structure of this message
type LBusmonInd struct {
	AdditionalInformationLength uint8
	AdditionalInformation       []*CEMIAdditionalInformation
	DataFrame                   *LDataFrame
	Crc                         *uint8
	Parent                      *CEMI
}

// The corresponding interface
type ILBusmonInd interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *LBusmonInd) MessageCode() uint8 {
	return 0x2B
}

func (m *LBusmonInd) InitializeParent(parent *CEMI) {
}

func NewLBusmonInd(additionalInformationLength uint8, additionalInformation []*CEMIAdditionalInformation, dataFrame *LDataFrame, crc *uint8) *CEMI {
	child := &LBusmonInd{
		AdditionalInformationLength: additionalInformationLength,
		AdditionalInformation:       additionalInformation,
		DataFrame:                   dataFrame,
		Crc:                         crc,
		Parent:                      NewCEMI(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastLBusmonInd(structType interface{}) *LBusmonInd {
	castFunc := func(typ interface{}) *LBusmonInd {
		if casted, ok := typ.(LBusmonInd); ok {
			return &casted
		}
		if casted, ok := typ.(*LBusmonInd); ok {
			return casted
		}
		if casted, ok := typ.(CEMI); ok {
			return CastLBusmonInd(casted.Child)
		}
		if casted, ok := typ.(*CEMI); ok {
			return CastLBusmonInd(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *LBusmonInd) GetTypeName() string {
	return "LBusmonInd"
}

func (m *LBusmonInd) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *LBusmonInd) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Simple field (additionalInformationLength)
	lengthInBits += 8

	// Array field
	if len(m.AdditionalInformation) > 0 {
		for _, element := range m.AdditionalInformation {
			lengthInBits += element.LengthInBits()
		}
	}

	// Simple field (dataFrame)
	lengthInBits += m.DataFrame.LengthInBits()

	// Optional Field (crc)
	if m.Crc != nil {
		lengthInBits += 8
	}

	return lengthInBits
}

func (m *LBusmonInd) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func LBusmonIndParse(io utils.ReadBuffer) (*CEMI, error) {

	// Simple Field (additionalInformationLength)
	additionalInformationLength, _additionalInformationLengthErr := io.ReadUint8(8)
	if _additionalInformationLengthErr != nil {
		return nil, errors.Wrap(_additionalInformationLengthErr, "Error parsing 'additionalInformationLength' field")
	}

	// Array field (additionalInformation)
	// Length array
	additionalInformation := make([]*CEMIAdditionalInformation, 0)
	_additionalInformationLength := additionalInformationLength
	_additionalInformationEndPos := io.GetPos() + uint16(_additionalInformationLength)
	for io.GetPos() < _additionalInformationEndPos {
		_item, _err := CEMIAdditionalInformationParse(io)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'additionalInformation' field")
		}
		additionalInformation = append(additionalInformation, _item)
	}

	// Simple Field (dataFrame)
	dataFrame, _dataFrameErr := LDataFrameParse(io)
	if _dataFrameErr != nil {
		return nil, errors.Wrap(_dataFrameErr, "Error parsing 'dataFrame' field")
	}

	// Optional Field (crc) (Can be skipped, if a given expression evaluates to false)
	var crc *uint8 = nil
	if CastLDataFrame(dataFrame).Child.NotAckFrame() {
		_val, _err := io.ReadUint8(8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'crc' field")
		}
		crc = &_val
	}

	// Create a partially initialized instance
	_child := &LBusmonInd{
		AdditionalInformationLength: additionalInformationLength,
		AdditionalInformation:       additionalInformation,
		DataFrame:                   dataFrame,
		Crc:                         crc,
		Parent:                      &CEMI{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *LBusmonInd) Serialize(io utils.WriteBuffer) error {
	ser := func() error {

		// Simple Field (additionalInformationLength)
		additionalInformationLength := uint8(m.AdditionalInformationLength)
		_additionalInformationLengthErr := io.WriteUint8(8, (additionalInformationLength))
		if _additionalInformationLengthErr != nil {
			return errors.Wrap(_additionalInformationLengthErr, "Error serializing 'additionalInformationLength' field")
		}

		// Array Field (additionalInformation)
		if m.AdditionalInformation != nil {
			for _, _element := range m.AdditionalInformation {
				_elementErr := _element.Serialize(io)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'additionalInformation' field")
				}
			}
		}

		// Simple Field (dataFrame)
		_dataFrameErr := m.DataFrame.Serialize(io)
		if _dataFrameErr != nil {
			return errors.Wrap(_dataFrameErr, "Error serializing 'dataFrame' field")
		}

		// Optional Field (crc) (Can be skipped, if the value is null)
		var crc *uint8 = nil
		if m.Crc != nil {
			crc = m.Crc
			_crcErr := io.WriteUint8(8, *(crc))
			if _crcErr != nil {
				return errors.Wrap(_crcErr, "Error serializing 'crc' field")
			}
		}

		return nil
	}
	return m.Parent.SerializeParent(io, m, ser)
}

func (m *LBusmonInd) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "additionalInformationLength":
				var data uint8
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.AdditionalInformationLength = data
			case "additionalInformation":
			arrayLoop:
				for {
					token, err = d.Token()
					switch token.(type) {
					case xml.StartElement:
						tok := token.(xml.StartElement)
						var dt *CEMIAdditionalInformation
						if err := d.DecodeElement(&dt, &tok); err != nil {
							return err
						}
						m.AdditionalInformation = append(m.AdditionalInformation, dt)
					default:
						break arrayLoop
					}
				}
			case "dataFrame":
				var dt *LDataFrame
				if err := d.DecodeElement(&dt, &tok); err != nil {
					if err == io.EOF {
						continue
					}
					return err
				}
				m.DataFrame = dt
			case "crc":
				// When working with pointers we need to check for an empty element
				var dataString string
				if err := d.DecodeElement(&dataString, &tok); err != nil {
					return err
				}
				if dataString != "" {
					atoi, err := strconv.Atoi(dataString)
					if err != nil {
						return err
					}
					data := uint8(atoi)
					m.Crc = &data
				}
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

func (m *LBusmonInd) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := e.EncodeElement(m.AdditionalInformationLength, xml.StartElement{Name: xml.Name{Local: "additionalInformationLength"}}); err != nil {
		return err
	}
	if err := e.EncodeToken(xml.StartElement{Name: xml.Name{Local: "additionalInformation"}}); err != nil {
		return err
	}
	for _, arrayElement := range m.AdditionalInformation {
		if err := e.EncodeElement(arrayElement, xml.StartElement{Name: xml.Name{Local: "additionalInformation"}}); err != nil {
			return err
		}
	}
	if err := e.EncodeToken(xml.EndElement{Name: xml.Name{Local: "additionalInformation"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.DataFrame, xml.StartElement{Name: xml.Name{Local: "dataFrame"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.Crc, xml.StartElement{Name: xml.Name{Local: "crc"}}); err != nil {
		return err
	}
	return nil
}

func (m LBusmonInd) String() string {
	return string(m.Box("", 120))
}

func (m LBusmonInd) Box(name string, width int) utils.AsciiBox {
	boxName := "LBusmonInd"
	if name != "" {
		boxName += "/" + name
	}
	childBoxer := func() []utils.AsciiBox {
		boxes := make([]utils.AsciiBox, 0)
		// Simple field (case simple)
		// uint8 can be boxed as anything with the least amount of space
		boxes = append(boxes, utils.BoxAnything("AdditionalInformationLength", m.AdditionalInformationLength, -1))
		// Array Field (additionalInformation)
		if m.AdditionalInformation != nil {
			// Complex array base type
			arrayBoxes := make([]utils.AsciiBox, 0)
			for _, _element := range m.AdditionalInformation {
				arrayBoxes = append(arrayBoxes, utils.BoxAnything("", _element, width-2))
			}
			boxes = append(boxes, utils.BoxBox("AdditionalInformation", utils.AlignBoxes(arrayBoxes, width-4), 0))
		}
		// Complex field (case complex)
		boxes = append(boxes, m.DataFrame.Box("dataFrame", width-2))
		// Optional Field (crc) (Can be skipped, if the value is null)
		var crc *uint8 = nil
		if m.Crc != nil {
			crc = m.Crc
			// uint8 can be boxed as anything with the least amount of space
			boxes = append(boxes, utils.BoxAnything("Crc", *(crc), -1))
		}
		return boxes
	}
	return m.Parent.BoxParent(boxName, width, childBoxer)
}
