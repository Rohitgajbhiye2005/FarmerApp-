import 'package:flutter/material.dart';

class InternetException extends StatelessWidget {
  const InternetException({super.key});

  @override
  Widget build(BuildContext context) {
    return Column(
      children: [
        SizedBox(height: MediaQuery.sizeOf(context).height * .15),
        Icon(Icons.cloud_off,
        color:Colors.red,
        size: 50,
        )
      ],
    );
  }
}