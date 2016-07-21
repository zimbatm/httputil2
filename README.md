httputil2 - more go net/http extensions
=======================================

[![Build Status](https://travis-ci.org/zimbatm/httputil2.svg?branch=master)](https://travis-ci.org/zimbatm/httputil2)

The httputil2 package contains common handlers and utilities to use in
combination with the net/http package from go's standard library.


Recommended handler order
-------------------------

RequestIDMiddleware - LogMiddleware - GzipMiddleware - RecoveryMiddleware - Application code
