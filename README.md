# polis
server configuration management

Its a server configuration management tool which relies on concepts of state machine to achieve final desired state from initial state of the server

Following are the steps to use this tool:
```
cd $HOME
mkdir -p go/github.com/adkhare
git clone https://github.com/adkhare/polis.git
cd polis
./bootstrap.sh
go run main.go
```

This should execute everything that is configured in `main.go` - `GetPolisStruct` function

Following is the structure of the configuration examples for 3 main modules:
Package:
```
"apache2_package": { // ID of the configuration which is unique across the whole configuration
    ModuleType: "Package", // Type of the Module
    Ensure:     true, // Ensure to either apply/unapply the module
    Module: polis.Package{ // Specific Module level configuration
        Name: "apache2", // Name of the package
    },
},
```

File:
```
"apache2_index_php": {
    ModuleType: "File",
    Ensure:     true,
    Triggers:   "apache2_service",
    Module: polis.File{
        Path: "/var/www/html/index.php", // Path of the file
        Contents: `<?php

header("Content-Type: text/plain");

echo "Hello, world!\n";`, // Contents of the file
        Owner: "root", // Owner of the file
        Group: "root", // Group of the file
        Perm:  0644, // Mode of the file
    },
},
```

Service:
```
"apache2_service": {
    ModuleType:    "Service",
    Ensure:        true,
    TriggerAction: "restart",
    Module: polis.Service{
        Name: "apache2", // Name of the service
    },
},
```