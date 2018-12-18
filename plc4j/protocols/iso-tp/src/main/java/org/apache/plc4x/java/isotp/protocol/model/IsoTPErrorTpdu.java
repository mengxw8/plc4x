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
import org.apache.plc4x.java.isotp.protocol.model.types.RejectCause;
import org.apache.plc4x.java.isotp.protocol.model.types.TpduCode;

import java.util.LinkedList;
import java.util.List;

public class IsoTPErrorTpdu extends IsoTPTpdu {

    private final short destinationReference;
    private final RejectCause rejectCause;

    public IsoTPErrorTpdu(short destinationReference, RejectCause rejectCause, ByteBuf userData) {
        super(TpduCode.TPDU_ERROR, new LinkedList<>(), userData);
        this.destinationReference = destinationReference;
        this.rejectCause = rejectCause;
    }

    public IsoTPErrorTpdu(short destinationReference, RejectCause rejectCause, List<Parameter> parameters, ByteBuf userData) {
        super(TpduCode.TPDU_ERROR, parameters, userData);
        this.destinationReference = destinationReference;
        this.rejectCause = rejectCause;
    }

    public short getDestinationReference() {
        return destinationReference;
    }

    public RejectCause getRejectCause() {
        return rejectCause;
    }

    @Override
    protected byte getTpduHeaderLength() {
        return 3;
    }

    @Override
    protected void serializeTpduHeader(ByteBuf out) {
        out.writeShort(getDestinationReference());
        out.writeByte(getRejectCause().getCode());
    }

    static IsoTPErrorTpdu decode(TpduCode tpduCode, ByteBuf userData) {
        short destinationReference = userData.readShort();
        RejectCause rejectCause = RejectCause.valueOf(userData.readByte());
        return new IsoTPErrorTpdu(destinationReference, rejectCause, userData);
    }

}
