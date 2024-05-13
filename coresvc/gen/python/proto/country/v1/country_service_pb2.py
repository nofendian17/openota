# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/country/v1/country_service.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
from proto.country.v1 import country_pb2 as proto_dot_country_dot_v1_dot_country__pb2
from buf.validate import validate_pb2 as buf_dot_validate_dot_validate__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n&proto/country/v1/country_service.proto\x12\x0f\x63ore.country.v1\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1eproto/country/v1/country.proto\x1a\x1b\x62uf/validate/validate.proto\"-\n\x0eGetByIDRequest\x12\x1b\n\x02id\x18\x01 \x01(\tB\x0b\xbaH\x08\xc8\x01\x01r\x03\xb0\x01\x01R\x02id\"4\n\x10GetByCodeRequest\x12 \n\x04\x63ode\x18\x01 \x01(\tB\x0c\xbaH\t\xc8\x01\x01r\x04\x10\x03\x18\x03R\x04\x63ode\"\xd0\x02\n\rCreateRequest\x12 \n\x04\x63ode\x18\x01 \x01(\tB\x0c\xbaH\t\xc8\x01\x01r\x04\x10\x03\x18\x03R\x04\x63ode\x12!\n\x04name\x18\x02 \x01(\tB\r\xbaH\n\xc8\x01\x01r\x05\x10\x01\x18\xff\x01R\x04name\x12+\n\nphone_code\x18\x03 \x01(\tB\x0c\xbaH\t\xc8\x01\x01r\x04\x10\x02\x18\x05R\tphoneCode\x12\'\n\x07\x63\x61pital\x18\x04 \x01(\tB\r\xbaH\n\xc8\x01\x01r\x05\x10\x05\x18\xff\x01R\x07\x63\x61pital\x12\"\n\x08latitude\x18\x05 \x01(\x01\x42\x06\xbaH\x03\xc8\x01\x01R\x08latitude\x12$\n\tlongitude\x18\x06 \x01(\x01\x42\x06\xbaH\x03\xc8\x01\x01R\tlongitude\x12\x31\n\rcurrency_code\x18\x07 \x01(\tB\x0c\xbaH\t\xc8\x01\x01r\x04\x10\x03\x18\x03R\x0c\x63urrencyCode\x12\'\n\tis_active\x18\x08 \x01(\x08\x42\n\xbaH\x07\xc8\x01\x01j\x02\x08\x01R\x08isActive\"\x99\x03\n\rUpdateRequest\x12\x1b\n\x02id\x18\x01 \x01(\tB\x0b\xbaH\x08\xc8\x01\x01r\x03\xb0\x01\x01R\x02id\x12 \n\x04\x63ode\x18\x02 \x01(\tB\x0c\xbaH\t\xc8\x01\x01r\x04\x10\x03\x18\x03R\x04\x63ode\x12!\n\x04name\x18\x03 \x01(\tB\r\xbaH\n\xc8\x01\x01r\x05\x10\x01\x18\xff\x01R\x04name\x12+\n\nphone_code\x18\x04 \x01(\tB\x0c\xbaH\t\xc8\x01\x01r\x04\x10\x02\x18\x05R\tphoneCode\x12\'\n\x07\x63\x61pital\x18\x05 \x01(\tB\r\xbaH\n\xc8\x01\x01r\x05\x10\x05\x18\xff\x01R\x07\x63\x61pital\x12\"\n\x08latitude\x18\x06 \x01(\x01\x42\x06\xbaH\x03\xc8\x01\x01R\x08latitude\x12$\n\tlongitude\x18\x07 \x01(\x01\x42\x06\xbaH\x03\xc8\x01\x01R\tlongitude\x12\x31\n\rcurrency_code\x18\x08 \x01(\tB\x0c\xbaH\t\xc8\x01\x01r\x04\x10\x03\x18\x03R\x0c\x63urrencyCode\x12\'\n\tis_active\x18\t \x01(\x08\x42\n\xbaH\x07\xc8\x01\x01j\x02\x08\x01R\x08isActive\x12*\n\nprecedence\x18\n \x01(\x03\x42\n\xbaH\x07\xc8\x01\x01\"\x02(\x05R\nprecedence\",\n\rDeleteRequest\x12\x1b\n\x02id\x18\x01 \x01(\tB\x0b\xbaH\x08\xc8\x01\x01r\x03\xb0\x01\x01R\x02id\"E\n\x0fGetByIDResponse\x12\x32\n\x07\x63ountry\x18\x01 \x01(\x0b\x32\x18.core.country.v1.CountryR\x07\x63ountry\"G\n\x11GetByCodeResponse\x12\x32\n\x07\x63ountry\x18\x01 \x01(\x0b\x32\x18.core.country.v1.CountryR\x07\x63ountry\"H\n\x0eGetAllResponse\x12\x36\n\tcountries\x18\x01 \x03(\x0b\x32\x18.core.country.v1.CountryR\tcountries2\xc7\x03\n\x0e\x43ountryService\x12N\n\x07GetByID\x12\x1f.core.country.v1.GetByIDRequest\x1a .core.country.v1.GetByIDResponse\"\x00\x12T\n\tGetByCode\x12!.core.country.v1.GetByCodeRequest\x1a\".core.country.v1.GetByCodeResponse\"\x00\x12\x43\n\x06GetAll\x12\x16.google.protobuf.Empty\x1a\x1f.core.country.v1.GetAllResponse\"\x00\x12\x42\n\x06\x43reate\x12\x1e.core.country.v1.CreateRequest\x1a\x16.google.protobuf.Empty\"\x00\x12\x42\n\x06Update\x12\x1e.core.country.v1.UpdateRequest\x1a\x16.google.protobuf.Empty\"\x00\x12\x42\n\x06\x44\x65lete\x12\x1e.core.country.v1.DeleteRequest\x1a\x16.google.protobuf.Empty\"\x00\x42\xa6\x01\n\x13\x63om.core.country.v1B\x13\x43ountryServiceProtoH\x02P\x01Z\x1aproto/country/v1;countryv1\xa2\x02\x03\x43\x43X\xaa\x02\x0f\x43ore.Country.V1\xca\x02\x0f\x43ore\\Country\\V1\xe2\x02\x1b\x43ore\\Country\\V1\\GPBMetadata\xea\x02\x11\x43ore::Country::V1b\x06proto3')

_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, globals())
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'proto.country.v1.country_service_pb2', globals())
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'\n\023com.core.country.v1B\023CountryServiceProtoH\002P\001Z\032proto/country/v1;countryv1\242\002\003CCX\252\002\017Core.Country.V1\312\002\017Core\\Country\\V1\342\002\033Core\\Country\\V1\\GPBMetadata\352\002\021Core::Country::V1'
  _GETBYIDREQUEST.fields_by_name['id']._options = None
  _GETBYIDREQUEST.fields_by_name['id']._serialized_options = b'\272H\010\310\001\001r\003\260\001\001'
  _GETBYCODEREQUEST.fields_by_name['code']._options = None
  _GETBYCODEREQUEST.fields_by_name['code']._serialized_options = b'\272H\t\310\001\001r\004\020\003\030\003'
  _CREATEREQUEST.fields_by_name['code']._options = None
  _CREATEREQUEST.fields_by_name['code']._serialized_options = b'\272H\t\310\001\001r\004\020\003\030\003'
  _CREATEREQUEST.fields_by_name['name']._options = None
  _CREATEREQUEST.fields_by_name['name']._serialized_options = b'\272H\n\310\001\001r\005\020\001\030\377\001'
  _CREATEREQUEST.fields_by_name['phone_code']._options = None
  _CREATEREQUEST.fields_by_name['phone_code']._serialized_options = b'\272H\t\310\001\001r\004\020\002\030\005'
  _CREATEREQUEST.fields_by_name['capital']._options = None
  _CREATEREQUEST.fields_by_name['capital']._serialized_options = b'\272H\n\310\001\001r\005\020\005\030\377\001'
  _CREATEREQUEST.fields_by_name['latitude']._options = None
  _CREATEREQUEST.fields_by_name['latitude']._serialized_options = b'\272H\003\310\001\001'
  _CREATEREQUEST.fields_by_name['longitude']._options = None
  _CREATEREQUEST.fields_by_name['longitude']._serialized_options = b'\272H\003\310\001\001'
  _CREATEREQUEST.fields_by_name['currency_code']._options = None
  _CREATEREQUEST.fields_by_name['currency_code']._serialized_options = b'\272H\t\310\001\001r\004\020\003\030\003'
  _CREATEREQUEST.fields_by_name['is_active']._options = None
  _CREATEREQUEST.fields_by_name['is_active']._serialized_options = b'\272H\007\310\001\001j\002\010\001'
  _UPDATEREQUEST.fields_by_name['id']._options = None
  _UPDATEREQUEST.fields_by_name['id']._serialized_options = b'\272H\010\310\001\001r\003\260\001\001'
  _UPDATEREQUEST.fields_by_name['code']._options = None
  _UPDATEREQUEST.fields_by_name['code']._serialized_options = b'\272H\t\310\001\001r\004\020\003\030\003'
  _UPDATEREQUEST.fields_by_name['name']._options = None
  _UPDATEREQUEST.fields_by_name['name']._serialized_options = b'\272H\n\310\001\001r\005\020\001\030\377\001'
  _UPDATEREQUEST.fields_by_name['phone_code']._options = None
  _UPDATEREQUEST.fields_by_name['phone_code']._serialized_options = b'\272H\t\310\001\001r\004\020\002\030\005'
  _UPDATEREQUEST.fields_by_name['capital']._options = None
  _UPDATEREQUEST.fields_by_name['capital']._serialized_options = b'\272H\n\310\001\001r\005\020\005\030\377\001'
  _UPDATEREQUEST.fields_by_name['latitude']._options = None
  _UPDATEREQUEST.fields_by_name['latitude']._serialized_options = b'\272H\003\310\001\001'
  _UPDATEREQUEST.fields_by_name['longitude']._options = None
  _UPDATEREQUEST.fields_by_name['longitude']._serialized_options = b'\272H\003\310\001\001'
  _UPDATEREQUEST.fields_by_name['currency_code']._options = None
  _UPDATEREQUEST.fields_by_name['currency_code']._serialized_options = b'\272H\t\310\001\001r\004\020\003\030\003'
  _UPDATEREQUEST.fields_by_name['is_active']._options = None
  _UPDATEREQUEST.fields_by_name['is_active']._serialized_options = b'\272H\007\310\001\001j\002\010\001'
  _UPDATEREQUEST.fields_by_name['precedence']._options = None
  _UPDATEREQUEST.fields_by_name['precedence']._serialized_options = b'\272H\007\310\001\001\"\002(\005'
  _DELETEREQUEST.fields_by_name['id']._options = None
  _DELETEREQUEST.fields_by_name['id']._serialized_options = b'\272H\010\310\001\001r\003\260\001\001'
  _GETBYIDREQUEST._serialized_start=149
  _GETBYIDREQUEST._serialized_end=194
  _GETBYCODEREQUEST._serialized_start=196
  _GETBYCODEREQUEST._serialized_end=248
  _CREATEREQUEST._serialized_start=251
  _CREATEREQUEST._serialized_end=587
  _UPDATEREQUEST._serialized_start=590
  _UPDATEREQUEST._serialized_end=999
  _DELETEREQUEST._serialized_start=1001
  _DELETEREQUEST._serialized_end=1045
  _GETBYIDRESPONSE._serialized_start=1047
  _GETBYIDRESPONSE._serialized_end=1116
  _GETBYCODERESPONSE._serialized_start=1118
  _GETBYCODERESPONSE._serialized_end=1189
  _GETALLRESPONSE._serialized_start=1191
  _GETALLRESPONSE._serialized_end=1263
  _COUNTRYSERVICE._serialized_start=1266
  _COUNTRYSERVICE._serialized_end=1721
# @@protoc_insertion_point(module_scope)
