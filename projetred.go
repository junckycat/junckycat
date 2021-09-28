package main

import (
	"fmt"
)

func main() {

func bienvenue(bvn string, ch1 int, ch2 int) {
	fmt.Println("
	
     ______                                                  
    | ___ (_)                                                 
    | |_/ /_  ___ _ ____   _____ _ __  _   _  ___             
    | ___ \ |/ _ \ '_ \ \ / / _ \ '_ \| | | |/ _ \            
    | |_/ / |  __/ | | \ V /  __/ | | | |_| |  __/            
    \____/|_|\___|_| |_|\_/ \___|_| |_|\__,_|\___|            
                                                          
                                                          
                                                          
                                                          
                    ___ _   _ _ __                          
                   / __| | | | '__|                         
                   \__ \ |_| | |                            
                   |___/\__,_|_|                            
                                                          
                                                          
______ ___________     _   _ _____ _   _  _   _ ___________ 
|  ___|  _  | ___ \   | | | |  _  | \ | || \ | |  _  | ___ \
| |_  | | | | |_/ /   | |_| | | | |  \| ||  \| | | | | |_/ /
|  _| | | | |    /    |  _  | | | | . ` || . ` | | | |    / 
| |   \ \_/ / |\ \    | | | \ \_/ / |\  || |\  \ \_/ / |\ \ 
\_|    \___/\_| \_|   \_| |_/\___/\_| \_/\_| \_/\___/\_| \_|
                                                          
                                                          ")
}

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

func perso() {
	var p1 Perso
	p1.Init("oroshi", "Samurai", []string{"Sabre", "Bouclier"}, 1, 1000, 150)
	fmt.Println("/_____/_____/_____/_____/_____/_____/_____/_____/")
	p1.DisplayInfo()
	fmt.Println("/_____/_____/_____/_____/_____/_____/_____/_____/")
	var p2 Perso
	p2.Init("raider", "viking", []string{"Hache,", "Potion de vie"}, 1, 1500, 300)
	p2.DisplayInfo()
	fmt.Println("")
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

func Bienvenue() {
	var p1 Perso
	fmt.Println("Bienvenue sur for honor golang!")
	p1.Init("oroshi", "Samurai", []string{"sabre", "bouclier"}, 1, 1000, 150)
	fmt.Println("/_____/_____/_____/_____/_____/_____/_____/_____/")
	fmt.Println("Personnage 1 :\nNom :", p1.nom, "\nClasse :", p1.classe)
	fmt.Println("/_____/_____/_____/_____/_____/_____/_____/_____/")
	var p2 Perso
	p2.Init("raider", "viking", []string{"hache", "potion"}, 1, 1500, 300)
	fmt.Println("Personnage 2 :\nNom :", p2.nom, "\nClasse", p2.classe)
	fmt.Println("/_____/_____/_____/_____/_____/_____/_____/_____/")
	choice := 1
	var pc Perso
	if choice == 1 {
		pc.Init(p1.nom, p1.classe, p1.inv, p1.niveau, p1.pv, p1.pva)
	} else if choice == 2 {
		pc.Init(p2.nom, p2.classe, p2.inv, p2.niveau, p2.pv, p2.pva)
	}
	fmt.Println("Vous avez choisi le personnage :")
	fmt.Println("Nom :", pc.nom, "\nClasse :", pc.classe)
	fmt.Println("/_____/_____/_____/_____/_____/_____/_____/_____/")
	fmt.Println("Menu")
	fmt.Println("/_____/_____/_____/_____/_____/_____/_____/_____/")
	fmt.Println("1 : Afficher les caratÃ©ristiques")
	fmt.Println("2 : Afficher l'inventaire du personnage")
	fmt.Println("3 : Afficher le marcher du IA")
	fmt.Println("4 : Combatre le IA")
	fmt.Println("5 : Quitter le jeux")
	menuChoice := 1
	switch menuChoice {
	case 1:
		pc.DisplayInfo()
	case 2:
		fmt.Println("Afficher l'inventaire du personnage")
	case 3:
		fmt.Println("Afficher le marcher du IA")
	case 4:
		fmt.Println("Combatre le IA")
	default:
		fmt.Println("Veuillez entrez un 1 ou 2")
	case 5:
		fmt.Println("Quitter le jeux")
	}
}

func Menu() {
}

func (p *Perso) dead() {
	if p.pva <= 0 {
		p.pva = p.pv / 2
		fmt.Println("Vous etes mort, vous réaparaissait avec la moitié de vos points de vies")
		Menu()
	}
}

func (p *Perso)marchand() {
    fmt.Println("1 : potion de soin (+50 HP)")
	fmt.Println("2 : potion de poison ( -10 hp par seconde pendant 3 secondes")
	fmt.Println("si le joueur tape 1")
	p.addinventory("potion de soin ")
	fmt.Println("si le joueur tape 2")
	p.addinventory("potion de poison")
}

func (p *Perso) addinventory(item string) {
	p.inv = append(p.inv, item)
}

func (p *Perso) PoisonPot(item string) {
	RemoveInventory(p.inventaire, item)
	for i := 1 ; i <= 3 ; i++ {
		p.pva - 10 r
		fmt.Println(p.pva,"/",p.pv)
		time.sleep(1 * Time.seconde)
		dead()
	}
}

func (p *Perso) RemoveInventory(item string) {
	p.inventaire = Remove(p.inventaire, item)
}
