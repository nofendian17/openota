// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: proto/airline/v1/airline_service.proto

// Protobuf Java Version: 3.25.1
package com.core.airline.v1;

public interface UpdateRequestOrBuilder extends
    // @@protoc_insertion_point(interface_extends:core.airline.v1.UpdateRequest)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <code>string id = 1 [json_name = "id", (.buf.validate.field) = { ... }</code>
   * @return The id.
   */
  java.lang.String getId();
  /**
   * <code>string id = 1 [json_name = "id", (.buf.validate.field) = { ... }</code>
   * @return The bytes for id.
   */
  com.google.protobuf.ByteString
      getIdBytes();

  /**
   * <code>string code = 2 [json_name = "code", (.buf.validate.field) = { ... }</code>
   * @return The code.
   */
  java.lang.String getCode();
  /**
   * <code>string code = 2 [json_name = "code", (.buf.validate.field) = { ... }</code>
   * @return The bytes for code.
   */
  com.google.protobuf.ByteString
      getCodeBytes();

  /**
   * <code>string name = 3 [json_name = "name", (.buf.validate.field) = { ... }</code>
   * @return The name.
   */
  java.lang.String getName();
  /**
   * <code>string name = 3 [json_name = "name", (.buf.validate.field) = { ... }</code>
   * @return The bytes for name.
   */
  com.google.protobuf.ByteString
      getNameBytes();

  /**
   * <code>string logo = 4 [json_name = "logo", (.buf.validate.field) = { ... }</code>
   * @return The logo.
   */
  java.lang.String getLogo();
  /**
   * <code>string logo = 4 [json_name = "logo", (.buf.validate.field) = { ... }</code>
   * @return The bytes for logo.
   */
  com.google.protobuf.ByteString
      getLogoBytes();

  /**
   * <code>bool is_active = 5 [json_name = "isActive", (.buf.validate.field) = { ... }</code>
   * @return The isActive.
   */
  boolean getIsActive();

  /**
   * <code>int64 precedence = 6 [json_name = "precedence", (.buf.validate.field) = { ... }</code>
   * @return The precedence.
   */
  long getPrecedence();
}