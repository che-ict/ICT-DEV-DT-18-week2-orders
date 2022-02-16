# Bug

De applicatie berekent de bruto prijs (incl. BTW) van de orders in het gegeven json bestand. Er gaat alleen iets mis: de totaal prijs is veel te hoog in sommige gevallen!

# Improvement

Er zit een veld voor korting op de order: pas de korting toe op de totaal prijs van elk order. Als de korting 7% is, dan heeft het veld `korting` de waarde `.07`.

$$totaal = (1 - korting) \sum_{r \in Regels} r_{aantal} \times r_{prijs} \times r_{btw}$$

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