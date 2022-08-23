package main

import (
	"fmt"
	"github.com/adkhare/polis/internal/polis"
	"log"
)

func main() {
	polis := map[string]polis.Polis{
		"apache2_package": {
			ModuleType: "Package",
			Ensure:     true,
			Module: polis.Package{
				Name: "apache2",
			},
		},
		"apache2_config": {
			ModuleType: "File",
			Ensure:     true,
			Triggers:   "apache2_service",
			Module: polis.File{
				Path: "/var/www/html/hello.php",
				Contents: `
<?php

header("Content-Type: text/plain");

echo "Hello, world!\n";`,
				Owner: "root",
				Group: "root",
				Perm:  0666,
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
	}

	for id, p := range polis {
		fmt.Printf("Starting %s\n", id)
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
