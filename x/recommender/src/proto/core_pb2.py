# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/core.proto
# Protobuf Python Version: 4.25.0
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x10proto/core.proto\x12\x04\x63ore\"\x15\n\x07Trigger\x12\n\n\x02id\x18\x01 \x01(\t\"+\n\x08Response\x12\x0f\n\x07message\x18\x01 \x01(\t\x12\x0e\n\x06status\x18\x02 \x01(\x08\"\x1a\n\x07Request\x12\x0f\n\x07message\x18\x01 \x01(\t2\xa8\x01\n\x06\x45ngine\x12.\n\x0bHealthCheck\x12\r.core.Request\x1a\x0e.core.Response\"\x00\x12\x35\n\x12ProcessContentBase\x12\r.core.Trigger\x1a\x0e.core.Response\"\x00\x12\x37\n\x14ProcessCollaborative\x12\r.core.Trigger\x1a\x0e.core.Response\"\x00\x42\x07Z\x05/coreb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'proto.core_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
  _globals['DESCRIPTOR']._options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z\005/core'
  _globals['_TRIGGER']._serialized_start=26
  _globals['_TRIGGER']._serialized_end=47
  _globals['_RESPONSE']._serialized_start=49
  _globals['_RESPONSE']._serialized_end=92
  _globals['_REQUEST']._serialized_start=94
  _globals['_REQUEST']._serialized_end=120
  _globals['_ENGINE']._serialized_start=123
  _globals['_ENGINE']._serialized_end=291
# @@protoc_insertion_point(module_scope)
