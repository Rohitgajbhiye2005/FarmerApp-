class AppException implements Exception{
  final _message;
  final _prefix;
  AppException([this._message,this._prefix]);

  @override
  String toString() {
    return '$_message$_prefix';
  }
}

class NoInternetException extends AppException{
  NoInternetException([String? message]):super(message,'No internet');
}

class UnAuthException extends AppException{
  UnAuthException([String? message]):super(message,'No internet');
}

class RequesttimeoutException extends AppException{
  RequesttimeoutException([String? message]):super(message,'No internet');
}

class FetchDataException extends AppException{
  FetchDataException([String? message]):super(message,'');
}