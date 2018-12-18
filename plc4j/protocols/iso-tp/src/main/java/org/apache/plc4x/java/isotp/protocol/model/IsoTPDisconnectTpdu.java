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
import org.apache.plc4x.java.isotp.protocol.model.types.DisconnectReason;
import org.apache.plc4x.java.isotp.protocol.model.types.TpduCode;

import java.util.List;

public abstract class IsoTPDisconnectTpdu extends IsoTPTpdu {

    private final short destinationReference;
    private final short sourceReference;

    public IsoTPDisconnectTpdu(TpduCode tpduCode, short destinationReference, short sourceReference, List<Parameter> parameters, ByteBuf userData) {
        super(tpduCode, parameters, userData);
        this.destinationReference = destinationReference;
        this.sourceReference = sourceReference;
    }

    public short getDestinationReference() {
        return destinationReference;
    }

    public short getSourceReference() {
        return sourceReference;
    }

    @Override
    protected byte getTpduHeaderLength() {
        return 4;
    }

    @Override
    protected void serializeTpduHeader(ByteBuf out) {
        out.writeShort(getDestinationReference());
        out.writeShort(getSourceReference());
    }

    static IsoTPDisconnectTpdu decode(TpduCode tpduCode, ByteBuf userData) {
        short destinationReference = userData.readShort();
        short sourceReference = userData.readShort();
        if (tpduCode == TpduCode.DISCONNECT_REQUEST) {
            DisconnectReason disconnectReason = DisconnectReason.valueOf(userData.readByte());
            return new IsoTPDisconnectRequestTpdu(
                destinationReference, sourceReference, disconnectReason, userData);
        } else {  // TpduCode.DISCONNECT_CONFIRM
            return new IsoTPDisconnectConfirmTpdu(
                destinationReference, sourceReference, userData);
        }
    }

}
