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

In order to make changes to the configuration, please change the struct object in main.go:GetPolisStruct function.

Following were the requirements and explanation of how configuration drives those changes:
1. If your tool has dependencies not available on a standard Ubuntu instance you may include a bootstrap.sh program to resolve them
    - ./bootstrap.sh - installs go1.19 using tar and adds the go binary to the PATH and also sets the GOPATH and GOROOT
2. Tool must provide abstraction for files, packages and service
    - `Module` is an interface which is part of `Polis` struct that is implemented implicitly by `File`, `Package` and `Service` structs
    by implementing required methods
3. Triggering capability between different modules
    - `Polis` struct configures capability for providing
        - `Triggers` which defines which ID of the module to trigger
        - `TriggerAction` which defines what is the action that needs to be executed when the Module is triggered
    Example: in the above example configurations, `apache2_index_php` triggers `apache2_index_php`. Which means when the file `index.php` is created/updated/deleted, `apache2` service will be restarted
4. Tool must be idempotent
    - `Module` interface provides function of `Check` which is used in `Apply` function to ensure that changes are made only if `Check` is `false` that facilitates the tool being idempotent
5. All current configurations that are set in `GetPolisStruct` in `main.go` will ensure to install everything, configure the `index.php` and ensure to restart the `apache2` server
6. This code currently assumes that the new hosts are exactly the same config as that of `54.221.50.166` and `52.90.125.150`