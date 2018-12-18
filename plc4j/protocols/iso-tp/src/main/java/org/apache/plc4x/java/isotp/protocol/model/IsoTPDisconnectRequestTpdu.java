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

import java.util.LinkedList;
import java.util.List;

public class IsoTPDisconnectRequestTpdu extends IsoTPDisconnectTpdu {

    private final DisconnectReason disconnectReason;

    public IsoTPDisconnectRequestTpdu(short destinationReference, short sourceReference, DisconnectReason disconnectReason, ByteBuf userData) {
        super(TpduCode.DISCONNECT_REQUEST, destinationReference, sourceReference, new LinkedList<>(), userData);
        this.disconnectReason = disconnectReason;
    }

    public IsoTPDisconnectRequestTpdu(short destinationReference, short sourceReference, DisconnectReason disconnectReason, List<Parameter> parameters, ByteBuf userData) {
        super(TpduCode.DISCONNECT_REQUEST, destinationReference, sourceReference, parameters, userData);
        this.disconnectReason = disconnectReason;
    }

    public DisconnectReason getDisconnectReason() {
        return disconnectReason;
    }

    @Override
    protected byte getTpduHeaderLength() {
        return (byte) (super.getTpduHeaderLength() + 1);
    }

    @Override
    protected void serializeTpduHeader(ByteBuf out) {
        super.serializeTpduHeader(out);

        // The additional attributes a disconnection request has.
        out.writeByte(disconnectReason.getCode());
    }

}
