package main

import "fmt"

/*
*go 语言没有构造函数一说，所以一般会定义NewXXX函数来初始化相关类。
*NewXXX函数返回接口时就是简单工厂模式，也就是说go语言的一般推荐做法就是简单工厂
 */

//在这个simplefactory包中只有API接口和NewAPI函数为包外可见，封装了实现的细节。

//API is  interfaced
type API interface {
	Say(name string) string
}

//hiAPI is one of API implement
type hiAPI struct{}

//Say hi to name
func (*hiAPI) Say(name string) string {
	return fmt.Sprintf("Hi, %s", name)
}

//helloAPI ia another API implement
type helloAPI struct {}

//Say hello to name
func (*helloAPI) Say(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}

//NewAPI return API instance by type
func NewAPI(t int) API {
	if t == 1 {
		return &hiAPI{}
	} else if t == 2 {
		return &helloAPI{}
	}
	return nil
}
