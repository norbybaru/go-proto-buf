# frozen_string_literal: true
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: user/v1/user.proto

require 'google/protobuf'


descriptor_data = "\n\x12user/v1/user.proto\x12\x07user.v1\"V\n\x04User\x12\x12\n\x04uuid\x18\x01 \x01(\tR\x04uuid\x12\x1b\n\tfull_name\x18\x02 \x01(\tR\x08\x66ullName\x12\x1d\n\nbirth_year\x18\x03 \x01(\x03R\tbirthYearB:Z8github.com/norbybaru/go-proto-buf/gren/go/user/v1;userpbb\x06proto3"

pool = Google::Protobuf::DescriptorPool.generated_pool

begin
  pool.add_serialized_file(descriptor_data)
rescue TypeError
  # Compatibility code: will be removed in the next major version.
  require 'google/protobuf/descriptor_pb'
  parsed = Google::Protobuf::FileDescriptorProto.decode(descriptor_data)
  parsed.clear_dependency
  serialized = parsed.class.encode(parsed)
  file = pool.add_serialized_file(serialized)
  warn "Warning: Protobuf detected an import path issue while loading generated file #{__FILE__}"
  imports = [
  ]
  imports.each do |type_name, expected_filename|
    import_file = pool.lookup(type_name).file_descriptor
    if import_file.name != expected_filename
      warn "- #{file.name} imports #{expected_filename}, but that import was loaded as #{import_file.name}"
    end
  end
  warn "Each proto file must use a consistent fully-qualified name."
  warn "This will become an error in the next major version."
end

module User
  module V1
    User = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("user.v1.User").msgclass
  end
end