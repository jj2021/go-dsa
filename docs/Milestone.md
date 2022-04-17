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
of the signatory, integrity of the signed data, and non-repudiation.

## Methedology

# References