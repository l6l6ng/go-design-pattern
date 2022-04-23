package main

/*
*适配器模式用于转换以龟速接口适配另一种接口
*实际使用中Adapter一般为接口，并且使用工厂函数生成示例。
*Adapter中匿名组合Adapter接口，所以Adapter类也拥有SpecificRequest实例方法，又因为Go语言中非入侵式接口特征，其实Adapter也适配Adaptee接口
 */

// Target 是适配器的目标接口
type Target interface {
	Request() string
}

// Adaptee 是被适配的目标接口
type Adaptee interface {
	SpecificRequest() string
}

func NewAdaptee() Adaptee {
	return &adapteeImpl{}
}

//adapteeImpl 是被适配的目标类
type adapteeImpl struct {
}

func (*adapteeImpl) SpecificRequest() string {
	return "adaptee method"
}

func NewAdapter(adaptee Adaptee) Target {
	return &adapter{
		Adaptee: adaptee,
	}
}

//Adapter 是转换Adaptee为Target接口的适配器
type adapter struct {
	Adaptee
}

func (a *adapter) Request() string {
	return a.SpecificRequest()
}
