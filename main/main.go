package main

import(
	"fmt"
	"HomeWork/customers_/customers"
)

type CustomerManagement struct{

	customers customers.Customers

	b bool
	s string
}

func (c *CustomerManagement) MainPage(){
	var index int
	for {
		fmt.Println("\n----首页面-----")
		fmt.Println("1 添加")
		fmt.Println("2 修改")
		fmt.Println("3 删除")
		fmt.Println("4 客户列表")
		fmt.Println("5 退出")
		fmt.Print("请选择：")
		fmt.Scanln(&index)
		switch index{
		case 1:
			c.add()
		case 2:
			c.amend()
		case 3:
			c.delete()
		case 4:
			c.all()
		case 5:
			c.st()
		default:
			fmt.Println("输入错误。")
		}
		if c.b {
			break
		}
	}
	fmt.Println("退出了~")
}
func (c *CustomerManagement)st(){
	fmt.Println("确定要退出吗？(Y/N)")	
	for{
		fmt.Scanln(&c.s)	
		if c.s=="Y"||c.s=="y"||c.s=="N"||c.s=="n" {
			break
		}
		fmt.Println("请重新输入(Y/N)")
	}
	if c.s=="Y"||c.s=="y"{
		c.b=true
	}
}
func (c *CustomerManagement)all(){
	fmt.Println("number\tname\tsex\tage\tphone\te_mail")
	sli:=c.customers.All()
	for i:=0;i<len(sli);i++{ 
		fmt.Println(sli[i].Shuchu())
	}
	fmt.Println("输出完成")
}

func (c *CustomerManagement)add(){
	var (	
	na string
	s string
	a int
	p string
	e string
	)
	fmt.Println("请输入name")
	fmt.Scanln(&na)
	fmt.Println("请输入sex")
	fmt.Scanln(&s)
	fmt.Println("请输入age")
	fmt.Scanln(&a)
	fmt.Println("请输入phone")
	fmt.Scanln(&p)
	fmt.Println("请输入e_mail")
	fmt.Scanln(&e)

	c.customers.Add(na,s,a,p,e)
}
func (c *CustomerManagement)delete(){
	fmt.Println("请输入想删除的编号：")
	n:=-1
	fmt.Scanln(&n)
	q:=c.customers.Delete(n)
	if q==-1{
		fmt.Println("没有找到此编号")
	}else{
		fmt.Println("删除成功")
	}
}
func (c *CustomerManagement)amend(){
	n:=-1
	fmt.Print("请输入你想修改的编号：")
	fmt.Scanln(&n)
	i:=c.customers.Amend(n)
	if i==-1{
		fmt.Println("没有找到这个编号")
		return
	}

	fmt.Print("请输入修改后的name：")
	fmt.Scanln(&c.customers.CustomersSlice[i].Name)
	fmt.Print("请输入修改后的sex：")
	fmt.Scanln(&c.customers.CustomersSlice[i].Sex)
	fmt.Print("请输入修改后的age：")
	fmt.Scanln(&c.customers.CustomersSlice[i].Age)
	fmt.Print("请输入修改后的phone：")
	fmt.Scanln(&c.customers.CustomersSlice[i].Phone)
	fmt.Print("请输入修改后的e_mail：")
	fmt.Scanln(&c.customers.CustomersSlice[i].E_mail)
	fmt.Println("修改完成")
	fmt.Println("number\tname\tsex\tage\tphone\te_mail")
	fmt.Println(c.customers.CustomersSlice[i].Shuchu())
	
}

func main(){
	var c CustomerManagement 
	c.customers = customers.NewCustomers() 
 
	c.MainPage() 
}
