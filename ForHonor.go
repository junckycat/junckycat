package main

import (
	"fmt"
 "time"
)

type Perso struct {
	nom    string
	classe string
	niveau int
	pv     int
	pva    int
	inv    []string
}

func (c *Perso) Init(nom string, class string, inv []string, niveau int, pv int, pva int) {
	c.nom = nom
	c.classe = class
	c.niveau = niveau
	c.pv = pv
	c.pva = pva
	c.inv = inv
}

func main() {
	var p1 Perso
	p1.Init("oroshi", "Samurai", []string{"", ""}, 1, 1000, 150)
	fmt.Println("----------------")
	p1.DisplayInfo()
	fmt.Println("----------------")
	var p2 Perso
	p2.Init("raider", "viking", []string{"hache,", "potion de vie"}, 1, 1500, 300)
	p2.DisplayInfo()
	fmt.Println("----------------")

}

func (p Perso) DisplayInfo() {
	fmt.Println("Nom :", p.nom)
	fmt.Println("classe :", p.classe)
	fmt.Println("niveau :", p.niveau)
	fmt.Println("point de vie actuel :", p.pva)
	fmt.Println("point de vie :", p.pv)
	fmt.Println("inventaire :", p.inv)
}

func (p Perso) DisplayInventory() {
	if len(p.inv) == 0 {
		fmt.Println("inventaire vide")
	}
	for i := 0; i < len(p.inv); i++ {
		fmt.Println("->", p.inv[i], "x1")
	}
}

func (p *Perso) takePot() {
	for _, lettre := range p.inv {
		if lettre == "potion de vie" && p.pv <= p.pv-50 {
			p.pv += 50
			p.inv[len(p.inv)-1] = ""
		}
	}
}

func (p *Perso) dead() {
	if p.pva <= 0 {
		p.pva = p.pv / 2
		fmt.Println("Vous etes mort, vous réaparaissait avec la moitié de vos points de vies")
		Menu()
	}
}

func (p *Perso)marchand() {
    fmt.Println("1 : potion de vie (+50 HP)")
	fmt.Println("2 : potion de poison ( -10 hp par seconde pendant 3 secondes")
    p.addinventory("Potion de soin")
}

func (p *Perso) addinventory(item string) {
	p.inventaire = append(p.inventaire, item)
   }
}

func (p *Perso) PoisonPot(item string) {
	RemoveInventory(p.inventaire, item)
	for i := 1 ; i <= 3 ; i++ {
		p.pva - 10 
		fmt.Println(p.pva,"/",p.pv)
		time.sleep(1 * Time.seconde)
		dead()
	}
}

func (p *Perso) RemoveInventory(item string) {
	p.inventaire = Remove(p.inventaire, item)
}

