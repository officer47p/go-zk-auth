# Zero-Knowledge Authentication (ZK-Auth)

A Go implementation of a zero-knowledge proof system for authentication, based on discrete logarithm problems and Schnorr-like protocols.

## Overview

This project implements a zero-knowledge authentication scheme that allows a prover to demonstrate knowledge of a secret without revealing the secret itself. The implementation uses modular exponentiation and discrete logarithm problems as the foundation for cryptographic security.

## Mathematical Foundation

The protocol is based on the following mathematical operations:

### Key Components

- **Public Parameters**: `α` (alpha), `β` (beta), `p` (large prime), `q` (order)
- **Secret**: `x` (private key known only to the prover)
- **Public Keys**: `y₁ = α^x mod p`, `y₂ = β^x mod p`
- **Challenge**: `c` (random challenge from verifier)
- **Random Value**: `k` (random value chosen by prover)

### Protocol Flow

1. **Commitment Phase**: Prover computes commitments:
   - `r₁ = α^k mod p`
   - `r₂ = β^k mod p`

2. **Challenge Phase**: Verifier sends a random challenge `c`

3. **Response Phase**: Prover computes response:
   - `s = (k - c·x) mod q`

4. **Verification Phase**: Verifier checks:
   - `r₁ ≟ α^s · y₁^c mod p`
   - `r₂ ≟ β^s · y₂^c mod p`

## Functions

### `Exponentiate(n, exp, p)`
Computes modular exponentiation: `n^exp mod p`

### `Solve(k, c, x, q)`
Computes the response value in the zero-knowledge proof:
- If `k ≥ c·x`: returns `(k - c·x) mod q`
- If `k < c·x`: returns `(k - c·x + q) mod q` (handles negative modular arithmetic)

### `Verify(r1, r2, y1, y2, alpha, beta, s, c, p)`
Verifies the zero-knowledge proof by checking:
- `r₁ = α^s · y₁^c mod p`
- `r₂ = β^s · y₂^c mod p`

Returns `true` if both conditions are satisfied, `false` otherwise.

## Usage Example

```go
package main

import "github.com/shopspring/decimal"

func main() {
    // Public parameters
    alpha := decimal.NewFromInt(4)
    beta := decimal.NewFromInt(9)
    p := decimal.NewFromInt(23)  // Prime modulus
    q := decimal.NewFromInt(11)  // Order
    
    // Secret and random values
    x := decimal.NewFromInt(6)   // Secret key
    k := decimal.NewFromInt(7)   // Random nonce
    c := decimal.NewFromInt(4)   // Challenge
    
    // Compute public keys
    y1 := alpha.Pow(x).Mod(p)  // y1 = α^x mod p
    y2 := beta.Pow(x).Mod(p)   // y2 = β^x mod p
    
    // Compute commitments
    r1 := alpha.Pow(k).Mod(p)  // r1 = α^k mod p
    r2 := beta.Pow(k).Mod(p)   // r2 = β^k mod p
    
    // Compute response
    s := Solve(k, c, x, q)     // s = (k - c*x) mod q
    
    // Verify the proof
    isValid := Verify(r1, r2, y1, y2, alpha, beta, s, c, p)
    // isValid should be true for correct proof
}
```

## Security Properties

- **Zero-Knowledge**: The verifier learns nothing about the secret `x`
- **Soundness**: A malicious prover without knowledge of `x` cannot convince the verifier (except with negligible probability)
- **Completeness**: An honest prover with knowledge of `x` will always convince the verifier

## Testing

Run the test suite to verify the implementation:

```bash
go test -v
```

The test includes:
- Verification of a valid proof with the correct secret
- Verification that a proof fails with an incorrect secret

## Dependencies

- [shopspring/decimal](https://github.com/shopspring/decimal): For precise decimal arithmetic

## Installation

```bash
go mod tidy
```

## Mathematical Note

This implementation handles negative modular arithmetic correctly in the `Solve` function. When `k < c·x`, the result `(k - c·x)` would be negative, so we add `q` to ensure the result is in the correct range `[0, q-1]`.

## Security Considerations

- Use cryptographically secure random number generators for `k` and challenges in production
- Ensure proper parameter selection (large primes, appropriate group orders)
- This is a educational/demonstration implementation - use established cryptographic libraries for production systems

## License

This project is for educational and research purposes.