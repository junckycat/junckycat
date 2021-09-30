package main

import (
	"bufio"
	"fmt"
	"os"
)

func magee(bvn string, ch1 int, ch2 int) {
	fmt.Println("Bienvenue sur For honnor GOLANG")
}

type Perso struct {
	nom    string
	classe string
	niveau int
	pv     int
	pva    int
	inv    []string
	skill  []string
	coins  int
}

func (c *Perso) Init(nom string, class string, inv []string, niveau int, pv int, pva int, skill []string, coins int) {
	c.nom = nom
	c.classe = class
	c.niveau = niveau
	c.pv = pv
	c.pva = pva
	c.inv = inv
	c.skill = skill
	c.coins = coins
}

func main() {
	var p1 Perso
	fmt.Println("Bienvenue sur for honor golang!")
	p1.Init("oroshi", "Samurai", []string{" Sabre, bouclier"}, 1, 1000, 150, []string{"coup de poing"}, 100)
	fmt.Println("===============")
	fmt.Println("Personnage 1 :\nNom :", p1.nom, "\nClasse :", p1.classe)
	fmt.Println("===============")
	var p2 Perso
	p2.Init("raider", "viking", []string{" Hache, potion de soin"}, 1, 1500, 300, []string{"coup de poing"}, 100)
	fmt.Println("===============")
	fmt.Println("Personnage 2 :\nNom :", p2.nom, "\nClasse", p2.classe)
	fmt.Println("===============")
	fmt.Println("Choisissez votre personnage :")
	var choice string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	choice = scanner.Text()
	var pc Perso
	if choice == "1" {
		pc.Init(p1.nom, p1.classe, p1.inv, p1.niveau, p1.pv, p1.pva, p1.skill, p1.coins)
		fmt.Println(pc.nom)
	} else if choice == "2" {
		pc.Init(p2.nom, p2.classe, p2.inv, p2.niveau, p2.pv, p2.pva, p2.skill, p1.coins)
		fmt.Println(pc)
	} else {
		fmt.Println("Veuillez tapez 1 ou 2")
	}
	scanner1 := bufio.NewScanner(os.Stdin)

	fmt.Println("Vous avez choisi le personnage :")
	fmt.Println("Nom :", pc.nom, "\nClasse :", pc.classe, "skill:", pc.skill)
	for {
		fmt.Println("===============")
		fmt.Println("Menu")
		fmt.Println("===============")
		fmt.Println("1 : Afficher les caratéristiques")
		fmt.Println("2 : Afficher l'inventaire du personnage")
		fmt.Println("3 : Afficher le marché du IA")
		fmt.Println("4 : Combattre le IA")
		fmt.Println("5 : Quitter le jeu")
		scanner1.Scan()
		menuChoice := scanner1.Text()
		switch menuChoice {
		case "1":
			pc.DisplayInfo()
		case "2":
			fmt.Println("Afficher l'inventaire du personnage")
			pc.DisplayInventory()
		case "3":
			fmt.Println("Afficher le marcher du IA")
			pc.marchand()
		case "4":
			fmt.Println("Afficher les skills")
		case "5":
			fmt.Println("Combatre le IA")
		case "6":
			fmt.Println("Quitter le jeux")
		default:
			fmt.Println("Veuillez entrez un chiffre entre 1 et 6")
		}
	}
}
func (p Perso) DisplayInfo() {
	fmt.Println("Nom :", p.nom)
	fmt.Println("classe :", p.classe)
	fmt.Println("niveau :", p.niveau)
	fmt.Println("point de vie actuel :", p.pva)
	fmt.Println("point de vie :", p.pv)
	fmt.Println("inventaire :", p.inv)
	fmt.Println("skills :", p.skill)
	fmt.Println("coins:", p.coins)
}
func (p Perso) DisplayInventory() {
	if len(p.inv) == 0 {
		fmt.Println("inventaire vide")
	}
	for i := 0; i < len(p.inv); i++ {
		fmt.Println("->", p.inv[i])
	}

}
func (p *Perso) takePot() {
	for _, lettre := range p.inv {
		if lettre == "potion de soin" && p.pv <= p.pv-50 {
			p.pv += 50
			p.inv[len(p.inv)-1] = "x1"
		}
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
func (p *Perso) marchand() {
	scanner1 := bufio.NewScanner(os.Stdin)
	fmt.Println("1 : potion de soin (100 coins)")
	fmt.Println("2 : potion de poison (100 coins)")
	fmt.Println("3 : livre de sort (250 coins)")
	scanner1.Scan()
	choice := scanner1.Text()
	if choice == "1" {
		if p.coins >= 100 {
			fmt.Println("Vous avez acheté une potion de soin")
			p.addinventory("Potion de soin")
			p.coins -= 100
		}
	}
	if choice == "2" {
		if p.coins >= 100 {
			fmt.Println("Vous avez acheté une potion de poison")
			p.addinventory("Potion de poison")
			p.coins -= 100
		} else {
			fmt.Println("Vous n'avez pas assez d'argents, travaillez")
		}
	}
	if choice == "3" {
		if p.coins <= 250 {
			fmt.Println("Vous venez d'acheter le Livre de sorts")
			p.addinventory("Livre de sorts")
			p.coins -= 250
		} else {
			fmt.Println("Vous n'avez pas assez d'argents, Travaillez")
		}
	}
}

func (p *Perso) addinventory(item string) {
	p.inv = append(p.inv, item)
}

func (p *Perso) spellbook(skill []string) {
	p.skill = append(p.skill, "Boule de feu")
}
func (p *Perso) Coins(coins int) {
}
