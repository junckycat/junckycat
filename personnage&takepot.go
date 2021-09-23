package main

import "fmt"

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
	p2.Init("raider", "viking", []string{"hache,", "potion de"}, 1, 1500, 300)
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
		}
	}
}
