// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: proto/city/v1/city_service.proto

// Protobuf Java Version: 3.25.1
package com.core.city.v1;

public interface CreateRequestOrBuilder extends
    // @@protoc_insertion_point(interface_extends:core.city.v1.CreateRequest)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <code>string name = 1 [json_name = "name", (.buf.validate.field) = { ... }</code>
   * @return The name.
   */
  java.lang.String getName();
  /**
   * <code>string name = 1 [json_name = "name", (.buf.validate.field) = { ... }</code>
   * @return The bytes for name.
   */
  com.google.protobuf.ByteString
      getNameBytes();

  /**
   * <code>string state_id = 2 [json_name = "stateId", (.buf.validate.field) = { ... }</code>
   * @return The stateId.
   */
  java.lang.String getStateId();
  /**
   * <code>string state_id = 2 [json_name = "stateId", (.buf.validate.field) = { ... }</code>
   * @return The bytes for stateId.
   */
  com.google.protobuf.ByteString
      getStateIdBytes();

  /**
   * <code>double latitude = 3 [json_name = "latitude", (.buf.validate.field) = { ... }</code>
   * @return The latitude.
   */
  double getLatitude();

  /**
   * <code>double longitude = 4 [json_name = "longitude", (.buf.validate.field) = { ... }</code>
   * @return The longitude.
   */
  double getLongitude();

  /**
   * <code>bool is_active = 5 [json_name = "isActive", (.buf.validate.field) = { ... }</code>
   * @return The isActive.
   */
  boolean getIsActive();
}
