package main

import (
	"log"
)

func main() {
	var args *Args = &Args{}
	args.parse()
	issuer, err := getIssuer(args)
	if err != nil {
		_, err = sign(issuer, args)
	}
	if err != nil {
		log.Fatal(err)
	}
}
