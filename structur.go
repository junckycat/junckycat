package main

import (
	"bufio"
	"fmt"
	"os"
)

func bienvenue(bvn string, ch1 int, ch2 int) {
	fmt.Println("Bienvenue sur forhonor Golang!")
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

func main() {
	var p1 Perso
	fmt.Println("Bienvenue sur for honor golang!")
	p1.Init("oroshi", "Samurai", []string{"sabre", "bouclier"}, 1, 1000, 150)
	fmt.Println("===============")
	fmt.Println("Personnage 1 :\nNom :", p1.nom, "\nClasse :", p1.classe)
	fmt.Println("===============")
	var p2 Perso
	p2.Init("raider", "viking", []string{"hache", "potion"}, 1, 1500, 300)
	fmt.Println("Personnage 2 :\nNom :", p2.nom, "\nClasse", p2.classe)
	fmt.Println("===============")
	fmt.Println("Choisissez votre personnage :")
	var choice string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	choice = scanner.Text()
	var pc Perso
	if choice == "1" {
		pc.Init(p1.nom, p1.classe, p1.inv, p1.niveau, p1.pv, p1.pva)
	} else if choice == "2" {
		pc.Init(p2.nom, p2.classe, p2.inv, p2.niveau, p2.pv, p2.pva)
	} else {
		fmt.Println("Veuillez tapez 1 ou 2")
	}
	fmt.Println("Vous avez choisi le personnage :")
	fmt.Println("Nom :", pc.nom, "\nClasse :", pc.classe)
	fmt.Println("===============")
	fmt.Println("Menu")
	fmt.Println("===============")
	fmt.Println("1 : Afficher les caratéristiques")
	fmt.Println("2 : Afficher l'inventaire du personnage")
	fmt.Println("3 : Afficher le marché du IA")
	fmt.Println("4 : Combattre le IA")
	fmt.Println("5 : Quitter le jeu")
	scanner1 := bufio.NewScanner(os.Stdin)
	scanner1.Scan()
	menuChoice := scanner1.Text()
	switch menuChoice {
	case "1":
		pc.DisplayInfo()
	case "2":
		fmt.Println("Afficher l'inventaire du personnage")
	case "3":
		fmt.Println("Afficher le marché du IA")
	case "4":
		fmt.Println("Combattre le IA")
	case "5":
		fmt.Println("Quitter le jeu")
	default:
		fmt.Println("Veuillez entrez un 1 ou 5")

	}
}

func (p Perso) DisplayInfo() {
	fmt.Println("Nom :", p.nom)
	fmt.Println("classe :", p.classe)
	fmt.Println("niveau :", p.niveau)
	fmt.Println("point de vie actuel :", p.pva)
	fmt.Println("point de vie :", p.pv)
	fmt.Println("inventaire :", p.inv)

}

// func Menu() {
// 	fmt.Println("Entrez votre nom")
// 	scanner := bufio.NewScanner(os.Stdin)
// 	for scanner.Scan() {ol
// 		fmt.Println("Votre nom est", scanner.Text())
// 	}
// }
