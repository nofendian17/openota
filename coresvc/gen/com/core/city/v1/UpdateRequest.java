// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: proto/city/v1/city_service.proto

// Protobuf Java Version: 3.25.1
package com.core.city.v1;

/**
 * Protobuf type {@code core.city.v1.UpdateRequest}
 */
public final class UpdateRequest extends
    com.google.protobuf.GeneratedMessageV3 implements
    // @@protoc_insertion_point(message_implements:core.city.v1.UpdateRequest)
    UpdateRequestOrBuilder {
private static final long serialVersionUID = 0L;
  // Use UpdateRequest.newBuilder() to construct.
  private UpdateRequest(com.google.protobuf.GeneratedMessageV3.Builder<?> builder) {
    super(builder);
  }
  private UpdateRequest() {
    id_ = "";
    name_ = "";
    stateId_ = "";
  }

  @java.lang.Override
  @SuppressWarnings({"unused"})
  protected java.lang.Object newInstance(
      UnusedPrivateParameter unused) {
    return new UpdateRequest();
  }

  public static final com.google.protobuf.Descriptors.Descriptor
      getDescriptor() {
    return com.core.city.v1.CityServiceProto.internal_static_core_city_v1_UpdateRequest_descriptor;
  }

  @java.lang.Override
  protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internalGetFieldAccessorTable() {
    return com.core.city.v1.CityServiceProto.internal_static_core_city_v1_UpdateRequest_fieldAccessorTable
        .ensureFieldAccessorsInitialized(
            com.core.city.v1.UpdateRequest.class, com.core.city.v1.UpdateRequest.Builder.class);
  }

  public static final int ID_FIELD_NUMBER = 1;
  @SuppressWarnings("serial")
  private volatile java.lang.Object id_ = "";
  /**
   * <code>string id = 1 [json_name = "id", (.buf.validate.field) = { ... }</code>
   * @return The id.
   */
  @java.lang.Override
  public java.lang.String getId() {
    java.lang.Object ref = id_;
    if (ref instanceof java.lang.String) {
      return (java.lang.String) ref;
    } else {
      com.google.protobuf.ByteString bs = 
          (com.google.protobuf.ByteString) ref;
      java.lang.String s = bs.toStringUtf8();
      id_ = s;
      return s;
    }
  }
  /**
   * <code>string id = 1 [json_name = "id", (.buf.validate.field) = { ... }</code>
   * @return The bytes for id.
   */
  @java.lang.Override
  public com.google.protobuf.ByteString
      getIdBytes() {
    java.lang.Object ref = id_;
    if (ref instanceof java.lang.String) {
      com.google.protobuf.ByteString b = 
          com.google.protobuf.ByteString.copyFromUtf8(
              (java.lang.String) ref);
      id_ = b;
      return b;
    } else {
      return (com.google.protobuf.ByteString) ref;
    }
  }

  public static final int NAME_FIELD_NUMBER = 2;
  @SuppressWarnings("serial")
  private volatile java.lang.Object name_ = "";
  /**
   * <code>string name = 2 [json_name = "name", (.buf.validate.field) = { ... }</code>
   * @return The name.
   */
  @java.lang.Override
  public java.lang.String getName() {
    java.lang.Object ref = name_;
    if (ref instanceof java.lang.String) {
      return (java.lang.String) ref;
    } else {
      com.google.protobuf.ByteString bs = 
          (com.google.protobuf.ByteString) ref;
      java.lang.String s = bs.toStringUtf8();
      name_ = s;
      return s;
    }
  }
  /**
   * <code>string name = 2 [json_name = "name", (.buf.validate.field) = { ... }</code>
   * @return The bytes for name.
   */
  @java.lang.Override
  public com.google.protobuf.ByteString
      getNameBytes() {
    java.lang.Object ref = name_;
    if (ref instanceof java.lang.String) {
      com.google.protobuf.ByteString b = 
          com.google.protobuf.ByteString.copyFromUtf8(
              (java.lang.String) ref);
      name_ = b;
      return b;
    } else {
      return (com.google.protobuf.ByteString) ref;
    }
  }

  public static final int STATE_ID_FIELD_NUMBER = 3;
  @SuppressWarnings("serial")
  private volatile java.lang.Object stateId_ = "";
  /**
   * <code>string state_id = 3 [json_name = "stateId", (.buf.validate.field) = { ... }</code>
   * @return The stateId.
   */
  @java.lang.Override
  public java.lang.String getStateId() {
    java.lang.Object ref = stateId_;
    if (ref instanceof java.lang.String) {
      return (java.lang.String) ref;
    } else {
      com.google.protobuf.ByteString bs = 
          (com.google.protobuf.ByteString) ref;
      java.lang.String s = bs.toStringUtf8();
      stateId_ = s;
      return s;
    }
  }
  /**
   * <code>string state_id = 3 [json_name = "stateId", (.buf.validate.field) = { ... }</code>
   * @return The bytes for stateId.
   */
  @java.lang.Override
  public com.google.protobuf.ByteString
      getStateIdBytes() {
    java.lang.Object ref = stateId_;
    if (ref instanceof java.lang.String) {
      com.google.protobuf.ByteString b = 
          com.google.protobuf.ByteString.copyFromUtf8(
              (java.lang.String) ref);
      stateId_ = b;
      return b;
    } else {
      return (com.google.protobuf.ByteString) ref;
    }
  }

  public static final int LATITUDE_FIELD_NUMBER = 4;
  private double latitude_ = 0D;
  /**
   * <code>double latitude = 4 [json_name = "latitude", (.buf.validate.field) = { ... }</code>
   * @return The latitude.
   */
  @java.lang.Override
  public double getLatitude() {
    return latitude_;
  }

  public static final int LONGITUDE_FIELD_NUMBER = 5;
  private double longitude_ = 0D;
  /**
   * <code>double longitude = 5 [json_name = "longitude", (.buf.validate.field) = { ... }</code>
   * @return The longitude.
   */
  @java.lang.Override
  public double getLongitude() {
    return longitude_;
  }

  public static final int IS_ACTIVE_FIELD_NUMBER = 7;
  private boolean isActive_ = false;
  /**
   * <code>bool is_active = 7 [json_name = "isActive", (.buf.validate.field) = { ... }</code>
   * @return The isActive.
   */
  @java.lang.Override
  public boolean getIsActive() {
    return isActive_;
  }

  public static final int PRECEDENCE_FIELD_NUMBER = 6;
  private long precedence_ = 0L;
  /**
   * <code>int64 precedence = 6 [json_name = "precedence", (.buf.validate.field) = { ... }</code>
   * @return The precedence.
   */
  @java.lang.Override
  public long getPrecedence() {
    return precedence_;
  }

  public static com.core.city.v1.UpdateRequest parseFrom(
      java.nio.ByteBuffer data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static com.core.city.v1.UpdateRequest parseFrom(
      java.nio.ByteBuffer data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static com.core.city.v1.UpdateRequest parseFrom(
      com.google.protobuf.ByteString data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static com.core.city.v1.UpdateRequest parseFrom(
      com.google.protobuf.ByteString data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static com.core.city.v1.UpdateRequest parseFrom(byte[] data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static com.core.city.v1.UpdateRequest parseFrom(
      byte[] data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static com.core.city.v1.UpdateRequest parseFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static com.core.city.v1.UpdateRequest parseFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input, extensionRegistry);
  }

  public static com.core.city.v1.UpdateRequest parseDelimitedFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input);
  }

  public static com.core.city.v1.UpdateRequest parseDelimitedFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input, extensionRegistry);
  }
  public static com.core.city.v1.UpdateRequest parseFrom(
      com.google.protobuf.CodedInputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static com.core.city.v1.UpdateRequest parseFrom(
      com.google.protobuf.CodedInputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input, extensionRegistry);
  }

  @java.lang.Override
  public Builder newBuilderForType() { return newBuilder(); }
  public static Builder newBuilder() {
    return DEFAULT_INSTANCE.toBuilder();
  }
  public static Builder newBuilder(com.core.city.v1.UpdateRequest prototype) {
    return DEFAULT_INSTANCE.toBuilder().mergeFrom(prototype);
  }
  @java.lang.Override
  public Builder toBuilder() {
    return this == DEFAULT_INSTANCE
        ? new Builder() : new Builder().mergeFrom(this);
  }

  @java.lang.Override
  protected Builder newBuilderForType(
      com.google.protobuf.GeneratedMessageV3.BuilderParent parent) {
    Builder builder = new Builder(parent);
    return builder;
  }
  /**
   * Protobuf type {@code core.city.v1.UpdateRequest}
   */
  public static final class Builder extends
      com.google.protobuf.GeneratedMessageV3.Builder<Builder> implements
      // @@protoc_insertion_point(builder_implements:core.city.v1.UpdateRequest)
      com.core.city.v1.UpdateRequestOrBuilder {
    public static final com.google.protobuf.Descriptors.Descriptor
        getDescriptor() {
      return com.core.city.v1.CityServiceProto.internal_static_core_city_v1_UpdateRequest_descriptor;
    }

    @java.lang.Override
    protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
        internalGetFieldAccessorTable() {
      return com.core.city.v1.CityServiceProto.internal_static_core_city_v1_UpdateRequest_fieldAccessorTable
          .ensureFieldAccessorsInitialized(
              com.core.city.v1.UpdateRequest.class, com.core.city.v1.UpdateRequest.Builder.class);
    }

    // Construct using com.core.city.v1.UpdateRequest.newBuilder()
    private Builder() {

    }

    private Builder(
        com.google.protobuf.GeneratedMessageV3.BuilderParent parent) {
      super(parent);

    }
    @java.lang.Override
    public Builder clear() {
      super.clear();
      bitField0_ = 0;
      id_ = "";
      name_ = "";
      stateId_ = "";
      latitude_ = 0D;
      longitude_ = 0D;
      isActive_ = false;
      precedence_ = 0L;
      return this;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.Descriptor
        getDescriptorForType() {
      return com.core.city.v1.CityServiceProto.internal_static_core_city_v1_UpdateRequest_descriptor;
    }

    @java.lang.Override
    public com.core.city.v1.UpdateRequest getDefaultInstanceForType() {
      return com.core.city.v1.UpdateRequest.getDefaultInstance();
    }

    @java.lang.Override
    public com.core.city.v1.UpdateRequest build() {
      com.core.city.v1.UpdateRequest result = buildPartial();
      if (!result.isInitialized()) {
        throw newUninitializedMessageException(result);
      }
      return result;
    }

    @java.lang.Override
    public com.core.city.v1.UpdateRequest buildPartial() {
      com.core.city.v1.UpdateRequest result = new com.core.city.v1.UpdateRequest(this);
      if (bitField0_ != 0) { buildPartial0(result); }
      onBuilt();
      return result;
    }

    private void buildPartial0(com.core.city.v1.UpdateRequest result) {
      int from_bitField0_ = bitField0_;
      if (((from_bitField0_ & 0x00000001) != 0)) {
        result.id_ = id_;
      }
      if (((from_bitField0_ & 0x00000002) != 0)) {
        result.name_ = name_;
      }
      if (((from_bitField0_ & 0x00000004) != 0)) {
        result.stateId_ = stateId_;
      }
      if (((from_bitField0_ & 0x00000008) != 0)) {
        result.latitude_ = latitude_;
      }
      if (((from_bitField0_ & 0x00000010) != 0)) {
        result.longitude_ = longitude_;
      }
      if (((from_bitField0_ & 0x00000020) != 0)) {
        result.isActive_ = isActive_;
      }
      if (((from_bitField0_ & 0x00000040) != 0)) {
        result.precedence_ = precedence_;
      }
    }

    @java.lang.Override
    public Builder clone() {
      return super.clone();
    }
    @java.lang.Override
    public Builder setField(
        com.google.protobuf.Descriptors.FieldDescriptor field,
        java.lang.Object value) {
      return super.setField(field, value);
    }
    @java.lang.Override
    public Builder clearField(
        com.google.protobuf.Descriptors.FieldDescriptor field) {
      return super.clearField(field);
    }
    @java.lang.Override
    public Builder clearOneof(
        com.google.protobuf.Descriptors.OneofDescriptor oneof) {
      return super.clearOneof(oneof);
    }
    @java.lang.Override
    public Builder setRepeatedField(
        com.google.protobuf.Descriptors.FieldDescriptor field,
        int index, java.lang.Object value) {
      return super.setRepeatedField(field, index, value);
    }
    @java.lang.Override
    public Builder addRepeatedField(
        com.google.protobuf.Descriptors.FieldDescriptor field,
        java.lang.Object value) {
      return super.addRepeatedField(field, value);
    }
    private int bitField0_;

    private java.lang.Object id_ = "";
    /**
     * <code>string id = 1 [json_name = "id", (.buf.validate.field) = { ... }</code>
     * @return The id.
     */
    public java.lang.String getId() {
      java.lang.Object ref = id_;
      if (!(ref instanceof java.lang.String)) {
        com.google.protobuf.ByteString bs =
            (com.google.protobuf.ByteString) ref;
        java.lang.String s = bs.toStringUtf8();
        id_ = s;
        return s;
      } else {
        return (java.lang.String) ref;
      }
    }
    /**
     * <code>string id = 1 [json_name = "id", (.buf.validate.field) = { ... }</code>
     * @return The bytes for id.
     */
    public com.google.protobuf.ByteString
        getIdBytes() {
      java.lang.Object ref = id_;
      if (ref instanceof String) {
        com.google.protobuf.ByteString b = 
            com.google.protobuf.ByteString.copyFromUtf8(
                (java.lang.String) ref);
        id_ = b;
        return b;
      } else {
        return (com.google.protobuf.ByteString) ref;
      }
    }
    /**
     * <code>string id = 1 [json_name = "id", (.buf.validate.field) = { ... }</code>
     * @param value The id to set.
     * @return This builder for chaining.
     */
    public Builder setId(
        java.lang.String value) {
      if (value == null) { throw new NullPointerException(); }
      id_ = value;
      bitField0_ |= 0x00000001;
      onChanged();
      return this;
    }
    /**
     * <code>string id = 1 [json_name = "id", (.buf.validate.field) = { ... }</code>
     * @return This builder for chaining.
     */
    public Builder clearId() {
      id_ = getDefaultInstance().getId();
      bitField0_ = (bitField0_ & ~0x00000001);
      onChanged();
      return this;
    }
    /**
     * <code>string id = 1 [json_name = "id", (.buf.validate.field) = { ... }</code>
     * @param value The bytes for id to set.
     * @return This builder for chaining.
     */
    public Builder setIdBytes(
        com.google.protobuf.ByteString value) {
      if (value == null) { throw new NullPointerException(); }
      checkByteStringIsUtf8(value);
      id_ = value;
      bitField0_ |= 0x00000001;
      onChanged();
      return this;
    }

    private java.lang.Object name_ = "";
    /**
     * <code>string name = 2 [json_name = "name", (.buf.validate.field) = { ... }</code>
     * @return The name.
     */
    public java.lang.String getName() {
      java.lang.Object ref = name_;
      if (!(ref instanceof java.lang.String)) {
        com.google.protobuf.ByteString bs =
            (com.google.protobuf.ByteString) ref;
        java.lang.String s = bs.toStringUtf8();
        name_ = s;
        return s;
      } else {
        return (java.lang.String) ref;
      }
    }
    /**
     * <code>string name = 2 [json_name = "name", (.buf.validate.field) = { ... }</code>
     * @return The bytes for name.
     */
    public com.google.protobuf.ByteString
        getNameBytes() {
      java.lang.Object ref = name_;
      if (ref instanceof String) {
        com.google.protobuf.ByteString b = 
            com.google.protobuf.ByteString.copyFromUtf8(
                (java.lang.String) ref);
        name_ = b;
        return b;
      } else {
        return (com.google.protobuf.ByteString) ref;
      }
    }
    /**
     * <code>string name = 2 [json_name = "name", (.buf.validate.field) = { ... }</code>
     * @param value The name to set.
     * @return This builder for chaining.
     */
    public Builder setName(
        java.lang.String value) {
      if (value == null) { throw new NullPointerException(); }
      name_ = value;
      bitField0_ |= 0x00000002;
      onChanged();
      return this;
    }
    /**
     * <code>string name = 2 [json_name = "name", (.buf.validate.field) = { ... }</code>
     * @return This builder for chaining.
     */
    public Builder clearName() {
      name_ = getDefaultInstance().getName();
      bitField0_ = (bitField0_ & ~0x00000002);
      onChanged();
      return this;
    }
    /**
     * <code>string name = 2 [json_name = "name", (.buf.validate.field) = { ... }</code>
     * @param value The bytes for name to set.
     * @return This builder for chaining.
     */
    public Builder setNameBytes(
        com.google.protobuf.ByteString value) {
      if (value == null) { throw new NullPointerException(); }
      checkByteStringIsUtf8(value);
      name_ = value;
      bitField0_ |= 0x00000002;
      onChanged();
      return this;
    }

    private java.lang.Object stateId_ = "";
    /**
     * <code>string state_id = 3 [json_name = "stateId", (.buf.validate.field) = { ... }</code>
     * @return The stateId.
     */
    public java.lang.String getStateId() {
      java.lang.Object ref = stateId_;
      if (!(ref instanceof java.lang.String)) {
        com.google.protobuf.ByteString bs =
            (com.google.protobuf.ByteString) ref;
        java.lang.String s = bs.toStringUtf8();
        stateId_ = s;
        return s;
      } else {
        return (java.lang.String) ref;
      }
    }
    /**
     * <code>string state_id = 3 [json_name = "stateId", (.buf.validate.field) = { ... }</code>
     * @return The bytes for stateId.
     */
    public com.google.protobuf.ByteString
        getStateIdBytes() {
      java.lang.Object ref = stateId_;
      if (ref instanceof String) {
        com.google.protobuf.ByteString b = 
            com.google.protobuf.ByteString.copyFromUtf8(
                (java.lang.String) ref);
        stateId_ = b;
        return b;
      } else {
        return (com.google.protobuf.ByteString) ref;
      }
    }
    /**
     * <code>string state_id = 3 [json_name = "stateId", (.buf.validate.field) = { ... }</code>
     * @param value The stateId to set.
     * @return This builder for chaining.
     */
    public Builder setStateId(
        java.lang.String value) {
      if (value == null) { throw new NullPointerException(); }
      stateId_ = value;
      bitField0_ |= 0x00000004;
      onChanged();
      return this;
    }
    /**
     * <code>string state_id = 3 [json_name = "stateId", (.buf.validate.field) = { ... }</code>
     * @return This builder for chaining.
     */
    public Builder clearStateId() {
      stateId_ = getDefaultInstance().getStateId();
      bitField0_ = (bitField0_ & ~0x00000004);
      onChanged();
      return this;
    }
    /**
     * <code>string state_id = 3 [json_name = "stateId", (.buf.validate.field) = { ... }</code>
     * @param value The bytes for stateId to set.
     * @return This builder for chaining.
     */
    public Builder setStateIdBytes(
        com.google.protobuf.ByteString value) {
      if (value == null) { throw new NullPointerException(); }
      checkByteStringIsUtf8(value);
      stateId_ = value;
      bitField0_ |= 0x00000004;
      onChanged();
      return this;
    }

    private double latitude_ ;
    /**
     * <code>double latitude = 4 [json_name = "latitude", (.buf.validate.field) = { ... }</code>
     * @return The latitude.
     */
    @java.lang.Override
    public double getLatitude() {
      return latitude_;
    }
    /**
     * <code>double latitude = 4 [json_name = "latitude", (.buf.validate.field) = { ... }</code>
     * @param value The latitude to set.
     * @return This builder for chaining.
     */
    public Builder setLatitude(double value) {

      latitude_ = value;
      bitField0_ |= 0x00000008;
      onChanged();
      return this;
    }
    /**
     * <code>double latitude = 4 [json_name = "latitude", (.buf.validate.field) = { ... }</code>
     * @return This builder for chaining.
     */
    public Builder clearLatitude() {
      bitField0_ = (bitField0_ & ~0x00000008);
      latitude_ = 0D;
      onChanged();
      return this;
    }

    private double longitude_ ;
    /**
     * <code>double longitude = 5 [json_name = "longitude", (.buf.validate.field) = { ... }</code>
     * @return The longitude.
     */
    @java.lang.Override
    public double getLongitude() {
      return longitude_;
    }
    /**
     * <code>double longitude = 5 [json_name = "longitude", (.buf.validate.field) = { ... }</code>
     * @param value The longitude to set.
     * @return This builder for chaining.
     */
    public Builder setLongitude(double value) {

      longitude_ = value;
      bitField0_ |= 0x00000010;
      onChanged();
      return this;
    }
    /**
     * <code>double longitude = 5 [json_name = "longitude", (.buf.validate.field) = { ... }</code>
     * @return This builder for chaining.
     */
    public Builder clearLongitude() {
      bitField0_ = (bitField0_ & ~0x00000010);
      longitude_ = 0D;
      onChanged();
      return this;
    }

    private boolean isActive_ ;
    /**
     * <code>bool is_active = 7 [json_name = "isActive", (.buf.validate.field) = { ... }</code>
     * @return The isActive.
     */
    @java.lang.Override
    public boolean getIsActive() {
      return isActive_;
    }
    /**
     * <code>bool is_active = 7 [json_name = "isActive", (.buf.validate.field) = { ... }</code>
     * @param value The isActive to set.
     * @return This builder for chaining.
     */
    public Builder setIsActive(boolean value) {

      isActive_ = value;
      bitField0_ |= 0x00000020;
      onChanged();
      return this;
    }
    /**
     * <code>bool is_active = 7 [json_name = "isActive", (.buf.validate.field) = { ... }</code>
     * @return This builder for chaining.
     */
    public Builder clearIsActive() {
      bitField0_ = (bitField0_ & ~0x00000020);
      isActive_ = false;
      onChanged();
      return this;
    }

    private long precedence_ ;
    /**
     * <code>int64 precedence = 6 [json_name = "precedence", (.buf.validate.field) = { ... }</code>
     * @return The precedence.
     */
    @java.lang.Override
    public long getPrecedence() {
      return precedence_;
    }
    /**
     * <code>int64 precedence = 6 [json_name = "precedence", (.buf.validate.field) = { ... }</code>
     * @param value The precedence to set.
     * @return This builder for chaining.
     */
    public Builder setPrecedence(long value) {

      precedence_ = value;
      bitField0_ |= 0x00000040;
      onChanged();
      return this;
    }
    /**
     * <code>int64 precedence = 6 [json_name = "precedence", (.buf.validate.field) = { ... }</code>
     * @return This builder for chaining.
     */
    public Builder clearPrecedence() {
      bitField0_ = (bitField0_ & ~0x00000040);
      precedence_ = 0L;
      onChanged();
      return this;
    }
    @java.lang.Override
    public final Builder setUnknownFields(
        final com.google.protobuf.UnknownFieldSet unknownFields) {
      return super.setUnknownFields(unknownFields);
    }

    @java.lang.Override
    public final Builder mergeUnknownFields(
        final com.google.protobuf.UnknownFieldSet unknownFields) {
      return super.mergeUnknownFields(unknownFields);
    }


    // @@protoc_insertion_point(builder_scope:core.city.v1.UpdateRequest)
  }

  // @@protoc_insertion_point(class_scope:core.city.v1.UpdateRequest)
  private static final com.core.city.v1.UpdateRequest DEFAULT_INSTANCE;
  static {
    DEFAULT_INSTANCE = new com.core.city.v1.UpdateRequest();
  }

  public static com.core.city.v1.UpdateRequest getDefaultInstance() {
    return DEFAULT_INSTANCE;
  }

  private static final com.google.protobuf.Parser<UpdateRequest>
      PARSER = new com.google.protobuf.AbstractParser<UpdateRequest>() {
    @java.lang.Override
    public UpdateRequest parsePartialFrom(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      Builder builder = newBuilder();
      try {
        builder.mergeFrom(input, extensionRegistry);
      } catch (com.google.protobuf.InvalidProtocolBufferException e) {
        throw e.setUnfinishedMessage(builder.buildPartial());
      } catch (com.google.protobuf.UninitializedMessageException e) {
        throw e.asInvalidProtocolBufferException().setUnfinishedMessage(builder.buildPartial());
      } catch (java.io.IOException e) {
        throw new com.google.protobuf.InvalidProtocolBufferException(e)
            .setUnfinishedMessage(builder.buildPartial());
      }
      return builder.buildPartial();
    }
  };

  public static com.google.protobuf.Parser<UpdateRequest> parser() {
    return PARSER;
  }

  @java.lang.Override
  public com.google.protobuf.Parser<UpdateRequest> getParserForType() {
    return PARSER;
  }

  @java.lang.Override
  public com.core.city.v1.UpdateRequest getDefaultInstanceForType() {
    return DEFAULT_INSTANCE;
  }

}

