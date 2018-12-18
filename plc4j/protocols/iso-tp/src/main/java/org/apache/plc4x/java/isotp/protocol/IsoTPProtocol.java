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

package org.apache.plc4x.java.isotp.protocol;

import io.netty.buffer.Unpooled;
import io.netty.channel.ChannelHandler;
import io.netty.channel.ChannelHandlerContext;
import org.apache.plc4x.java.api.exceptions.PlcProtocolException;
import org.apache.plc4x.java.api.exceptions.PlcProtocolPayloadTooBigException;
import org.apache.plc4x.java.base.events.ConnectEvent;
import org.apache.plc4x.java.isoontcp.protocol.IsoOnTcpProtocol;
import org.apache.plc4x.java.isoontcp.protocol.model.IsoOnTcpMessage;
import org.apache.plc4x.java.isotp.protocol.events.IsoTPConnectedEvent;
import org.apache.plc4x.java.isotp.protocol.model.IsoTPConnectionConfirmTpdu;
import org.apache.plc4x.java.isotp.protocol.model.IsoTPConnectionRequestTpdu;
import org.apache.plc4x.java.isotp.protocol.model.IsoTPTpdu;
import org.apache.plc4x.java.isotp.protocol.model.params.TpduSizeParameter;
import org.apache.plc4x.java.isotp.protocol.model.params.TsapParameterCalled;
import org.apache.plc4x.java.isotp.protocol.model.params.TsapParameterCalling;
import org.apache.plc4x.java.isotp.protocol.model.types.ProtocolClass;
import org.apache.plc4x.java.isotp.protocol.model.types.TpduSize;
import org.apache.plc4x.java.spi.BaseIntermediateProtocol;
import org.apache.plc4x.java.spi.ModelIO;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.util.Arrays;
import java.util.List;

public class IsoTPProtocol extends BaseIntermediateProtocol<IsoTPTpdu, IsoOnTcpMessage> {

    private static final Logger logger = LoggerFactory.getLogger(IsoTPProtocol.class);

    private short callingTsapId;
    private short calledTsapId;
    private TpduSize tpduSize;

    public IsoTPProtocol(short callingTsapId, short calledTsapId, TpduSize tpduSize) {
        this.callingTsapId = callingTsapId;
        this.calledTsapId = calledTsapId;
        this.tpduSize = tpduSize;
    }

    @Override
    protected ModelIO<IsoTPTpdu, IsoOnTcpMessage> getRootModelIo() {
        return new IsoTPTpdu.ModelIO();
    }

    /**
     * If the IsoTP protocol is used on top of the ISO on TCP protocol, then as soon as the pipeline receives the
     * request to connect, an IsoTP connection request TPDU must be sent in order to initialize the connection.
     *
     * @param ctx the current protocol layers context
     * @param evt the event
     * @throws Exception throws an exception if something goes wrong internally
     */
    @Override
    public void userEventTriggered(ChannelHandlerContext ctx, Object evt) throws Exception {
        ChannelHandler prevHandler = getPrevChannelHandler(ctx);

        // If the connection has just been established, start setting up the connection
        // by sending a connection request to the plc.
        if ((prevHandler instanceof IsoOnTcpProtocol) && (evt instanceof ConnectEvent)) {
            logger.debug("ISO Transport Protocol Sending Connection Request");
            // Open the session on ISO Transport Protocol first.
            IsoTPConnectionRequestTpdu connectionRequest = new IsoTPConnectionRequestTpdu(
                (short) 0x0000, (short) 0x000F, ProtocolClass.CLASS_0,
                Arrays.asList(
                    new TsapParameterCalled(calledTsapId),
                    new TsapParameterCalling(callingTsapId),
                    new TpduSizeParameter(tpduSize)),
                Unpooled.buffer());
            ctx.channel().writeAndFlush(connectionRequest);
        } else {
            super.userEventTriggered(ctx, evt);
        }
    }

    ////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
    // Encoding
    ////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

    @Override
    protected void handlePreEncoding(ChannelHandlerContext ctx, IsoTPTpdu isoTPTpdu) throws PlcProtocolException {
        // Check if the message would exceed the negotiated maximum size.
        if(isoTPTpdu.getSerializedLength() > tpduSize.getValue()) {
            throw new PlcProtocolPayloadTooBigException(
                "iso-tp", tpduSize.getValue(), isoTPTpdu.getSerializedLength(), isoTPTpdu);
        }
    }

    @Override
    protected void encode(ChannelHandlerContext ctx, IsoTPTpdu msg, List<Object> out) throws Exception {
        super.encode(ctx, msg, out);
    }

    ////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
    // Decoding
    ////////////////////////////////////////////////////////////////////////////////////////////////////////////////////


    @Override
    protected void decode(ChannelHandlerContext ctx, IsoOnTcpMessage msg, List<Object> out) throws Exception {
        super.decode(ctx, msg, out);
    }

    @Override
    protected void handlePostDecoding(ChannelHandlerContext ctx, IsoTPTpdu isoTPTpdu, IsoOnTcpMessage isoOnTcpMessage) {
        // If the incoming message was a connection confirmation, take the called-tsap-id and
        // tpdu-size values from that and save them for future usage.
        if (isoTPTpdu instanceof IsoTPConnectionConfirmTpdu) {
            isoTPTpdu.getParameter(TsapParameterCalled.class).ifPresent(
                calledTsapParameter -> calledTsapId = calledTsapParameter.getTsapId());
            isoTPTpdu.getParameter(TpduSizeParameter.class).ifPresent(
                tpduSizeParameter -> tpduSize = tpduSizeParameter.getTpduSize());

            // Send an event to the pipeline telling the other protocol implementations that
            // they can start doing their thing.
            ctx.channel().pipeline().fireUserEventTriggered(new IsoTPConnectedEvent());
        }
    }

}
