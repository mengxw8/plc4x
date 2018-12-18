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

/**
 * Base class for calling and called TSAPs
 */
public abstract class TsapParameter implements Parameter {

    private final short tsapId;

    public TsapParameter(short tsapId) {
        this.tsapId = tsapId;
    }

    public short getTsapId() {
        return tsapId;
    }

    @Override
    public byte getLength() {
        return 4;
    }

    @Override
    public void serialize(ByteBuf out) {
        out.writeByte(getType().getCode());
        // Output the size of the rest of the header (Total size of the header - 2)
        out.writeByte(getLength() - 2);
        out.writeShort(getTsapId());
    }

    public static TsapParameter decode(ByteBuf in, ParameterCode parameterCode) throws PlcProtocolException {
        // Skip the length.
        in.skipBytes(1);
        short tsapId = in.readShort();
        switch (parameterCode) {
            case CALLING_TSAP:
                return new TsapParameterCalling(tsapId);
            case CALLED_TSAP:
                return new TsapParameterCalled(tsapId);
            default:
                // This shouldn't happen if it's called from the Parameter.decode method.
                throw new PlcProtocolException("Internal Error");
        }
    }

}
