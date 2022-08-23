package main

import (
	"fmt"
	"github.com/adkhare/polis/internal/polis"
	"log"
)

func main() {
	polis := GetPolisStruct()

	for id, p := range polis {
		fmt.Printf("Starting Module: %s\n", id)
		nextId, err := p.Execute()
		if err != nil {
			if err != nil {
				log.Fatal(err)
			}
		}
		if nextId != "" {

			if triggeredPolis, found := polis[nextId]; found {
				_, err := triggeredPolis.Module.TriggerExec(triggeredPolis.TriggerAction)
				if err != nil {
					log.Fatal(err)
				}
			}
		}
	}
}

func GetPolisStruct() map[string]polis.Polis {
	return map[string]polis.Polis{
		"apache2_package": {
			ModuleType: "Package",
			Ensure:     true,
			Module: polis.Package{
				Name: "apache2",
			},
		},
		"php_package": {
			ModuleType: "Package",
			Ensure:     true,
			Module: polis.Package{
				Name: "php",
			},
		},
		"libapache2-mod-php_package": {
			ModuleType: "Package",
			Ensure:     true,
			Module: polis.Package{
				Name: "libapache2-mod-php",
			},
		},
		"php-cli_package": {
			ModuleType: "Package",
			Ensure:     true,
			Module: polis.Package{
				Name: "php-cli",
			},
		},
		"php-cgi_package": {
			ModuleType: "Package",
			Ensure:     true,
			Module: polis.Package{
				Name: "php-cgi",
			},
		},
		"apache2_index_php": {
			ModuleType: "File",
			Ensure:     true,
			Triggers:   "apache2_service",
			Module: polis.File{
				Path: "/var/www/html/index.php",
				Contents: `<?php

header("Content-Type: text/plain");

echo "Hello, world!\n";`,
				Owner: "root",
				Group: "root",
				Perm:  0644,
			},
		},
		"apache2_hello_php": {
			ModuleType: "File",
			Ensure:     false, // Unapplies
			Triggers:   "apache2_service",
			Module: polis.File{
				Path:  "/var/www/html/hello.php",
				Owner: "root",
				Group: "root",
				Perm:  0644,
			},
		},
		"apache2_service": {
			ModuleType:    "Service",
			Ensure:        true,
			TriggerAction: "restart",
			Module: polis.Service{
				Name: "apache2",
			},
		},
		"php-mysql_package": {
			ModuleType: "Package",
			Ensure:     true,
			Module: polis.Package{
				Name: "php-mysql",
			},
		},
	}
}
