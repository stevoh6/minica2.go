# `minica2` v1.0.0

* © 2016–2022 Jacob Hoffman-Andrews, Fredrick R. Brennan &lt;copypaste@kittens.ph&gt;, and Minica(2) Project Authors

Minica2 is a simple CA intended for use in situations where the CA operator
also operates each host where a certificate will be used. It automatically
generates both a key and a certificate when asked to produce a certificate.
It does not offer OCSP or CRL services. Minica2 is appropriate, for instance,
for generating certificates for RPC systems or microservices.

On first run, minica2 will generate a keypair and a root certificate in the
current directory, and will reuse that same keypair and root certificate
unless they are deleted.

On each run, minica2 will generate a new keypair and sign an end-entity (leaf)
certificate for that keypair. The certificate will contain a list of DNS names
and/or IP addresses from the command line flags. The key and certificate are
placed in a new directory whose name is chosen as the first domain name from
the certificate, or the first IP address if no domain names are present. It
will not overwrite existing keys or certificates.

The certificate will have a validity of 2 years and 30 days.

## Installation

First, install the [Go tools](https://golang.org/dl/) and set up your `$GOPATH`.
Then, run:

`go install github.com/ctrlcctrlv/minica2.go@latest`

If you prefer to compile manually:

```bash
cd # $HOME/Workspace
git clone https://github.com/ctrlcctrlv/minica2.go.git
cd minica2.go
make install
```

Mac OS users could alternatively use Homebrew: `brew install minica2`

## Example usage

```
# Generate a root key and cert in minica2-key.pem, and minica2.pem, then
# generate and sign an end-entity key and cert, storing them in ./foo.com/
$ minica2 --domain foo.com

# Wildcard
$ minica2 --domain '*.foo.com'
```

## Usage
```
Usage of minica2:

Minica is a simple CA intended for use in situations where the CA operator
also operates each host where a certificate will be used. It automatically
generates both a key and a certificate when asked to produce a certificate.
It does not offer OCSP or CRL services. Minica is appropriate, for instance,
for generating certificates for RPC systems or microservices.

On first run, minica will generate a keypair and a root certificate in the
current directory, and will reuse that same keypair and root certificate
unless they are deleted.

On each run, minica will generate a new keypair and sign an end-entity (leaf)
certificate for that keypair. The certificate will contain a list of DNS names
and/or IP addresses from the command line flags. The key and certificate are
placed in a new directory whose name is chosen as the first domain name from
the certificate, or the first IP address if no domain names are present. It
will not overwrite existing keys or certificates.

  -address value
    	Issuer's address
  -ca-cert string
    	Root certificate filename, PEM encoded. (default "minica.pem")
  -ca-key string
    	Root private key filename, PEM encoded. (default "minica-key.pem")
  -country value
    	Issuer's country
  -domain value
    	Comma separated domain names to include as Server Alternative Names.
  -ip-address value
    	Comma separated IP addresses to include as Server Alternative Names.
  -issuer value
    	Issuing organization common name
  -locality value
    	Issuer's locality (i.e., city)
  -organization value
    	Issuing organization
  -postal-code value
    	Issuer's postal code (in 🇺🇸 called a “ZIP” code)
  -unit value
    	Issuing unit of organization (e.g., IT)
```

## License

```
MIT License

Copyright (c) 2016 Jacob Hoffman-Andrews
Copyright (c) 2022 Fredrick R. Brennan <copypaste@kittens.ph>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```
