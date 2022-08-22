# polis
server configuration management

Its a server configuration management tool which relies on concepts of state machine to achieve final desired state from initial state of the server

Following is the example configuration yaml:

```
polis:
    package:
        id: apache2_package
        name: apache2
        ensure: present
    file:
        id: index_file
        name: /var/www/index.php
        contents: "<?php

header(\"Content-Type: text/plain\");

echo \"Hello, world!\n\";

"
        owner: root
        group: root
        mode: 0666
        ensure: present
        triggers: apache2_service
    service:
        id: apache2_service
        name: apache2
        ensure: present
        triggerAction: restart
```

Above configuration file should be able to achieve following capabilites:
1. Install a package if does not exist
2. Create a file with given contents and metadata (owner, group, mode)
3. Starts a service if the service is not running
4. Creates a dependency configuration so that a state is executed if it depends on other state
