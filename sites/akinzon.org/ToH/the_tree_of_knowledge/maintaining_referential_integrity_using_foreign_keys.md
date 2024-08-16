# Maintaining [Referential Integrity](referential_integrity) using [Foreign Keys](foreign_keys)

Up: 
Brother(s):
TARGET DECK

The behaviour of foreign keys can be masterfully controlled using the keywords:

`ON DELETE` - Describes what happens when the record that contains the foreign key is deleted

`ON UPDATE` - Describes what happens when the record that contain the foreign key is updated

There are 4 ways to describe the required behaviour:
	 - `CASCADE` - anything that happens to the referenced record, happens to the referencing record (delete booking $\implies$ delete member)
	 - `RESTRICT` - No changes to referencing record happen
	 - `SET DEFAULT` - Sets the referencing record's foreign key column value to the default
	 - `SET NULL` Sets the referencing record's foreign key column value to `NULL`




































#### Why:
#### How:









