package elemental

import (
	"crypto/tls"

	"github.com/opentracing/opentracing-go"
)

// A ClaimsHolder is the interface of a structure that can hold
// Identity Claims (as in a JWT).
type ClaimsHolder interface {
	SetClaims([]string)
	GetClaims() []string
	GetClaimsMap() map[string]string
}

// A TokenHolder is the interface of a structure that can hold a token.
type TokenHolder interface {
	GetToken() string
	TLSConnectionState() *tls.ConnectionState
}

// A SpanHolder is the interface of a structure that can holde a tracing span.
type SpanHolder interface {
	Span() opentracing.Span
	NewChildSpan(string) opentracing.Span
}

// An SessionHolder is both a ClaimsHolder and a TokenHolder.
type SessionHolder interface {
	ClaimsHolder
	TokenHolder
}
