
# Zusatzschaltflächen für PROFFIX in Go

In diesem Repo sind einige funktionierende Beispiel von Zusatzschaltflächen für PROFFIX in Go.

Mehr Infos über [Go / Golang](https://golang.org) bzw. [PROFFIX](http://www.proffix.net) 

# Übersicht

## Google Maps
Zeigt die gewünschte Adresse aus PROFFIX in Google Maps an. Es ist keine weitere Konfiguration erforderlich.

## Email versenden
Versendet ein Email an die ausgewählte Adresse. Die Konfiguration erfolgt mit dem File config.json (Beispiel config.json.ex umbenennen und editieren). 

Die config.json **muss auf derselben Ebene** sein wie die Executable.

# Installation

 1.  [Go installieren](https://golang.org/dl/)
 2. Repo herunterladen / Klonen
 3. In Unterverzeichnis navigieren (z.B.  /google-maps)
 4. Mit `go build`wird eine Executable generiert
 5. Diese kann über Zusatzfelder in PROFFIX angesteuert werden
