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

# References