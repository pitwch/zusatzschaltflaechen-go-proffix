package main

import (
	"fmt"
	"gopkg.in/ini.v1"
	"os"
)

func main() {

	// Get Temp - Filepath
	filepath := os.Args[1]
	pxzu, err := ini.Load(filepath)

	if err != nil {
		fmt.Printf("Error reading Zusatzfeld: %v", err)
	}

	// Set some return values
	pxzu.Section("Return").Key("AktModul").SetValue("pxAuftrag")
	pxzu.Section("Return").Key("SelKeys").SetValue("600847")
	pxzu.Section("Return").Key("SelKey").SetValue("600847")
	
	// Test another variant
	pxzu.Section("Settings").Key("SelKey").SetValue("600847")
	pxzu.Section("Settings").Key("SelKeys").SetValue("600847")
	
	// Parameter isn't documented in PROFFIX Docs ?
	pxzu.Section("Msg").Key("Test").SetValue("Demo")

	// Save back to Temp - File
	err = pxzu.SaveTo(filepath)
	if err != nil {
		fmt.Printf("Error reading Zusatzfeld: %v", err)
	}

}
