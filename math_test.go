package main

import "testing"

func TestSoma(t *testing.T) {
	total := soma(3, 6)
	if total != 9 {
		t.Errorf("Resultado da soma é inválido: Resultado esperado: %d, Resultado obtido: %d", 9, total)
	}
}
