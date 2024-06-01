package main

//导入cmd包
import "xiudong/cli/cmd"

func main() {
	//执行cmd包中的Execute()方法
	//这个方法是在root.go里定义的
	//这个方法实际上是一个将子命令添加到root命令中的方法
	cmd.Execute()
}
