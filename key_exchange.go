package main

import (
	//	"log"
	"github.com/smowafy/rlwe-kex-go/gaussian"
	"github.com/smowafy/rlwe-kex-go/polynomial"
)

func ServerKeyExchange(a polynomial.Polynomial) (polynomial.Polynomial, polynomial.Polynomial) {
	s := polynomial.NewSmallPolynomial()
	e := polynomial.NewSmallPolynomial()

	b := (a.Multiply(s)).Add(e)

	return s, b
}

func ClientKeyExchange(a polynomial.Polynomial, b polynomial.Polynomial) (polynomial.Polynomial, polynomial.LongPolynomial, polynomial.Polynomial, polynomial.Polynomial) {
	sprime := polynomial.NewSmallPolynomial()
	eprime := polynomial.NewSmallPolynomial()

	bprime := (a.Multiply(sprime)).Add(eprime)

	edprime := polynomial.NewSmallPolynomial()

	bSprime := b.Multiply(sprime)

	v := bSprime.Add(edprime)

	vdbl := v.ErrorDouble(gaussian.NewRandomGenerator())

	c := vdbl.CrossRound()

	return sprime, vdbl, bprime, c
}
