# strength [![Go Reference](https://pkg.go.dev/badge/github.com/pchchv/strength.svg)](https://pkg.go.dev/github.com/pchchv/strength) [![Go Report Card](https://goreportcard.com/badge/github.com/pchchv/strength)](https://goreportcard.com/report/github.com/pchchv/strength)
**Validating password strength in Go without transitive dependencies.**

![XKCD Passwords](https://imgs.xkcd.com/comics/password_strength.png)

## Usage

```go
package main

import (
    "github.com/pchchv/strength"
)

func main(){
    const minEntropyBits = 60

    // entropy is a float64, representing the strength in base 2 (bits)
    entropy := strength.GetEntropy("a longer password")

    // if the password has sufficient entropy, err is nil,
    // otherwise a formatted error message is given explaining
    // how to increase the strength of the password
    // (safe to show to the client)
    err := strength.Validate("password", minEntropyBits)
}
```

## What Entropy Value Should I Use?

![entropy](https://external-preview.redd.it/rhdADIZYXJM2FxqNf6UOFqU5ar0VX3fayLFpKspN8uI.png?auto=webp&s=9c142ebb37ed4c39fb6268c1e4f6dc529dcb4282)

Keep in mind that attackers most likely don't just brute-force passwords, and if you need protection against regular passwords or [PWNed passwords](https://haveibeenpwned.com/), you'll have to do extra work. This library is lightweight, does not load large datasets and does not access external services.

## How It Works

First, we determine the "base" number. The base is a sum of the different "character sets" found in the password.

We've *arbitrarily* chosen the following character sets:

* 26 lowercase letters
* 26 uppercase letters
* 10 digits
* 5 replacement characters - `!@$&*`
* 5 seperator characters - `_-., `
* 22 less common special characters - `"#%'()+/:;<=>?[\]^{|}~`


Using at least one character from each set your base number will be 94: `26+26+10+5+5+22 = 94`

Every unique character that doesn't match one of those sets will add `1` to the base.

If you only use, for example, lowercase letters and numbers, your base will be 36: `26+10 = 36`.

After we have calculated a base, the total number of brute-force-guesses is found using the following formulae: `base^length`

A password using base 26 with 7 characters would require `26^7`, or `8031810176` guesses.

Once we know the number of guesses it would take, we can calculate the actual entropy in bits using `log2(guesses)`. That calculation is done in log space in practice to avoid numeric overflow.

### Additional Safety

We try to err on the side of reporting *less* entropy rather than *more*.

#### Same Character

With repeated characters like `aaaaaaaaaaaaa`, or `111222`, we modify the length of the sequence to count as no more than `2`.

* `aaaa` has length 2
* `111222` has length 4

#### Common Sequences

Common sequences of length three or greater count as length `2`.

* `12345` has length 2
* `765432` has length 2
* `abc` has length 2
* `qwerty` has length 2

The sequences are checked from back->front and front->back. Here are the sequences we've implemented so far, and they're case-insensitive:

* `0123456789`
* `qwertyuiop`
* `asdfghjkl`
* `zxcvbnm`
* `abcdefghijklmnopqrstuvwxyz`