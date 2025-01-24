package main

import "fmt"

// readUserInput fragt den Benutzer nach der Anzahl der Würfe und der Anzahl der Würfel.
// Die Funktion liefert beide Werte zurück.
func readUserInput() (int, int) {
	var d, n int
	fmt.Print("Gib die Anzahl der Würfe ein: ")
	fmt.Scan(&d)
	fmt.Println("Gib die Anzahl der Würfe ein: ")
	fmt.Scan(&n)
	return d, n
}

// printDiceStatistics berechnet die Statistik für die Würfelwürfe und gibt sie aus.
func printDiceStatistics(rollResults []int) {
	// TODO
}

// main kombiniert die anderen Funktionen zu einem Programm.
func main() {
	// TODO
}
