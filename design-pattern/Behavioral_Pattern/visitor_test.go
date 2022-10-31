/*
@Time : 2022/10/31 21:15
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : visitor_test
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package Behavioral_Pattern

import (
	"fmt"
	"testing"
)

type Visitor interface {
	Visit(Customer)
}

type Customer interface {
	Accept(Visitor)
}

type EnterpriseCustomer struct {
	name string
}

type CustomerCol struct {
	customers [] Customer
}

func (c *CustomerCol) Add (customer Customer) {
	c.customers = append(c.customers, customer)
}

func (c *CustomerCol) Accept(visitor Visitor) {
	for _, customer := range  c.customers {
		customer.Accept(visitor)
	}
}

func NewEnterpriseCustomer(name string) *EnterpriseCustomer {
	return &EnterpriseCustomer{name: name}
}

func (c *EnterpriseCustomer) Accept(visitor Visitor) {
	visitor.Visit(c)
}

type IndividualCustomer struct {
	name string
}

func NewIndividualCustomer(name string) *IndividualCustomer {
	return &IndividualCustomer{
		name: name,
	}
}

func (c *IndividualCustomer) Accept(visitor Visitor) {
	visitor.Visit(c)
}

type ServiceRequestVisitor struct {}

func (*ServiceRequestVisitor) Visit(customer Customer) {
	switch c := customer.(type) {
	case *EnterpriseCustomer:
		fmt.Printf("saving enterprise customer %s\n", c.name)
	case *IndividualCustomer:
		fmt.Printf("serving individual customer %s\n", c.name)
	}
}

type AnalysisVisitor struct {}

func (*AnalysisVisitor) Visit(customer Customer) {
	switch c := customer.(type) {
		case *EnterpriseCustomer:
			fmt.Printf("analysis enterprise customer %s\n", c.name)
	}
}

func TestVisitor(t *testing.T) {
	c1 := &CustomerCol{}
	c1.Add(NewEnterpriseCustomer("AAA"))
	c1.Add(NewEnterpriseCustomer("BBB"))
	c1.Add(NewIndividualCustomer("CCC"))
	c1.Accept(&ServiceRequestVisitor{})

	c2 := &CustomerCol{}
	c2.Add(NewEnterpriseCustomer("AAA"))
	c2.Add(NewIndividualCustomer("BBB"))
	c2.Add(NewEnterpriseCustomer("CCC"))
	c2.Accept(&AnalysisVisitor{})

}