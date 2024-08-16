# Primary Keys

Up: [Keys in SQL](keys_in_sql)
Brother(s):
TARGET DECK

Primary keys in SQL are used to ensure the uniqueness of every row.
	- Primary keys can be a single attribute or [Composite Keys](composite_keys)
	- If not provided, SQL implementations (like [SQLite](sqlite)) can provide an auto-generated `rowid`

Primary Keys are added at [Creating Tables in SQL|table creation](creating_tables_in_sql|table_creation), and cannot be altered in any way afterwards.

Primary keys are declared in two ways:

```SQL
Field field_type PRIMARY KEY
```

or 

```SQL
PRIMARY KEY(field1, field2)
```































#### Why:
#### How:









