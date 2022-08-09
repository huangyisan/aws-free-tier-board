package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"runtime"
	"strings"
	"trojan-dashboard/pkg/configs"
)

var (
	Success = NewMsg(0, "success")
	Failed  = NewMsg(-1, "failed")
	Info    = NewMsg(1, "info")
	Warn    = NewMsg(2, "warn")
	skip    int
)

type Msg struct {
	code int
	msg  string
}

func NewMsg(code int, msg string) *Msg {
	return &Msg{
		code: code,
		msg:  msg,
	}
}

func (e *Msg) Msgf(detail string) {
	switch e.code {
	case 0:
		logrus.Infof("code: %d, msg: %s, detail: %s", e.code, e.msg, detail)
	case -1:
		logrus.Errorf("code: %d, msg: %s, detail: %s", e.code, e.msg, detail)
	case 1:
		logrus.Infof("code: %d, msg: %s, detail: %s", e.code, e.msg, detail)
	case 2:
		logrus.Warnf("code: %d, msg: %s, detail: %s", e.code, e.msg, detail)
	}
}

func InitLog() {
	r := configs.LogrusConf
	lvl, _ := logrus.ParseLevel(r.LogLevel)
	logrus.SetLevel(lvl)
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors: r.DisableColors,
		FullTimestamp: r.ReportCaller,
	})
	logrus.SetReportCaller(r.ReportCaller)
	logrus.SetFormatter(&logrus.JSONFormatter{
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			// InitLog可以确保只调用一次, skip是并发安全的
			if skip == 0 {
				// 获取当前路径
				_, filename, _, ok := runtime.Caller(0)
				if ok {
					for l := 1; l < 20; l++ {
						_, file, _, _ := runtime.Caller(l)
						// 判断是否到了当前方法堆栈
						if strings.Index(file, filename) != -1 {
							// 则上一层调用堆栈为本层+1
							skip = l + 1
							break
						}
					}
				}
			}

			_, file, line, _ := runtime.Caller(skip)
			// 切割多余前缀
			p, _ := os.Getwd()
			return "", fmt.Sprintf("%s:%d", strings.TrimPrefix(file, p), line)
		},
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyFile: "caller",
		},
	})

	if r.EnableFile {
		file, err := os.OpenFile(r.LogFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err == nil {
			logrus.SetOutput(file)
		} else {
			Failed.Msgf("failed to log to file, using default stderr")
		}
	}
}
