// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: proto/airline/v1/airline.proto

// Protobuf Java Version: 3.25.1
package com.core.airline.v1;

public final class AirlineProto {
  private AirlineProto() {}
  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistryLite registry) {
  }

  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistry registry) {
    registerAllExtensions(
        (com.google.protobuf.ExtensionRegistryLite) registry);
  }
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_core_airline_v1_Airline_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_core_airline_v1_Airline_fieldAccessorTable;

  public static com.google.protobuf.Descriptors.FileDescriptor
      getDescriptor() {
    return descriptor;
  }
  private static  com.google.protobuf.Descriptors.FileDescriptor
      descriptor;
  static {
    java.lang.String[] descriptorData = {
      "\n\036proto/airline/v1/airline.proto\022\017core.a" +
      "irline.v1\032\033google/protobuf/empty.proto\032\037" +
      "google/protobuf/timestamp.proto\"\210\002\n\007Airl" +
      "ine\022\016\n\002id\030\001 \001(\tR\002id\022\022\n\004name\030\002 \001(\tR\004name\022" +
      "\022\n\004code\030\003 \001(\tR\004code\022\022\n\004logo\030\004 \001(\tR\004logo\022" +
      "\033\n\tis_active\030\005 \001(\010R\010isActive\022\036\n\npreceden" +
      "ce\030\006 \001(\003R\nprecedence\0229\n\ncreated_at\030\007 \001(\013" +
      "2\032.google.protobuf.TimestampR\tcreatedAt\022" +
      "9\n\nupdated_at\030\010 \001(\0132\032.google.protobuf.Ti" +
      "mestampR\tupdatedAtB\237\001\n\023com.core.airline." +
      "v1B\014AirlineProtoH\002P\001Z\032proto/airline/v1;a" +
      "irlinev1\242\002\003CAX\252\002\017Core.Airline.V1\312\002\017Core\\" +
      "Airline\\V1\342\002\033Core\\Airline\\V1\\GPBMetadata" +
      "\352\002\021Core::Airline::V1b\006proto3"
    };
    descriptor = com.google.protobuf.Descriptors.FileDescriptor
      .internalBuildGeneratedFileFrom(descriptorData,
        new com.google.protobuf.Descriptors.FileDescriptor[] {
          com.google.protobuf.EmptyProto.getDescriptor(),
          com.google.protobuf.TimestampProto.getDescriptor(),
        });
    internal_static_core_airline_v1_Airline_descriptor =
      getDescriptor().getMessageTypes().get(0);
    internal_static_core_airline_v1_Airline_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_core_airline_v1_Airline_descriptor,
        new java.lang.String[] { "Id", "Name", "Code", "Logo", "IsActive", "Precedence", "CreatedAt", "UpdatedAt", });
    com.google.protobuf.EmptyProto.getDescriptor();
    com.google.protobuf.TimestampProto.getDescriptor();
  }

  // @@protoc_insertion_point(outer_class_scope)
}
