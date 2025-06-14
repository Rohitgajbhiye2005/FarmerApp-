import 'dart:async';
import 'dart:convert';
import 'dart:io';

import 'package:farmer_app/data/Exception/app_exception.dart';
import 'package:farmer_app/data/network/base_api_service.dart';
import 'package:http/http.dart'as http;

class NetworkApiService implements BaseApiService{

  @override
  Future<dynamic> getApi(String url) async{
    dynamic jsonResponse;
    try {
      final response=await http.get(Uri.parse(url)).timeout(Duration(seconds: 50));
     jsonResponse=returnResponse(response);
    }on SocketException{
      throw NoInternetException();
    }
    on TimeoutException{
      throw RequesttimeoutException();
    }
    return jsonResponse;
  }

  @override
  Future<dynamic> postApi(String url,var data) async{
 dynamic jsonResponse;
    try {
      final response=await http.post(Uri.parse(url),body: data).timeout(Duration(seconds: 50));
     jsonResponse=returnResponse(response);
    }on SocketException{
      throw NoInternetException();
    }
    on TimeoutException{
      throw RequesttimeoutException();
    }
    return jsonResponse;
  }


  dynamic returnResponse(http.Response response){
    switch(response){
      // ignore: constant_pattern_never_matches_value_type
      case 200:
        dynamic jsonResponse=jsonDecode(response.body);
        return jsonResponse;
      // ignore: constant_pattern_never_matches_value_type
      case 400:
        dynamic jsonResponse=jsonDecode(response.body);
        return jsonResponse;
      // ignore: constant_pattern_never_matches_value_type
      case 500:
        throw FetchDataException('error communicationg with server ${response.statusCode}');
    }
  }

}