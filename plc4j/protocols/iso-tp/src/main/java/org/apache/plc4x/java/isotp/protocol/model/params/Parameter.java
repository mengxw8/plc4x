/*
Licensed to the Apache Software Foundation (ASF) under one
or more contributor license agreements.  See the NOTICE file
distributed with this work for additional information
regarding copyright ownership.  The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License.  You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, either express or implied.  See the License for the
specific language governing permissions and limitations
under the License.
*/
package org.apache.plc4x.java.isotp.protocol.model.params;

import io.netty.buffer.ByteBuf;
import org.apache.plc4x.java.api.exceptions.PlcProtocolException;
import org.apache.plc4x.java.isotp.protocol.model.types.ParameterCode;

public interface Parameter {

    ParameterCode getType();

    byte getLength();

    void serialize(ByteBuf out);

    static Parameter decode(ByteBuf out) throws PlcProtocolException {
        byte parameterCodeByte = out.readByte();
        ParameterCode parameterCode = ParameterCode.valueOf(parameterCodeByte);
        if (parameterCode == null) {
            throw new PlcProtocolException(
                String.format("Unknown ISO-TO parameter code %02x", parameterCodeByte));
        }

        switch (parameterCode) {
            case CALLING_TSAP:
            case CALLED_TSAP:
                return TsapParameter.decode(out, parameterCode);
            case CHECKSUM:
                return ChecksumParameter.decode(out);
            case DISCONNECT_ADDITIONAL_INFORMATION:
                return DisconnectAdditionalInformationParameter.decode(out);
            case TPDU_SIZE:
                return TpduSizeParameter.decode(out);
            default:
                throw new PlcProtocolException(
                    String.format("Unimplemented ISO-TO parameter code %02x", parameterCodeByte));
        }
    }

}
