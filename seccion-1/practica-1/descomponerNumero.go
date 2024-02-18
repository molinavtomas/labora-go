package main

import "fmt"

func main() {
	// Ejemplos de uso
	X1 := 120
	s1, s2, s3, s4, s5 := descomponerNumero(X1)
	fmt.Printf("Para X = %d, S1 = %d, S2 = %d, S3 = %d, S4 = %d, S5 = %d\n", X1, s1, s2, s3, s4, s5)

	X2 := 860
	s1, s2, s3, s4, s5 = descomponerNumero(X2)
	fmt.Printf("Para X = %d, S1 = %d, S2 = %d, S3 = %d, S4 = %d, S5 = %d\n", X2, s1, s2, s3, s4, s5)
}

func descomponerNumero(X int) (S1, S2, S3, S4, S5 int) {
	limiteS1 := 50
	limiteS2 := 50
	limiteS3 := 600
	limiteS4 := 800
	S5 = 0

	if X > limiteS1 {
		S1 = limiteS1
		X -= limiteS1
	} else {
		S1 = X
		X = 0
	}

	if X > limiteS2 {
		S2 = limiteS2
		X -= limiteS2
	} else {
		S2 = X
		X = 0
	}

	if X > limiteS3 {
		S3 = limiteS3
		X -= limiteS3
	} else {
		S3 = X
		X = 0
	}

	if X > limiteS4 {
		S4 = limiteS4
		X -= limiteS4
	} else {
		S4 = X
		X = 0
	}

	S5 = X

	return S1, S2, S3, S4, S5
}
