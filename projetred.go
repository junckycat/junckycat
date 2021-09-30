package main

import (
	"bufio"
	"fmt"
	"os"
)

func magee(bvn string, ch1 int, ch2 int) {
	fmt.Println(`
	 ______________________________
	/                            
   |   |      ----             ---- |.
	 |       ---     --------       |.
	   |      -- BIENVENUE	--		|.
	   |	------   SUR	-------	|.
	   |	  -- FOR HONNOR		---	|.
	   |	--	EN GOLANG   --      |.
	   |          -------           |.
	   |   --    ----   --      --  |.
	   |                            |.
	   |    ---       ------   -    |.
	   |   _________________________|___
  	   |  /                            /.
    	\_____________________________/.		`)
}

type Perso struct {
	nom    string
	classe string
	niveau int
	pv     int
	pva    int
	inv    []string
	invMax int
	skill  []string
	coins  int
}

func (c *Perso) Init(nom string, class string, inv []string, invMax int, niveau int, pv int, pva int, skill []string, coins int) {
	c.nom = nom
	c.classe = class
	c.niveau = niveau
	c.pv = pv
	c.pva = pva
	c.inv = inv
	c.invMax = invMax
	c.skill = skill
	c.coins = coins
}

func main() {
	var p1 Perso
	fmt.Println(`	
	 ______________________________
	/                              
    |      ----                      ---|.
	 |       ---     --------       |.
	   |      -- BIENVENUE	      --|.
	   |	------   SUR	      --|.
	   |	  -- FOR HONNOR ---	|.
	   |	--	EN GOLANG   --  |.
	   |          -------           |.
	   |   --    ----   --      --  |.
	   |                            |.
	   |    ---       ------   -    |.
	   |   _________________________|_
  	   |  /                          /.
    	\____________________________/.		`)
	p1.Init("Oroshi", "Samurai", []string{" Sabre, bouclier"}, 10, 1, 1000, 150, []string{"coup de poing"}, 100)
	fmt.Println("		")
	fmt.Println("		================")
	fmt.Println("		")
	fmt.Println("	        Personnage numéro 1 \n		Nom :", p1.nom, "\n		Classe :", p1.classe, "\n 	        Niveau :", p1.niveau)
	fmt.Println("		")
	fmt.Println("		================")
	var p2 Perso
	p2.Init("Raider", "Viking", []string{" Hache, potion de soin"}, 10, 1, 1500, 300, []string{"coup de poing"}, 100)
	fmt.Println("		")
	fmt.Println("		Personnage numéro 2 \n		Nom :", p2.nom, "\n		Classe :", p2.classe, "\n 	        Niveau :", p2.niveau)
	fmt.Println("		")
	fmt.Println("		================")
	fmt.Println("		")
	fmt.Println("Choisissez votre personnage :")
	var choice string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	choice = scanner.Text()
	var pc Perso
	if choice == "1" {
		pc.Init(p1.nom, p1.classe, p1.inv, p1.invMax, p1.niveau, p1.pv, p1.pva, p1.skill, p1.coins)
		fmt.Println(pc.nom)
	} else if choice == "2" {
		pc.Init(p2.nom, p2.classe, p2.inv, p2.invMax, p2.niveau, p2.pv, p2.pva, p2.skill, p1.coins)
		fmt.Println(pc)
	} else {
		fmt.Println("Veuillez tapez 1 ou 2")
	}
	scanner1 := bufio.NewScanner(os.Stdin)

	fmt.Println("Vous avez choisi le personnage :")
	fmt.Println("Nom :", pc.nom, "\nClasse :", pc.classe, "skill:", pc.skill)
	for {
		fmt.Println("				")
		fmt.Println("				")
		fmt.Println("		==================================")
		fmt.Println("			       Menu")
		fmt.Println("		==================================")
		fmt.Println("				")
		fmt.Println("		1 : Afficher les caratéristiques")
		fmt.Println("				")
		fmt.Println("		==================================")
		fmt.Println("				")
		fmt.Println("		2 : Afficher l'inventaire du personnage")
		fmt.Println("				")
		fmt.Println("		==================================")
		fmt.Println("				")
		fmt.Println("		3 : Rentrer dans la Taverne")
		fmt.Println("				")
		fmt.Println("		==================================")
		fmt.Println("				")
		fmt.Println("		4 : Combattre le IA")
		fmt.Println("				")
		fmt.Println("		==================================")
		fmt.Println("				")
		fmt.Println("		5 : Quitter le jeu")
		fmt.Println("				")
		fmt.Println("		==================================")
		scanner1.Scan()
		menuChoice := scanner1.Text()
		switch menuChoice {
		case "1":
			pc.DisplayInfo()
		case "2":
			fmt.Println("Afficher l'inventaire du personnage")
			pc.DisplayInventory()
		case "3":
			fmt.Println("Rentrer dans la Taverne")
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
	fmt.Println("                    ")
	fmt.Println("======Potions=======")
	fmt.Println("                    ")
	fmt.Println("'1': Potion de soin (10 coins)")
	fmt.Println("'2': Potion de poison (10 coins)")
	fmt.Println("                    ")
	fmt.Println("======Sortileges=======")
	fmt.Println("                    ")
	fmt.Println("'3': Boule de Feu (25 coins)")
	fmt.Println("                    ")
	fmt.Println("======Chasseur=======")
	fmt.Println("                    ")
	fmt.Println("'4':  Fourrure de Loup (25 coins)")
	fmt.Println("'5':  Cuir de Sanglier (25 coins)")
	fmt.Println("'6':  Plume de Corbeau (25 coins)")
	fmt.Println("'7':  Peau de Troll (25 coins)")
	fmt.Println("                    ")
	fmt.Println("==========================")
	scanner1.Scan()
	choice := scanner1.Text()
	if choice == "1" {
		if p.coins >= 10 {
			if p.addinventory("Potion de soin") {
				fmt.Println("Vous avez acheté une potion de soin")
				p.coins -= 10
			}
		}
	}
	if choice == "2" {
		if p.coins >= 10 {
			if p.addinventory("Potion de Poison") {
				fmt.Println("Vous avez acheté une potion de Poison")
				p.coins -= 10
			} else {
				fmt.Println("Vous n'avez pas assez d'argents, Travaillez")
			}
		}
	}
	if choice == "3" {
		if p.coins >= 25 {
			if p.addinventory("Sort : Boule de feu") {
				fmt.Println("Vous avez Appris un sort 'Boule de feu'")
				p.coins -= 25
			} else {
				fmt.Println("Vous n'avez pas assez d'argents, Travaillez")
			}
		}
	}
	if choice == "4" {
		if p.coins >= 25 {
			fmt.Println("Vous avez acheté une fourrure de Loup")
			p.addinventory("Fourrure de Loup")
			p.coins -= 25

		} else {
			fmt.Println("Vous n'avez pas assez d'argents, travaillez")
		}
	}
	if choice == "5" {
		if p.coins >= 25 {
			if p.addinventory("Cuir de sanglier") {
				fmt.Println("Vous avez acheté un Cuir de sanglier")
				p.coins -= 25
			} else {
				fmt.Println("Vous n'avez pas assez d'argents, Travaillez")
			}
		}
	}
	if choice == "6" {
		if p.coins >= 25 {
			if p.addinventory("Plume de Corbeau") {
				fmt.Println("Vous avez acheté une Plume de Corbeau")
				p.coins -= 25
			} else {
				fmt.Println("Vous n'avez pas assez d'argents, Travaillez")
			}
		}
	}
	if choice == "7" {
		if p.coins >= 25 {
			if p.addinventory("Peau de Troll") {
				fmt.Println("Vous avez acheté une Peau de Troll")
				p.coins -= 25
			} else {
				fmt.Println("Vous n'avez pas assez d'argents, Travaillez")
			}
		}
	}
}
func (p *Perso) addinventory(item string) bool {
	if len(p.inv) > 10 {
		fmt.Println("Vous n'avez plus assez de place dans votre sac.")
		return false
	} else {
		p.inv = append(p.inv, item)
		return true
	}
}
func (p *Perso) spellbook(skill []string) {
	p.skill = append(p.skill, "Boule de feu")
}
func (p *Perso) Coins(coins int) {
}
func (p *Perso) InvMax(inv int) {
}
