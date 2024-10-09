import 'package:flutter/material.dart';
import 'util/base.dart';

void main() {
  runApp(MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    return BaseAppLayout(
      child: Scaffold(
        appBar: AppBar(
          title: Text('Main Page'),
        ),
        child: Text('Welcome to Dragonscale!'),
      ),
    );
  }
}
