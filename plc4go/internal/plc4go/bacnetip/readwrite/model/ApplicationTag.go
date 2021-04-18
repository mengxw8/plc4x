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

type ApplicationTag int8

type IApplicationTag interface {
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

const (
	ApplicationTag_NULL                     ApplicationTag = 0x0
	ApplicationTag_BOOLEAN                  ApplicationTag = 0x1
	ApplicationTag_UNSIGNED_INTEGER         ApplicationTag = 0x2
	ApplicationTag_SIGNED_INTEGER           ApplicationTag = 0x3
	ApplicationTag_REAL                     ApplicationTag = 0x4
	ApplicationTag_DOUBLE                   ApplicationTag = 0x5
	ApplicationTag_OCTET_STRING             ApplicationTag = 0x6
	ApplicationTag_CHARACTER_STRING         ApplicationTag = 0x7
	ApplicationTag_BIT_STRING               ApplicationTag = 0x8
	ApplicationTag_ENUMERATED               ApplicationTag = 0x9
	ApplicationTag_DATE                     ApplicationTag = 0xA
	ApplicationTag_TIME                     ApplicationTag = 0xB
	ApplicationTag_BACNET_OBJECT_IDENTIFIER ApplicationTag = 0xC
)

var ApplicationTagValues []ApplicationTag

func init() {
	ApplicationTagValues = []ApplicationTag{
		ApplicationTag_NULL,
		ApplicationTag_BOOLEAN,
		ApplicationTag_UNSIGNED_INTEGER,
		ApplicationTag_SIGNED_INTEGER,
		ApplicationTag_REAL,
		ApplicationTag_DOUBLE,
		ApplicationTag_OCTET_STRING,
		ApplicationTag_CHARACTER_STRING,
		ApplicationTag_BIT_STRING,
		ApplicationTag_ENUMERATED,
		ApplicationTag_DATE,
		ApplicationTag_TIME,
		ApplicationTag_BACNET_OBJECT_IDENTIFIER,
	}
}

func ApplicationTagByValue(value int8) ApplicationTag {
	switch value {
	case 0x0:
		return ApplicationTag_NULL
	case 0x1:
		return ApplicationTag_BOOLEAN
	case 0x2:
		return ApplicationTag_UNSIGNED_INTEGER
	case 0x3:
		return ApplicationTag_SIGNED_INTEGER
	case 0x4:
		return ApplicationTag_REAL
	case 0x5:
		return ApplicationTag_DOUBLE
	case 0x6:
		return ApplicationTag_OCTET_STRING
	case 0x7:
		return ApplicationTag_CHARACTER_STRING
	case 0x8:
		return ApplicationTag_BIT_STRING
	case 0x9:
		return ApplicationTag_ENUMERATED
	case 0xA:
		return ApplicationTag_DATE
	case 0xB:
		return ApplicationTag_TIME
	case 0xC:
		return ApplicationTag_BACNET_OBJECT_IDENTIFIER
	}
	return 0
}

func ApplicationTagByName(value string) ApplicationTag {
	switch value {
	case "NULL":
		return ApplicationTag_NULL
	case "BOOLEAN":
		return ApplicationTag_BOOLEAN
	case "UNSIGNED_INTEGER":
		return ApplicationTag_UNSIGNED_INTEGER
	case "SIGNED_INTEGER":
		return ApplicationTag_SIGNED_INTEGER
	case "REAL":
		return ApplicationTag_REAL
	case "DOUBLE":
		return ApplicationTag_DOUBLE
	case "OCTET_STRING":
		return ApplicationTag_OCTET_STRING
	case "CHARACTER_STRING":
		return ApplicationTag_CHARACTER_STRING
	case "BIT_STRING":
		return ApplicationTag_BIT_STRING
	case "ENUMERATED":
		return ApplicationTag_ENUMERATED
	case "DATE":
		return ApplicationTag_DATE
	case "TIME":
		return ApplicationTag_TIME
	case "BACNET_OBJECT_IDENTIFIER":
		return ApplicationTag_BACNET_OBJECT_IDENTIFIER
	}
	return 0
}

func CastApplicationTag(structType interface{}) ApplicationTag {
	castFunc := func(typ interface{}) ApplicationTag {
		if sApplicationTag, ok := typ.(ApplicationTag); ok {
			return sApplicationTag
		}
		return 0
	}
	return castFunc(structType)
}

func (m ApplicationTag) LengthInBits() uint16 {
	return 4
}

func (m ApplicationTag) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func ApplicationTagParse(io utils.ReadBuffer) (ApplicationTag, error) {
	val, err := io.ReadInt8(4)
	if err != nil {
		return 0, nil
	}
	return ApplicationTagByValue(val), nil
}

func (e ApplicationTag) Serialize(io utils.WriteBuffer) error {
	err := io.WriteInt8(4, int8(e))
	return err
}

func (m *ApplicationTag) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			*m = ApplicationTagByName(string(tok))
		}
	}
}

func (m ApplicationTag) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := e.EncodeElement(m.String(), start); err != nil {
		return err
	}
	return nil
}

func (e ApplicationTag) name() string {
	switch e {
	case ApplicationTag_NULL:
		return "NULL"
	case ApplicationTag_BOOLEAN:
		return "BOOLEAN"
	case ApplicationTag_UNSIGNED_INTEGER:
		return "UNSIGNED_INTEGER"
	case ApplicationTag_SIGNED_INTEGER:
		return "SIGNED_INTEGER"
	case ApplicationTag_REAL:
		return "REAL"
	case ApplicationTag_DOUBLE:
		return "DOUBLE"
	case ApplicationTag_OCTET_STRING:
		return "OCTET_STRING"
	case ApplicationTag_CHARACTER_STRING:
		return "CHARACTER_STRING"
	case ApplicationTag_BIT_STRING:
		return "BIT_STRING"
	case ApplicationTag_ENUMERATED:
		return "ENUMERATED"
	case ApplicationTag_DATE:
		return "DATE"
	case ApplicationTag_TIME:
		return "TIME"
	case ApplicationTag_BACNET_OBJECT_IDENTIFIER:
		return "BACNET_OBJECT_IDENTIFIER"
	}
	return ""
}

func (e ApplicationTag) String() string {
	return e.name()
}

func (m ApplicationTag) Box(s string, i int) utils.AsciiBox {
	boxName := "ApplicationTag"
	if s != "" {
		boxName += "/" + s
	}
	return utils.BoxString(boxName, fmt.Sprintf("%#0*x %s", 1, int8(m), m.name()), -1)
}
