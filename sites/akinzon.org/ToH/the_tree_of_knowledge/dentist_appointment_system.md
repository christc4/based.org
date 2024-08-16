# Dentist appointment system

Up: [Comp1216 Coursework](comp1216_coursework)
Brother(s):
TARGET DECK

![Pasted image 20240308071751.png](pasted_image_20240308071751.png)
 obviously a class for each bullet point here

possibly a `WorkingDay` class to model each working day, to act as a container for the `WorkingHour` class

`WorkingHour` class has number to keep track of slots available, which can be incremented or decremented based on certain cases, perhaps a slot has been freed in an hour because someone has had to cancel (increment) or since someone has made a booking, a slot has been taken up (decrement).

is the `DentistPractice` the main central big boi class I've been looking for?

`Dentist`s, `Patient`s and `PracticeAdministrator`s all have a name and a number, so will inherit those two attributes and 4 getters/setters from a `Person` class

`Treatment` class? has a `type` string for the treatment, could argue that each treatment requires certain consumables

relation between `Treatment` and `Consumable`: treatments **use a** consumable/ set of consumables, (binary relation)

Defo need a way to update treatments from required to received

Have absence on `WorkingDays` and `WorkingHours`

`WorkingDays` have `WorkingHours` (composition)

`WorkingHours` have `appointments` (aggregation)

























#### Why:
#### How:









