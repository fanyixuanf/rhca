/*
@Time : 2022/10/19 21:49
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : prototype_test
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package creation_pattern

import (
	"fmt"
	"testing"
)

type clone interface {
	clone() clone
}

type prototypeManager struct {
	prototype map[string]clone
}

func newPrototypeManager() *prototypeManager {
	return &prototypeManager{
		prototype: make(map[string]clone),
	}
}

func (p *prototypeManager)get(name string) clone {
	return p.prototype[name].clone()
}

func (p *prototypeManager)set(name string, prototype clone) {
	p.prototype[name] = prototype
}

type protetype1 struct {
	name string
}

func (p *protetype1)clone() clone {
	pc := *p
	return &pc
}

func TestPrototype(t *testing.T){
	manager := newPrototypeManager()
	p := &protetype1{name: "prototype"}
	manager.set("prototype", p)
	p1 := manager.get("prototype")

	p2 := p1.clone()

	fmt.Println(&p1)
	fmt.Println(&p2)

	if p1 != p2 {
		t.Log("p1 != p2")
	}

}