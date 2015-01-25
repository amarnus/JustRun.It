# JustRun.It

[Site](http://gophergala.justrun.it)

![Screenshot](https://raw.githubusercontent.com/gophergala/JustRun.It/master/static/screenshot.png)

## Motivation

JustRun.It is a web application that allows users to write, run and share code snippets in any language* from a web browser A.K.A JSBin for server-side programming languages.

Users are freed from the hassle of setting up a language runtime just for writing a short block of code in it. Developers can now share code along with its runnable execution context with their peers.

## Supported Programming Languages

- Python
- Ruby
- PHP
- Javascript (NodeJS runtime)

## Features

- Docker-powered Isolated execution environments for all snippet runs.
- Allow users to author snippets in a editor that supports Syntax Highlighting.
- Terminal Emulator which pipes the running snippet's STDOUT/STDERR streams to the browser. 
- Ability to *run* snippets and see the results, as they arrive, on the terminal. 
- Automatic dependency detection and installation for Python, Ruby and NodeJS snippets.
- Allow users to manually list their dependencies for PHP snippets.
- Allow users to fork other users' snippets as a starting point for their own.
- Allow users to see lint their snippets.
- Allow users to change the editor theme and persist the change across sessions.

## Architecture

JustRun.It is composed of three main components:

### [Potts](https://github.com/gophergala/JustRun.It/tree/master/potts)

Potts is a REST Server written in Go that manages JustRun.It snippets and their metadata.

### [Fury](https://github.com/gophergala/JustRun.It/tree/master/fury)

Fury is a REST Server written in Go that executes JustRun.It snippets in Docker containers.

It uses Go channels extensively to pipe the STDOUT/STDERR stream of the running snippet to the user's browser via WebSockets.

### [Stark](https://github.com/gophergala/JustRun.It/tree/master/stark)

Stark is a AngularJS-powered client application.
