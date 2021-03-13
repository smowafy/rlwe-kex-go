package main

import(
//	"log"
)

func ServerKeyExchange(a Polynomial) (Polynomial, Polynomial) {
	s := NewSmallPolynomial()
	e := NewSmallPolynomial()

	b := (a.Multiply(s)).Add(e)

	return s, b
}

func ClientKeyExchange(a Polynomial, b Polynomial) (Polynomial, LongPolynomial, Polynomial, Polynomial) {
	sprime := NewSmallPolynomial()
	eprime := NewSmallPolynomial()

	bprime := (a.Multiply(sprime)).Add(eprime)

	edprime := NewSmallPolynomial()

	bSprime := b.Multiply(sprime)

	v := bSprime.Add(edprime)

	vdbl := v.ErrorDouble(NewRandomGenerator())

	c := vdbl.CrossRound()

	return sprime, vdbl, bprime, c
}
