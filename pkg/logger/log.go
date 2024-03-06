package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

var (
	log *zap.Logger
)

func init() {
	initLogger()
}
func initLogger() {

	//获取编码器,NewJSONEncoder()输出json格式，NewConsoleEncoder()输出普通文本格式
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder //指定时间格式
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoder := zapcore.NewConsoleEncoder(encoderConfig)

	//文件writeSyncer
	fileWriteSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "/mnt/data/nfs/myx/heteroflow/logs/error.log", //日志文件存放目录
		MaxSize:    1024,                                          //文件大小限制,单位MB
		MaxBackups: 5,                                             //最大保留日志文件数量
		MaxAge:     15,                                            //日志文件保留天数
		Compress:   false,                                         //是否压缩处理
	})
	// fileCore := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(fileWriteSyncer, zapcore.AddSync(os.Stdout)), zapcore.DebugLevel) //第三个及之后的参数为写入文件的日志级别,ErrorLevel模式只记录error级别的日志
	fileCore := zapcore.NewCore(encoder, fileWriteSyncer, zapcore.InfoLevel) //第三个及之后的参数为写入文件的日志级别,ErrorLevel模式只记录error级别的日志

	log = zap.New(fileCore, zap.AddCaller()) //AddCaller()为显示文件名和行号
}

func Info(msg string, fields ...zapcore.Field) {
	log.Info(msg, fields...)
}
func Error(msg string, fields ...zapcore.Field) {
	log.Error(msg, fields...)
}
func Debug(msg string, fields ...zapcore.Field) {
	log.Debug(msg, fields...)
}
func Fatal(msg string, fields ...zapcore.Field) {
	log.Fatal(msg, fields...)
}
func Panic(msg string, fields ...zapcore.Field) {
	log.Panic(msg, fields...)
}
