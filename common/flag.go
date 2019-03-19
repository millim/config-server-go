package common

import "flag"

//InitParams 初始化使用的内容
type Params struct {
	DBPath   string
	LogLevel string
	LogFile  string
	Port     string
	Host     string
}

//FlagParams 初始化参数
func buildParams() Params {
	m := Params{}
	flag.StringVar(&m.DBPath, "dbPath", "", "This is db path, default: ./server_ + GIN_MODE")
	flag.StringVar(&m.LogLevel, "logLevel", "debug", "panic fatal error warn info debug trace")
	flag.StringVar(&m.LogFile, "logFile", "", "set print log file, default os.Stdout")
	flag.StringVar(&m.Host, "host", "0.0.0.0", "server host")
	flag.StringVar(&m.Port, "port", "3000", "server port")
	flag.Parse()
	return m
}

var FlagParams Params

func init() {
	FlagParams = buildParams()
}
