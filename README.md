# RLWE Key Exchange in Go

An implementation of an unauthenticated key exchange algorithm presented by the paper
["Joppe W. Bos, Craig Costello, Michael Naehrig and Douglas Stebila. Post-quantum
key exchange for the TLS protocol from the ring learning with errors problem,
August 2018."](https://ia.cr/2014/599)

The algorithm is based on the
[Ring-Learning-With-Error](https://en.wikipedia.org/wiki/Ring_learning_with_errors)
problem, which is conjectured to be hard and quantum-resistant.

The implementation is supposed to be constant time (https://en.wikipedia.org/wiki/Timing_attack).

Also this C implementation https://github.com/dstebila/rlwekex provided by the
authors of the paper was used as a reference for some parts, such as the lookup
table for the gaussian sampling (as would be mentioned later)
and some bitwise utility functions.

## Why?

Since [lattice-based](https://en.wikipedia.org/wiki/Lattice-based_cryptography)
and [learning-with-errors](https://en.wikipedia.org/wiki/Learning_with_errors)
cryptography seems very promising, specially the facts that it is quantum-resistant
(so far), and that it presents more opportunities for much more powerful tools like
homomorphic encryption. I thought starting here would be a good introduction to
the topic, and possibly used as a basis to build homomorphic encryption with
the same primitives.

[Homomorphic encryption](https://en.wikipedia.org/wiki/Homomorphic_encryption)
provides this wonderful ability to do computations on
encrypted data, which could have amazing implications and applications that are
based on the fact that you can delegate computations to an untrusted
third-party, without them being able to infer anything about the data being
processed.

So the goal is to extend on this to include homomorphic encryption and
eventually build applications on top of these cryptosystems. I am personally
very interested in the possibility it presents to push for the decentralization
of the web and democratizing the supply of compute resources.

### API

At the moment the primitives are implemented for a key exchange, a good example
would be `key_exchange_test.go`.

Enhancing the API to be more friendly than the status quo would be ongoing.

### Gaussian sampling

The sampling was implemented following the paper, which shows an algorithm for
sampling from a distribution extremely similar to a one-dimensional discrete gaussian
distribution with standard deviation `8 / sqrt(2*PI)`.

The values for the precomputed lookup table were referenced and double-checked
against the values in https://github.com/dstebila/rlwekex in `rlwe_table.h`.

### Polynomial arithmetic

All the polynomial arithmetic is done in the [ring](https://en.wikipedia.org/wiki/Ring_%28mathematics%29)
of polynomials modulo (X^1024 + 1) with coefficients modulo (2^32 - 1). (https://en.wikipedia.org/wiki/Polynomial_ring)

Polynomial multiplication is implemented via an [FFT](https://en.wikipedia.org/wiki/Fast_Fourier_transform)-like algorithm based on the
paper ["Henri J. Nussbaumer. Fast Polynomial Transform Algorithms for Digital
Convolution. IEEE Transactions on Acoustics, Speech and Signal Processing, April
1980."](https://ieeexplore.ieee.org/document/1163372)
The paper describes an algorithm for performing a [number-theoretic discrete
fourier transform](https://en.wikipedia.org/wiki/Discrete_Fourier_transform_(general)#Number-theoretic_transform)
which is FFT-like, specifically in the ring of polynomials
mod (X^n + 1) where n is a power of 2.

In summary the algorithm exploits two properties of the ring:
1. X is a 2n-th principal root of unity in the ring. (since X^2n is 1 mod (X^n + 1))
2. Multiplying a polynomial by X corresponds to circular-shifting the polynomial
by 1 place and flipping the sign of the overflow terms.

### Exchange-specific subroutines

Apart from the general RLWE operations, the key exchange uses a reconciliation
mechanism for the two parties to agree on a common key. This mechanism is
explained very nicely in Section 3 of the paper ["Lattice Cryptography for the
Internet. Chris Peikert, 2014"](https://web.eecs.umich.edu/~cpeikert/pubs/suite.pdf).

Consists mainly of 3 components:
- Modular rounding function
- Cross rounding function
- Reconciliation function

The algorithm relies on the following premises:
1. If an element v is uniformly random, then the modular rounding of v is uniformly
random given the cross rounding of v.
2. Given an element w which is sufficiently close to v and the cross rounding of
v, we can recover the modular rounding of v, which is the agreed-on key between
the two parties.


## Contributions, criticism and suggestions

Since this is my first attempt at implementing a crypto library, it could
potentially contain a bunch of mistakes and vulnerabilities, either in
the theory and the understanding of the cryptosystem itself or
the implementation.

Contributions, suggestions and guidance are more than welcome!

Also a deep-dive write-up for explaining the algorithm, some of the relevant
mathematics behind it and other topics that help understanding lattice-based
crypto and (R)LWE generally, and this algorithm specifically, is in progress.
The write-up aims at reducing the friction to understand and get involved in
this very interesting area, hopefully bringing much more ideas
and perspectives and adding more driving force.
