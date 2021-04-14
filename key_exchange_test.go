package main

import (
	"github.com/smowafy/rlwe-kex-go/polynomial"
	"reflect"
	"testing"
)

func TestServerClientKeyExchange(t *testing.T) {
	cases := 100

	for i := 0; i < cases; i++ {
		a := polynomial.NewRandomPolynomial()

		s, b := ServerKeyExchange(a)
		_, vdbl, bprime, c := ClientKeyExchange(a, b)

		clientCompute := vdbl.ModularRound()

		//server
		bPrimeS := bprime.Multiply(s)

		twoBprimeS := bPrimeS.Double()

		serverCompute := twoBprimeS.Reconciliate(c)

		if !reflect.DeepEqual(serverCompute, clientCompute) {
			t.Errorf("key exchange failed\n")
		}
	}
}

func BenchmarkServerClientKeyExchange(t *testing.B) {
	a := polynomial.NewRandomPolynomial()

	s, b := ServerKeyExchange(a)
	_, vdbl, bprime, c := ClientKeyExchange(a, b)

	vdbl.ModularRound()

	//server
	bPrimeS := bprime.Multiply(s)

	twoBprimeS := bPrimeS.Double()

	twoBprimeS.Reconciliate(c)
}
