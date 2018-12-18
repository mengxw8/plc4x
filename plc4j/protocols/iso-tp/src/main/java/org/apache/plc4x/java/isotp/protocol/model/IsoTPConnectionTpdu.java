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
package org.apache.plc4x.java.isotp.protocol.model;

import io.netty.buffer.ByteBuf;
import org.apache.plc4x.java.isotp.protocol.model.params.Parameter;
import org.apache.plc4x.java.isotp.protocol.model.types.ProtocolClass;
import org.apache.plc4x.java.isotp.protocol.model.types.TpduCode;

import java.util.List;

public abstract class IsoTPConnectionTpdu extends IsoTPTpdu {

    private final short destinationReference;
    private final short sourceReference;
    private final ProtocolClass protocolClass;

    public IsoTPConnectionTpdu(TpduCode tpduCode, short destinationReference, short sourceReference, ProtocolClass protocolClass, List<Parameter> parameters, ByteBuf userData) {
        super(tpduCode, parameters, userData);
        this.destinationReference = destinationReference;
        this.sourceReference = sourceReference;
        this.protocolClass = protocolClass;
    }

    public short getDestinationReference() {
        return destinationReference;
    }

    public short getSourceReference() {
        return sourceReference;
    }

    public ProtocolClass getProtocolClass() {
        return protocolClass;
    }

    @Override
    protected byte getTpduHeaderLength() {
        return 5;
    }

    @Override
    protected void serializeTpduHeader(ByteBuf out) {
        out.writeShort(getDestinationReference());
        out.writeShort(getSourceReference());
        out.writeByte(getProtocolClass().getCode());
    }

    static IsoTPConnectionTpdu decode(TpduCode tpduCode, ByteBuf userData) {
        short destinationReference = userData.readShort();
        short sourceReference = userData.readShort();
        ProtocolClass protocolClass = ProtocolClass.valueOf(userData.readByte());

        if (tpduCode == TpduCode.CONNECTION_REQUEST) {
            return new IsoTPConnectionRequestTpdu(destinationReference, sourceReference, protocolClass, userData);

        } else { // TpduCode.CONNECTION_CONFIRM
            return new IsoTPConnectionConfirmTpdu(destinationReference, sourceReference, protocolClass, userData);
        }
    }

}
