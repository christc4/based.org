# Synchronisation Rules of Thumb

Up: [Thread Synchronisation](thread_synchronisation)
Brother(s):
TARGET DECK

Always [Thread Locks|lock](thread_locks|lock) when updating an [Objects|object's](objects|object's) shared [Variables|fields](variables|fields)

Always lock when accessing fields that may be updated

Never lock when invoking [methods](methods) on other objects.

































#### Why:
#### How:









