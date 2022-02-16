# Fork

Spreek met een persoon af om samen te werken aan deze opdracht. Neem de volgende stappen:

1. **Eén van jullie** klikt rechtsboven op `Fork`. Hiermee kopieert hij/zij de repository naar zijn/haar eigen omgeving. 
2. Vervolgens gaat die persoon naar `Settings` &gt; `Collaborators`. Bij `Manage access` voeg je de ander persoon toe, zodat jullie beiden kunnen werken aan deze repo. 
3. Ga **beiden** in jullie github repo naar `Code` (eerste tab) en klik op de groene knop `Code`, een dropdown verschijnt.
4. Controleer dat `HTTPS` is geselecteerd, en klik op het `kopieer` icoontje naast de url
5. Open een terminal in je Visual Code omgeving en zorg dat je in je DT directory zit
6. Type: `git clone <plak hier de gekopieerde url>`
7. Open deze nieuw aangemaakt folder (`ICT-DEV-DT-18-week2-orders`) in Visual Code

# Bug

De applicatie berekent de bruto prijs (incl. BTW) van de orders in het gegeven json bestand. Er gaat alleen iets mis: de totaal prijs is veel te hoog in sommige gevallen!

# Tests

Schrijf tests voor de volgende scenarios:

 - `CalculateTotalOrderPrice(...)` voor hardcoded order met 0 regels geeft 0 terug
 - `CalculateTotalOrderPrice(...)` voor hardcoded order met 1 hardcoded regel geeft juiste prijs terug
 - `CalculateTotalOrderPrice(...)` voor hardcoded order met 2 hardcoded regels geeft juiste prijs terug
 - `CalculateTotal(...)` voor hardcoded lijst van 2 orders geeft juiste prijs terug
 - `CalculateTotal(...)` voor hardcoded lijst van lege lijst orders geeft 0 terug

Denk ook aan:

 - `korting` van meer dan 1 geeft foutmelding (of wellicht andere grens afgesproken met product owner)
 - `btw` anders dan 0, 9 of 21 procent geeft foutmelding
 - negatieve `prijs` geeft foutmelding
 - negatieve `aantal` geeft foutmelding
 - welke kun je nog meer bedenken?

# Improvement

Er zit een veld voor korting op de order: pas de korting toe op de totaal prijs van elk order. Als de korting 7% is, dan heeft het veld `korting` de waarde `.07`.

$$totaal = (1 - korting) \sum_{r \in Regels} r_{aantal} \times r_{prijs} \times r_{btw}$$

> Bovenstaande formule wordt wel goed weergegeven in Visual Code preview, maar niet op github.

Schrijf ook een test om deze improvement te bewaken.

# Json

Json is gegenereerd middels `https://json-generator.com/` met de volgende definitie:

```js
[
  '{{repeat(50)}}',
  {
    id: '{{index()}}',
    organisatie: '{{company()}}',
    regels: [
      '{{repeat(10, 20)}}',
      {
        nummer: '{{index()}}',
        btw: '{{random(21, 9, 0)}}%',
        aantal: '{{integer(1, 10)}}',
        prijs: '{{floating(10, 500, 2)}}',
        omschrijving: '{{lorem(3, "word")}}'
      }
    ],
    korting: '{{floating(0, .2, 2)}}',
    definitief: '{{bool()}}',
    datum: '{{date(new Date(2021, 0, 1), new Date(), "YYYY-MM-ddThh:mm:ssZ")}}'
  }
]
```