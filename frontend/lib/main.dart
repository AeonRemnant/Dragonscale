import 'package:flutter/material.dart';
import 'util/base.dart';

void main() {
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    return BaseAppLayout(
      child: Scaffold(
        appBar: AppBar(
          title: const Text('Main Page'),
        ),
      ),
    );
  }
}
