package logger

import "encoding/json"

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository}
}

func (s *service) Write(p []byte) (n int, err error) {
	var data LoggerData

	json.Unmarshal(p, &data)

	return len(p), s.repository.AddLoggerActivity(&data)
}
