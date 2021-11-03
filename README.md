# vigenere-cypher

Perform a simple Vinegere cipher substitution of ASCII alphabetical characters. 
Ignores special ASCII characters, and non ASCII characters.

## How to run:
```bash
# download the binary
wget https://github.com/nevilleomangi/vigenere-cypher/releases/download/v0.1/vigenere
chmod u+x ./vigenere
./vigenere

# OR

git clone https://github.com/nevilleomangi/vigenere-cypher
# git clone git@github.com:nevilleomangi/vigenere-cypher.git 
cd vigenere-cypher
go run .
```
## Usage examples:
```bash
./vigenere enc -k apples "lorem ipsum dolor sit amet" # ldgpq aphjx hgldg dml abte
./vigenere dec -k apples "ldgpq aphjx hgldg dml abte" # lorem ipsum dolor sit amet
```
## Tests:
```
go test -v vigenere
```

## The Vigenere-cypher

Each letter in the message is shifted by a different amount, based on the key. The key is repeated if it is shorter than the message. 

The Vigenere cipher can be broken through [Kasiski Analysis](https://en.wikipedia.org/wiki/Kasiski_examination).
