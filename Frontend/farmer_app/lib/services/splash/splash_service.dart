import 'dart:async';

import 'package:farmer_app/config/routes/routes_name.dart';
import 'package:flutter/widgets.dart';

class SplashService {


  void isLogin(BuildContext context){
    Timer(const Duration(seconds: 3),
      ()=>Navigator.pushNamedAndRemoveUntil(context, RoutesName.loginScreen, (route)=>false)
    );
  }
}