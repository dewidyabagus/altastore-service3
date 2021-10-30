package logger

type Service interface {
	Write(p []byte) (n int, err error)
}

type Repository interface {
	AddLoggerActivity(data *LoggerData) error
}
