package request

// Copy from https://github.com/aws/aws-sdk-go
// May have been modified by Byteplus.

import (
	"io"
	"time"

	"github.com/byteplus-sdk/byteplus-go-sdk-v2/byteplus/bytepluserr"
)

var timeoutErr = bytepluserr.New(
	ErrCodeResponseTimeout,
	"read on byteplusbody has reached the timeout limit",
	nil,
)

type readResult struct {
	n   int
	err error
}

// timeoutReadCloser will handle byteplusbody reads that take too long.
// We will return a ErrReadTimeout error if a timeout occurs.
type timeoutReadCloser struct {
	reader   io.ReadCloser
	duration time.Duration
}

// Read will spin off a goroutine to call the reader's Read method. We will
// select on the timer's channel or the read's channel. Whoever completes first
// will be returned.
func (r *timeoutReadCloser) Read(b []byte) (int, error) {
	timer := time.NewTimer(r.duration)
	c := make(chan readResult, 1)

	go func() {
		n, err := r.reader.Read(b)
		timer.Stop()
		c <- readResult{n: n, err: err}
	}()

	select {
	case data := <-c:
		return data.n, data.err
	case <-timer.C:
		return 0, timeoutErr
	}
}

func (r *timeoutReadCloser) Close() error {
	return r.reader.Close()
}

const (
	// HandlerResponseTimeout is what we use to signify the name of the
	// response timeout handler.
	HandlerResponseTimeout = "ResponseTimeoutHandler"
)

// adaptToResponseTimeoutError is a handler that will replace any top level error
// to a ErrCodeResponseTimeout, if its child is that.
func adaptToResponseTimeoutError(req *Request) {
	if err, ok := req.Error.(bytepluserr.Error); ok {
		aerr, ok := err.OrigErr().(bytepluserr.Error)
		if ok && aerr.Code() == ErrCodeResponseTimeout {
			req.Error = aerr
		}
	}
}

// WithResponseReadTimeout is a request option that will wrap the byteplusbody in a timeout read closer.
// This will allow for per read timeouts. If a timeout occurred, we will return the
// ErrCodeResponseTimeout.
//
//	svc.PutObjectWithContext(ctx, params, request.WithTimeoutReadCloser(30 * time.Second)
func WithResponseReadTimeout(duration time.Duration) Option {
	return func(r *Request) {

		var timeoutHandler = NamedHandler{
			HandlerResponseTimeout,
			func(req *Request) {
				req.HTTPResponse.Body = &timeoutReadCloser{
					reader:   req.HTTPResponse.Body,
					duration: duration,
				}
			}}

		// remove the handler so we are not stomping over any new durations.
		r.Handlers.Send.RemoveByName(HandlerResponseTimeout)
		r.Handlers.Send.PushBackNamed(timeoutHandler)

		r.Handlers.Unmarshal.PushBack(adaptToResponseTimeoutError)
		r.Handlers.UnmarshalError.PushBack(adaptToResponseTimeoutError)
	}
}
