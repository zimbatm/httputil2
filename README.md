httputil2 - more go net/http extensions
=======================================

[![Build Status](https://travis-ci.org/zimbatm/httputil2.svg?branch=master)](https://travis-ci.org/zimbatm/httputil2)

The httputil2 package contains common handlers and utilities to use in
combination with the net/http package from go's standard library.

Just pick and choose whatever is useful to you!

## Recommended middleware order

From the list of avaiable middlewares, here is the recommended order:

* RequestIDMiddleware (left-most)
* LogMiddleware
* GzipMiddleware
* RecoveryMiddleware
* Application code

## Technical documentation

Find all the goodies over at:

https://godoc.org/github.com/zimbatm/httputil2

