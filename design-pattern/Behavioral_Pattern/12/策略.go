/*
@Time : 2022/10/12 21:46
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : 策略
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package main

import "fmt"

//// 抽象层
//type WeaponStrategy interface {
//	UseWeapon()
//}
//
//type Ak47 struct {
//
//}
//
//// 实现类
//func (ak *Ak47)UseWeapon() {
//	fmt.Println("use ak47 ......")
//}
//
//type Bishou struct {
//
//}
//
//func (bs *Bishou) UseWeapon() {
//	fmt.Println("use Bishou ....")
//}
//
//type Hero struct {
//	strategy WeaponStrategy
//}
//
//func (h *Hero)SetWeaponStrategy(s WeaponStrategy) {
//	h.strategy = s
//}
//
//func (h *Hero) Fight() {
//	h.strategy.UseWeapon()
//}
//
//func main() {
//	hero := Hero{}
//	hero.SetWeaponStrategy(new(Ak47))
//	hero.Fight()
//
//	hero.SetWeaponStrategy(new(Bishou))
//	hero.Fight()
//}


type SellStrategy interface {
	GetPrice(price float64) float64
}

type StrategyA struct {

}

func (a *StrategyA) GetPrice(price float64) float64 {
	fmt.Println("AAAAA.....")
	return price * 0.8
}

type StrategyB struct {

}

func (b *StrategyB) GetPrice(price float64) float64 {
	fmt.Println("BBBB.....")
	if price >= 200 {
		 price -= 100
	}
	return price
}

type Goods struct {
	price float64
	Strategy SellStrategy
}

func (g *Goods) SetStrategy(s SellStrategy) {
	g.Strategy = s
}

func (g *Goods) SellPrice() float64 {
	return g.Strategy.GetPrice(g.price)
}

func main() {
	g1 := Goods{
		price: 200.00,
	}
	g1.SetStrategy(new(StrategyA))
	fmt.Println(g1.SellPrice())
	
	g1.SetStrategy(new(StrategyB))
	fmt.Println(g1.SellPrice())
}

