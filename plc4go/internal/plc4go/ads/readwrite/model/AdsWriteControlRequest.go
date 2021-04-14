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
	"encoding/hex"
	"encoding/xml"
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
	"io"
	"strings"
)

// Code generated by build-utils. DO NOT EDIT.

// The data-structure of this message
type AdsWriteControlRequest struct {
	AdsState    uint16
	DeviceState uint16
	Data        []int8
	Parent      *AdsData
}

// The corresponding interface
type IAdsWriteControlRequest interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *AdsWriteControlRequest) CommandId() CommandId {
	return CommandId_ADS_WRITE_CONTROL
}

func (m *AdsWriteControlRequest) Response() bool {
	return false
}

func (m *AdsWriteControlRequest) InitializeParent(parent *AdsData) {
}

func NewAdsWriteControlRequest(adsState uint16, deviceState uint16, data []int8) *AdsData {
	child := &AdsWriteControlRequest{
		AdsState:    adsState,
		DeviceState: deviceState,
		Data:        data,
		Parent:      NewAdsData(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastAdsWriteControlRequest(structType interface{}) *AdsWriteControlRequest {
	castFunc := func(typ interface{}) *AdsWriteControlRequest {
		if casted, ok := typ.(AdsWriteControlRequest); ok {
			return &casted
		}
		if casted, ok := typ.(*AdsWriteControlRequest); ok {
			return casted
		}
		if casted, ok := typ.(AdsData); ok {
			return CastAdsWriteControlRequest(casted.Child)
		}
		if casted, ok := typ.(*AdsData); ok {
			return CastAdsWriteControlRequest(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *AdsWriteControlRequest) GetTypeName() string {
	return "AdsWriteControlRequest"
}

func (m *AdsWriteControlRequest) LengthInBits() uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Simple field (adsState)
	lengthInBits += 16

	// Simple field (deviceState)
	lengthInBits += 16

	// Implicit Field (length)
	lengthInBits += 32

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}

func (m *AdsWriteControlRequest) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func AdsWriteControlRequestParse(io *utils.ReadBuffer) (*AdsData, error) {

	// Simple Field (adsState)
	adsState, _adsStateErr := io.ReadUint16(16)
	if _adsStateErr != nil {
		return nil, errors.Wrap(_adsStateErr, "Error parsing 'adsState' field")
	}

	// Simple Field (deviceState)
	deviceState, _deviceStateErr := io.ReadUint16(16)
	if _deviceStateErr != nil {
		return nil, errors.Wrap(_deviceStateErr, "Error parsing 'deviceState' field")
	}

	// Implicit Field (length) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	length, _lengthErr := io.ReadUint32(32)
	_ = length
	if _lengthErr != nil {
		return nil, errors.Wrap(_lengthErr, "Error parsing 'length' field")
	}

	// Array field (data)
	// Count array
	data := make([]int8, length)
	for curItem := uint16(0); curItem < uint16(length); curItem++ {
		_item, _err := io.ReadInt8(8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'data' field")
		}
		data[curItem] = _item
	}

	// Create a partially initialized instance
	_child := &AdsWriteControlRequest{
		AdsState:    adsState,
		DeviceState: deviceState,
		Data:        data,
		Parent:      &AdsData{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *AdsWriteControlRequest) Serialize(io utils.WriteBuffer) error {
	ser := func() error {

		// Simple Field (adsState)
		adsState := uint16(m.AdsState)
		_adsStateErr := io.WriteUint16(16, (adsState))
		if _adsStateErr != nil {
			return errors.Wrap(_adsStateErr, "Error serializing 'adsState' field")
		}

		// Simple Field (deviceState)
		deviceState := uint16(m.DeviceState)
		_deviceStateErr := io.WriteUint16(16, (deviceState))
		if _deviceStateErr != nil {
			return errors.Wrap(_deviceStateErr, "Error serializing 'deviceState' field")
		}

		// Implicit Field (length) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
		length := uint32(uint32(len(m.Data)))
		_lengthErr := io.WriteUint32(32, (length))
		if _lengthErr != nil {
			return errors.Wrap(_lengthErr, "Error serializing 'length' field")
		}

		// Array Field (data)
		if m.Data != nil {
			for _, _element := range m.Data {
				_elementErr := io.WriteInt8(8, _element)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'data' field")
				}
			}
		}

		return nil
	}
	return m.Parent.SerializeParent(io, m, ser)
}

func (m *AdsWriteControlRequest) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "adsState":
				var data uint16
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.AdsState = data
			case "deviceState":
				var data uint16
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.DeviceState = data
			case "data":
				var _encoded string
				if err := d.DecodeElement(&_encoded, &tok); err != nil {
					return err
				}
				_decoded, err := hex.DecodeString(_encoded)
				_len := len(_decoded)
				if err != nil {
					return err
				}
				m.Data = utils.ByteArrayToInt8Array(_decoded[0:_len])
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

func (m *AdsWriteControlRequest) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := e.EncodeElement(m.AdsState, xml.StartElement{Name: xml.Name{Local: "adsState"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.DeviceState, xml.StartElement{Name: xml.Name{Local: "deviceState"}}); err != nil {
		return err
	}
	_encodedData := hex.EncodeToString(utils.Int8ArrayToByteArray(m.Data))
	_encodedData = strings.ToUpper(_encodedData)
	if err := e.EncodeElement(_encodedData, xml.StartElement{Name: xml.Name{Local: "data"}}); err != nil {
		return err
	}
	return nil
}

func (m AdsWriteControlRequest) String() string {
	return string(m.Box("AdsWriteControlRequest", utils.DefaultWidth*2))
}

func (m AdsWriteControlRequest) Box(name string, width int) utils.AsciiBox {
	if name == "" {
		name = "AdsWriteControlRequest"
	}
	boxes := make([]utils.AsciiBox, 0)
	boxes = append(boxes, utils.BoxAnything("AdsState", m.AdsState, width-2))
	boxes = append(boxes, utils.BoxAnything("DeviceState", m.DeviceState, width-2))
	boxes = append(boxes, utils.BoxAnything("Data", m.Data, width-2))
	return utils.BoxBox(name, utils.AlignBoxes(boxes, width-2), 0)
}
