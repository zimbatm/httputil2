httputil2 - more go net/http extensions
=======================================

The httputil2 package contains common handlers and utilities to use in
combination with the net/http package from go's standard library.


Recommended handler order
-------------------------

RequestIDMiddleware - LogMiddleware - GzipMiddleware - RecoveryMiddleware - Application code
