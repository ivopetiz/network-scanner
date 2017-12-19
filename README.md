# Go-Portscanner

TCP Network Port Scanner written in Go, nmap style. 

## Installation

```sh
 $ 
 git clone https://github.com/ivopetiz/portscanner.git portscanner
 $ cd portscanner
 $ go build portscanner.go portslist.go
```

## Examples

Different ways to use Go-Portscanner.

#### Get help

```sh
 $ portscanner -h
 Usage of portscanner:
   -A   Scans all ports from port 1 to 1024
   -p  	Port or ports to scan (default "80")
```

Shows the output above, presenting the available options.

#### Localhost Portscanner

```sh
$ portscanner 
```

Scans all local machine ports, from 1 to 1024.

#### Network Machine Full Portscan

```sh
$ portscanner -p 1-65535 210.67.210.76
```

Scan all TCP ports from a network machine.

#### Network Discovery

```sh
$ portscanner -p 21,80 192.168.0.1-192.168.0.254
```

Looks for HTTP and FTP servers on 192.168.0.0/24.

#### Ping Servers

```sh
$ portscanner -p 22 113.213.200.101 4.0.75.4 84.188.238.94 11.2.224.214 153.194.246.247
```

Checks if SSH servers are up.

## TODO
* Improve performance.
* Read IP:Port pair from file.
* Tests