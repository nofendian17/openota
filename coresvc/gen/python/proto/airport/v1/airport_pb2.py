# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/airport/v1/airport.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1eproto/airport/v1/airport.proto\x12\x0f\x63ore.airport.v1\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"\x84\x03\n\x07\x41irport\x12\x0e\n\x02id\x18\x01 \x01(\tR\x02id\x12\x12\n\x04\x63ode\x18\x02 \x01(\tR\x04\x63ode\x12\x12\n\x04name\x18\x03 \x01(\tR\x04name\x12\x17\n\x07\x63ity_id\x18\x04 \x01(\tR\x06\x63ityId\x12\x1a\n\x08latitude\x18\x05 \x01(\x01R\x08latitude\x12\x1c\n\tlongitude\x18\x06 \x01(\x01R\tlongitude\x12\x1f\n\x0bis_domestic\x18\x07 \x01(\x08R\nisDomestic\x12\x1a\n\x08timezone\x18\x08 \x01(\tR\x08timezone\x12\x1b\n\tis_active\x18\t \x01(\x08R\x08isActive\x12\x1e\n\nprecedence\x18\n \x01(\x03R\nprecedence\x12\x39\n\ncreated_at\x18\x0b \x01(\x0b\x32\x1a.google.protobuf.TimestampR\tcreatedAt\x12\x39\n\nupdated_at\x18\x0c \x01(\x0b\x32\x1a.google.protobuf.TimestampR\tupdatedAtB\x9f\x01\n\x13\x63om.core.airport.v1B\x0c\x41irportProtoH\x02P\x01Z\x1aproto/airport/v1;airportv1\xa2\x02\x03\x43\x41X\xaa\x02\x0f\x43ore.Airport.V1\xca\x02\x0f\x43ore\\Airport\\V1\xe2\x02\x1b\x43ore\\Airport\\V1\\GPBMetadata\xea\x02\x11\x43ore::Airport::V1b\x06proto3')

_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, globals())
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'proto.airport.v1.airport_pb2', globals())
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'\n\023com.core.airport.v1B\014AirportProtoH\002P\001Z\032proto/airport/v1;airportv1\242\002\003CAX\252\002\017Core.Airport.V1\312\002\017Core\\Airport\\V1\342\002\033Core\\Airport\\V1\\GPBMetadata\352\002\021Core::Airport::V1'
  _AIRPORT._serialized_start=114
  _AIRPORT._serialized_end=502
# @@protoc_insertion_point(module_scope)