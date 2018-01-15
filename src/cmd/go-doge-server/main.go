// Code generated by go-swagger; DO NOT EDIT.

package main

import (
	"log"
	"os"
	"flag"

	loads "github.com/go-openapi/loads"
	flags "github.com/jessevdk/go-flags"

	"github.com/hunterBhough/go-doge/src/restapi"
	"github.com/hunterBhough/go-doge/src/restapi/operations"
)

// This file was generated by the swagger tool.
// Make sure not to overwrite this file after you generated it because all your edits would be lost!
var portFlag = flag.Int("port", 3000, "Port to run this service on")

func main() {

	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewGoDogeAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	server.Port = *portFlag
	parser := flags.NewParser(server, flags.Default)
	parser.ShortDescription = "Go-DOGE"
	parser.LongDescription = "An api for moving data between a form and dogechain through a bitmask"

	server.ConfigureFlags()
	for _, optsGroup := range api.CommandLineOptionsGroups {
		_, err := parser.AddGroup(optsGroup.ShortDescription, optsGroup.LongDescription, optsGroup.Options)
		if err != nil {
			log.Fatalln(err)
		}
	}

	if _, err := parser.Parse(); err != nil {
		code := 1
		if fe, ok := err.(*flags.Error); ok {
			if fe.Type == flags.ErrHelp {
				code = 0
			}
		}
		os.Exit(code)
	}

	server.ConfigureAPI()

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}

}