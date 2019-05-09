package public

import (
	"fmt"
	"os"
	"runtime/debug"
	"time"
)

func PanicHandler() {
	if err := recover(); err != nil {
		exeName := os.Args[0] //获取程序名称

		now := time.Now()  //获取当前时间
		pid := os.Getpid() //获取进程ID

		time_str := now.Format("20060102150405")                          //设定时间格式
		fname := fmt.Sprintf("%s-%d-%s-dump.log", exeName, pid, time_str) //保存错误信息文件名:程序名-进程ID-当前时间（年月日时分秒）
		fmt.Println("dump to file ", fname)

		f, errfile := os.Create(fname)
		if errfile != nil {
			return
		}
		defer f.Close()
		f.WriteString(fmt.Sprintf("%v\n", err)) //输出panic信息
		f.WriteString("========\n")
		f.WriteString(string(debug.Stack())) //输出堆栈信息
		panic(err)
	}
}
