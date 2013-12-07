package httputil2

// HTTP methods
const (
	OPTIONS = "OPTIONS"
	GET     = "GET"
	HEAD    = "HEAD"
	POST    = "POST"
	PUT     = "PUT"
	DELETE  = "DELETE"
	TRACE   = "TRACE"
	CONNECT = "CONNECT"
)

// HTTP headers
//
// Everything is from RFC2616. Keep-Alive is not in the main list but kept for
// HTTP/1.0 compatibility.
const (
	HeaderAccept             = "Accept"
	HeaderAcceptCharset      = "Accept-Charset"
	HeaderAcceptEncoding     = "Accept-Encoding"
	HeaderAcceptLanguage     = "Accept-Language"
	HeaderAcceptRanges       = "Accept-Ranges"
	HeaderAge                = "Age"
	HeaderAllow              = "Allow"
	HeaderAuthorization      = "Authorization"
	HeaderCacheControl       = "Cache-Control"
	HeaderConnection         = "Connection"
	HeaderContentEncoding    = "Content-Encoding"
	HeaderContentLanguage    = "Content-Language"
	HeaderContentLength      = "Content-Length"
	HeaderContentLocation    = "Content-Location"
	HeaderContentMD5         = "Content-MD5"
	HeaderContentRange       = "Content-Range"
	HeaderContentType        = "Content-Type"
	HeaderDate               = "Date"
	HeaderETag               = "ETag"
	HeaderExpect             = "Expect"
	HeaderExpires            = "Expires"
	HeaderFrom               = "From"
	HeaderHost               = "Host"
	HeaderIfMatch            = "If-Match"
	HeaderIfModifiedSince    = "If-Modified-Since"
	HeaderIfNoneMatch        = "If-None-Match"
	HeaderIfRange            = "If-Range"
	HeaderIfUnmodifiedSince  = "If-Unmodified-Since"
	HeaderLastModified       = "Last-Modified"
	HeaderLocation           = "Location"
	HeaderMaxForwards        = "Max-Forwards"
	HeaderPragma             = "Pragma"
	HeaderRange              = "Range"
	HeaderReferer            = "Referer"
	HeaderRetryAfter         = "Retry-After"
	HeaderServer             = "Server"
	HeaderKeepAlive          = "Keep-Alive"
	HeaderProxyAuthenticate  = "Proxy-Authenticate"
	HeaderProxyAuthorization = "Proxy-Authorization"
	HeaderTE                 = "TE"
	HeaderTrailer            = "Trailer"
	HeaderTransferEncoding   = "Transfer-Encoding"
	HeaderUpgrade            = "Upgrade"
	HeaderUserAgent          = "User-Agent"
	HeaderVary               = "Vary"
	HeaderVia                = "Via"
	HeaderWarning            = "Warning"
	HeaderWWWAuthenticate    = "WWW-Authenticate"
)

// See http://www.w3.org/Protocols/rfc2616/rfc2616-sec13.html#sec13.5.1
func HopByHopHeaders() []string {
	return []string{
		HeaderConnection,
		HeaderKeepAlive,
		HeaderProxyAuthenticate,
		HeaderProxyAuthorization,
		HeaderTE,
		HeaderTrailer,
		HeaderTransferEncoding,
		HeaderUpgrade,
	}
}
