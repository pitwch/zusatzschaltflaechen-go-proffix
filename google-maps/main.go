package main

import (
	"os"
	"fmt"
	"gopkg.in/ini.v1"
	"github.com/skratchdot/open-golang/open"
	"bytes"
	"net/url"
)



func main() {
	//cfg, err := ini.Load("C:\\test.tmp")
	pxzu, err := ini.Load(os.Args[1])

	if err != nil {
		fmt.Printf("Error reading File: %v", err)
		os.Exit(1)
	}

	plz := pxzu.Section("Fields").Key("PLZ").String()
	ort := pxzu.Section("Fields").Key("Ort").String()
	land := pxzu.Section("Fields").Key("LandPRO").String()
	strasse := pxzu.Section("Fields").Key("dfsStrasse").String()

	list := []string{"https://www.google.ch/maps/place/",strasse,"+",plz,"+",ort," ",land}

	var str bytes.Buffer
	for _, l := range list {
		str.WriteString(l)
	}

	var urlmaps = url.PathEscape(str.String())

	fmt.Println("Redirect to ",urlmaps)
	open.Run(urlmaps)

}


