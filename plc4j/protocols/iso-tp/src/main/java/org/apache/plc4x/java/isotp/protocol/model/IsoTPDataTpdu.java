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
import org.apache.plc4x.java.base.messages.PlcProtocolMessage;
import org.apache.plc4x.java.isotp.protocol.model.params.Parameter;
import org.apache.plc4x.java.isotp.protocol.model.types.TpduCode;

import java.util.LinkedList;
import java.util.List;

public class IsoTPDataTpdu extends IsoTPTpdu {

    private final boolean eot;
    private final byte tpduRef;

    public IsoTPDataTpdu(boolean eot, byte tpduRef, ByteBuf userData) {
        this(eot, tpduRef, new LinkedList<>(), userData, null);
    }

    public IsoTPDataTpdu(boolean eot, byte tpduRef, List<Parameter> parameters, ByteBuf userData) {
        this(eot, tpduRef, parameters, userData, null);
    }

    public IsoTPDataTpdu(boolean eot, byte tpduRef, List<Parameter> parameters, ByteBuf userData, PlcProtocolMessage parent) {
        super(TpduCode.DATA, parameters, userData, parent);
        this.eot = eot;
        this.tpduRef = tpduRef;
    }

    public boolean isEot() {
        return eot;
    }

    public byte getTpduRef() {
        return tpduRef;
    }

    @Override
    public byte getTpduHeaderLength() {
        return 1;
    }

    @Override
    protected void serializeTpduHeader(ByteBuf out) {
        // EOT (Bit 8 = 1) / TPDU (All other bits 0)
        out.writeByte((byte) (getTpduRef() | (isEot() ? 0x80 : 0x00)));
    }

    static IsoTPDataTpdu decode(TpduCode tpduCode, ByteBuf userData) {
        byte tmp = userData.readByte();
        // Bit 8 is the EOT indicator (1 = last TPDU)
        boolean eot = (tmp & 0x80) == 0x80;
        // The rest is simply a 7 bit number identifying the current request.
        byte tpduRef = (byte) (tmp & 0x7F);
        return new IsoTPDataTpdu(eot, tpduRef, userData);
    }

}
