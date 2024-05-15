// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: proto/state/v1/state_service.proto

// Protobuf Java Version: 3.25.1
package com.core.state.v1;

public final class StateServiceProto {
  private StateServiceProto() {}
  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistryLite registry) {
  }

  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistry registry) {
    registerAllExtensions(
        (com.google.protobuf.ExtensionRegistryLite) registry);
  }
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_core_state_v1_GetByIDRequest_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_core_state_v1_GetByIDRequest_fieldAccessorTable;
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_core_state_v1_CreateRequest_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_core_state_v1_CreateRequest_fieldAccessorTable;
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_core_state_v1_UpdateRequest_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_core_state_v1_UpdateRequest_fieldAccessorTable;
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_core_state_v1_DeleteRequest_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_core_state_v1_DeleteRequest_fieldAccessorTable;
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_core_state_v1_GetByIDResponse_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_core_state_v1_GetByIDResponse_fieldAccessorTable;
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_core_state_v1_GetAllResponse_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_core_state_v1_GetAllResponse_fieldAccessorTable;

  public static com.google.protobuf.Descriptors.FileDescriptor
      getDescriptor() {
    return descriptor;
  }
  private static  com.google.protobuf.Descriptors.FileDescriptor
      descriptor;
  static {
    java.lang.String[] descriptorData = {
      "\n\"proto/state/v1/state_service.proto\022\rco" +
      "re.state.v1\032\033google/protobuf/empty.proto" +
      "\032\032proto/state/v1/state.proto\032\033buf/valida" +
      "te/validate.proto\"-\n\016GetByIDRequest\022\033\n\002i" +
      "d\030\001 \001(\tB\013\272H\010r\003\260\001\001\310\001\001R\002id\"\321\001\n\rCreateReque" +
      "st\022!\n\004name\030\001 \001(\tB\r\272H\nr\005\020\001\030\377\001\310\001\001R\004name\022*\n" +
      "\ncountry_id\030\002 \001(\tB\013\272H\010r\003\260\001\001\310\001\001R\tcountryI" +
      "d\022\"\n\010latitude\030\003 \001(\001B\006\272H\003\310\001\001R\010latitude\022$\n" +
      "\tlongitude\030\004 \001(\001B\006\272H\003\310\001\001R\tlongitude\022\'\n\ti" +
      "s_active\030\005 \001(\010B\n\272H\007j\002\010\001\310\001\001R\010isActive\"\232\002\n" +
      "\rUpdateRequest\022\033\n\002id\030\001 \001(\tB\013\272H\010r\003\260\001\001\310\001\001R" +
      "\002id\022!\n\004name\030\002 \001(\tB\r\272H\nr\005\020\001\030\377\001\310\001\001R\004name\022*" +
      "\n\ncountry_id\030\003 \001(\tB\013\272H\010r\003\260\001\001\310\001\001R\tcountry" +
      "Id\022\"\n\010latitude\030\004 \001(\001B\006\272H\003\310\001\001R\010latitude\022$" +
      "\n\tlongitude\030\005 \001(\001B\006\272H\003\310\001\001R\tlongitude\022\'\n\t" +
      "is_active\030\007 \001(\010B\n\272H\007j\002\010\001\310\001\001R\010isActive\022*\n" +
      "\nprecedence\030\006 \001(\003B\n\272H\007\"\002(\000\310\001\001R\nprecedenc" +
      "e\",\n\rDeleteRequest\022\033\n\002id\030\001 \001(\tB\013\272H\010r\003\260\001\001" +
      "\310\001\001R\002id\"=\n\017GetByIDResponse\022*\n\005state\030\001 \001(" +
      "\0132\024.core.state.v1.StateR\005state\">\n\016GetAll" +
      "Response\022,\n\006states\030\001 \003(\0132\024.core.state.v1" +
      ".StateR\006states2\343\002\n\014StateService\022J\n\007GetBy" +
      "ID\022\035.core.state.v1.GetByIDRequest\032\036.core" +
      ".state.v1.GetByIDResponse\"\000\022A\n\006GetAll\022\026." +
      "google.protobuf.Empty\032\035.core.state.v1.Ge" +
      "tAllResponse\"\000\022@\n\006Create\022\034.core.state.v1" +
      ".CreateRequest\032\026.google.protobuf.Empty\"\000" +
      "\022@\n\006Update\022\034.core.state.v1.UpdateRequest" +
      "\032\026.google.protobuf.Empty\"\000\022@\n\006Delete\022\034.c" +
      "ore.state.v1.DeleteRequest\032\026.google.prot" +
      "obuf.Empty\"\000B\226\001\n\021com.core.state.v1B\021Stat" +
      "eServiceProtoH\002P\001Z\026proto/state/v1;statev" +
      "1\242\002\003CSX\252\002\rCore.State.V1\312\002\rCore\\State\\V1\342" +
      "\002\031Core\\State\\V1\\GPBMetadata\352\002\017Core::Stat" +
      "e::V1b\006proto3"
    };
    descriptor = com.google.protobuf.Descriptors.FileDescriptor
      .internalBuildGeneratedFileFrom(descriptorData,
        new com.google.protobuf.Descriptors.FileDescriptor[] {
          com.google.protobuf.EmptyProto.getDescriptor(),
          com.core.state.v1.StateProto.getDescriptor(),
          com.buf.validate.ValidateProto.getDescriptor(),
        });
    internal_static_core_state_v1_GetByIDRequest_descriptor =
      getDescriptor().getMessageTypes().get(0);
    internal_static_core_state_v1_GetByIDRequest_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_core_state_v1_GetByIDRequest_descriptor,
        new java.lang.String[] { "Id", });
    internal_static_core_state_v1_CreateRequest_descriptor =
      getDescriptor().getMessageTypes().get(1);
    internal_static_core_state_v1_CreateRequest_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_core_state_v1_CreateRequest_descriptor,
        new java.lang.String[] { "Name", "CountryId", "Latitude", "Longitude", "IsActive", });
    internal_static_core_state_v1_UpdateRequest_descriptor =
      getDescriptor().getMessageTypes().get(2);
    internal_static_core_state_v1_UpdateRequest_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_core_state_v1_UpdateRequest_descriptor,
        new java.lang.String[] { "Id", "Name", "CountryId", "Latitude", "Longitude", "IsActive", "Precedence", });
    internal_static_core_state_v1_DeleteRequest_descriptor =
      getDescriptor().getMessageTypes().get(3);
    internal_static_core_state_v1_DeleteRequest_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_core_state_v1_DeleteRequest_descriptor,
        new java.lang.String[] { "Id", });
    internal_static_core_state_v1_GetByIDResponse_descriptor =
      getDescriptor().getMessageTypes().get(4);
    internal_static_core_state_v1_GetByIDResponse_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_core_state_v1_GetByIDResponse_descriptor,
        new java.lang.String[] { "State", });
    internal_static_core_state_v1_GetAllResponse_descriptor =
      getDescriptor().getMessageTypes().get(5);
    internal_static_core_state_v1_GetAllResponse_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_core_state_v1_GetAllResponse_descriptor,
        new java.lang.String[] { "States", });
    com.google.protobuf.ExtensionRegistry registry =
        com.google.protobuf.ExtensionRegistry.newInstance();
    registry.add(com.buf.validate.ValidateProto.field);
    com.google.protobuf.Descriptors.FileDescriptor
        .internalUpdateFileDescriptor(descriptor, registry);
    com.google.protobuf.EmptyProto.getDescriptor();
    com.core.state.v1.StateProto.getDescriptor();
    com.buf.validate.ValidateProto.getDescriptor();
  }

  // @@protoc_insertion_point(outer_class_scope)
}
