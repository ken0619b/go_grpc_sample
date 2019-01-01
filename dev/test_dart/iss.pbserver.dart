///
//  Generated code. Do not modify.
//  source: iss.proto
///
// ignore_for_file: non_constant_identifier_names,library_prefixes,unused_import

import 'dart:async' as $async;

import 'package:protobuf/protobuf.dart';

import 'iss.pb.dart';
import 'iss.pbjson.dart';

export 'iss.pb.dart';

abstract class LocDetectorServiceBase extends GeneratedService {
  $async.Future<LocReply> getLoc(ServerContext ctx, LocRequest request);

  GeneratedMessage createRequest(String method) {
    switch (method) {
      case 'GetLoc': return new LocRequest();
      default: throw new ArgumentError('Unknown method: $method');
    }
  }

  $async.Future<GeneratedMessage> handleCall(ServerContext ctx, String method, GeneratedMessage request) {
    switch (method) {
      case 'GetLoc': return this.getLoc(ctx, request);
      default: throw new ArgumentError('Unknown method: $method');
    }
  }

  Map<String, dynamic> get $json => LocDetectorServiceBase$json;
  Map<String, Map<String, dynamic>> get $messageJson => LocDetectorServiceBase$messageJson;
}

