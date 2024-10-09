import 'package:flutter/material.dart';

class BaseAppLayout extends StatefulWidget {
  final Widget? child;

  const BaseAppLayout({super.key, this.child});

  @override
  BaseAppLayoutState createState() => BaseAppLayoutState();
}

class BaseAppLayoutState extends State<BaseAppLayout> {
  ThemeMode _themeMode = ThemeMode.system;

  void _toggleTheme() {
    setState(() {
      _themeMode = _themeMode == ThemeMode.light ? ThemeMode.dark : ThemeMode.light;
    });
  }

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      theme: ThemeData(
        appBarTheme: AppBarTheme(
          color: Color.fromARGB(255, 0, 122, 255),
        ),
      ),
      darkTheme: ThemeData(
        appBarTheme: AppBarTheme(
          color: Color.fromARGB(255, 0, 0, 0),
        ),
      ),
      themeMode: _themeMode,
      home: Scaffold(
        appBar: AppBar(
          title: const Text('Base App Layout'),
          actions: [
            IconButton(
              icon: Icon(_themeMode == ThemeMode.light ? Icons.dark_mode : Icons.light_mode),
              onPressed: _toggleTheme,
            ),
          ],
        ),
        body: widget.child,
      ),
    );
  }
}