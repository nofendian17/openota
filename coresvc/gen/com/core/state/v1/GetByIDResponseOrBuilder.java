// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: proto/state/v1/state_service.proto

// Protobuf Java Version: 3.25.1
package com.core.state.v1;

public interface GetByIDResponseOrBuilder extends
    // @@protoc_insertion_point(interface_extends:core.state.v1.GetByIDResponse)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <code>.core.state.v1.State state = 1 [json_name = "state"];</code>
   * @return Whether the state field is set.
   */
  boolean hasState();
  /**
   * <code>.core.state.v1.State state = 1 [json_name = "state"];</code>
   * @return The state.
   */
  com.core.state.v1.State getState();
  /**
   * <code>.core.state.v1.State state = 1 [json_name = "state"];</code>
   */
  com.core.state.v1.StateOrBuilder getStateOrBuilder();
}