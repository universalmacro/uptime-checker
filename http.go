package uptimechecker

import (
	"context"
	"crypto/tls"
	"errors"
	"net/http"
	"time"
)

type HttpChecker struct {
	Host   string
	Method string
	Port   string
	Https  bool
}

type HttpResult struct {
	StatusCode   int
	Timeout      bool
	SslExpiresOn time.Time
}

func (h *HttpChecker) Url() string {
	return h.Host + ":" + h.Port
}

func (h *HttpChecker) Check(ctx context.Context) (HttpResult, error) {
	conn, err := tls.Dial("tcp", h.Url(), nil)
	if err != nil {
		return HttpResult{}, errors.New("Server doesn't support SSL certificate err: " + err.Error())
	}
	expiry := conn.ConnectionState().PeerCertificates[0].NotAfter
	client := &http.Client{}
	req, _ := http.NewRequest(h.Method, h.Url(), nil)
	resp, err := client.Do(req)
	if err != nil {
		return HttpResult{Timeout: true}, err
	}
	defer resp.Body.Close()
	return HttpResult{StatusCode: resp.StatusCode, SslExpiresOn: expiry}, nil
}
