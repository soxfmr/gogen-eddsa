# gogen-eddsa
Generate a fresh Ed25519 key pairs in PKCS#8 / PKIX specification

# Usage

```shell
go run main.go mysign 

2021/09/16 17:32:13 Ed25519 Private Key: <REMOVED> 
2021/09/16 17:32:13 Ed25519 Public Key: <REMOVED>
2021/09/16 17:32:13 The Ed25519 Private key was saved in PKCS#8 specification at mysign.key
2021/09/16 17:32:13 The Ed25519 Public key was saved in PKIX/X.509 - SubjectPublicKeyInfo specification at mysign.pub
```
