James Jenkins

MSCS 630L

April 17, 2022

# Milestone: Digital Signature Algorithm Exploration

## Abstract

The Digital Signature Algorithm (DSA) is an algorithm that allows users to digitally
sign electronic assets or verify a DSA signature. The algorithm is secure 
enough that it is included in the Federal Information Processing Standard 
(FIPS), allowing the algorithm to be used to sign and verify sensitive
government data. This paper will describe a custom implemetation of the DSA 
algorithm and what can be learned from implementing this algorithm from scratch.
This milestone will 1) Introduce the concepts and motivation for implementing the
Digital Signature Algorithm 2) Give an overview of the DSA algorithm and its 
properties 3) Discuss the work that has been completed so far 4) Discuss planned
features.

## Introduction

Cryptography is an essential part of modern-day digital communications. 
Cryptographic techniques are not just used to hide sensitive data, they
may also be used to create a system for identity in the digital world. 
After learning about data encryption techniques and implementing the AES 
algorithm, I wanted to explore another side of cryptography. The DSA 
algorithm does not encrypt any data. It instead uses cryptographic 
techiniques for the sole purpose of ensuring data integrity, identity, 
and non-repudiation. Without these three properties, users would not be
able to establish trust in the digital world. Having to establish trust 
by traditional means would slow down the use of government and legal 
applications and could create large security vulnerabilities.

The paper will give an overview of the algorithm and its properties, then describe
the implementation and results.
The implementation described in this paper closely follows the FIPS 186-4 
Digital Signature Standard. FIPS 186-4 standardizes approved implementations
of digital signature algorithms for use with data owned by the Federal Government.

## Related Work

Digital signatures provide three essential properties. They provide the identity 
of the signatory, integrity of the signed data, and non-repudiation. In 1991, NIST
created the Digital Signature Standard to standardize the implementation and use of 
digital signature algorithms for use with data owned by the Federal Government. 
The Digital Signature Standard outlines three acceptable algorithms for digitally 
validating the integrity of signed binary data and the identity of the signatory. 
The three algorithms that are described as of FIPS 186-4 are DSA, RSA, and ECDSA. 

As defined in FIPS 186-4, there are four main steps in the DSA algorithm: 
1. Generate domain parameters
2. Generate the public/private key pair
3. Generate the signature
4. Verify the DSA signature 

There are two aspects that ensure the security of the algorithm. The first 
aspect is the hash function used to create a message digest. During the 
signature generation and signature verification, a message digest is created
by a hash function to ensure the integrity of the data. If a weak hash function
is utilized at these stages, the algorithm is opened to attacks that could 
affect the integrity of the signed data. Although the SHA-1 algorithm has been
declared insecure since 2005, the implementation in this paper still uses it
as its purpose is to explore the inner workings of the algorithm, not to
achieve perfect security. Collisions in the hash algorithm would enable an 
attacker to theoretically modify the signed content such that the resulting 
message digest would be identical to the message digest of the original 
content.

The second aspect that ensures the algorithm's security is the discrete
logarithm problem. A public/private key pair has some necessary properties,
the keys must be mathematically related and the private key must not be able
to be derived from the public key. In the DSA algorithm, the public key is 
derived from the private key through the process of modular exponentiation. 
This is can be quickly calculated by a computer through the fast modular 
exponentiation algorithm because the algorithm relies on the binary 
representation of integers. The reverse operation of modular exponentiation
is called the discrete logarithm. In order to find the discrete logarithm
of a prime modulous, a program would require a lookup table holding *n* values
where *n* is equal to *ceil(sqrt(p-1))*. For an example, the table needed to 
solve a discrete logarithm with a modulous of 53 would require storing 7 
values. Solving the discrete logarithm problem for a sufficiently large modulous
would require vast amounts of memory to store the lookup table. Using just a 
512 bit prime number would require storing approximately 115,792,089,237,316,195,423,570,985,008,687,907,853,269,984,665,640,564,039,457,584,007,913,129,639,936 values in
a lookup table. With modern computing power, it would take an impractical amount
of time and memory to solve this. Using probablistic algorithms,
such as the Miller-Rabin primality test, we can generate prime numbers that 
are much larger than 512 bits.

## Methedology

To increase focus on the implementation of the DSA algorithm, a CLI application
was created. The application was created in the Go language, allowing it to 
be compiled and executed on any system of the user's choosing. As of this
milestone, a command called 'keygen' was created that generates the domain
parameters and generates a public/private key pair, completing two of the 
four steps outlined in the process above. 

### Parameter Generation

There are three parameters that must be generated in the DSA process. These
three parameters are named *p, q,* and *g*. The lengths of parameters *p* and
*q* must be part of an approved set of lengths specified in FIPS 186-4. In 
order to be compatible with the SHA-1 hash function, the lengths of 1024 bits
and 160 bits were chosen for *p* and *q* respectively. These values are chosen
by a random number generator and checked for primality with the Miller-Rabin
primality check. The only other restriction on these numbers is that *q* must
be a divisor of *p-1*. This is ensured by first generating the random prime
*q*, then generating a random prime *R* and deriving *p* from the formula below.

*q* | *p*-1

*p*-1 = *qx*+0

*p* = *qx*+1

*R* = *qx*+*r*

*R*-*p* = *y*

(*qx*+*r*)-(*qx*+1) = *y*

*r*-1 = *y*

*R*-*y* = *p*

If the computed *p* value does not pass the Miller-Rabin primality test, it is
generated again from a new random *R* value until it passes.

The value *g* (also known as the generator) is generated such that it is in a 
subgroup of *q* in the multiplicative group GF(*p*) and 1 < *g* < *p*.

### Key Pair Generation

Compared to the parameter generation, the public/private key pair generation
is relatively straight forward. The private key *x* is randomly generated 
such that 1 < *x* < *q*. The public key *y* is then generated using modular
exponentiation, using the formula *y* = *g*^*x* mod *p*. As discussed earlier,
the discrete logarithm problem ensures that the public and private keys are 
mathematically related and that the private key cannot (reasonably) be derived 
from the public key.

### Future Work

To achieve a fully functional application, the work that remains is to 
implement the signing and verification functions. The public and private 
keys will be saved to separate files once generated. The user will then 
be able to specify these files when running the sign and verify commands.
An implementation of the SHA-1 hash function is also required to create
a message digest when signing and verifying. This will ideally be 
implemented in the application. An outside implementation of the SHA-1 
hash function will only be utilized if it is absolutely necessary for 
having the application fully functional and complete on time. 

# References

[1] https://nvlpubs.nist.gov/nistpubs/FIPS/NIST.FIPS.186-4.pdf
[2] https://www.makeuseof.com/introduction-to-digital-signature-algorithm/
[3] https://www.lifewire.com/what-is-sha-1-2626011
[4] https://www.khanacademy.org/computing/computer-science/cryptography/modern-crypt/v/discrete-logarithm-problem
[5] https://www.khanacademy.org/computing/computer-science/cryptography/modarithmetic/a/fast-modular-exponentiation