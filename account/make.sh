#!/bin/bash

# protoc -I=../ -I=../third_party -I=../third_party/github.com/gogo/protobuf -I=../third_party/github.com/gogo/googleapis -I=. --gogofast_out=plugins=grpc:. api.proto account.proto
# protoc -I=../ -I=../third_party -I=../third_party/github.com/gogo/protobuf -I=../third_party/github.com/gogo/googleapis -I=. --gin-xproto_out=:. http.proto
# sed -i "s/account\.//g" http.gin.go


# sleep 3

# go generate

protoc -I=../ -I=../third_party -I=../third_party/github.com/gogo/protobuf -I=../third_party/github.com/gogo/googleapis -I=. --gogofast_out=plugins=grpc,Mprotobuf/google/protobuf/any.proto=github.com/gogo/protobuf/types,Mprotobuf/google/protobuf/api.proto=github.com/gogo/protobuf/types,Mprotobuf/google/protobuf/descriptor.proto=github.com/gogo/protobuf/types,Mprotobuf/google/protobuf/duration.proto=github.com/gogo/protobuf/types,Mprotobuf/google/protobuf/empty.proto=github.com/gogo/protobuf/types,Mprotobuf/google/protobuf/field_mask.proto=github.com/gogo/protobuf/types,Mprotobuf/google/protobuf/source_context.proto=github.com/gogo/protobuf/types,Mprotobuf/google/protobuf/struct.proto=github.com/gogo/protobuf/types,Mprotobuf/google/protobuf/timestamp.proto=github.com/gogo/protobuf/types,Mprotobuf/google/protobuf/type.proto=github.com/gogo/protobuf/types,Mprotobuf/google/protobuf/wrappers.proto=github.com/gogo/protobuf/types,Mgoogle/api/annotations.proto=github.com/gogo/googleapis/google/api:./ ./*.proto

protoc -I=../ -I=../third_party -I=../third_party/github.com/gogo/protobuf -I=../third_party/github.com/gogo/googleapis -I=. --gin_out=Mprotobuf/google/protobuf/any.proto=github.com/gogo/protobuf/types,Mprotobuf/google/protobuf/api.proto=github.com/gogo/protobuf/types,Mprotobuf/google/protobuf/descriptor.proto=github.com/gogo/protobuf/types,Mprotobuf/google/protobuf/duration.proto=github.com/gogo/protobuf/types,Mprotobuf/google/protobuf/empty.proto=github.com/gogo/protobuf/types,Mprotobuf/google/protobuf/field_mask.proto=github.com/gogo/protobuf/types,Mprotobuf/google/protobuf/source_context.proto=github.com/gogo/protobuf/types,Mprotobuf/google/protobuf/struct.proto=github.com/gogo/protobuf/types,Mprotobuf/google/protobuf/timestamp.proto=github.com/gogo/protobuf/types,Mprotobuf/google/protobuf/type.proto=github.com/gogo/protobuf/types,Mprotobuf/google/protobuf/wrappers.proto=github.com/gogo/protobuf/types,Mgoogle/api/annotations.proto=github.com/gogo/googleapis/google/api:./ ./api.proto
