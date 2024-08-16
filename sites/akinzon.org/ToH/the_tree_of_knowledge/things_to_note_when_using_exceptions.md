# Things to Note When Using Exceptions

Up: [Further Exceptions](further_exceptions)
Brother(s):
TARGET DECK

Since exceptions are [classes](classes), they can inherit from other classes.

Therefore, some exceptions can be covered by more general ones.

Catching these general ones at the top of a `try...catch` block stop any of its [subclasses](subclasses) down below from being caught.

Furthermore, catching the base exception (`Exception`) is bad practice; you must be specific and methodical about which exceptions are caught and which can be deferred.



































#### Why:
#### How:









