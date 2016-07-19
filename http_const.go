package httputil2

import (
	"net/http"
)

// HTTP methods
const (
	CONNECT = http.MethodConnect
	DELETE  = http.MethodDelete
	GET     = http.MethodGet
	HEAD    = http.MethodHead
	OPTIONS = http.MethodOptions
	PATCH   = http.MethodPatch // RFC 5789
	POST    = http.MethodPost
	PUT     = http.MethodPut
	TRACE   = http.MethodTrace
)

// HTTP headers
//
// Everything is from RFC2616 + CORS + RFC 5741 + De-facto extensions
const (
	HeaderAccept                  = "Accept"
	HeaderAcceptCharset           = "Accept-Charset"
	HeaderAcceptEncoding          = "Accept-Encoding"
	HeaderAcceptLanguage          = "Accept-Language"
	HeaderAcceptRanges            = "Accept-Ranges"
	HeaderAcceptPatch             = "Accept-Patch" // RFC 5789
	HeaderAge                     = "Age"
	HeaderAllow                   = "Allow"
	HeaderAllowPatch              = "Allow-Patch" // RFC 5741
	HeaderAuthorization           = "Authorization"
	HeaderCacheControl            = "Cache-Control"
	HeaderConnection              = "Connection"
	HeaderContentDisposition      = "Content-Disposition" // RFC 6266
	HeaderContentEncoding         = "Content-Encoding"
	HeaderContentLanguage         = "Content-Language"
	HeaderContentLength           = "Content-Length"
	HeaderContentLocation         = "Content-Location"
	HeaderContentRange            = "Content-Range"
	HeaderContentType             = "Content-Type"
	HeaderCookie                  = "Cookie"
	HeaderDate                    = "Date"
	HeaderETag                    = "ETag"
	HeaderExpect                  = "Expect"
	HeaderExpires                 = "Expires"
	HeaderFrom                    = "From"
	HeaderHost                    = "Host"
	HeaderIfMatch                 = "If-Match"
	HeaderIfModifiedSince         = "If-Modified-Since"
	HeaderIfNoneMatch             = "If-None-Match"
	HeaderIfRange                 = "If-Range"
	HeaderIfUnmodifiedSince       = "If-Unmodified-Since"
	HeaderKeepAlive               = "Keep-Alive"
	HeaderLastModified            = "Last-Modified"
	HeaderLocation                = "Location"
	HeaderMaxForwards             = "Max-Forwards"
	HeaderPragma                  = "Pragma"
	HeaderProxyAuthenticate       = "Proxy-Authenticate"
	HeaderProxyAuthorization      = "Proxy-Authorization"
	HeaderRange                   = "Range"
	HeaderReferer                 = "Referer"
	HeaderRetryAfter              = "Retry-After"
	HeaderServer                  = "Server"
	HeaderSetCookie               = "Set-Cookie"
	HeaderStrictTransportSecurity = "Strict-Transport-Security" // HSTS
	HeaderTE                      = "TE"
	HeaderTrailer                 = "Trailer"
	HeaderTransferEncoding        = "Transfer-Encoding"
	HeaderUpgrade                 = "Upgrade"
	HeaderUserAgent               = "User-Agent"
	HeaderVary                    = "Vary"
	HeaderVia                     = "Via"
	HeaderWWWAuthenticate         = "WWW-Authenticate"
	HeaderWarning                 = "Warning"
	HeaderXForwardedFor           = "X-Forwarded-For"
	HeaderXForwardedProto         = "X-Forwarded-Proto"
	HeaderXFrameOptions           = "X-Frame-Options" // RFC 7034
	HeaderXRequestID              = "X-Request-ID"
)

// See http://www.w3.org/Protocols/rfc2616/rfc2616-sec13.html#sec13.5.1
var HopByHopHeaders = []string{
	HeaderConnection,
	HeaderKeepAlive,
	HeaderProxyAuthenticate,
	HeaderProxyAuthorization,
	HeaderTE,
	HeaderTrailer,
	HeaderTransferEncoding,
	HeaderUpgrade,
}

var RequestHeaders = []string{
	HeaderAccept,
	HeaderAcceptCharset,
	HeaderAcceptEncoding,
	HeaderAcceptLanguage,
	HeaderAcceptRanges,
	HeaderAuthorization,
	HeaderCacheControl,
	HeaderConnection,
	HeaderContentLength,
	HeaderContentType,
	HeaderCookie,
	HeaderDate,
	HeaderFrom,
	HeaderHost,
	HeaderIfMatch,
	HeaderIfModifiedSince,
	HeaderIfNoneMatch,
	HeaderIfRange,
	HeaderIfUnmodifiedSince,
	HeaderMaxForwards,
	HeaderMaxForwards,
	HeaderPragma,
	HeaderProxyAuthorization,
	HeaderRange,
	HeaderTE,
	HeaderUpgrade,
	HeaderUserAgent,
	HeaderVia,
	HeaderWarning,
	HeaderXForwardedFor,
	HeaderXForwardedProto,
	HeaderXRequestID,
}

var ResponseHeaders = []string{
	HeaderAcceptPatch,
	HeaderAcceptRanges,
	HeaderAge,
	HeaderAllow,
	HeaderCacheControl,
	HeaderConnection,
	HeaderContentDisposition,
	HeaderContentEncoding,
	HeaderContentLanguage,
	HeaderContentLength,
	HeaderContentLocation,
	HeaderContentRange,
	HeaderContentType,
	HeaderDate,
	HeaderETag,
	HeaderExpires,
	HeaderLastModified,
	HeaderLocation,
	HeaderPragma,
	HeaderProxyAuthenticate,
	HeaderRange,
	HeaderTrailer,
	HeaderTransferEncoding,
	HeaderVary,
	HeaderVia,
	HeaderWWWAuthenticate,
	HeaderWarning,
	HeaderXFrameOptions,
	HeaderAccessControlAllowMethods,
	HeaderRetryAfter,
	HeaderServer,
	HeaderSetCookie,
	HeaderStrictTransportSecurity,
}

const (
	StatusContinue           = 100
	StatusSwitchingProtocols = 101

	StatusOK                   = 200
	StatusCreated              = 201
	StatusAccepted             = 202
	StatusNonAuthoritativeInfo = 203
	StatusNoContent            = 204
	StatusResetContent         = 205
	StatusPartialContent       = 206

	StatusMultipleChoices   = 300
	StatusMovedPermanently  = 301
	StatusFound             = 302
	StatusSeeOther          = 303
	StatusNotModified       = 304
	StatusUseProxy          = 305
	StatusTemporaryRedirect = 307

	StatusBadRequest                   = 400
	StatusUnauthorized                 = 401
	StatusPaymentRequired              = 402
	StatusForbidden                    = 403
	StatusNotFound                     = 404
	StatusMethodNotAllowed             = 405
	StatusNotAcceptable                = 406
	StatusProxyAuthRequired            = 407
	StatusRequestTimeout               = 408
	StatusConflict                     = 409
	StatusGone                         = 410
	StatusLengthRequired               = 411
	StatusPreconditionFailed           = 412
	StatusRequestEntityTooLarge        = 413
	StatusRequestURITooLong            = 414
	StatusUnsupportedMediaType         = 415
	StatusRequestedRangeNotSatisfiable = 416
	StatusExpectationFailed            = 417
	StatusTeapot                       = 418
	StatusPreconditionRequired         = 428
	StatusTooManyRequests              = 429
	StatusRequestHeaderFieldsTooLarge  = 431
	StatusUnavailableForLegalReasons   = 451

	StatusInternalServerError           = 500
	StatusNotImplemented                = 501
	StatusBadGateway                    = 502
	StatusServiceUnavailable            = 503
	StatusGatewayTimeout                = 504
	StatusHTTPVersionNotSupported       = 505
	StatusNetworkAuthenticationRequired = 511
)
