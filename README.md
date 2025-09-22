# 📇 Mini CRM en Go

Un projet en **Golang** qui permet de gérer des contacts en ligne de commande.  

## Fonctionnalités

- Afficher un menu en boucle
- Ajouter un contact (ID, Nom, Email)
- Lister tous les contacts
- Supprimer un contact par son ID
- Mettre à jour un contact
- Quitter l’application
- Ajouter un contact directement avec des **flags** (ex: `-id -nom -email`)

---

## Concepts utilisés

- `map` pour stocker les contacts  
- Boucle infinie `for {}`  
- `switch` pour gérer le menu  
- `if err != nil` pour vérifier les erreurs  
- `comma ok idiom` pour vérifier si un contact existe  
- `strconv.Atoi` pour convertir une chaîne en entier  
- `bufio.Scanner` et `os.Stdin` pour lire l’entrée utilisateur  

## Installation

Cloner le projet :

```bash
git clone https://github.com/ton-pseudo/mini-crm.git
cd mini-crm
