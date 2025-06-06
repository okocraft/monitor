package ctxlib

import (
	"context"
	"log/slog"
	"net"
	"net/http"
	"runtime"
	"time"

	"github.com/okocraft/monitor/internal/domain/auth"
	"github.com/okocraft/monitor/internal/domain/user"
)

type HTTPAccessLog struct {
	Timestamp     time.Time
	Method        string
	URL           string
	ContentLength int64
	Proto         string
	Host          string
	RemoteAddr    string
	UserAgent     string
	RequestURI    string
	Referer       string

	AuthMethod auth.Method

	UserID user.ID

	Response *HTTPResponse
}

func (a *HTTPAccessLog) FromRequest(r *http.Request) {
	a.Timestamp = time.Now()
	a.Method = r.Method
	a.Proto = r.Proto
	a.ContentLength = r.ContentLength
	a.Host = r.Host
	if r.URL != nil {
		a.URL = r.URL.String()
	}
	a.RemoteAddr = r.RemoteAddr
	a.UserAgent = r.UserAgent()
	a.RequestURI = r.RequestURI
	a.Referer = r.Referer()
}

func (a *HTTPAccessLog) ToAttr() slog.Attr {
	attrs := make([]any, 0, 4)

	attrs = append(attrs, slog.Group("request",
		slog.String("timestamp", a.Timestamp.Format(time.RFC3339)),
		slog.String("method", a.Method),
		slog.String("url", a.URL),
		slog.String("full_url", a.Host+a.RequestURI)),
		slog.Int64("content_length", a.ContentLength),
		slog.String("proto", a.Proto),
		slog.String("host", a.Host),
		slog.String("remote_addr", a.RemoteAddr),
		slog.String("user_agent", a.UserAgent),
		slog.String("request_uri", a.RequestURI),
		slog.String("referer", a.Referer),
	)

	attrs = append(attrs, slog.Int("auth_method", int(a.AuthMethod)))
	attrs = append(attrs, slog.Int("user_id", int(a.UserID)))

	if a.Response != nil {
		attrs = append(attrs, a.Response.ToAttr())
		attrs = append(attrs, slog.Float64("took", float64(a.Response.FinishedAt.Sub(a.Timestamp).Nanoseconds())/1e6))
	}

	return slog.Group("http_access", attrs...)
}

func (a *HTTPAccessLog) GetIP() net.IP {
	host, _, err := net.SplitHostPort(a.RemoteAddr)
	if err != nil {
		return []byte{}
	}

	ip := net.ParseIP(host)
	if ip == nil {
		return []byte{}
	}

	if ipv4 := ip.To4(); ipv4 != nil {
		return ipv4
	}

	return ip.To16()
}

func InitHTTPAccessLog(ctx context.Context) (context.Context, *HTTPAccessLog) {
	a := HTTPAccessLog{}
	ctx = context.WithValue(ctx, accessLogKey, &a)
	return ctx, &a
}

func GetHTTPAccessLog(ctx context.Context) *HTTPAccessLog {
	a, ok := ctx.Value(accessLogKey).(*HTTPAccessLog)
	if ok {
		return a
	}
	return &HTTPAccessLog{}
}

func GetHTTPResponse(ctx context.Context) *HTTPResponse {
	a := GetHTTPAccessLog(ctx)
	if a.Response == nil {
		a.Response = &HTTPResponse{}
	}
	return a.Response
}

type HTTPResponse struct {
	StatusCode int
	Error      error

	Handler HTTPHandlerInfo

	ResponseSize int
	FinishedAt   time.Time
}

func (res *HTTPResponse) ToAttr() slog.Attr {
	attrs := make([]any, 0, 6)

	attrs = append(attrs, slog.String("timestamp", res.FinishedAt.Format(time.RFC3339)))
	attrs = append(attrs, slog.Int("status_code", res.StatusCode))
	attrs = append(attrs, slog.Int("size", res.ResponseSize))

	if res.Error != nil {
		attrs = append(attrs, slog.String("error", res.Error.Error()))
	}

	attrs = append(attrs, slog.Group("handler",
		slog.String("func_name", res.Handler.FuncName),
		slog.String("file", res.Handler.File),
		slog.Int("line", res.Handler.Line),
	))

	return slog.Group("response", attrs...)
}

func (res *HTTPResponse) WriteHandlerInfo(skip int) {
	pc, file, line, ok := runtime.Caller(skip)
	if !ok {
		return
	}

	fn := runtime.FuncForPC(pc)
	if fn == nil {
		res.Handler = HTTPHandlerInfo{
			File: file,
			Line: line,
		}
		return
	}

	res.Handler = HTTPHandlerInfo{
		FuncName: fn.Name(),
		File:     file,
		Line:     line,
	}
}

type HTTPHandlerInfo struct {
	FuncName string
	File     string
	Line     int
}
