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
import io.netty.buffer.Unpooled;
import org.apache.plc4x.java.api.exceptions.PlcProtocolException;
import org.apache.plc4x.java.base.messages.PlcProtocolMessage;
import org.apache.plc4x.java.base.messages.PlcRawMessage;
import org.apache.plc4x.java.isoontcp.protocol.model.IsoOnTcpMessage;
import org.apache.plc4x.java.isotp.protocol.model.params.Parameter;
import org.apache.plc4x.java.isotp.protocol.model.types.TpduCode;
import org.apache.plc4x.java.spi.Model;

import java.util.List;
import java.util.Optional;

public abstract class IsoTPTpdu extends PlcRawMessage implements Model<IsoOnTcpMessage> {

    private final TpduCode tpduCode;
    private final List<Parameter> parameters;

    public IsoTPTpdu(TpduCode tpduCode, List<Parameter> parameters, ByteBuf userData) {
        this(tpduCode, parameters, userData, null);
    }

    public IsoTPTpdu(TpduCode tpduCode, List<Parameter> parameters, ByteBuf userData, PlcProtocolMessage parent) {
        super(userData, parent);
        this.tpduCode = tpduCode;
        this.parameters = parameters;
    }

    public TpduCode getTpduCode() {
        return tpduCode;
    }

    public List<Parameter> getParameters() {
        return parameters;
    }

    public <T extends Parameter> Optional<T> getParameter(Class<T> parameterClass) {
        if (parameters != null) {
            for (Parameter parameter : parameters) {
                if (parameter.getClass() == parameterClass) {
                    return Optional.of(parameterClass.cast(parameter));
                }
            }
        }
        return Optional.empty();
    }

    void addParameter(Parameter parameter) {
        parameters.add(parameter);
    }

    @Override
    public int getSerializedLength() throws PlcProtocolException {
        return getHeaderLength() + getUserData().readableBytes();
    }

    @Override
    public IsoOnTcpMessage serialize() throws PlcProtocolException {
        return new ModelIO().encode(this);
    }

    protected abstract byte getTpduHeaderLength();

    protected abstract void serializeTpduHeader(ByteBuf out);

    protected byte getHeaderLength() {
        // Size of the header itself plus the bytes for size and tpdu code.
        byte size = (byte) (getTpduHeaderLength() + 2);
        for (Parameter parameter : getParameters()) {
            size += parameter.getLength();
        }
        return size;
    }

    public static class ModelIO implements org.apache.plc4x.java.spi.ModelIO<IsoTPTpdu, IsoOnTcpMessage> {

        @Override
        public IsoOnTcpMessage encode(IsoTPTpdu model) throws PlcProtocolException {
            ByteBuf buf = Unpooled.buffer();

            // Header length indicator field (The length byte doesn't count)
            buf.writeByte(model.getHeaderLength() - 1);
            // TPDU Code (First 4 bits), Initial Credit Allocation (Second 4 bits)
            buf.writeByte(model.getTpduCode().getCode());

            // Have the TPDUs encode their custom headers.
            model.serializeTpduHeader(buf);

            // Serialize all the header-parameters.
            for (Parameter parameter : model.getParameters()) {
                parameter.serialize(buf);
            }

            // Add the user-data itself.
            buf.writeBytes(model.getUserData());

            return new IsoOnTcpMessage(buf);
        }

        @Override
        public IsoTPTpdu decode(IsoOnTcpMessage in) throws PlcProtocolException {
            if(in == null) {
                return null;
            }
            ByteBuf userData = in.getUserData();
            if (userData.writerIndex() < 1) {
                return null;
            }

            int packetStart = userData.readerIndex();
            byte headerLength = userData.readByte();
            int headerEnd = packetStart + headerLength;
            TpduCode tpduCode = TpduCode.valueOf(userData.readByte());

            IsoTPTpdu tpdu;
            switch (tpduCode) {
                case CONNECTION_REQUEST:
                case CONNECTION_CONFIRM:
                    tpdu = IsoTPConnectionTpdu.decode(tpduCode, userData);
                    break;
                case DATA:
                    tpdu = IsoTPDataTpdu.decode(tpduCode, userData);
                    break;
                case DISCONNECT_REQUEST:
                case DISCONNECT_CONFIRM:
                    tpdu = IsoTPDisconnectTpdu.decode(tpduCode, userData);
                    break;
                case TPDU_ERROR:
                    tpdu = IsoTPErrorTpdu.decode(tpduCode, userData);
                    break;
                default:
                    throw new PlcProtocolException(String.format("Tpdu Code %s not implemented", tpduCode.name()));
            }

            // Read variable header parameters
            while (userData.readerIndex() < headerEnd) {
                tpdu.addParameter(Parameter.decode(userData));
            }

            return tpdu;
        }
    }

}
