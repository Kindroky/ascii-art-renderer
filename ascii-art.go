package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Fonction qui prend une chaîne en argument et renvoie une slice de runes
func convertRunes(chaine string) []rune {
	return []rune(chaine)
}

// Fonction pour lire les lignes du fichier "standard.txt"
func readFile(chemin string) []string {
	fichier, err := os.Open(chemin)
	if err != nil {
		log.Fatal(err)
	}
	defer fichier.Close()

	var lignes []string
	scanner := bufio.NewScanner(fichier)
	for scanner.Scan() {
		lignes = append(lignes, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lignes
}

// Fonction pour calculer les indices des runes dans le fichier "standard.txt"
func getRunesIndex(runes []rune) []int {
	var indices []int
	for _, r := range runes {
		index := (int(r) - 32) * 9
		indices = append(indices, index)
	}
	return indices
}

// Fonction récursive pour afficher les lignes intercalées pour chaque rune
func recursiveGetLines(indices []int, lignes []string, ligneActuelle int) {
	if ligneActuelle >= 9 {
		return // Cas de base : arrêt après 9 lignes
	}

	// Affichage de la ligneActuelle pour chaque rune
	for _, index := range indices {
		if index+ligneActuelle < len(lignes) {
			fmt.Printf("%s ", lignes[index+ligneActuelle])
		}
	}
	fmt.Println()

	// Appel récursif pour la ligne suivante
	recursiveGetLines(indices, lignes, ligneActuelle+1)
}

func main() {
	// Vérification de la présence d'un argument
	if len(os.Args) < 2 {
		fmt.Println("Veuillez fournir une chaîne de caractères en argument.")
		os.Exit(1)
	}

	// Récupérer la chaîne de caractères passée en argument
	chaine := os.Args[1]

	// Remplacer les occurrences de \n par des sauts de ligne réels
	chaine = strings.ReplaceAll(chaine, `\n`, "\n")

	// Diviser la chaîne en lignes
	segments := strings.Split(chaine, "\n")

	// Lire le fichier "standard.txt" et récupérer les lignes
	lignes := readFile("standard.txt")

	// Afficher chaque segment avec un espacement approprié
	for i, segment := range segments {
		// Si le segment est vide, cela correspond à un saut de ligne \n
		if segment == "" {
			if i < len(segments)-1 {
				fmt.Println() // Imprimer un seul saut de ligne pour le \n
			}
			continue
		}

		// Convertir la chaîne en runes
		runes := convertRunes(segment)

		// Calculer les indices des runes dans le fichier
		indices := getRunesIndex(runes)

		// Afficher les lignes pour chaque rune de manière récursive
		recursiveGetLines(indices, lignes, 0)

		// Supprimer l'appel à fmt.Println() ici pour éviter l'impression de plusieurs lignes vides
	}
}
