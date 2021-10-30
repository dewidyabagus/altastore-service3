package logger

import "time"

type LoggerData struct {
	Group      string    `json:"group"`
	LatencyNs  int       `json:"latencyns"`
	Level      string    `json:"level"`
	Method     string    `json:"method"`
	Msg        string    `json:"msg"`
	RemoteAddr string    `json:"remoteaddr"`
	Time       time.Time `json:"time"`
	Uri        string    `json:"uri"`
}
