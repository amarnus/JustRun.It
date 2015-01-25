# JustRunit Fury Service

Snippet execution manager in all Suit instances

## Terminology

`Suit` - An instance running snippets

`sid` - Session ID of execution

## API Overview

### [Run](#run-1)

* [`POST /run/complete`](#post-runcomplete)
* [`POST /run`](#post-run)

### [Lint](#lint-1)

* [`POST /lint/complete`](#post-lintcomplete)

### [Websocket](#websocket-1)

* [`ws /ws/io`](#ws-wsio)

## API Reference

### Run

#### POST /run/complete

Run a snippet, wait for it to complete, and collect output from STDOUT and STDERR

##### input parameters

`language` - **string** - A language supported by JustRunIt

`uid` - **string** - UID of the script to execute 

`sid` - **string** - Session ID of the execution 

`snippet` - **string** - Code snippet to execute 

##### example request

	$ curl -H "Content-Type: application/json" -X POST -d \
            '{
               "language": "python",
               "uid": "python_uid_23432adfs2",
               "sid": "python_session_254fdagt",
               "snippet": "print \"random\""
             }' http://localhost:3000/run/complete | json_xs

##### example output

```javascript
{
   "status" : 1,
   "result" : [
      "New python executable in env/bin/python2",
      "Not overwriting existing python script env/bin/python (you must use env/bin/python2)",
      "Installing setuptools, pip...done.",
      "Running virtualenv with interpreter /usr/bin/python2",
      "You must give at least one requirement to install (see \"pip help install\")",
      "random",
      "",
      "Stderr",
      ""
   ]
}
```

#### POST /run

Start a snippet and pipe all its output to STDOUT/STDERR to a websocket requesting the same session id 

##### input parameters

`language` - **string** - A language supported by JustRunIt

`uid` - **string** - UID of the script to execute 

`sid` - **string** - Session ID of the execution 

`snippet` - **string** - Code snippet to execute 

##### example request

	$ curl -H "Content-Type: application/json" -X POST -d \
            '{
               "language": "python",
               "uid": "python_uid_23432adfs2",
               "sid": "python_session_254fdagt",
               "snippet": "print \"random\""
             }' http://localhost:3000/run | json_xs

##### example output

```javascript
{
   "status" : 1,
}
```

### Lint

#### POST /lint/complete

Lint check a snippet, wait for it to complete, and collect output from STDOUT and STDERR

##### input parameters

`language` - **string** - A language supported by JustRunIt

`uid` - **string** - UID of the script to execute 

`sid` - **string** - Session ID of the execution 

`snippet` - **string** - Code snippet to execute 

##### example request

	$ curl -H "Content-Type: application/json" -X POST -d \
            '{
               "language": "python",
               "uid": "python_uid_23432adfs2",
               "sid": "python_session_254fdagt",
               "snippet": "print \"random\""
             }' http://localhost:3000/lint/complete | json_xs

##### example output

```javascript
{
   "status" : 0,
   "result" : [
      "Stderr",
      "code:2:17: invalid syntax",
      "function subtract(number1, number2) {",
      "                 ^ ",
   ]
}
```

### Websocket

#### ws /ws/io

Websocket to provide IO to a running snippet

A single session can run multiple operations. All operations' IO
are routed to the session's websocket by the session ID.

The listener needs to just listen to this server's websocket
at its Session ID and provide any input need to its snippet in
the same websocket

##### input parameters

`id` - **string** - Session ID to listen to

##### example request

	$ /usr/lib/node_modules/ws/wscat/bin/wscat -c ws://localhost:3000/ws/io --origin http://localhost:3000
	$ > {"id":"python_session_5678"}

##### example output

A single websocket stream event

```javascript
{
   "id" : "python_session_5678",
   "data" : "00524 lines in /env/lib/python2.7/site-packages/setuptools/tests/test_sdist.py"
}
```

