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

