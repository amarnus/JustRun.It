# JustRun.It

Share and execute code snippets in different programming languages

## Terminology

`Snippet` - A short piece of code in a programming language 

`run` - Run the code snippet 

`lint` - Lint check the code snippet 

`install` - Install module dependencies used by the code snippet for the current execution
context. Not saved across sessions

## Supported Languages

- Python
- Ruby
- PHP
- Node.js

## Features

- Share code snippets by URL
- Private snippets
- Execute snippets live
- Automatic dependency install for all languages except PHP
- STDOUT/STDERR streamed to the browser
- Language-specific Syntax Highlighting
- Editor feature: Ability to change theme
- PHP deps UI
- Forking
- Linting code snippet

## Architecture

This web service has 3 main components

### Potts

REST Server written in Go to serve snippets and schedule execution

[Code](https://github.com/gophergala/JustRun.It/tree/master/potts)

### Fury

REST Server written in Go to execute snippets in isolated docker environments.

Uses Go channels extensively to route STDOUT/STDERR from different snippet
executions to the appropriate websocket connection

[API Reference](https://github.com/gophergala/JustRun.It/tree/master/fury)

### Stark

Web application written in AngularJS featuring snippet editor, real time 
output streaming, lint checking and forking among many other features

[code](https://github.com/gophergala/JustRun.It/tree/master/stark)

