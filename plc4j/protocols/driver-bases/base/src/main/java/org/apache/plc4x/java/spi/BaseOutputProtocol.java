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

package org.apache.plc4x.java.spi;

import io.netty.buffer.ByteBuf;
import io.netty.buffer.ByteBufUtil;
import io.netty.channel.ChannelHandlerContext;
import org.apache.plc4x.java.api.exceptions.PlcProtocolException;
import org.apache.plc4x.java.base.PlcByteToMessageCodec;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.util.List;

/**
 * Output protocols produce pure byte output and are typically the lowest level
 * of any protocol hierarchy.
 *
 * @param <HI_TYPE> the type of objects this protocol offers to higher level protocols.
 */
public abstract class BaseOutputProtocol<HI_TYPE extends Model<ByteBuf>> extends PlcByteToMessageCodec<HI_TYPE> {

    private static final Logger logger = LoggerFactory.getLogger(BaseOutputProtocol.class);

    protected abstract ModelIO<HI_TYPE, ByteBuf> getRootModelIo();

    ////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
    // Encoding
    ////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

    @Override
    protected void encode(ChannelHandlerContext ctx, HI_TYPE in, ByteBuf out) throws Exception {
        try {
            out.writeBytes(in.serialize());
        } catch (PlcProtocolException e) {
            exceptionCaught(ctx, e);
        }
    }

    ////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
    // Decoding
    ////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

    @Override
    protected void decode(ChannelHandlerContext ctx, ByteBuf in, List<Object> out) throws Exception {
        if(logger.isTraceEnabled()) {
            logger.trace("Got Data: {}", ByteBufUtil.hexDump(in));
        }
        try {
            HI_TYPE decodedMessage = getRootModelIo().decode(in);
            if(decodedMessage != null) {
                out.add(decodedMessage);
            }
        } catch (PlcProtocolException e) {
            exceptionCaught(ctx, e);
        }
    }

}
