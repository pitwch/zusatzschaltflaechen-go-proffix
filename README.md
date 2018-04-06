
# Zusatzschaltflächen für PROFFIX in Go

In diesem Repo sind einige funktionierende Beispiele von Zusatzschaltflächen für PROFFIX in Go.

Mehr Infos über [Go / Golang](https://golang.org) bzw. [PROFFIX](http://www.proffix.net) 

# Übersicht

## Google Maps
Zeigt die gewünschte Adresse aus PROFFIX in Google Maps an. Es ist keine weitere Konfiguration erforderlich.

## Email versenden
Versendet ein Email an die ausgewählte Adresse. Die Konfiguration erfolgt mit dem File config.json (Beispiel config.json.ex umbenennen und editieren). 

Die config.json **muss auf derselben Ebene** sein wie die Executable.

# Installation

## Variante 1: Vorerstellter Release downloaden


 1.  [Aktuellsten Relase herunterladen](https://github.com/pitwch/zusatzschaltflaechen-go-proffix/releases/latest)
 2.  ZIP - File entpacken, EXE unter sinnvollem Pfad speichern
 3.  **Fakultativ**: Wenn eine Config benötigt wird (z.B. Send-Email) diese editieren und am selben Ort wie die EXE ablegen.  
 4.  Zusatzfeld **Schaltfläche** in PROFFIX auf EXE verlinken
 

## Variante 2: Release selber erstellen

 1.  [Go installieren](https://golang.org/dl/)
 2. Repo herunterladen / Klonen
 3. In Unterverzeichnis navigieren (z.B.  /google-maps)
 4. Mit `go build`wird eine Executable generiert
 5. Diese kann über Zusatzfelder in PROFFIX angesteuert werden
