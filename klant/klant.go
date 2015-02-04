package klant

import (
	"TAT/bestelling"
	//	"fmt"
	"gopkg.in/qml.v1"
)

type Klant struct {
	AchterNaam   string
	VoorNaam     string
	BezorgAdress Adres
	FactuurAdres Adres
	Betaaltype   int
	Bestellingen []bestelling.Bestelling
	qml.Object
}

type Adres struct {
	Land        string
	Postcode    string
	Stad        string
	Straat      string
	Huisnummer  string
	Opmerkingen string
}
