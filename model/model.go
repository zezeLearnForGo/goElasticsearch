package model

type Span struct {
	TraceId string `json:"trace_id"`
	SpanId  string `json:"span_id"`
	Ref     []*Ref `json:"ref"`
}

type Ref struct {
	TraceId string `json:"trace_id"`
	SpanId  string `json:"span_id"`
	Type    string `json:"type"`
}
