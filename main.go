package main

import (
	"bufio" "flag" "fmt" "os" "strconv" "strings"
)

type Contact struct {
ID    int
	Nom   string
	Email string
}

var contacts = make(map[int]Contact)

var scanner = bufio.NewScanner(os.Stdin)

func main() {

	idFlag := flag.Int("id", 0, "ID du contact")
	nomFlag := flag.String("nom", "", "Nom du contact")
	emailFlag := flag.String("email", "", "Email du contact")
	flag.Parse()

	if *idFlag != 0 && *nomFlag != "" && *emailFlag != "" {
		contacts[*idFlag] = Contact{ID: *idFlag, Nom: *nomFlag, Email: *emailFlag}
		fmt.Println("Contact ajouté grâce aux flags !")
	}

	for {
		fmt.Println("\n=== Mini CRM ===")
		fmt.Println("1. Ajouter un contact")
		fmt.Println("2. Lister les contacts")
		fmt.Println("3. Supprimer un contact")
		fmt.Println("4. Mettre à jour un contact")
		fmt.Println("5. Quitter")
		fmt.Print("Votre choix : ")

		scanner.Scan()
		choix := scanner.Text()

		switch choix {
		case "1":
			ajouterContact()
		case "2":
			listerContacts()
		case "3":
			supprimerContact()
		case "4":
			mettreAJourContact()
		case "5":
			fmt.Println("Fin du programme")
			return
		default:
			fmt.Println("Choix invalide")
		}
	}
}


func ajouterContact() {
	fmt.Print("ID: ")
	scanner.Scan()
	idStr := scanner.Text()
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("ID invalide")
		return
	}

	if _, existe := contacts[id]; existe {
		fmt.Println("Un contact avec cet ID existe déjà")
		return
	}

	fmt.Print("Nom: ")
	scanner.Scan()
	nom := scanner.Text()

	fmt.Print("Email: ")
	scanner.Scan()
	email := scanner.Text()

	contacts[id] = Contact{ID: id, Nom: nom, Email: email}
	fmt.Println("Contact ajouté")
}

func listerContacts() {
	if len(contacts) == 0 {
		fmt.Println("📭 Aucun contact enregistré")
		return
	}

	fmt.Println("\n--- Liste des contacts ---")
	for _, c := range contacts {
		fmt.Printf("ID: %d | Nom: %s | Email: %s\n", c.ID, c.Nom, c.Email)
	}
}

func supprimerContact() {
	fmt.Print("ID du contact à supprimer: ")
	scanner.Scan()
	idStr := scanner.Text()
	id, err := strconv.Atoi(strings.TrimSpace(idStr))
	if err != nil {
		fmt.Println("ID invalide")
		return
	}

	if _, ok := contacts[id]; ok {
		delete(contacts, id)
		fmt.Println( "Contact supprimé")
	} else {
		fmt.Println(" Contact introuvable")
	}
}

func mettreAJourContact() {
	fmt.Print("ID du contact à mettre à jour: ")
	scanner.Scan()
	idStr := scanner.Text()
	id, err := strconv.Atoi(strings.TrimSpace(idStr))
	if err != nil {
		fmt.Println(" ID invalide")
		return
	}

	if c, ok := contacts[id]; ok {
		fmt.Printf("Nouveau nom (laisser vide pour garder '%s') : ", c.Nom)
		scanner.Scan()
		nom := scanner.Text()
		if nom == "" {
			nom = c.Nom
		}

		fmt.Printf("Nouvel email (laisser vide pour garder '%s') : ", c.Email)
		scanner.Scan()
		email := scanner.Text()
		if email == "" {
			email = c.Email
		}

		contacts[id] = Contact{ID: id, Nom: nom, Email: email}
		fmt.Println("Contact mis à jour")
	} else {
		fmt.Println("Contact introuvable")
	}
}
