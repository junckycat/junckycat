package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const (
	cuirDeSanglier = "Cuir de Sanglier"
	plumeDeCorbarc = "Plume de Corbarc"
	fourrureDeloup = "Fourrure de Loup"
	peaudetroll    = "Peau de Troll"
)

var pot int
var pot2 int
var menu int
var retour int
var pc Personage

type Personage struct {
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
type Equipement struct {
	Chapeau string
	Tunique string
	Bottes  string
}
type Barbare struct {
	name   string
	lp     int
	lpmax  int
	attack int
}

func mainn() {
	var p1 Personage

	fmt.Printf("Vous allez être redirigé vers le menu principal")
	time.Sleep(1 * time.Second)
	fmt.Printf(". ")
	time.Sleep(1 * time.Second)
	fmt.Printf(". ")
	time.Sleep(1 * time.Second)
	fmt.Printf(". \n")
	time.Sleep(1 * time.Second)
	p1.menu()
}
func magee(bvn string, ch1 int, ch2 int) {
	fmt.Println()
}
func (c *Personage) Init(nom string, class string, inv []string, invMax int, niveau int, pv int, pva int, skill []string, coins int) {
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
	var p1 Personage
	fmt.Println(`	
	#######  #####   ######            ##   ##    #####   ##   ##  ##   ##   #####   ######
	##   #  ##   ##   ##  ##           ##   ##  ##   ##  ###  ##  ###  ##  ##   ##   ##  ##
	## #    ##   ##   ##  ##           ##   ##  ##   ##  #### ##  #### ##  ##   ##   ##  ##
	####    ##   ##   #####            #######  ##   ##  ## ####  ## ####  ##   ##   #####
	## #    ##   ##   ## ##            ##   ##  ##   ##  ##  ###  ##  ###  ##   ##   ## ##
	##      ##   ##   ##  ##           ##   ##  ##   ##  ##   ##  ##   ##  ##   ##   ##  ##
	####     #####   #### ##           ##   ##   #####   ##   ##  ##   ##   #####   #### ##
							
												EN GOLANG	
      
		`)

	time.Sleep(3 * time.Second)

	p1.Init("Oroshi", "Samurai", []string{" Sabre, bouclier"}, 10, 1, 1000, 150, []string{"coup de poing"}, 10000)
	fmt.Println("		")
	fmt.Println("			░▒░▒▓█░▒▓█▓▒░▒▓█▓▒░▒▓░▒░▒▓█░▒▓█▓▒░▒▓█▓▒░▒▓░▒░▒▓█░▒▓█▓▒░▒▓█▓▒░▒▓")
	fmt.Println("		")
	fmt.Println("	        				Personnage numéro 1 \n						Nom :", p1.nom, "\n						Classe :", p1.classe, "\n 	        				Niveau :", p1.niveau)
	fmt.Println("		")
	fmt.Println("			░▒░▒▓█░▒▓█▓▒░▒▓█▓▒░▒▓░▒░▒▓█░▒▓█▓▒░▒▓█▓▒░▒▓░▒░▒▓█░▒▓█▓▒░▒▓█▓▒░▒▓")
	var p2 Personage
	p2.Init("Raider", "Viking", []string{" Hache, potion de soin"}, 10, 1, 1500, 300, []string{"coup de poing"}, 100)
	fmt.Println("		")
	fmt.Println("						Personnage numéro 2 \n						Nom :", p2.nom, "\n						Classe :", p2.classe, "\n 	        				Niveau :", p2.niveau)
	fmt.Println("		")
	fmt.Println("			░▒░▒▓█░▒▓█▓▒░▒▓█▓▒░▒▓░▒░▒▓█░▒▓█▓▒░▒▓█▓▒░▒▓░▒░▒▓█░▒▓█▓▒░▒▓█▓▒░▒▓")
	fmt.Println("		")
	fmt.Println("		")
	fmt.Println("		")
	fmt.Println("-----------------------------")
	fmt.Println("Choisissez votre personnage :")
	fmt.Println("-----------------------------")
	var choice string
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		choice = scanner.Text()
		if choice == "1" {
			pc.Init(p1.nom, p1.classe, p1.inv, p1.invMax, p1.niveau, p1.pv, p1.pva, p1.skill, p1.coins)
			fmt.Println(pc.nom)
			break
		} else if choice == "2" {
			pc.Init(p2.nom, p2.classe, p2.inv, p2.invMax, p2.niveau, p2.pv, p2.pva, p2.skill, p1.coins)
			fmt.Println(pc)
			break
		} else {
			fmt.Println("Veuillez tapez 1 ou 2")
		}
	}
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
		fmt.Println("		4 : Parler au forgeron")
		fmt.Println("				")
		fmt.Println("		==================================")
		fmt.Println("				")
		fmt.Println("		5 : Combattre le IA")
		fmt.Println("				")
		fmt.Println("		==================================")
		fmt.Println("		6 : Combattre Barbe noir")
		fmt.Println("				")
		fmt.Println("		==================================")
		fmt.Println("		7 : Combattre le Barbarre")
		fmt.Println("				")
		fmt.Println("		==================================")
		fmt.Println("		8 : Combattre le Dragon")
		fmt.Println("				")
		fmt.Println("		==================================")
		fmt.Println("		9 : Combattre Baltazar")
		fmt.Println("				")
		fmt.Println("		==================================")
		fmt.Println("		0 : Quittez le jeux")
		fmt.Println("				")
		fmt.Println("		==================================")
		var menuchoice string
		fmt.Scanln(&menuchoice)

		switch menuchoice {
		case "1":
			pc.DisplayInfo()

		case "2":
			fmt.Println("Afficher l'inventaire du personnage")
			pc.DisplayInventory()
		case "3":
			fmt.Println("Rentrer dans la Taverne")
			pc.Marchand()
		case "4":
			pc.Forgeron()
		case "5":
			fmt.Println("Combatre le IA")
			pc.TrainingFight()
		case "6":
			fmt.Println("Combatre le Barbe noir")
			pc.TheFirst()
		case "7":
			fmt.Println("Combatre le Barbarre")
			pc.TheSecond()
		case "8":
			fmt.Println("Combatre le Dragon")
			pc.TheThird()
		case "9":
			fmt.Println("Combatre Baltazar")
			pc.TheFourth()
		case "0":
			fmt.Println("Quittez le jeux")
		case "cheat":
			pc.coins += 100000
		default:
			fmt.Println("Veuillez entrez un chiffre entre 1 et 6")

		}
	}
}
func (p *Personage) DisplayInfo() {
	fmt.Println("Nom :", p.nom)
	fmt.Println("classe :", p.classe)
	fmt.Println("niveau :", p.niveau)
	fmt.Println("point de vie actuel :", p.pva)
	fmt.Println("point de vie :", p.pv)
	fmt.Println("inventaire :", p.inv)
	fmt.Println("skills :", p.skill)
	fmt.Println("coins:", p.coins)
}
func (p *Personage) menu() {
}
func (p *Personage) DisplayInventory() {
	if len(p.inv) == 0 {
		fmt.Println("inventaire vide")
	}
	for i := 0; i < len(p.inv); i++ {
		fmt.Println(i, "->", p.inv[i])

	}
}
func (p *Personage) AddInventory(item string, price int) {
	if p.coins >= price {
		p.coins -= price
		p.inv = append(p.inv, item)
	} else {
		fmt.Println("Tu n'a pas assez d'argent pour acheter cet objet")
	}
}
func (p *Personage) RemoveInventory(item string) {
	var index int = -1
	for i := range p.inv {
		if p.inv[i] == item {
			index = i
		}
	}
	if index != -1 {
		p.inv = append(p.inv[:index], p.inv[index+1:]...)
	}
}
func (p *Personage) takePot() {
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	for _, letter := range p.inv {
		if letter == "Potion de vie" {
			if p.pva <= (p.pv - 30) {
				p.pva += 30
				fmt.Println("Vous avez gagner 30 points de vie!")
				p.RemoveInventory("Potion de vie")
				break
			} else if p.pva > (p.pva-50) && p.pva < p.pv {
				p.pva = p.pv
				p.RemoveInventory("Potion de vie")
				break
			} else {
				fmt.Println("vous ne pouvez pas utiliser la potion")
				break
			}
		}
	}
}
func (p *Personage) PoisonPot() {
	for _, letter := range p.inv {
		if letter == "Potion de poison" {
			fmt.Println("Vous buvez la potion de poison, ouch")
			p.RemoveInventory("Potion de poison")
			time.Sleep(1 * time.Second)
			fmt.Println(p.pva, "/", p.pv, "PV")
			p.pva -= 10
			fmt.Println(p.pva, "/", p.pv, "PV")
			time.Sleep(1 * time.Second)
			p.pva -= 10
			fmt.Println(p.pva, "/", p.pv, "PV")
			time.Sleep(1 * time.Second)
			p.pva -= 10
			fmt.Println(p.pva, "/", p.pv, "PV")
		}
	}
}

func (pc *Personage) Marchand() {
	scanner1 := bufio.NewScanner(os.Stdin)
	fmt.Println("============================================")
	fmt.Println("Il vous reste :", pc.coins, "Coins à dépenser")
	fmt.Println("============================================")
	fmt.Println("                    ")
	fmt.Println("======Potions=======")
	fmt.Println("                    ")
	fmt.Println("[1]: Potion de soin (10 coins)")
	fmt.Println("[2]: Potion de poison (10 coins)")
	fmt.Println("                    ")
	fmt.Println("======Sortileges=======")
	fmt.Println("                    ")
	fmt.Println("[3]: Boule de Feu (25 coins)")
	fmt.Println("                    ")
	fmt.Println("======Chasseur=======")
	fmt.Println("                    ")
	fmt.Println("[4]:  Fourrure de Loup (25 coins)")
	fmt.Println("[5]:  Cuir de Sanglier (25 coins)")
	fmt.Println("[6]:  Plume de Corbac (25 coins)")
	fmt.Println("[7]:  Peau de Troll (25 coins)")
	fmt.Println("                    ")
	fmt.Println("======Autres=======")
	fmt.Println("                    ")
	fmt.Println("[8]:  Sac XL ( 20 places ) ( 50 coins )")
	fmt.Println("===================")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("[0] Pour retourner au menu")
	scanner1.Scan()
	choice := scanner1.Text()
	if choice == "1" {
		if pc.coins >= 10 {
			if pc.addinventory("Potion de soin") {
				fmt.Println("Vous avez acheté une potion de soin")
				pc.coins -= 10
				pc.Marchand()
			} else if pc.coins <= 10 {
				fmt.Println("Vous n'avez pas assez d'argents, Travaillez")
			}
		}
	}
	if choice == "2" {
		if pc.coins >= 10 {
			if pc.addinventory("Potion de Poison") {
				fmt.Println("Vous avez acheté une potion de Poison")
				pc.coins -= 10
				pc.Marchand()
			} else if pc.coins <= 10 {
				fmt.Println("Vous n'avez pas assez d'argents, Travaillez")
			}
		}
	}
	if choice == "3" {
		if pc.coins >= 25 {
			if pc.addinventory("Sort : Boule de feu") {
				fmt.Println("Vous avez Appris un sort 'Boule de feu'")
				pc.coins -= 25
				pc.Marchand()
			} else if pc.coins <= 25 {
				fmt.Println("Vous n'avez pas assez d'argents, Travaillez")
			}
		}
	}
	if choice == "4" {
		if pc.coins >= 25 {
			fmt.Println("Vous avez acheté une fourrure de Loup")
			pc.addinventory(fourrureDeloup)
			pc.coins -= 25
			pc.Marchand()
		} else if pc.coins <= 25 {
			fmt.Println("Vous n'avez pas assez d'argents, Travaillez")
		}
	}
	if choice == "5" {
		if pc.coins >= 25 {
			if pc.addinventory(cuirDeSanglier) {
				fmt.Println("Vous avez acheté un Cuir de sanglier")
				pc.coins -= 25
				pc.Marchand()
			} else if pc.coins <= 25 {
				fmt.Println("Vous n'avez pas assez d'argents, Travaillez")
			}
		}
	}
	if choice == "6" {
		if pc.coins >= 25 {
			if pc.addinventory(plumeDeCorbarc) {
				fmt.Println("Vous avez acheté une Plume de Corbarc")
				pc.coins -= 25
				pc.Marchand()
			} else if pc.coins <= 25 {
				fmt.Println("Vous n'avez pas assez d'argents, Travaillez")
			}
		}
	}
	if choice == "7" {
		if pc.coins >= 25 {
			if pc.addinventory(peaudetroll) {
				fmt.Println("Vous avez acheté une Peau de Troll")
				pc.coins -= 25
				pc.Marchand()
			} else if pc.coins <= 25 {
				fmt.Println("Vous n'avez pas assez d'argents, Travaillez")
			}
		}
	}
	if choice == "8" {
		if pc.coins >= 50 {
			fmt.Println("Vous avez Agrandis votre sac")
			pc.coins -= 50
			pc.invMax = 20
			pc.Marchand()
		} else if pc.coins <= 50 {
			fmt.Println("Vous n'avez pas assez d'argents, Travaillez")
		}
	}
}
func (p *Personage) Dead() { // fonction dead//
	if p.pva == 0 {
		fmt.Printf("Vous êtes mort. \n")
		time.Sleep(2 * time.Second)
		fmt.Printf("Vous allez réaparaitre\n")
		time.Sleep(2 * time.Second)
		fmt.Printf("BRZZZZZZZZZZZZZZZZZZZZ")
		time.Sleep(1 * time.Second)
		fmt.Printf(". ")
		time.Sleep(1 * time.Second)
		fmt.Printf(". ")
		time.Sleep(1 * time.Second)
		fmt.Printf(". \n")
		time.Sleep(1 * time.Second)
		p.pva = p.pv / 2
		p.DisplayInfo()
	}
}
func (p *Personage) addinventory(item string) bool { // ajouez a linventaire//
	if len(p.inv) > 10 {
		fmt.Println("Vous n'avez plus assez de place dans votre sac.")
		return false
	} else {
		p.inv = append(p.inv, item)
		return true
	}
}
func (p *Personage) spellbook(skill []string) {
	p.skill = append(p.skill, "Boule de feu")
}
func (p *Personage) Coins(coins int) {
}
func (p *Personage) InvMax(inv int) {
}
func (p *Personage) upgradeInventorySlot(inv int) {
	{

	}
}

//Partie forgeron//
func (p *Personage) Forgeron() {
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	fmt.Println("------ Vous entrez dans la forge ------")
	fmt.Println("\n-- Construire un Chapeau de l'aventurier (1) --")
	fmt.Println("il faut 1 plume de corbeau et 1 cuir de sanglier")
	fmt.Println("\n--construire une Tunique de l'Aventurier (2) --")
	fmt.Println("il faut 2 Fourrure de Loup et 1 Peau de Troll")
	fmt.Println("\n-- Construire les bottes de l'aventurier (3) --")
	fmt.Println("il faut 1 Fourrure de Loup et 1 Cuir de Sanglier")
	fmt.Printf("\n-- Retourner au menu principal (4) --\n")
	fmt.Println("\nVotre inventaire:")
	for i := 0; i < len(p.inv); i++ {
		if p.inv[i] != " " {
			fmt.Println("---]", p.inv[i], "[---")
		}
	}
	var enclume int
	fmt.Scanln(&enclume)
	switch enclume {
	case 1:
		fmt.Println(p.Checkinv("Plume de Corbac"))
		fmt.Println(p.Checkinv(cuirDeSanglier))
		fmt.Println(p.Checkinv("Chapeau de l'aventurier"))
		if p.Checkinv(plumeDeCorbarc) && p.Checkinv(cuirDeSanglier) && !p.Checkinv("Chapeau de l'aventurier") {
			p.AddInventory("Chapeau de l'aventurier", 5)
			p.RemoveInventory(plumeDeCorbarc)
			p.RemoveInventory(cuirDeSanglier)
			fmt.Println("Vous avez reçu le Chapeau de l'aventurier")
			time.Sleep(2 * time.Second)
			p.Forgeron()
		} else {
			fmt.Println("Tu n'a pas les item requis idiot.")
			time.Sleep(2 * time.Second)
			p.Forgeron()
		}
	case 2:
		if p.Checkinv(fourrureDeloup) && p.Checkinv(peaudetroll) {
			p.RemoveInventory(fourrureDeloup)
			if p.Checkinv(fourrureDeloup) {
				p.RemoveInventory(fourrureDeloup)
				p.RemoveInventory(peaudetroll)
				p.AddInventory("Tunique de l'Aventurier", 5)
				fmt.Println("Vous avez reçu la Tunique de l'aventurier")
				time.Sleep(2 * time.Second)
				p.Forgeron()
			} else {
				p.AddInventory(fourrureDeloup, 5)
			}
		} else {
			fmt.Println("Tu possède deja cette item idiot.")
			time.Sleep(2 * time.Second)
			p.Forgeron()
		}
	case 3:
		if p.Checkinv(fourrureDeloup) && p.Checkinv(cuirDeSanglier) {
			p.RemoveInventory(fourrureDeloup)
			p.RemoveInventory(cuirDeSanglier)
			p.AddInventory("Bottes de lAventurier", 0)
			fmt.Println("Vous avez reçu des Bottes de laventurier")
			time.Sleep(2 * time.Second)
			p.Forgeron()
		} else {
			fmt.Println("Tu possède deja cette item idiot.")
			time.Sleep(2 * time.Second)
			p.Forgeron()
		}
	case 4:
		p.menu()
	}
}
func (p *Personage) Checkinv(item string) bool { // verif linventaire//
	var founditem bool
	for _, letter := range p.inv {
		if letter == item {
			founditem = true
		}
	}
	if founditem {
		return true
	} else {
		return false
	}
}

func (p *Personage) katana() bool {
	var katana bool
	for _, letter := range p.inv {
		if letter == "Katana des dieux" {
			katana = true
		} else {
			katana = false
		}
	}
	return katana
}

var m Barbare

func (m *Barbare) InitGoblin(name string, lpmax int, attack int) {
	m.name = name
	m.lpmax = lpmax
	m.lp = lpmax / 2
	m.attack = attack
}
func (p *Personage) GoblinPattern(m *Barbare, att int) {
	p.pva -= m.attack * att
	fmt.Println(m.name, "attaque", pc.nom, "et lui inflige", m.attack, " points de dégâts, il lui reste", p.pva, "PV")

}

func (p *Personage) CharTurn(m *Barbare) {
	var choice int
	var damage int = 5
	fmt.Println("Attaquer (1)")
	fmt.Println("Utiliser un objet (2)")
	fmt.Scanln(&choice)
	switch choice {
	case 1:
		if p.katana() == true {
			m.lp = 0
		} else if pc.nom == "Utilisateur" {
			m.lp -= damage * 10
		} else {
			m.lp -= damage
		}
		fmt.Println(pc.nom, "utilise Coup de poing et inflige", damage, "points de dégâts à", m.name, "il lui reste", m.lp, "PV")
	case 2:
		for i := 0; i < len(p.inv); i++ {
			if p.inv[i] != " " {
				fmt.Println("---]", p.inv[i], "[---")

			}
		}
		fmt.Println("Prendre une potion de soin (1)")
		fmt.Println("Prendre une potion de poison (2)")
		fmt.Println("Retour (3)")
		var use int
		fmt.Scanln(&use)
		switch use {
		case 1:
			p.takePot()
		case 2:
			p.PoisonPot()
		case 3:
			p.CharTurn(m)
		}
	}
}

func (p *Personage) TrainingFight() {
	var e1 Barbare
	var e2 Barbare
	n := rand.Intn(10)
	if n == 9 {
		fmt.Println("Un Barbarre veut se filler")
		time.Sleep(3 * time.Second)
		var turn int
		e2.InitGoblin("Barbarre", 80, 5)
		for i := 0; i <= 9999; i++ {
			turn++
			os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
			fmt.Println("Tour", turn)
			time.Sleep(1 * time.Second)
			fmt.Println("C'est a ton tour!")
			p.CharTurn(&e2)
			if e2.lp <= 0 {
				var mimic string = "Vous avez vaincu le Barbarre "
				for _, letter := range mimic {
					time.Sleep(25 * time.Millisecond)
					fmt.Println(letter)
				}
				time.Sleep(1 * time.Second)
				var strange1 string = "vous le chercher"
				var strange2 string = " et vous obtennez "
				var strange3 string = "un objet étrange"
				for _, letter := range strange1 {
					time.Sleep(100 * time.Millisecond)
					fmt.Println(letter)
				}
				for _, letter := range strange2 {
					time.Sleep(200 * time.Millisecond)
					fmt.Println(letter)
				}
				for _, letter := range strange3 {
					time.Sleep(400 * time.Millisecond)
					fmt.Println(letter)
				}
				time.Sleep(3 * time.Second)
				p.AddInventory("Katana", 0)
				break
			}
			time.Sleep(1 * time.Second)
			fmt.Println("\nC'est à lui !")
			p.GoblinPattern(&e2, 1)
			if p.pva <= 0 {
				fmt.Println(e2.name, "vous a battu")
				p.Dead()
				break
			}
			time.Sleep(3 * time.Second)
		}
	} else {
		e1.InitGoblin("Gobelin des rues", 40, 5)
		fmt.Println("Le", e1.name, "estprêt à se battre")
		time.Sleep(3 * time.Second)
		var turn int
		for i := 0; i <= 9999; i++ {
			turn++
			os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
			fmt.Println("Tour", turn)
			time.Sleep(1 * time.Second)
			fmt.Println("C'est au joueur !")
			p.CharTurn(&e1)
			if e1.lp <= 0 {
				fmt.Println("\nVous avez vaincu", e1.name)
				time.Sleep(2 * time.Second)
				break
			}
			time.Sleep(1 * time.Second)
			fmt.Println("\nC'est à l'ennemi !")
			if turn%3 == 3 || turn == 3 {
				p.GoblinPattern(&e1, 2)
			} else {
				p.GoblinPattern(&e1, 1)
			}
			if p.pva <= 0 {
				fmt.Println(e1.name, "vous a battu")
				p.Dead()
				break
			}
			time.Sleep(3 * time.Second)
		}
	}

}

func (p *Personage) TheFirst() {
	var e3 Barbare
	e3.InitGoblin("Barbe noir", 80, 8)
	var turn int
	for i := 0; i <= 9999; i++ {
		turn++

		os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
		fmt.Println("Tour", turn)
		time.Sleep(1 * time.Second)
		fmt.Println("C'est à toi !")
		p.CharTurn(&e3)
		if e3.lp <= 0 {
			fmt.Println("\nVous avez vaincu", e3.name)
			time.Sleep(2 * time.Second)
			p.TheSecond()
			break
		}
		time.Sleep(1 * time.Second)
		fmt.Println("\nC'est à l'ennemi !")
		p.GoblinPattern(&e3, 1)
		if p.pva <= 0 {
			fmt.Println(e3.name, "vous a battu")
			p.Dead()
			break
		}
		time.Sleep(3 * time.Second)
	}
}

func (p *Personage) TheSecond() {
	var e4 Barbare
	e4.InitGoblin("barbarre", 100, 15)
	var turn2 int
	for i := 0; i <= 9999; i++ {
		turn2++

		os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
		fmt.Println("Tour", turn2)
		time.Sleep(1 * time.Second)
		fmt.Println("C'est à toi !")
		p.CharTurn(&e4)
		if e4.lp <= 0 {
			fmt.Println("\nVous avez vaincu", e4.name)
			time.Sleep(2 * time.Second)
			p.TheThird()
			break
		}
		time.Sleep(1 * time.Second)
		fmt.Println("\nC'est à l'ennemi !")
		p.GoblinPattern(&e4, 1)
		if p.pva <= 0 {
			fmt.Println(e4.name, "vous a battu")
			p.Dead()
			break
		}
		time.Sleep(3 * time.Second)
	}
}

func (p *Personage) TheThird() {
	var e5 Barbare
	e5.InitGoblin("Dragon", 150, 40)
	var turn3 int
	for i := 0; i <= 9999; i++ {
		turn3++

		os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
		fmt.Println("Tour", turn3)
		time.Sleep(1 * time.Second)
		fmt.Println("C'est à toi!")
		p.CharTurn(&e5)
		if e5.lp <= 0 {
			fmt.Println("\nVous avez vaincu", e5.name)
			time.Sleep(2 * time.Second)
			p.TheFourth()
			break
		}
		time.Sleep(1 * time.Second)
		fmt.Println("\nC'est à l'ennemi !")
		p.GoblinPattern(&e5, 1)
		if p.pva <= 0 {
			fmt.Println(e5.name, "vous a battu")
			p.Dead()
			break
		}
		time.Sleep(3 * time.Second)
	}
}

func (p *Personage) TheFourth() {
	var e6 Barbare
	e6.InitGoblin("Baltazar", 200, 66)
	var turn4 int
	for i := 0; i <= 9999; i++ {
		turn4++

		os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
		fmt.Println("Tour", turn4)
		time.Sleep(1 * time.Second)
		fmt.Println("C'est à toi !")
		p.CharTurn(&e6)
		if e6.lp <= 0 {
			fmt.Println("\nVous avez vaincu", e6.name)
			time.Sleep(2 * time.Second)
			break
		}
		time.Sleep(1 * time.Second)
		fmt.Println("\nC'est à l'ennemi !")
		p.GoblinPattern(&e6, 1)
		if p.pva <= 0 {
			fmt.Println(e6.name, "vous a battu")
			p.Dead()
			break
		}
		time.Sleep(3 * time.Second)
	}
}
