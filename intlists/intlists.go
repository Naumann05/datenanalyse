package intlists

// Min berechnet das Minimum einer Liste von Integer-Werten.
// Funktioniert nur für nicht-leere Listen.
func Min(values []int) int {
	min := values[0]
	for i := 0; i+1 <= len(values); i++ {
		if values[i] < min {
			min = values[i]
		}
	}

	return min
}

// Max berechnet das Maximum einer Liste von Integer-Werten.
// Funktioniert nur für nicht-leere Listen.
func Max(values []int) int {
	max := values[0]
	for i := 0; i+1 <= len(values); i++ {
		if values[i] > max {
			max = values[i]
		}
	}

	return max
}

// ValueRange erwartet eine Liste mit ganzzahligen Werten.
// Die Funktion liefert eine lückenlose Liste mit allen Zahlen zwischen
// dem Minimum und dem Maximum der Messreihe.
func ValueRange(values []int) []int {
	result := []int{}
	max := Max(values)
	min := Min(values)
	for i := 0; i+min <= max; i++ {
		result = append(result, i+min)
	}

	return result

}

// Sum erwartet eine Liste mit ganzzahligen Werten.
// Die Funktion liefert die Summe aller Werte.
func Sum(values []int) int {
	sum := 0
	for i := 0; i < len(values); i++ {
		sum = sum + values[i]
	}
	return sum
}

// Product erwartet eine Liste mit ganzzahligen Werten.
// Die Funktion liefert das Produkt aller Werte.
func Product(values []int) int {
	product := 1
	for i := 0; i < len(values); i++ {
		product = product * values[i]
	}
	return product
}
