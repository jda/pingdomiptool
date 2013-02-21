package main

import (
	"fmt"
	"os"
	"text/template"
	"encoding/xml"
	"flag"
	"net/http"
)

const pingdomFeed = "https://my.pingdom.com/probes/feed"
const ver = "1.0"

const defaultTmpl = "{{range .Hosts}}{{.IP}}\n{{end}}"

type Host struct {
	Title   string `xml:"title"`
	Descr   string `xml:"description"`
	IP      string `xml:"pingdom ip"`
	Host    string `xml:"pingdom hostname"`
	Country string `xml:"pingdom country"`
	City    string `xml:"pingdom city"`
	Status  string `xml:"pingdom state"`
}

type TmplHosts struct {
	Hosts []Host
}

func main() {
	tmplfn := flag.String("tmpl", "", "Template file")
	outfn := flag.String("out", "", "Output file (STDOUT if not used)")
	flag.Parse()

	// Get feed
	client := &http.Client{}
	req, err := http.NewRequest("GET", pingdomFeed, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not fetch Pingdom feed: %v\n", err)
		os.Exit(2)
	}

	req.Header.Set("User-Agent", "pingdomiptool "+ver)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error fetching Pingdom feed: %v\n", err)
		os.Exit(2)
	}
	defer resp.Body.Close()

	var hosts []Host
	
	decoder := xml.NewDecoder(resp.Body)
	for {
		token, _ := decoder.Token()
		if token == nil {
			break
		}

		switch s := token.(type) {
		case xml.StartElement:
			if s.Name.Local == "item" {
				var h Host
				decoder.DecodeElement(&h, &s)
				hosts = append(hosts, h)
			}
		}
	}
	
	tHosts := TmplHosts{Hosts: hosts}
	tmpl := template.New("hoststmpl")
	if *tmplfn != "" {
		tmpl = template.New(*tmplfn)
		template.Must(tmpl.ParseFiles(*tmplfn))
	} else {
		template.Must(tmpl.Parse(defaultTmpl))
	}
	
	out := os.Stdout
	if *outfn != "" {
		out, err := os.Create(*outfn)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening output file: %v\n", err)
			os.Exit(2)
		}
		defer out.Close()
	}
	
	err = tmpl.Execute(out, tHosts)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error applying template: %v\n", err)
		os.Exit(2)
	}
}
