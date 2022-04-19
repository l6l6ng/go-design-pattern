package main

import "fmt"

/*
* 外观模式
* API 为facade 模块的外观接口，大部分代理使用此接口简化对facade类的访问
* facade模块同时暴露了a和b两个Module的NewXXX和interface，其他代码如果需要使用细节功能时可以调用。
 */

type AModuleAPI interface {
	TestA() string
}

type aModuleTmpl struct {
}

func (*aModuleTmpl) TestA() string {
	return "A module running"
}

func NewAModuleAPI() AModuleAPI {
	return &aModuleTmpl{}
}

type BModuleAPI interface {
	TestB() string
}

type bModuleTmpl struct {
}

func (*bModuleTmpl) TestB() string {
	return "B module running"
}

func NewBModuleAPI() BModuleAPI {
	return &bModuleTmpl{}
}

// API is facade interface of facade package
type API interface {
	Test() string
}

//facade implement
type apiTmpl struct {
	a AModuleAPI
	b BModuleAPI
}

func (a *apiTmpl) Test() string {
	aRet := a.a.TestA()
	bRet := a.b.TestB()
	return fmt.Sprintf("%s\n%s", aRet, bRet)
}

func NewAPI() API {
	return &apiTmpl{
		a: NewAModuleAPI(),
		b: NewBModuleAPI(),
	}
}
