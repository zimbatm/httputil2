package httputil2

import (
	"net/http"
	"time"
	"strings"
	"fmt"
)

const (
	CommonLogFormat  = `%h %l %u %t "%r" %>s %b`
	CommonDateFormat = "02/Jan/2006:15:04:05 -0700"
)

// Implements the Apache Common Log Format formatting for the LogHandler
// See: https://httpd.apache.org/docs/1.3/logs.html
// Example:
//   CommonLogFormatter(CommonLogFormat)
func CommonLogFormatter(format string) LogFormatter {
	return &commonLogFormatter{format}
}

type commonLogFormatter struct {
	format string
}

func (self *commonLogFormatter) RequestLog(r *http.Request, start time.Time) string {
	return ""
}

func (self *commonLogFormatter) ResponseLog(r *http.Request, start time.Time, status int, bytes int) string {
	line := self.format
	// %h is the IP address of the remote host
	line = strings.Replace(line, "%h", r.RemoteAddr, -1)
	// %l is the ident and will not be implemented
	line = strings.Replace(line, "%l", "-", -1)
	// TODO: %u is the REMOTE_USER
	line = strings.Replace(line, "%u", "-", -1)
	// %t 10/Oct/2000:13:55:36 -0700
	line = strings.Replace(line, "%t", start.Format(CommonDateFormat), -1)
	// %r GET /apache_pb.gif HTTP/1.0
	line = strings.Replace(line, "%r",
		fmt.Sprintf("%s %s %s", r.Method, r.RequestURI, r.Proto), -1)
	// %>s HTTP status
	line = strings.Replace(line, "%>s", fmt.Sprintf("%d", status), -1)

	b := "-"
	B := fmt.Sprintf("%d", bytes)
	if bytes > 0 {
		b = B
	}
	// %b size of the body. - for no content
	line = strings.Replace(line, "%b", b, -1)
	// %B like %b. 0 for no content
	line = strings.Replace(line, "%B", B, -1)

	return line + "\n"
}
