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
package org.apache.plc4x.java.simulated.connection;

import org.apache.commons.lang3.tuple.Pair;
import org.apache.plc4x.java.api.model.PlcSubscriptionField;
import org.apache.plc4x.java.api.model.PlcSubscriptionHandle;
import org.apache.plc4x.java.api.value.*;
import org.apache.plc4x.java.simulated.field.SimulatedField;
import org.apache.plc4x.java.simulated.readwrite.io.DataItemIO;
import org.apache.plc4x.java.simulated.readwrite.types.SimulatedDataTypeSizes;
import org.apache.plc4x.java.spi.generation.ParseException;
import org.apache.plc4x.java.spi.generation.ReadBuffer;

import org.apache.plc4x.java.spi.model.DefaultPlcSubscriptionField;
import org.apache.plc4x.java.spi.values.IEC61131ValueHandler;
import org.apache.plc4x.java.simulated.field.SimulatedField;

import java.nio.charset.StandardCharsets;
import java.time.Duration;
import java.util.*;
import java.util.concurrent.*;
import java.util.function.Consumer;

/**
 * Test device storing its state in memory.
 * Values are stored in a HashMap.
 */
public class SimulatedDevice {

    private final Random random = new Random();

    private final String name;

    private final Map<SimulatedField, PlcValue> state = new HashMap<>();

    private final Map<PlcSubscriptionHandle, ScheduledFuture<?>> cyclicSubscriptions = new HashMap<>();

    private final Map<PlcSubscriptionHandle, Future<?>> eventSubscriptions = new HashMap<>();

    private final IdentityHashMap<PlcSubscriptionHandle, Pair<SimulatedField, Consumer<PlcValue>>> changeOfStateSubscriptions = new IdentityHashMap<>();

    private final ScheduledExecutorService scheduler = Executors.newScheduledThreadPool(1);

    private final ExecutorService pool = Executors.newCachedThreadPool();

    public SimulatedDevice(String name) {
        this.name = name;
    }

    public Optional<PlcValue> get(SimulatedField field) {
        Objects.requireNonNull(field);
        switch (field.getType()) {
            case STATE:
                return Optional.ofNullable(state.get(field));
            case RANDOM:
                return Optional.of(randomValue(field));
            case STDOUT:
                return Optional.empty();
        }
        throw new IllegalArgumentException("Unsupported field type: " + field.getType().name());
    }

    public void set(SimulatedField field, PlcValue value) {
        Objects.requireNonNull(field);
        switch (field.getType()) {
            case STATE:
                changeOfStateSubscriptions.values().stream()
                    .filter(pair -> pair.getKey().equals(field))
                    .map(Pair::getValue)
                    .forEach(baseDefaultPlcValueConsumer -> baseDefaultPlcValueConsumer.accept(value));
                state.put(field, value);
                return;
            case STDOUT:
                System.out.printf("TEST PLC STDOUT [%s]: %s%n", field.getName(), value.getString());
                return;
            case RANDOM:
                switch (field.getPlcDataType()) {
                    case "IEC61131_STRING":
                    case "IEC61131_WSTRING":
                        break;
                    default:
                        try {
                            DataItemIO.staticSerialize(value, field.getPlcDataType(), 1, false);
                        } catch (ParseException e) {
                            System.out.printf("Write failed");
                        }
                }
                System.out.printf("TEST PLC RANDOM [%s]: %s%n", field.getName(), value.getString());
                return;
        }
        throw new IllegalArgumentException("Unsupported field type: " + field.getType().name());
    }

    @SuppressWarnings("unchecked")
    private PlcValue randomValue(SimulatedField field) {
        Object result = null;

        Short fieldDataTypeSize = SimulatedDataTypeSizes.enumForValue(field.getPlcDataType()).getDataTypeSize();

        byte[] b = new byte[fieldDataTypeSize];
        new Random().nextBytes(b);

        ReadBuffer io = new ReadBuffer(b);
        try {
            return DataItemIO.staticParse(io, field.getPlcDataType(), 1);
        } catch (ParseException e) {
            return null;
        }

    }

    @Override
    public String toString() {
        return name;
    }

    public void addCyclicSubscription(Consumer<PlcValue> consumer, PlcSubscriptionHandle handle, PlcSubscriptionField plcField, Duration duration) {
        ScheduledFuture<?> scheduledFuture = scheduler.scheduleAtFixedRate(() -> {
            PlcValue baseDefaultPlcValue = state.get(((DefaultPlcSubscriptionField) plcField).getPlcField());
            if (baseDefaultPlcValue == null) {
                return;
            }
            consumer.accept(baseDefaultPlcValue);
        }, duration.toMillis(), duration.toMillis(), TimeUnit.MILLISECONDS);
        cyclicSubscriptions.put(handle, scheduledFuture);
    }

    public void addChangeOfStateSubscription(Consumer<PlcValue> consumer, PlcSubscriptionHandle handle, PlcSubscriptionField plcField) {
        changeOfStateSubscriptions.put(handle, Pair.of((SimulatedField) ((DefaultPlcSubscriptionField) plcField).getPlcField(), consumer));
    }

    public void addEventSubscription(Consumer<PlcValue> consumer, PlcSubscriptionHandle handle, PlcSubscriptionField plcField) {
        Future<?> submit = pool.submit(() -> {
            while (!Thread.currentThread().isInterrupted()) {
                PlcValue baseDefaultPlcValue = state.get(((DefaultPlcSubscriptionField) plcField).getPlcField());
                if (baseDefaultPlcValue == null) {
                    continue;
                }
                consumer.accept(baseDefaultPlcValue);
                try {
                    TimeUnit.SECONDS.sleep((long) (Math.random() * 10));
                } catch (InterruptedException ignore) {
                    Thread.currentThread().interrupt();
                    return;
                }
            }
        });

        eventSubscriptions.put(handle, submit);
    }

    public void removeHandles(Collection<? extends PlcSubscriptionHandle> internalPlcSubscriptionHandles) {
        internalPlcSubscriptionHandles.forEach(handle -> {
            ScheduledFuture<?> remove = cyclicSubscriptions.remove(handle);
            if (remove == null) {
                return;
            }
            remove.cancel(true);
        });
        internalPlcSubscriptionHandles.forEach(handle -> {
            Future<?> remove = eventSubscriptions.remove(handle);
            if (remove == null) {
                return;
            }
            remove.cancel(true);
        });
        internalPlcSubscriptionHandles.forEach(changeOfStateSubscriptions::remove);
    }
}
