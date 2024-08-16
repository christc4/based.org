# Software Modelling Lab 4

Up: [Software Modelling Labs](software_modelling_labs)
Brother(s):
TARGET DECK

****

`The customer selects “Buy weekly card” and is taken to the “Select destinations” screen`

boundary: `buy weekly card` button

control: mechanism that directs customer to `select destinations` screen

entity: user entity containing any user data entered thus far

****

`The customer selects a destination and the “Add to Basket” button is enabled.`

boundary: buttons to specify destination selected, add to basket

control: mechanism that allows customer to add to basket only if a destination is selected.

entity: user entity but now with a basket, (composition), and the basket that contains the ticket

****

`The customer selects “Add to Basket” and is taken to the “Basket” screen`

boundary: `add to basket` button, `basket` screen

control: mechanism that takes customer to `basket` screen, and also adds the item to the customer's basket.

entity: user entity, basket data

****

`The customer selects “Pay” and is take to the “Select payment type” screen`

boundary: `pay` button, `select payment type` screen

control: mechanism that takes customer to payment type screen

entity: user entity, basket data

****

`The customer selects “Credit Card” and enters his card`

boundary: actual physical area to enter your credit cart, `credit card` button

control: mechanism for reading the physical card, and displaying input screen

entity: user data with credit card details added for the sake of purchase, basket data

****
`The distributor validates the card and requests his PIN`

boundary: validation of the credit card, the `enter pin` screen

control: mechanism that directs customer to `enter pin` screen, distributor that validates the card

entity: user data, basket data

****

`The customer enters a PIN and the distributor validates it by reference to the bank`

boundary: input area for entering pin/`enter pin` screen

control: mechanism for validating pin, displaying pin screen

entity: user data possibly with bank balance info to ensure payment can happen as normal, basket data, perhaps bank because its data is being interacted with?

****

`The distributor takes payment and returns the credit card to the customer`

boundary: possible 'payment processing' screen

control: mechanism for displaying the waiting screen, subtracting ticket cost from bank balance

entity: user data with updated bank balance, basket data cleared since payment is made, receipt of payment created, bank data notified to subtract the cost of the ticket

****

`The weekly season card is printed and the customer takes it`

boundary: possible 'collect ticket' screen

control: mechanism for printing ticket, and displaying 'collect ticket' screen

entity: user data, receipt data





























#### Why:
#### How:









