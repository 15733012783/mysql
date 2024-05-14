package logger

import (
	"go.uber.org/zap"
)

func Logger() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()
	zap.ReplaceGlobals(logger)

	zap.S().Info("zap初始化完成")
}

//func Logger() {
//	logFile, err := os.Create("./" + time.Now().Format("20060102") + ".txt")
//	if err != nil {
//		fmt.Println(err)
//	}
//
//	loger := log.New(logFile, "test_", log.Ldate|log.Ltime|log.Lshortfile)
//	//Flags返回Logger的输出选项
//	fmt.Println(loger.Flags())
//
//	//SetFlags设置输出选项
//	loger.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
//
//	//返回输出前缀
//	fmt.Println(loger.Prefix())
//
//	//设置输出前缀
//	loger.SetPrefix("test_")
//
//	//输出一条日志
//	loger.Output(2, "打印一条日志信息")
//
//	fmt.Println(log.Flags())
//	//获取前缀
//	fmt.Printf(log.Prefix())
//
//}
