httputil2 - more go net/http extensions
=======================================


Recommended handler order
-------------------------

IdHandler - LogHandler - CacheHandler - GzipHandler - RecoverHandler - Application code
