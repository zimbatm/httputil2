# httputil2 [![Build Status](https://travis-ci.org/zimbatm/httputil2.svg?branch=master)](https://travis-ci.org/zimbatm/httputil2) [![GoDoc](https://godoc.org/github.com/zimbatm/httputil2?status.svg)](http://godoc.org/github.com/zimbatm/httputil2)

What should really be in golang's net/http/httputil package.

A collection of net/http utilities and middlewares to build your own stack.

Just pick and choose whatever is useful to you!

## Recommended middleware order

From the list of avaiable middlewares, here is the recommended order:

* RequestIDMiddleware (left-most)
* LogMiddleware
* CleanPathMiddleware
* GzipMiddleware
* RecoveryMiddleware
* Application code

## License

All files in this projects are licensed under the ISC. See the LICENSE file
for more details.
