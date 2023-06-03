# SSL Generate

## Easy way

Use generate.go

---

## Manual

### 1. Create a root CA certificate

- Create the root key
```bash
openssl ecparam -out root.key -name prime256v1 -genkey
```

- Create a Root Certificate and self-sign it
```bash
openssl req -new -sha256 -key root.key -out root.csr
```

- Generate the Root Certificate
```bash
openssl x509 -req -sha256 -days 365 -in root.csr -signkey root.key -out root.crt
```

### 2.Create a server certificate

- Create the certificate's key
```bash
openssl ecparam -out server.key -name prime256v1 -genkey
```

- Create the CSR (Certificate Signing Request)
```bash
openssl req -new -sha256 -key server.key -out server.csr
```

- Generate the certificate with the CSR and the key and sign it with the CA's root key

Replace \<\<DOMAIN>>
```bash

cat > "server.ext" <<EOF
authorityKeyIdentifier=keyid,issuer
basicConstraints=CA:FALSE
keyUsage = digitalSignature, nonRepudiation, keyEncipherment, dataEncipherment
subjectAltName = @alt_names

[alt_names]
DNS.1 = <<DOMAIN>>
EOF

openssl x509 -req -in server.csr -CA  root.crt -CAkey root.key -CAcreateserial -out server.crt -days 365 -sha256 -extfile server.ext

```
- Verify the created certificate
```bash
openssl x509 -in server.crt -text -noout
```
