package customer

import "fmt"

type Customer struct{
	Number int
	Name string
	Sex string
	Age int
	Phone string
	E_mail string
}
var number1 =0

func NewCustomer(na string,s string,a int,p string,e string)Customer{ 

	number1++ 

	return Customer{
		number1,
		na,
		s,
		a,
		p,
		e,
	}
}

func (c *Customer)Shuchu()string{
	return fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v\t",c.Number,c.Name,c.Sex,c.Age,c.Phone,c.E_mail)
}

func (c *Customer)Amend(){


}
