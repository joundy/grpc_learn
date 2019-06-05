// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var service_pb = require('./service_pb.js');

function serialize_proto_Request(arg) {
  if (!(arg instanceof service_pb.Request)) {
    throw new Error('Expected argument of type proto.Request');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_proto_Request(buffer_arg) {
  return service_pb.Request.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_proto_Response(arg) {
  if (!(arg instanceof service_pb.Response)) {
    throw new Error('Expected argument of type proto.Response');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_proto_Response(buffer_arg) {
  return service_pb.Response.deserializeBinary(new Uint8Array(buffer_arg));
}


var AddServiceService = exports.AddServiceService = {
  multiply: {
    path: '/proto.AddService/Multiply',
    requestStream: false,
    responseStream: false,
    requestType: service_pb.Request,
    responseType: service_pb.Response,
    requestSerialize: serialize_proto_Request,
    requestDeserialize: deserialize_proto_Request,
    responseSerialize: serialize_proto_Response,
    responseDeserialize: deserialize_proto_Response,
  },
};

exports.AddServiceClient = grpc.makeGenericClientConstructor(AddServiceService);
