# Foreign Keys

Up: [Primary Keys](primary_keys)
Brother(s):
TARGET DECK

Foreign keys help to establish [Referential Integrity](referential_integrity) between tables.

If a column in `Table1` references a column in `Table2`, that reference cannot be to something that does not yet exist.

For example, you cannot have scores in a `Scores` table for a student that doesn't exist in the `Student` table yet.

Foreign keys are declared at [Creating Tables in SQL|table creation](creating_tables_in_sql|table_creation), like so:

```SQLite
FOREIGN KEY (columnName) REFERENCES TableName(columnName)
```

[Maintaining Referential Integrity using Foreign Keys](maintaining_referential_integrity_using_foreign_keys)


































#### Why:
#### How:









