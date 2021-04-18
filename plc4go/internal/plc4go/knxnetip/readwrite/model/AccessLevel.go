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
	"fmt"
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"io"
)

// Code generated by build-utils. DO NOT EDIT.

type AccessLevel uint8

type IAccessLevel interface {
	Purpose() string
	NeedsAuthentication() bool
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

const (
	AccessLevel_Level0  AccessLevel = 0x0
	AccessLevel_Level1  AccessLevel = 0x1
	AccessLevel_Level2  AccessLevel = 0x2
	AccessLevel_Level3  AccessLevel = 0x3
	AccessLevel_Level15 AccessLevel = 0xF
)

var AccessLevelValues []AccessLevel

func init() {
	AccessLevelValues = []AccessLevel{
		AccessLevel_Level0,
		AccessLevel_Level1,
		AccessLevel_Level2,
		AccessLevel_Level3,
		AccessLevel_Level15,
	}
}

func (e AccessLevel) Purpose() string {
	switch e {
	case 0x0:
		{ /* '0x0' */
			return "system manufacturer"
		}
	case 0x1:
		{ /* '0x1' */
			return "product manufacturer"
		}
	case 0x2:
		{ /* '0x2' */
			return "configuration"
		}
	case 0x3:
		{ /* '0x3' */
			return "end-user"
		}
	case 0xF:
		{ /* '0xF' */
			return "read access"
		}
	default:
		{
			return ""
		}
	}
}

func (e AccessLevel) NeedsAuthentication() bool {
	switch e {
	case 0x0:
		{ /* '0x0' */
			return true
		}
	case 0x1:
		{ /* '0x1' */
			return true
		}
	case 0x2:
		{ /* '0x2' */
			return true
		}
	case 0x3:
		{ /* '0x3' */
			return false
		}
	case 0xF:
		{ /* '0xF' */
			return false
		}
	default:
		{
			return false
		}
	}
}
func AccessLevelByValue(value uint8) AccessLevel {
	switch value {
	case 0x0:
		return AccessLevel_Level0
	case 0x1:
		return AccessLevel_Level1
	case 0x2:
		return AccessLevel_Level2
	case 0x3:
		return AccessLevel_Level3
	case 0xF:
		return AccessLevel_Level15
	}
	return 0
}

func AccessLevelByName(value string) AccessLevel {
	switch value {
	case "Level0":
		return AccessLevel_Level0
	case "Level1":
		return AccessLevel_Level1
	case "Level2":
		return AccessLevel_Level2
	case "Level3":
		return AccessLevel_Level3
	case "Level15":
		return AccessLevel_Level15
	}
	return 0
}

func CastAccessLevel(structType interface{}) AccessLevel {
	castFunc := func(typ interface{}) AccessLevel {
		if sAccessLevel, ok := typ.(AccessLevel); ok {
			return sAccessLevel
		}
		return 0
	}
	return castFunc(structType)
}

func (m AccessLevel) LengthInBits() uint16 {
	return 4
}

func (m AccessLevel) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func AccessLevelParse(io utils.ReadBuffer) (AccessLevel, error) {
	val, err := io.ReadUint8(4)
	if err != nil {
		return 0, nil
	}
	return AccessLevelByValue(val), nil
}

func (e AccessLevel) Serialize(io utils.WriteBuffer) error {
	err := io.WriteUint8(4, uint8(e))
	return err
}

func (m *AccessLevel) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var token xml.Token
	var err error
	for {
		token, err = d.Token()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		switch token.(type) {
		case xml.CharData:
			tok := token.(xml.CharData)
			*m = AccessLevelByName(string(tok))
		}
	}
}

func (m AccessLevel) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := e.EncodeElement(m.String(), start); err != nil {
		return err
	}
	return nil
}

func (e AccessLevel) name() string {
	switch e {
	case AccessLevel_Level0:
		return "Level0"
	case AccessLevel_Level1:
		return "Level1"
	case AccessLevel_Level2:
		return "Level2"
	case AccessLevel_Level3:
		return "Level3"
	case AccessLevel_Level15:
		return "Level15"
	}
	return ""
}

func (e AccessLevel) String() string {
	return e.name()
}

func (m AccessLevel) Box(s string, i int) utils.AsciiBox {
	boxName := "AccessLevel"
	if s != "" {
		boxName += "/" + s
	}
	return utils.BoxString(boxName, fmt.Sprintf("%#0*x %s", 1, uint8(m), m.name()), -1)
}
