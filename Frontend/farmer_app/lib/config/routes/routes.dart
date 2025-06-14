import 'package:farmer_app/config/routes/routes_name.dart';
import 'package:farmer_app/views/splash/splash_screen.dart';
import 'package:flutter/material.dart';
class Routes {
  static Route<dynamic> generateRoute(RouteSettings settings){
    switch(settings.name){
      case RoutesName.splashScreen:
       return MaterialPageRoute(builder: (context)=>SplashScreen());
      default:
       return MaterialPageRoute(builder: (context){
        return const Scaffold(
          body: Center(
            child: Text('No routes is generated'),
          ),
        );
       });
    }
  }
}