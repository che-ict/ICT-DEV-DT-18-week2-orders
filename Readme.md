# Bug

De applicatie berekent de bruto prijs (incl. BTW) van de orders in het gegeven json bestand. Er gaat alleen iets mis: de totaal prijs is veel te hoog in sommige gevallen!

# Tests

## Uitvoeren van tests

Je kunt de tests op verschillende manieren uitvoeren.

Met het volgende commando voer je alle tests van de hele applicatie uit:

```bash
go test ./...
```

Met het volgende commando voer je de tests van een specifieke package uit:

```bash
go test che/dt/week2/orders/business
```

In VS Code kun je specifieke tests uitvoeren door op de 'play' knop links van de test functie te klikken, of op 'run test' of 'debug test'

## Test cases

De applicatie bevat twee functies in de `business` package die je goed kunt testen:

- `CalculateTotal`
- `CalculateTotalOrderPrice`

Voor `CalculateTotalOrderPrice` kun je de volgende tests schrijven:

- `TestCalculateTotalOrderPrice_MultipleOrderLines_GivesExpectedTotal` - Test dat een order met meerdere orderregels van verschillend prijs, aantal en btw tarief het juiste totaalbedrag teruggeeft
- `TestCalculateTotalOrderPrice_NoOrderLines_GivesTotalZero` - Test dat een order zonder order regels een totaal van 0 teruggeeft
- `TestCalculateTotalOrderPrice_SingleOrderLine_GivesExpectedTotal` - Test dat een order met eén orderregel het juiste totaalbedrag teruggeeft
- `TestCalculateTotalOrderPrice_OrderLineWithDiscount_GivesExpectedTotal` - Test dat een order met een orderregel met een kortingspercentage het juiste totaalbedrag teruggeeft
- `TestCalculateTotalOrderPrice_OrderLineWith21PercentVat_GivesExpectedTotal` - Test dat een order met een orderregel met een btw tarief van 21% het juiste totaalbedrag teruggeeft
- `TestCalculateTotalOrderPrice_OrderLineWith9PercentVat_GivesExpectedTotal` - Test dat een order met een orderregel met een btw tarief van 9% het juiste totaalbedrag teruggeeft
- `TestCalculateTotalOrderPrice_NegativeQuantity_GivesError` - Test dat een order met een orderregel met een negatieve aantal een error teruggeeft
- `TestCalculateTotalOrderPrice_NegativePrijs_GivesError` - Test dat een order met een orderregel met een negatieve prijs een error teruggeeft

> Bedenk zelf test cases voor de `CalculateTotalOrderPrice` functie

# Fork

Spreek met een persoon af om samen te werken aan deze opdracht. Neem de volgende stappen:

1. **Eén van jullie** klikt rechtsboven op `Fork`. Hiermee kopieert hij/zij de repository naar zijn/haar eigen omgeving. 
2. Vervolgens gaat die persoon naar `Settings` &gt; `Collaborators`. Bij `Manage access` voeg je de ander persoon toe, zodat jullie beiden kunnen werken aan deze repo. 
3. Ga **beiden** in jullie github repo naar `Code` (eerste tab) en klik op de groene knop `Code`, een dropdown verschijnt.
4. Controleer dat `HTTPS` is geselecteerd, en klik op het `kopieer` icoontje naast de url
5. Open een terminal in je Visual Code omgeving en zorg dat je in je DT directory zit
6. Type: `git clone <plak hier de gekopieerde url>`
7. Open deze nieuw aangemaakt folder (`ICT-DEV-DT-18-week2-orders`) in Visual Code

Vervolgens maakt 1 van jullie de improvement zoals beschreven in de improvement sectie hieronder. De ander hernoemt de variabele `sum` naar `totaal` in de functie `CalculateTotal(...)`. Degene die de improvement maakt hernoemt de variabele `sum` niet! Controleer dat jullie wijzingen hebben gemaakt op dezelfde regel.

Commit jullie werk en push het allebei naar jullie remote repo (github).

# Improvement

Er zit een veld voor korting op de order: pas de korting toe op de totaal prijs van elk order. Als de korting 7% is, dan heeft het veld `korting` de waarde `.07`.

$$totaal = (1 - korting) \sum_{r \in Regels} r_{aantal} \times r_{prijs} \times r_{btw}$$

> Bovenstaande formule wordt wel goed weergegeven in Visual Code preview, maar niet op github.

Schrijf ook een test om deze improvement te bewaken.

# Json

Json is gegenereerd middels [json-generator.com](https://json-generator.com/) met de volgende definitie:

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