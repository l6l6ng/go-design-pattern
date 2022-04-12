package main

import "fmt"

/*
* 抽象工厂模式用于生成族的工厂，所生成的对象是有关联的。
* 如果抽象工厂退化成生成的对象无关联则成为工厂函数模式
* 比如本例子中使用PDB和XML存储订单信息，抽象工厂分别能生成相关的主订单信息和订单详情信息。
* 如果业务逻辑中需要替换使用的时候只需改动工厂函数相关的类就能替换使用不同的存储方式了
 */

// OrderMainDAO 为订单主记录
type OrderMainDAO interface {
	SaveOrderMain()
}

// OrderDetailDAO 为订单详情记录
type OrderDetailDAO interface {
	SaveOrderDetail()
}

// DAOFactory 抽象模式工厂接口
type DAOFactory interface {
	CreateOrderMainDAO() OrderMainDAO
	CreateOrderDetailDAO() OrderDetailDAO
}

// RDBMainDAO 为关系型数据库的OrderMainDAO实现
type RDBMainDAO struct {
}

func (*RDBMainDAO) SaveOrderMain() {
	fmt.Println("rdb main sava")
}

// RDBDetailFactory 为关系型数据库的OrderDetail 实现
type RDBDetailFactory struct {
}

func (*RDBDetailFactory) SaveOrderDetail() {
	fmt.Println("rdb detail save")
}

// RDBDAOFactory 是RDB 抽象工厂实现
type RDBDAOFactory struct {
}

func (*RDBDAOFactory) CreateOrderMainDAO() OrderMainDAO {
	return &RDBMainDAO{}
}

func (*RDBDAOFactory) CreateOrderDetailDAO() OrderDetailDAO {
	return &RDBDetailFactory{}
}

// XMLMainDAO xml 存储
type XMLMainDAO struct {
}

func (*XMLMainDAO) SaveOrderMain() {
	fmt.Println("xml main save")
}

// XMLDetailDAO xml存储
type XMLDetailDAO struct {
}

func (*XMLDetailDAO) SaveOrderDetail() {
	fmt.Println("xml detail main")
}

// XMLDAOFactory 是RDB抽象工厂的实现
type XMLDAOFactory struct {
}

func (*XMLDAOFactory) CreateOrderMainDAO() OrderMainDAO {
	return &XMLMainDAO{}
}

func (*XMLDAOFactory) CreateOrderDetailDAO() OrderDetailDAO {
	return &XMLDetailDAO{}
}
