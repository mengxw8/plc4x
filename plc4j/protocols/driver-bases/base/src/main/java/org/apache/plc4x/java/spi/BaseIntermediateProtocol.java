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

import io.netty.channel.ChannelHandlerContext;
import org.apache.plc4x.java.api.exceptions.PlcProtocolException;
import org.apache.plc4x.java.base.PlcMessageToMessageCodec;

import java.util.List;

/**
 * Any intermediate protocol usually produces and consumes messages of a give type.
 *
 * @param <HI_TYPE> the type of objects this protocol offers to higher level protocols.
 * @param <LOW_TYPE> the type of objects this protocol offers to lower level protocols.
 */
public abstract class BaseIntermediateProtocol<HI_TYPE extends Model<LOW_TYPE>, LOW_TYPE> extends PlcMessageToMessageCodec<LOW_TYPE, HI_TYPE> {

    protected abstract ModelIO<HI_TYPE, LOW_TYPE> getRootModelIo();

    ////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
    // Encoding
    ////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

    @Override
    protected void encode(ChannelHandlerContext ctx, HI_TYPE msg, List<Object> out) throws Exception {
        // If the message is not intended for this level, msg is sometimes null
        // This is usually the case if PLC4X needs to send a low-level message.
        // Especially during connection setup such cases may occur.
        if (msg == null) {
            return;
        }

        // Have the message serialize itself.
        try {
            // Perform some checks before serializing.
            handlePreEncoding(ctx, msg);

            LOW_TYPE serialized = msg.serialize();

            // Perform some checks after serializing.
            handlePostEncoding(ctx, msg, serialized);

            if(serialized != null) {
                out.add(serialized);
            }
        } catch (PlcProtocolException e) {
            exceptionCaught(ctx, e);
        }
    }

    protected void handlePreEncoding(ChannelHandlerContext ctx, HI_TYPE hiType) throws PlcProtocolException {
        // Ignore per default ...
    }

    protected void handlePostEncoding(ChannelHandlerContext ctx, HI_TYPE hiType, LOW_TYPE lowType) throws PlcProtocolException {
        // Ignore per default ...
    }

    ////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
    // Decoding
    ////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

    @Override
    protected void decode(ChannelHandlerContext ctx, LOW_TYPE msg, List<Object> out) throws Exception {
        try {
            // Perform some checks before de-serializing.
            handlePreDecoding(ctx, msg);

            HI_TYPE decodedMessage = getRootModelIo().decode(msg);

            // Perform some checks after de-serializing.
            handlePostDecoding(ctx, decodedMessage, msg);

            if(decodedMessage != null) {
                out.add(decodedMessage);
            }
        } catch (PlcProtocolException e) {
            exceptionCaught(ctx, e);
        }
    }

    protected void handlePreDecoding(ChannelHandlerContext ctx, LOW_TYPE lowType) throws PlcProtocolException {
        // Ignore per default ...
    }

    protected void handlePostDecoding(ChannelHandlerContext ctx, HI_TYPE hiType, LOW_TYPE lowType) throws PlcProtocolException {
        // Ignore per default ...
    }

}
