package httputil2

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

const (
	CommonLogFormat  = `%h %l %u %t "%r" %>s %b`
	CommonDateFormat = "02/Jan/2006:15:04:05 -0700"
)

// Implements the Apache Common Log Format formatting for the LogMiddleware
// See: https://httpd.apache.org/docs/1.3/logs.html
// Example:
//   CommonLog(os.Stdout)
//
// FIXME: Tokenize the string upfront
func CommonLog(w io.Writer) *commonLog {
	return &commonLog{w, CommonLogFormat}
}

var _ HTTPLogger = new(commonLog)

type commonLog struct {
	w      io.Writer
	format string
}

// Changes the pattern to use to write the logs to
func (l *commonLog) SetFormat(format string) {
	l.format = format
}

func (_ *commonLog) LogRequest(r *http.Request, start time.Time) {
	return
}

func (l *commonLog) LogResponse(r *http.Request, start time.Time, status int, bytes int) {
	line := l.Format(r, start, status, bytes)
	l.w.Write([]byte(line))
}

func (l *commonLog) Format(r *http.Request, start time.Time, status int, bytes int) string {
	line := l.format
	// %h is the IP address of the remote host
	line = strings.Replace(line, "%h", r.RemoteAddr, -1)
	// %l is the ident and will not be implemented
	line = strings.Replace(line, "%l", "-", -1)
	// %u is the REMOTE_USER
	{
		var user string
		u := r.URL.User
		if u != nil && u.Username() != "" {
			user = u.Username()
		} else {
			user = "-"
		}
		line = strings.Replace(line, "%u", user, -1)
	}
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
