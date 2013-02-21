# Pingdom IP Tool
Turns [Pingdom](http://www.pingdom.com)'s [feed](https://my.pingdom.com/probes/feed) of monitoring hosts into arbitrary config files

## Install
You need to have Go: http://golang.org/doc/install

Run "go get github.com/jda/pingdomiptool"

Go will download, compile, and install pingdomiptool in your $GOPATH/bin folder.

## Usage
    jda@snow:~$ pingdomiptool -h
    Usage of pingdomiptool:
      -out="": Output file (STDOUT if not used)
      -tmpl="": Template file

If you run pingdomiptool with no arguments it will print a list of Pingdom IPs to standard out.

## Templates
    {{range .Hosts}}
    {{.Title}}: {{.Status}}
     {{.Descr}}
     {{.IP}}
     {{.Host}}
     {{.Country}}
     {{.City}}
    {{end}}

{{range .Hosts}} through {{end}} is repeated for every host. You can use any of the above attributes in your template.

There are a couple example templates included:
+ [IPTables commands](iptables.tmpl)
+ [Unbound Resolver access control](unbound.tmpl)
