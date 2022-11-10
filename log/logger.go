package log

type ILogger interface {
	LogT(args ...interface{})
	LogD(args ...interface{})
	LogI(args ...interface{})
	LogW(args ...interface{})
	LogE(args ...interface{})
	LogF(args ...interface{})
	LogP(args ...interface{})
	OpenLog(bool)
}
