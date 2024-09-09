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
		index := ((int(r) - 32) * 9) + 2
		indices = append(indices, index)
	}
	return indices
}

// Fonction récursive pour afficher les lignes intercalées pour chaque rune
func recursiveGetLines(indices []int, lignes []string, ligneActuelle int) {
	if ligneActuelle >= 8 {
		return //Cas de base : arrêt après 9 lignes
	}
	// Affichage de la ligneActuelle pour chaque rune
	for _, index := range indices {
		if index+ligneActuelle < len(lignes) {
			fmt.Printf("%s ", lignes[index+(ligneActuelle-1)])
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
	chaine := os.Args[1]
	chaine = strings.ReplaceAll(chaine, `\n`, "\n")
	segments := strings.Split(chaine, "\n")
	// Lire le fichier "standard.txt" et récupérer les lignes
	lignes := readFile("standard.txt")

	// Afficher chaque segment avec un espacement approprié
	for _, segment := range segments {
		if segment == "" {
			fmt.Print("\n") // Imprimer un seul saut de ligne pour le \n
		} else {
			runes := convertRunes(segment)

			// Calculer les indices des runes dans le fichier
			indices := getRunesIndex(runes)

			// Afficher les lignes pour chaque rune de manière récursive
			recursiveGetLines(indices, lignes, 0)
		}
	}
}
