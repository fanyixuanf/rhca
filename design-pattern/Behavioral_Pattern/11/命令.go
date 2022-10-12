/*
@Time : 2022/10/12 21:15
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : 命令
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package main

import "fmt"

// 抽象层

//type Doctor struct {
//}
//
//func (d *Doctor) treatNose() {
//	fmt.Println("treatNose")
//}
//
//func (d *Doctor) treatEye() {
//	fmt.Println("treatEye")
//}
//
//type CommandThreatEye struct {
//	doctor *Doctor
//}
//
//func (cmd *CommandThreatEye) Treat() {
//	cmd.doctor.treatEye()
//}
//
//type CommandThreatNose struct {
//	doctor *Doctor
//}
//
//func (cmd *CommandThreatNose) Treat() {
//	cmd.doctor.treatNose()
//}
//
//type Command interface {
//	Treat()
//}
//
//type Nurse struct {
//	CmdList []Command
//}
//
//func (n *Nurse)Notify() {
//	if n.CmdList == nil {
//		return
//	}
//	for _, cmd := range n.CmdList {
//		cmd.Treat()
//	}
//}
//
//
//func main() {
//	doctor := new(Doctor)
//	cmde := CommandThreatEye{doctor: doctor}
//	cmde.Treat()
//	fmt.Println("----------")
//    cmdn := CommandThreatNose{doctor: doctor}
//
//	nurse := new(Nurse)
//	nurse.CmdList = append(nurse.CmdList, &cmdn)
//	nurse.CmdList = append(nurse.CmdList, &cmde)
//	nurse.Notify()
//
//}


// 练习

// 抽象
type Cooker struct {
}

func (c *Cooker)MakeChuan() {
	fmt.Println("MakeChuan .....")
}

func (c *Cooker) MakeChi() {
	fmt.Println("MakeChi .....")
}

type Command interface {
	Make()
}

type CommandChuan struct {
	cooker *Cooker
}

func (cc *CommandChuan) Make() {
	cc.cooker.MakeChuan()
}

type CommandChi struct {
	cooker *Cooker
}

func (cc *CommandChi)Make() {
	cc.cooker.MakeChi()
}


type MM struct {
	cmdList []Command
}

func (m *MM) Notify() {
	if m.cmdList == nil {
		return
	}
	for _, cmd := range m.cmdList {
		cmd.Make()
	}
}


func main()  {
	cooker := new(Cooker)
	cmdc := CommandChuan{cooker: cooker}
	cmdcc := CommandChi{cooker: cooker}
	mm := new(MM)
	mm.cmdList = append(mm.cmdList, &cmdc)
	mm.cmdList = append(mm.cmdList, &cmdcc)
	mm.cmdList = append(mm.cmdList, &cmdc)
	mm.Notify()
}

