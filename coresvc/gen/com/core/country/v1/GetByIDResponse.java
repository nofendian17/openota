// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: proto/country/v1/country_service.proto

// Protobuf Java Version: 3.25.1
package com.core.country.v1;

/**
 * Protobuf type {@code core.country.v1.GetByIDResponse}
 */
public final class GetByIDResponse extends
    com.google.protobuf.GeneratedMessageV3 implements
    // @@protoc_insertion_point(message_implements:core.country.v1.GetByIDResponse)
    GetByIDResponseOrBuilder {
private static final long serialVersionUID = 0L;
  // Use GetByIDResponse.newBuilder() to construct.
  private GetByIDResponse(com.google.protobuf.GeneratedMessageV3.Builder<?> builder) {
    super(builder);
  }
  private GetByIDResponse() {
  }

  @java.lang.Override
  @SuppressWarnings({"unused"})
  protected java.lang.Object newInstance(
      UnusedPrivateParameter unused) {
    return new GetByIDResponse();
  }

  public static final com.google.protobuf.Descriptors.Descriptor
      getDescriptor() {
    return com.core.country.v1.CountryServiceProto.internal_static_core_country_v1_GetByIDResponse_descriptor;
  }

  @java.lang.Override
  protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internalGetFieldAccessorTable() {
    return com.core.country.v1.CountryServiceProto.internal_static_core_country_v1_GetByIDResponse_fieldAccessorTable
        .ensureFieldAccessorsInitialized(
            com.core.country.v1.GetByIDResponse.class, com.core.country.v1.GetByIDResponse.Builder.class);
  }

  private int bitField0_;
  public static final int COUNTRY_FIELD_NUMBER = 1;
  private com.core.country.v1.Country country_;
  /**
   * <code>.core.country.v1.Country country = 1 [json_name = "country"];</code>
   * @return Whether the country field is set.
   */
  @java.lang.Override
  public boolean hasCountry() {
    return ((bitField0_ & 0x00000001) != 0);
  }
  /**
   * <code>.core.country.v1.Country country = 1 [json_name = "country"];</code>
   * @return The country.
   */
  @java.lang.Override
  public com.core.country.v1.Country getCountry() {
    return country_ == null ? com.core.country.v1.Country.getDefaultInstance() : country_;
  }
  /**
   * <code>.core.country.v1.Country country = 1 [json_name = "country"];</code>
   */
  @java.lang.Override
  public com.core.country.v1.CountryOrBuilder getCountryOrBuilder() {
    return country_ == null ? com.core.country.v1.Country.getDefaultInstance() : country_;
  }

  public static com.core.country.v1.GetByIDResponse parseFrom(
      java.nio.ByteBuffer data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static com.core.country.v1.GetByIDResponse parseFrom(
      java.nio.ByteBuffer data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static com.core.country.v1.GetByIDResponse parseFrom(
      com.google.protobuf.ByteString data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static com.core.country.v1.GetByIDResponse parseFrom(
      com.google.protobuf.ByteString data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static com.core.country.v1.GetByIDResponse parseFrom(byte[] data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static com.core.country.v1.GetByIDResponse parseFrom(
      byte[] data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static com.core.country.v1.GetByIDResponse parseFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static com.core.country.v1.GetByIDResponse parseFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input, extensionRegistry);
  }

  public static com.core.country.v1.GetByIDResponse parseDelimitedFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input);
  }

  public static com.core.country.v1.GetByIDResponse parseDelimitedFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input, extensionRegistry);
  }
  public static com.core.country.v1.GetByIDResponse parseFrom(
      com.google.protobuf.CodedInputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static com.core.country.v1.GetByIDResponse parseFrom(
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
  public static Builder newBuilder(com.core.country.v1.GetByIDResponse prototype) {
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
   * Protobuf type {@code core.country.v1.GetByIDResponse}
   */
  public static final class Builder extends
      com.google.protobuf.GeneratedMessageV3.Builder<Builder> implements
      // @@protoc_insertion_point(builder_implements:core.country.v1.GetByIDResponse)
      com.core.country.v1.GetByIDResponseOrBuilder {
    public static final com.google.protobuf.Descriptors.Descriptor
        getDescriptor() {
      return com.core.country.v1.CountryServiceProto.internal_static_core_country_v1_GetByIDResponse_descriptor;
    }

    @java.lang.Override
    protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
        internalGetFieldAccessorTable() {
      return com.core.country.v1.CountryServiceProto.internal_static_core_country_v1_GetByIDResponse_fieldAccessorTable
          .ensureFieldAccessorsInitialized(
              com.core.country.v1.GetByIDResponse.class, com.core.country.v1.GetByIDResponse.Builder.class);
    }

    // Construct using com.core.country.v1.GetByIDResponse.newBuilder()
    private Builder() {
      maybeForceBuilderInitialization();
    }

    private Builder(
        com.google.protobuf.GeneratedMessageV3.BuilderParent parent) {
      super(parent);
      maybeForceBuilderInitialization();
    }
    private void maybeForceBuilderInitialization() {
      if (com.google.protobuf.GeneratedMessageV3
              .alwaysUseFieldBuilders) {
        getCountryFieldBuilder();
      }
    }
    @java.lang.Override
    public Builder clear() {
      super.clear();
      bitField0_ = 0;
      country_ = null;
      if (countryBuilder_ != null) {
        countryBuilder_.dispose();
        countryBuilder_ = null;
      }
      return this;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.Descriptor
        getDescriptorForType() {
      return com.core.country.v1.CountryServiceProto.internal_static_core_country_v1_GetByIDResponse_descriptor;
    }

    @java.lang.Override
    public com.core.country.v1.GetByIDResponse getDefaultInstanceForType() {
      return com.core.country.v1.GetByIDResponse.getDefaultInstance();
    }

    @java.lang.Override
    public com.core.country.v1.GetByIDResponse build() {
      com.core.country.v1.GetByIDResponse result = buildPartial();
      if (!result.isInitialized()) {
        throw newUninitializedMessageException(result);
      }
      return result;
    }

    @java.lang.Override
    public com.core.country.v1.GetByIDResponse buildPartial() {
      com.core.country.v1.GetByIDResponse result = new com.core.country.v1.GetByIDResponse(this);
      if (bitField0_ != 0) { buildPartial0(result); }
      onBuilt();
      return result;
    }

    private void buildPartial0(com.core.country.v1.GetByIDResponse result) {
      int from_bitField0_ = bitField0_;
      int to_bitField0_ = 0;
      if (((from_bitField0_ & 0x00000001) != 0)) {
        result.country_ = countryBuilder_ == null
            ? country_
            : countryBuilder_.build();
        to_bitField0_ |= 0x00000001;
      }
      result.bitField0_ |= to_bitField0_;
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

    private com.core.country.v1.Country country_;
    private com.google.protobuf.SingleFieldBuilderV3<
        com.core.country.v1.Country, com.core.country.v1.Country.Builder, com.core.country.v1.CountryOrBuilder> countryBuilder_;
    /**
     * <code>.core.country.v1.Country country = 1 [json_name = "country"];</code>
     * @return Whether the country field is set.
     */
    public boolean hasCountry() {
      return ((bitField0_ & 0x00000001) != 0);
    }
    /**
     * <code>.core.country.v1.Country country = 1 [json_name = "country"];</code>
     * @return The country.
     */
    public com.core.country.v1.Country getCountry() {
      if (countryBuilder_ == null) {
        return country_ == null ? com.core.country.v1.Country.getDefaultInstance() : country_;
      } else {
        return countryBuilder_.getMessage();
      }
    }
    /**
     * <code>.core.country.v1.Country country = 1 [json_name = "country"];</code>
     */
    public Builder setCountry(com.core.country.v1.Country value) {
      if (countryBuilder_ == null) {
        if (value == null) {
          throw new NullPointerException();
        }
        country_ = value;
      } else {
        countryBuilder_.setMessage(value);
      }
      bitField0_ |= 0x00000001;
      onChanged();
      return this;
    }
    /**
     * <code>.core.country.v1.Country country = 1 [json_name = "country"];</code>
     */
    public Builder setCountry(
        com.core.country.v1.Country.Builder builderForValue) {
      if (countryBuilder_ == null) {
        country_ = builderForValue.build();
      } else {
        countryBuilder_.setMessage(builderForValue.build());
      }
      bitField0_ |= 0x00000001;
      onChanged();
      return this;
    }
    /**
     * <code>.core.country.v1.Country country = 1 [json_name = "country"];</code>
     */
    public Builder mergeCountry(com.core.country.v1.Country value) {
      if (countryBuilder_ == null) {
        if (((bitField0_ & 0x00000001) != 0) &&
          country_ != null &&
          country_ != com.core.country.v1.Country.getDefaultInstance()) {
          getCountryBuilder().mergeFrom(value);
        } else {
          country_ = value;
        }
      } else {
        countryBuilder_.mergeFrom(value);
      }
      if (country_ != null) {
        bitField0_ |= 0x00000001;
        onChanged();
      }
      return this;
    }
    /**
     * <code>.core.country.v1.Country country = 1 [json_name = "country"];</code>
     */
    public Builder clearCountry() {
      bitField0_ = (bitField0_ & ~0x00000001);
      country_ = null;
      if (countryBuilder_ != null) {
        countryBuilder_.dispose();
        countryBuilder_ = null;
      }
      onChanged();
      return this;
    }
    /**
     * <code>.core.country.v1.Country country = 1 [json_name = "country"];</code>
     */
    public com.core.country.v1.Country.Builder getCountryBuilder() {
      bitField0_ |= 0x00000001;
      onChanged();
      return getCountryFieldBuilder().getBuilder();
    }
    /**
     * <code>.core.country.v1.Country country = 1 [json_name = "country"];</code>
     */
    public com.core.country.v1.CountryOrBuilder getCountryOrBuilder() {
      if (countryBuilder_ != null) {
        return countryBuilder_.getMessageOrBuilder();
      } else {
        return country_ == null ?
            com.core.country.v1.Country.getDefaultInstance() : country_;
      }
    }
    /**
     * <code>.core.country.v1.Country country = 1 [json_name = "country"];</code>
     */
    private com.google.protobuf.SingleFieldBuilderV3<
        com.core.country.v1.Country, com.core.country.v1.Country.Builder, com.core.country.v1.CountryOrBuilder> 
        getCountryFieldBuilder() {
      if (countryBuilder_ == null) {
        countryBuilder_ = new com.google.protobuf.SingleFieldBuilderV3<
            com.core.country.v1.Country, com.core.country.v1.Country.Builder, com.core.country.v1.CountryOrBuilder>(
                getCountry(),
                getParentForChildren(),
                isClean());
        country_ = null;
      }
      return countryBuilder_;
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


    // @@protoc_insertion_point(builder_scope:core.country.v1.GetByIDResponse)
  }

  // @@protoc_insertion_point(class_scope:core.country.v1.GetByIDResponse)
  private static final com.core.country.v1.GetByIDResponse DEFAULT_INSTANCE;
  static {
    DEFAULT_INSTANCE = new com.core.country.v1.GetByIDResponse();
  }

  public static com.core.country.v1.GetByIDResponse getDefaultInstance() {
    return DEFAULT_INSTANCE;
  }

  private static final com.google.protobuf.Parser<GetByIDResponse>
      PARSER = new com.google.protobuf.AbstractParser<GetByIDResponse>() {
    @java.lang.Override
    public GetByIDResponse parsePartialFrom(
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

  public static com.google.protobuf.Parser<GetByIDResponse> parser() {
    return PARSER;
  }

  @java.lang.Override
  public com.google.protobuf.Parser<GetByIDResponse> getParserForType() {
    return PARSER;
  }

  @java.lang.Override
  public com.core.country.v1.GetByIDResponse getDefaultInstanceForType() {
    return DEFAULT_INSTANCE;
  }

}
