///
//  Generated code. Do not modify.
//  source: iss.proto
///
// ignore_for_file: non_constant_identifier_names,library_prefixes,unused_import

const LocRequest$json = const {
  '1': 'LocRequest',
  '2': const [
    const {'1': 'message', '3': 1, '4': 1, '5': 9, '10': 'message'},
  ],
};

const LocReply$json = const {
  '1': 'LocReply',
  '2': const [
    const {'1': 'longitude', '3': 1, '4': 1, '5': 9, '10': 'longitude'},
    const {'1': 'latitude', '3': 2, '4': 1, '5': 9, '10': 'latitude'},
  ],
};

const LocDetectorServiceBase$json = const {
  '1': 'LocDetector',
  '2': const [
    const {'1': 'GetLoc', '2': '.iss.LocRequest', '3': '.iss.LocReply', '4': const {}},
  ],
};

const LocDetectorServiceBase$messageJson = const {
  '.iss.LocRequest': LocRequest$json,
  '.iss.LocReply': LocReply$json,
};

