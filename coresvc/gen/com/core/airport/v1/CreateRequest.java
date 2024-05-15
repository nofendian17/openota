// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: proto/airport/v1/airport_service.proto

// Protobuf Java Version: 3.25.1
package com.core.airport.v1;

/**
 * Protobuf type {@code core.airport.v1.CreateRequest}
 */
public final class CreateRequest extends
    com.google.protobuf.GeneratedMessageV3 implements
    // @@protoc_insertion_point(message_implements:core.airport.v1.CreateRequest)
    CreateRequestOrBuilder {
private static final long serialVersionUID = 0L;
  // Use CreateRequest.newBuilder() to construct.
  private CreateRequest(com.google.protobuf.GeneratedMessageV3.Builder<?> builder) {
    super(builder);
  }
  private CreateRequest() {
    code_ = "";
    name_ = "";
    cityId_ = "";
    timezone_ = "";
  }

  @java.lang.Override
  @SuppressWarnings({"unused"})
  protected java.lang.Object newInstance(
      UnusedPrivateParameter unused) {
    return new CreateRequest();
  }

  public static final com.google.protobuf.Descriptors.Descriptor
      getDescriptor() {
    return com.core.airport.v1.AirportServiceProto.internal_static_core_airport_v1_CreateRequest_descriptor;
  }

  @java.lang.Override
  protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internalGetFieldAccessorTable() {
    return com.core.airport.v1.AirportServiceProto.internal_static_core_airport_v1_CreateRequest_fieldAccessorTable
        .ensureFieldAccessorsInitialized(
            com.core.airport.v1.CreateRequest.class, com.core.airport.v1.CreateRequest.Builder.class);
  }

  public static final int CODE_FIELD_NUMBER = 1;
  @SuppressWarnings("serial")
  private volatile java.lang.Object code_ = "";
  /**
   * <code>string code = 1 [json_name = "code", (.buf.validate.field) = { ... }</code>
   * @return The code.
   */
  @java.lang.Override
  public java.lang.String getCode() {
    java.lang.Object ref = code_;
    if (ref instanceof java.lang.String) {
      return (java.lang.String) ref;
    } else {
      com.google.protobuf.ByteString bs = 
          (com.google.protobuf.ByteString) ref;
      java.lang.String s = bs.toStringUtf8();
      code_ = s;
      return s;
    }
  }
  /**
   * <code>string code = 1 [json_name = "code", (.buf.validate.field) = { ... }</code>
   * @return The bytes for code.
   */
  @java.lang.Override
  public com.google.protobuf.ByteString
      getCodeBytes() {
    java.lang.Object ref = code_;
    if (ref instanceof java.lang.String) {
      com.google.protobuf.ByteString b = 
          com.google.protobuf.ByteString.copyFromUtf8(
              (java.lang.String) ref);
      code_ = b;
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

  public static final int CITY_ID_FIELD_NUMBER = 3;
  @SuppressWarnings("serial")
  private volatile java.lang.Object cityId_ = "";
  /**
   * <code>string city_id = 3 [json_name = "cityId", (.buf.validate.field) = { ... }</code>
   * @return The cityId.
   */
  @java.lang.Override
  public java.lang.String getCityId() {
    java.lang.Object ref = cityId_;
    if (ref instanceof java.lang.String) {
      return (java.lang.String) ref;
    } else {
      com.google.protobuf.ByteString bs = 
          (com.google.protobuf.ByteString) ref;
      java.lang.String s = bs.toStringUtf8();
      cityId_ = s;
      return s;
    }
  }
  /**
   * <code>string city_id = 3 [json_name = "cityId", (.buf.validate.field) = { ... }</code>
   * @return The bytes for cityId.
   */
  @java.lang.Override
  public com.google.protobuf.ByteString
      getCityIdBytes() {
    java.lang.Object ref = cityId_;
    if (ref instanceof java.lang.String) {
      com.google.protobuf.ByteString b = 
          com.google.protobuf.ByteString.copyFromUtf8(
              (java.lang.String) ref);
      cityId_ = b;
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

  public static final int IS_DOMESTIC_FIELD_NUMBER = 6;
  private boolean isDomestic_ = false;
  /**
   * <code>bool is_domestic = 6 [json_name = "isDomestic", (.buf.validate.field) = { ... }</code>
   * @return The isDomestic.
   */
  @java.lang.Override
  public boolean getIsDomestic() {
    return isDomestic_;
  }

  public static final int TIMEZONE_FIELD_NUMBER = 7;
  @SuppressWarnings("serial")
  private volatile java.lang.Object timezone_ = "";
  /**
   * <code>string timezone = 7 [json_name = "timezone", (.buf.validate.field) = { ... }</code>
   * @return The timezone.
   */
  @java.lang.Override
  public java.lang.String getTimezone() {
    java.lang.Object ref = timezone_;
    if (ref instanceof java.lang.String) {
      return (java.lang.String) ref;
    } else {
      com.google.protobuf.ByteString bs = 
          (com.google.protobuf.ByteString) ref;
      java.lang.String s = bs.toStringUtf8();
      timezone_ = s;
      return s;
    }
  }
  /**
   * <code>string timezone = 7 [json_name = "timezone", (.buf.validate.field) = { ... }</code>
   * @return The bytes for timezone.
   */
  @java.lang.Override
  public com.google.protobuf.ByteString
      getTimezoneBytes() {
    java.lang.Object ref = timezone_;
    if (ref instanceof java.lang.String) {
      com.google.protobuf.ByteString b = 
          com.google.protobuf.ByteString.copyFromUtf8(
              (java.lang.String) ref);
      timezone_ = b;
      return b;
    } else {
      return (com.google.protobuf.ByteString) ref;
    }
  }

  public static final int IS_ACTIVE_FIELD_NUMBER = 8;
  private boolean isActive_ = false;
  /**
   * <code>bool is_active = 8 [json_name = "isActive", (.buf.validate.field) = { ... }</code>
   * @return The isActive.
   */
  @java.lang.Override
  public boolean getIsActive() {
    return isActive_;
  }

  public static com.core.airport.v1.CreateRequest parseFrom(
      java.nio.ByteBuffer data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static com.core.airport.v1.CreateRequest parseFrom(
      java.nio.ByteBuffer data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static com.core.airport.v1.CreateRequest parseFrom(
      com.google.protobuf.ByteString data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static com.core.airport.v1.CreateRequest parseFrom(
      com.google.protobuf.ByteString data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static com.core.airport.v1.CreateRequest parseFrom(byte[] data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static com.core.airport.v1.CreateRequest parseFrom(
      byte[] data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static com.core.airport.v1.CreateRequest parseFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static com.core.airport.v1.CreateRequest parseFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input, extensionRegistry);
  }

  public static com.core.airport.v1.CreateRequest parseDelimitedFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input);
  }

  public static com.core.airport.v1.CreateRequest parseDelimitedFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input, extensionRegistry);
  }
  public static com.core.airport.v1.CreateRequest parseFrom(
      com.google.protobuf.CodedInputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static com.core.airport.v1.CreateRequest parseFrom(
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
  public static Builder newBuilder(com.core.airport.v1.CreateRequest prototype) {
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
   * Protobuf type {@code core.airport.v1.CreateRequest}
   */
  public static final class Builder extends
      com.google.protobuf.GeneratedMessageV3.Builder<Builder> implements
      // @@protoc_insertion_point(builder_implements:core.airport.v1.CreateRequest)
      com.core.airport.v1.CreateRequestOrBuilder {
    public static final com.google.protobuf.Descriptors.Descriptor
        getDescriptor() {
      return com.core.airport.v1.AirportServiceProto.internal_static_core_airport_v1_CreateRequest_descriptor;
    }

    @java.lang.Override
    protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
        internalGetFieldAccessorTable() {
      return com.core.airport.v1.AirportServiceProto.internal_static_core_airport_v1_CreateRequest_fieldAccessorTable
          .ensureFieldAccessorsInitialized(
              com.core.airport.v1.CreateRequest.class, com.core.airport.v1.CreateRequest.Builder.class);
    }

    // Construct using com.core.airport.v1.CreateRequest.newBuilder()
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
      code_ = "";
      name_ = "";
      cityId_ = "";
      latitude_ = 0D;
      longitude_ = 0D;
      isDomestic_ = false;
      timezone_ = "";
      isActive_ = false;
      return this;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.Descriptor
        getDescriptorForType() {
      return com.core.airport.v1.AirportServiceProto.internal_static_core_airport_v1_CreateRequest_descriptor;
    }

    @java.lang.Override
    public com.core.airport.v1.CreateRequest getDefaultInstanceForType() {
      return com.core.airport.v1.CreateRequest.getDefaultInstance();
    }

    @java.lang.Override
    public com.core.airport.v1.CreateRequest build() {
      com.core.airport.v1.CreateRequest result = buildPartial();
      if (!result.isInitialized()) {
        throw newUninitializedMessageException(result);
      }
      return result;
    }

    @java.lang.Override
    public com.core.airport.v1.CreateRequest buildPartial() {
      com.core.airport.v1.CreateRequest result = new com.core.airport.v1.CreateRequest(this);
      if (bitField0_ != 0) { buildPartial0(result); }
      onBuilt();
      return result;
    }

    private void buildPartial0(com.core.airport.v1.CreateRequest result) {
      int from_bitField0_ = bitField0_;
      if (((from_bitField0_ & 0x00000001) != 0)) {
        result.code_ = code_;
      }
      if (((from_bitField0_ & 0x00000002) != 0)) {
        result.name_ = name_;
      }
      if (((from_bitField0_ & 0x00000004) != 0)) {
        result.cityId_ = cityId_;
      }
      if (((from_bitField0_ & 0x00000008) != 0)) {
        result.latitude_ = latitude_;
      }
      if (((from_bitField0_ & 0x00000010) != 0)) {
        result.longitude_ = longitude_;
      }
      if (((from_bitField0_ & 0x00000020) != 0)) {
        result.isDomestic_ = isDomestic_;
      }
      if (((from_bitField0_ & 0x00000040) != 0)) {
        result.timezone_ = timezone_;
      }
      if (((from_bitField0_ & 0x00000080) != 0)) {
        result.isActive_ = isActive_;
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

    private java.lang.Object code_ = "";
    /**
     * <code>string code = 1 [json_name = "code", (.buf.validate.field) = { ... }</code>
     * @return The code.
     */
    public java.lang.String getCode() {
      java.lang.Object ref = code_;
      if (!(ref instanceof java.lang.String)) {
        com.google.protobuf.ByteString bs =
            (com.google.protobuf.ByteString) ref;
        java.lang.String s = bs.toStringUtf8();
        code_ = s;
        return s;
      } else {
        return (java.lang.String) ref;
      }
    }
    /**
     * <code>string code = 1 [json_name = "code", (.buf.validate.field) = { ... }</code>
     * @return The bytes for code.
     */
    public com.google.protobuf.ByteString
        getCodeBytes() {
      java.lang.Object ref = code_;
      if (ref instanceof String) {
        com.google.protobuf.ByteString b = 
            com.google.protobuf.ByteString.copyFromUtf8(
                (java.lang.String) ref);
        code_ = b;
        return b;
      } else {
        return (com.google.protobuf.ByteString) ref;
      }
    }
    /**
     * <code>string code = 1 [json_name = "code", (.buf.validate.field) = { ... }</code>
     * @param value The code to set.
     * @return This builder for chaining.
     */
    public Builder setCode(
        java.lang.String value) {
      if (value == null) { throw new NullPointerException(); }
      code_ = value;
      bitField0_ |= 0x00000001;
      onChanged();
      return this;
    }
    /**
     * <code>string code = 1 [json_name = "code", (.buf.validate.field) = { ... }</code>
     * @return This builder for chaining.
     */
    public Builder clearCode() {
      code_ = getDefaultInstance().getCode();
      bitField0_ = (bitField0_ & ~0x00000001);
      onChanged();
      return this;
    }
    /**
     * <code>string code = 1 [json_name = "code", (.buf.validate.field) = { ... }</code>
     * @param value The bytes for code to set.
     * @return This builder for chaining.
     */
    public Builder setCodeBytes(
        com.google.protobuf.ByteString value) {
      if (value == null) { throw new NullPointerException(); }
      checkByteStringIsUtf8(value);
      code_ = value;
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

    private java.lang.Object cityId_ = "";
    /**
     * <code>string city_id = 3 [json_name = "cityId", (.buf.validate.field) = { ... }</code>
     * @return The cityId.
     */
    public java.lang.String getCityId() {
      java.lang.Object ref = cityId_;
      if (!(ref instanceof java.lang.String)) {
        com.google.protobuf.ByteString bs =
            (com.google.protobuf.ByteString) ref;
        java.lang.String s = bs.toStringUtf8();
        cityId_ = s;
        return s;
      } else {
        return (java.lang.String) ref;
      }
    }
    /**
     * <code>string city_id = 3 [json_name = "cityId", (.buf.validate.field) = { ... }</code>
     * @return The bytes for cityId.
     */
    public com.google.protobuf.ByteString
        getCityIdBytes() {
      java.lang.Object ref = cityId_;
      if (ref instanceof String) {
        com.google.protobuf.ByteString b = 
            com.google.protobuf.ByteString.copyFromUtf8(
                (java.lang.String) ref);
        cityId_ = b;
        return b;
      } else {
        return (com.google.protobuf.ByteString) ref;
      }
    }
    /**
     * <code>string city_id = 3 [json_name = "cityId", (.buf.validate.field) = { ... }</code>
     * @param value The cityId to set.
     * @return This builder for chaining.
     */
    public Builder setCityId(
        java.lang.String value) {
      if (value == null) { throw new NullPointerException(); }
      cityId_ = value;
      bitField0_ |= 0x00000004;
      onChanged();
      return this;
    }
    /**
     * <code>string city_id = 3 [json_name = "cityId", (.buf.validate.field) = { ... }</code>
     * @return This builder for chaining.
     */
    public Builder clearCityId() {
      cityId_ = getDefaultInstance().getCityId();
      bitField0_ = (bitField0_ & ~0x00000004);
      onChanged();
      return this;
    }
    /**
     * <code>string city_id = 3 [json_name = "cityId", (.buf.validate.field) = { ... }</code>
     * @param value The bytes for cityId to set.
     * @return This builder for chaining.
     */
    public Builder setCityIdBytes(
        com.google.protobuf.ByteString value) {
      if (value == null) { throw new NullPointerException(); }
      checkByteStringIsUtf8(value);
      cityId_ = value;
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

    private boolean isDomestic_ ;
    /**
     * <code>bool is_domestic = 6 [json_name = "isDomestic", (.buf.validate.field) = { ... }</code>
     * @return The isDomestic.
     */
    @java.lang.Override
    public boolean getIsDomestic() {
      return isDomestic_;
    }
    /**
     * <code>bool is_domestic = 6 [json_name = "isDomestic", (.buf.validate.field) = { ... }</code>
     * @param value The isDomestic to set.
     * @return This builder for chaining.
     */
    public Builder setIsDomestic(boolean value) {

      isDomestic_ = value;
      bitField0_ |= 0x00000020;
      onChanged();
      return this;
    }
    /**
     * <code>bool is_domestic = 6 [json_name = "isDomestic", (.buf.validate.field) = { ... }</code>
     * @return This builder for chaining.
     */
    public Builder clearIsDomestic() {
      bitField0_ = (bitField0_ & ~0x00000020);
      isDomestic_ = false;
      onChanged();
      return this;
    }

    private java.lang.Object timezone_ = "";
    /**
     * <code>string timezone = 7 [json_name = "timezone", (.buf.validate.field) = { ... }</code>
     * @return The timezone.
     */
    public java.lang.String getTimezone() {
      java.lang.Object ref = timezone_;
      if (!(ref instanceof java.lang.String)) {
        com.google.protobuf.ByteString bs =
            (com.google.protobuf.ByteString) ref;
        java.lang.String s = bs.toStringUtf8();
        timezone_ = s;
        return s;
      } else {
        return (java.lang.String) ref;
      }
    }
    /**
     * <code>string timezone = 7 [json_name = "timezone", (.buf.validate.field) = { ... }</code>
     * @return The bytes for timezone.
     */
    public com.google.protobuf.ByteString
        getTimezoneBytes() {
      java.lang.Object ref = timezone_;
      if (ref instanceof String) {
        com.google.protobuf.ByteString b = 
            com.google.protobuf.ByteString.copyFromUtf8(
                (java.lang.String) ref);
        timezone_ = b;
        return b;
      } else {
        return (com.google.protobuf.ByteString) ref;
      }
    }
    /**
     * <code>string timezone = 7 [json_name = "timezone", (.buf.validate.field) = { ... }</code>
     * @param value The timezone to set.
     * @return This builder for chaining.
     */
    public Builder setTimezone(
        java.lang.String value) {
      if (value == null) { throw new NullPointerException(); }
      timezone_ = value;
      bitField0_ |= 0x00000040;
      onChanged();
      return this;
    }
    /**
     * <code>string timezone = 7 [json_name = "timezone", (.buf.validate.field) = { ... }</code>
     * @return This builder for chaining.
     */
    public Builder clearTimezone() {
      timezone_ = getDefaultInstance().getTimezone();
      bitField0_ = (bitField0_ & ~0x00000040);
      onChanged();
      return this;
    }
    /**
     * <code>string timezone = 7 [json_name = "timezone", (.buf.validate.field) = { ... }</code>
     * @param value The bytes for timezone to set.
     * @return This builder for chaining.
     */
    public Builder setTimezoneBytes(
        com.google.protobuf.ByteString value) {
      if (value == null) { throw new NullPointerException(); }
      checkByteStringIsUtf8(value);
      timezone_ = value;
      bitField0_ |= 0x00000040;
      onChanged();
      return this;
    }

    private boolean isActive_ ;
    /**
     * <code>bool is_active = 8 [json_name = "isActive", (.buf.validate.field) = { ... }</code>
     * @return The isActive.
     */
    @java.lang.Override
    public boolean getIsActive() {
      return isActive_;
    }
    /**
     * <code>bool is_active = 8 [json_name = "isActive", (.buf.validate.field) = { ... }</code>
     * @param value The isActive to set.
     * @return This builder for chaining.
     */
    public Builder setIsActive(boolean value) {

      isActive_ = value;
      bitField0_ |= 0x00000080;
      onChanged();
      return this;
    }
    /**
     * <code>bool is_active = 8 [json_name = "isActive", (.buf.validate.field) = { ... }</code>
     * @return This builder for chaining.
     */
    public Builder clearIsActive() {
      bitField0_ = (bitField0_ & ~0x00000080);
      isActive_ = false;
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


    // @@protoc_insertion_point(builder_scope:core.airport.v1.CreateRequest)
  }

  // @@protoc_insertion_point(class_scope:core.airport.v1.CreateRequest)
  private static final com.core.airport.v1.CreateRequest DEFAULT_INSTANCE;
  static {
    DEFAULT_INSTANCE = new com.core.airport.v1.CreateRequest();
  }

  public static com.core.airport.v1.CreateRequest getDefaultInstance() {
    return DEFAULT_INSTANCE;
  }

  private static final com.google.protobuf.Parser<CreateRequest>
      PARSER = new com.google.protobuf.AbstractParser<CreateRequest>() {
    @java.lang.Override
    public CreateRequest parsePartialFrom(
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

  public static com.google.protobuf.Parser<CreateRequest> parser() {
    return PARSER;
  }

  @java.lang.Override
  public com.google.protobuf.Parser<CreateRequest> getParserForType() {
    return PARSER;
  }

  @java.lang.Override
  public com.core.airport.v1.CreateRequest getDefaultInstanceForType() {
    return DEFAULT_INSTANCE;
  }

}
