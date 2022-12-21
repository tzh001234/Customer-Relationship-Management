package customers

import (
	"Customer-Relationship-Management/customer"
) 


type Customers struct{
	CustomersSlice []customer.Customer
}

func (c *Customers)All() []customer.Customer{
	return c.CustomersSlice
}

func (c *Customers)Add(na string,s string,a int,p string,e string){
	ccc:=customer.NewCustomer(na,s,a,p,e)
	c.CustomersSlice=append(c.CustomersSlice,ccc)
} 
func (c *Customers)Delete(n int)int{
	q:=c.chazhaoxiabiao(n)
	if q!=-1{
		c.CustomersSlice=append(c.CustomersSlice[:q],c.CustomersSlice[q+1:]... )
	}
	return q
}
func (c *Customers)chazhaoxiabiao(n int)int{
	q:=-1
	for i:=0;i<len(c.CustomersSlice);i++{
		if c.CustomersSlice[i].Number==n{
			q=i 
		}
	}
	return q 
}

func (c *Customers)Amend(n int)int{
	
	return c.chazhaoxiabiao(n)

}


func NewCustomers()Customers{
	
	customers:=Customers{
	}
	c:=customer.NewCustomer("田子恒","男",18,"110","1186@qq.com")

	customers.CustomersSlice=append(customers.CustomersSlice,c)
	
	return customers
		
}

