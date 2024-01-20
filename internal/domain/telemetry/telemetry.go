package telemetry

type PayloadContext struct {
	Country string `json:"country"`
	OS      string `json:"os"`
}

func NewPayloadContext(country string, os string) PayloadContext {
	return PayloadContext{
		Country: country,
		OS:      os,
	}
}
