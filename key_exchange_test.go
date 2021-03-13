package main

import(
	"testing"
)

func TestServerClientKeyExchange(t *testing.T) {
	cases := 100

	for i := 0; i < cases; i++ {
		a := NewRandomPolynomial()

		s, b := ServerKeyExchange(a)
		_, vdbl, bprime, c := ClientKeyExchange(a, b)

		clientCompute := vdbl.ModularRound()


		//server
		bPrimeS := bprime.Multiply(s)

		twoBprimeS := bPrimeS.Double()

		serverCompute := twoBprimeS.Reconciliate(c)

		if serverCompute.Coefficients != clientCompute.Coefficients {
			diffCount := 0

			for i := range serverCompute.Coefficients {
				if serverCompute.Coefficients[i] != clientCompute.Coefficients[i] {
					diffCount++
				}
			}

			t.Errorf("diffCount = %v\n", diffCount)

			return
		}
	}
}
