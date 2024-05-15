// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: proto/healthcheck/healthcheck.proto

// Protobuf Java Version: 3.25.1
package com.core.healthcheck;

public interface HealthOrBuilder extends
    // @@protoc_insertion_point(interface_extends:core.healthcheck.Health)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <code>string version = 1 [json_name = "version"];</code>
   * @return The version.
   */
  java.lang.String getVersion();
  /**
   * <code>string version = 1 [json_name = "version"];</code>
   * @return The bytes for version.
   */
  com.google.protobuf.ByteString
      getVersionBytes();

  /**
   * <code>string uptime = 2 [json_name = "uptime"];</code>
   * @return The uptime.
   */
  java.lang.String getUptime();
  /**
   * <code>string uptime = 2 [json_name = "uptime"];</code>
   * @return The bytes for uptime.
   */
  com.google.protobuf.ByteString
      getUptimeBytes();

  /**
   * <code>string cpu = 3 [json_name = "cpu"];</code>
   * @return The cpu.
   */
  java.lang.String getCpu();
  /**
   * <code>string cpu = 3 [json_name = "cpu"];</code>
   * @return The bytes for cpu.
   */
  com.google.protobuf.ByteString
      getCpuBytes();

  /**
   * <code>string memory = 4 [json_name = "memory"];</code>
   * @return The memory.
   */
  java.lang.String getMemory();
  /**
   * <code>string memory = 4 [json_name = "memory"];</code>
   * @return The bytes for memory.
   */
  com.google.protobuf.ByteString
      getMemoryBytes();
}
