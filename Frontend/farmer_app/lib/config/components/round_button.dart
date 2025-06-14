import 'package:flutter/material.dart';

class RoundButton extends StatelessWidget {
  final String title;
  final VoidCallback onpress;
  // ignore: use_super_parameters
  const RoundButton({Key? key,required this.title,required this.onpress}):super(key: key);

  @override
  Widget build(BuildContext context) {
    return InkWell(
      onTap: onpress,
      child: Container(
        height: 50,
        width: 150,
        decoration: BoxDecoration(
          color: Colors.green,
          borderRadius: BorderRadius.circular(20)
        ),
        child: Center(
          child: Text(
            title
          ),
        ),
      ),
    );
  }
}