package main

import (
	"crypto/x509/pkix"
	"flag"
	"fmt"
	"net"
	"regexp"
	"strings"
)

type Args struct {
	caKey, caCert *string
	domains, ipAddresses ArgsArr
	org, iorg, iunit, icountry, ilocale, iaddress, ipostal ArgsArr
}

type ArgsArr struct {
	a []string
}

func (arr *ArgsArr) sumFlagFunc(arg string) error {
	arr.a = append(arr.a, arg)
	return nil
}

func (arr *ArgsArr) sumFlagFuncIP(arg string) error {
	if arg == "" {
		return nil
	}
	if net.ParseIP(arg) == nil {
		return fmt.Errorf("Invalid IP address %q\n", arg)
	}
	arr.a = append(arr.a, arg)
	return nil
}

func (arr *ArgsArr) sumFlagFuncDomain(arg string) error {
	if arg == "" {
		return nil
	}
	domainRe := regexp.MustCompile("^[A-Za-z0-9.*-]+$")
	if !domainRe.MatchString(arg) {
		return fmt.Errorf("Invalid domain name %q\n", arg)
	}
	arr.a = append(arr.a, arg)
	return nil
}

func (args *Args) assignIssuerFlags() {
	flag.Func("issuer", "Issuing organization common name", args.org.sumFlagFunc)
	flag.Func("organization", "Issuing organization", args.iorg.sumFlagFunc)
	flag.Func("unit", "Issuing unit of organization (e.g., IT)", args.iunit.sumFlagFunc)
	flag.Func("country", "Issuer's country", args.icountry.sumFlagFunc)
	flag.Func("locality", "Issuer's locality (i.e., city)", args.ilocale.sumFlagFunc)
	flag.Func("address", "Issuer's address", args.iaddress.sumFlagFunc)
	flag.Func("postal-code", "Issuer's postal code (in 🇺🇸 called a “ZIP” code)", args.ipostal.sumFlagFunc)
}

func (args *Args) assignStringFlags() {
	args.caKey = flag.String("ca-key", "minica-key.pem", "Root private key filename, PEM encoded.")
	args.caCert = flag.String("ca-cert", "minica.pem", "Root certificate filename, PEM encoded.")
}

func (args *Args) assignTargetFlags() {
	flag.Func("domains", "Comma separated domain names to include as Server Alternative Names.", args.domains.sumFlagFuncDomain)
	flag.Func("ip-addresses", "Comma separated IP addresses to include as Server Alternative Names.", args.ipAddresses.sumFlagFuncIP)
}

func (args *Args) assignFlags() {
	args.assignStringFlags()
	args.assignTargetFlags()
	args.assignIssuerFlags()
}

func (args *Args) parseIssuer() (pkix.Name) {
	return pkix.Name{
		CommonName: strings.Join(args.org.a, ","),
		Organization: args.iorg.a,
		OrganizationalUnit: args.iunit.a,
		Country: args.icountry.a,
		Locality: args.ilocale.a,
		StreetAddress: args.iaddress.a,
		PostalCode: args.ipostal.a,
	}
}
