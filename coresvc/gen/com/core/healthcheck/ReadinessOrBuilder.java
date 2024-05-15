// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: proto/healthcheck/healthcheck.proto

// Protobuf Java Version: 3.25.1
package com.core.healthcheck;

public interface ReadinessOrBuilder extends
    // @@protoc_insertion_point(interface_extends:core.healthcheck.Readiness)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <code>string status = 1 [json_name = "status"];</code>
   * @return The status.
   */
  java.lang.String getStatus();
  /**
   * <code>string status = 1 [json_name = "status"];</code>
   * @return The bytes for status.
   */
  com.google.protobuf.ByteString
      getStatusBytes();

  /**
   * <code>repeated .core.healthcheck.Check checks = 2 [json_name = "checks"];</code>
   */
  java.util.List<com.core.healthcheck.Check> 
      getChecksList();
  /**
   * <code>repeated .core.healthcheck.Check checks = 2 [json_name = "checks"];</code>
   */
  com.core.healthcheck.Check getChecks(int index);
  /**
   * <code>repeated .core.healthcheck.Check checks = 2 [json_name = "checks"];</code>
   */
  int getChecksCount();
  /**
   * <code>repeated .core.healthcheck.Check checks = 2 [json_name = "checks"];</code>
   */
  java.util.List<? extends com.core.healthcheck.CheckOrBuilder> 
      getChecksOrBuilderList();
  /**
   * <code>repeated .core.healthcheck.Check checks = 2 [json_name = "checks"];</code>
   */
  com.core.healthcheck.CheckOrBuilder getChecksOrBuilder(
      int index);
}
