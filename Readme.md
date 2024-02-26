# Hello world

Hoi allemaal

# Bug

De applicatie berekent de bruto prijs (incl. BTW) van de orders in het gegeven json bestand. Er gaat alleen iets mis: de totaal prijs is veel te hoog in sommige gevallen!

# Tests

## Uitvoeren van tests

Je kunt de tests op verschillende manieren uitvoeren.

### Via de IDE (GoLand)

In GoLand kan je een configuratie maken die i.p.v. `go run` een `go test` uitvoert. Kies hiervoor 'Go Test' in de configuratie, in plaats van 'Go Build' die je gewend bent.
Je kan nu alle tests runnen door op de groene play knop bovenin te drukken. Een enkele test runnen kan door op de play knop naast de test functie te drukken.

In VS Code kun je specifieke tests uitvoeren door op de 'play' knop links van de test functie te klikken, of op 'run test' of 'debug test'

### Via de terminal

Met het volgende commando voer je alle tests van de hele applicatie uit:

```bash
go test ./...
```

Met het volgende commando voer je de tests van een specifieke package uit:

```bash
go test che/dt/week2/orders/business
```

## Opdracht: Schrijf unit testen voor de onderstaande test cases

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

> Bedenk zelf test cases voor de `CalculateTotalOrderPrice` functie. Denk hierbij aan de happy path en sad paths van de functie.


