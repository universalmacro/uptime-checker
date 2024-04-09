package uptimechecker

import (
	"context"
	"testing"
)

func TestHttp(t *testing.T) {
	checker := &HttpChecker{
		Host: "www.google.com",
		Port: "443",
	}
	checker.Check(context.Background())
}
