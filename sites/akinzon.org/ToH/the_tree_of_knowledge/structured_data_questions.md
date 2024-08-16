# Structured Data Questions

Up: [Structured Data](structured_data)
Brother(s):
TARGET DECK
Leaves of the Tree of Ohara::Branch I::Semester II::1204 Data Management::Structured Data



What are the main [benefits of structured data](benefits_of_structured_data)? #flashcard 
Structured data is easier for a computer to process, which makes it easier to search through, aggregate/summarise, predict patterns, and create links between data.
<!--ID: 1709757362363-->





What are [namespaces in XML](namespaces_in_xml)? #flashcard 
Namespaces allow us to define different contexts under which XML tags are interpreted.
The same tag can be used, but interpreted differently due to distinct namespaces.
<!--ID: 1709757362385-->





What are the [Benefits of Having an XML Schema](benefits_of_having_an_xml_schema)? #flashcard 
XML schemas are good templates which can help us tell if there is an error in our XML, and it makes the task of parsing said XML easier for a computer, since there is only one structure!
<!--ID: 1709757362392-->




What are the differences between the two [XSD Types](xsd_types), SimpleType and ComplexType? #flashcard 
SimpleType can only contain text, must be empty, and cannot have attributes.
ComplexType can contain other elements, can be empty, and can have attributes.
<!--ID: 1709757362398-->




What the heck is [XPath](xpath)? #flashcard 
XPath is a query language that can be used to fetch parts of an XML document.
<!--ID: 1709757362404-->



What are [database systems](database_systems)? #flashcard 
Database systems are systems that allow for **large** amounts of data to be created, processed and managed efficiently!
<!--ID: 1710009352053-->




What can [database management systems](database_management_systems) do? #flashcard 
Database management systems can: use data models to define what a database is and use a high-level language to create and manage that database
Store large chunks of persistent data efficiently
Manage data transactions
Manage concurrent access
Manage access rights
Not lose all its data to a crash
<!--ID: 1710009352071-->




What is one of the key things a [Database Management Systems|DBMS](database_management_systems|dbms) guarantees, and what 2 things can it be split into? explain what they both are. #flashcard 
Database management systems allow for data independence, which is the idea that applications/users do not need to worry about how the database is **structured** (logical independence) and how the database is **stored** (physical independence)
<!--ID: 1710009352080-->





What is a [Data Models|data model](data_models|data_model)? Give an example of one. #flashcard 
Data models are mathematical concepts that can be used to describe how data is represented, organised and structured.
The relational model is an example of data model.
<!--ID: 1710009352088-->

What are [relation schemas](relation_schemas), their instances and [relational database schemas](relational_database_schemas)? #flashcard 
Relation schemas consist of the name of the relation, and an ordered sequence of $k$ attributes.
Their instances are actual tables 
  
<!--ID: 1710009352094-->


How can a [Relation Schemas|relation schema](relation_schemas|relation_schema) achieve [First Normal Form](first_normal_form)? #flashcard 
It can do so when each attribute is **atomic** and cannot be broken down further, and each **tuple** (each record in a table) has a [Primary Keys|primary key](primary_keys|primary_key).
<!--ID: 1711097892529-->


How can a [Relation Schemas|relation schema](relation_schemas|relation_schema) achieve [Second Normal Form](second_normal_form)? Explain what it is that stops a schema from being in Second Normal Form. #flashcard 
Second Normal Form is achieved when all partial dependencies are removed from a relation schema.
Partial Dependencies arise when attributes are [Functional Dependencies|functionally dependent](functional_dependencies|functionally_dependent) on only a part of a [Superkeys and Candidate Keys|candidate key](superkeys_and_candidate_keys|candidate_key), and not all of it.
<!--ID: 1711098301726-->

How can a [Relation Schemas|relation schema](relation_schemas|relation_schema) achieve [Third Normal Form](third_normal_form)? Explain what it is that stops a schema from being in Third Normal Form. #flashcard 
Third Normal Form can be achieved can be achieved by eliminating all transitive dependencies from a relation schema.
Transitive dependencies arise when one **non-prime** (not part of a candidate key) attribute is [Functional Dependencies|functionally dependent](functional_dependencies|functionally_dependent) on another non-prime attribute.
<!--ID: 1711098453211-->

How can a [Relation Schemas|relation schema](relation_schemas|relation_schema) achieve [Boyce-Codd Normal Form](boyce-codd_normal_form)? Explain what it is that stops a schema from being in Boyce-Codd Normal Form. #flashcard 
Boyce-Codd Normal Form is achieved when all trivial dependencies are removed from a table.
Trivial dependencies arise when one attribute is [Functional Dependencies|functionally dependent](functional_dependencies|functionally_dependent) on an attribute that *may be a key*, but is not a [Superkeys and Candidate Keys|candidate key](superkeys_and_candidate_keys|candidate_key).
Every determinant in a table that is in Boyce-Codd Normal form **must** be able to determine all other non-key attributes (i.e. it is a candidate key.)
<!--ID: 1711099026742-->


What are three [benefits of database normalisation](benefits_of_database_normalisation)? #flashcard 
Normalisation eliminates redundant data, which frees up space to store more data.
Normalisation makes querying faster than querying a 0NF database
Normalisation increases **data integrity** and reduces the likelihood of inconsistencies and anomalies.
<!--ID: 1711099158469-->

What are three [disadvantages of database normalisation](disadvantages_of_database_normalisation)? #flashcard 
Database normalisation involves a lot of table decomposition and results in much more tables.
More tables results in a more overall complex design.
More relationships are needed in the design since there are more tables too.
<!--ID: 1711099283851-->

What is [database denormalisation](database_denormalisation), and what are its benefits and downsides? #flashcard 
Database denormalisation is the process of partially reversing the process of normalisation.
Database denormalisation gives a slight boost in querying speed.
However, queries on a denormalised database need to be more careful and complex to maintain data integrity.
<!--ID: 1711099425204-->

What is a [Conceptual Models|conceptual model](conceptual_models|conceptual_model), a [Logical Modelling|logical model](logical_modelling|logical_model) and a [Physical Modelling|physical model](physical_modelling|physical_model) for databases? #flashcard 
Conceptual models solely focus on identifying key ideas and concepts, such as what the [Entities|entities](entities|entities) are and the [Entity Relationships|relationships](entity_relationships|relationships) between them.
Logical models refine conceptual models by identifying [Entity Attributes|attributes](entity_attributes|attributes) and their [Database Attribute Types|types](database_attribute_types|types).
Physical models implement logical models by defining a database structure, adding in actual tables with fields, and implementing relationships.
<!--ID: 1711100415116-->

What are 7 more [Benefits of SQL and Relational Databases|benefits of using SQL and relational databases](benefits_of_sql_and_relational_databases|benefits_of_using_sql_and_relational_databases)? #flashcard 
It is a simple model with simple querying
It's easy to access data from a SQL relational database
The [relational data model](relational_data_model) keep the database consistent
The organised, [Database Normalisation|normalised](database_normalisation|normalised) structure maintains data integrity and reduces duplicate data
Multiple users can access the database concurrently
SQL relational databases support access rights which boost security.
<!--ID: 1713866484213-->

What are 5 [limitations of SQL and relational databases](limitations_of_sql_and_relational_databases)? #flashcard 
**Data integrity** - maintaining data integrity adds overhead and becomes a bottleneck on performance in some systems and for some data
**Performance** - the performance of a relational database degrades with the size of the database
**Schema Design** - Schema designs are rigid and cannot easily accommodate a rapidly-evolving dataset
**Scalability** - It's not that easily to scale a relational database to make use of multiple servers
**Storage Efficiency** - The way relational databases store BLOBs is inefficient and quickly takes up space.
<!--ID: 1713866770417-->

Where would SQL and relational databases [The Benefits of SQL and Relational Databases, In Spite of the Limitations|still be used](the_benefits_of_sql_and_relational_databases,_in_spite_of_the_limitations|still_be_used) for a rapidly-evolving dataset and why? #flashcard 
Relational databases can still be used for metadata since it is much more suitable to SQL's rigid structure.
<!--ID: 1713866874110-->


How are [NoSQL](nosql) databases different to [SQL](sql) databases, and how are they similar? #flashcard 
NoSQL is different in the way that it doesn't follow [ACID](acid) or use [Relational Database Schemas|schemas](relational_database_schemas|schemas) as rigidly as SQL would.
However, SQL and SQL-like languages can be used with NoSQL.
<!--ID: 1713867038097-->

What are some examples of human-readable and machine-readable data? What are the differences? #flashcard 
Human-readable data includes text and images, is mostly unstructured and focuses on usability/readability
Machine-readable data includes markup and needs to be structured, is mostly not human-readable and focuses on efficiency.
<!--ID: 1716728594670-->


What are the three key aspects of SGML? #flashcard 
Tags - `<tag></tag>`
Attributes - `<tag attribute="value">`
Content - `<tag>text goes here</tag>`
<!--ID: 1716728594690-->


What are the 5 rules of [XML](xml) syntax? #flashcard 
All XML elements must have a closing tag
XML tags are case-sensitive
XML tags must be properly nested
XML documents must have a root element (everything else should be nested inside the root)
XML attribute values must be in quotes
<!--ID: 1716728594698-->


What is the difference between cardinality and modality in database modelling? #flashcard 
Cardinality is the maximum number of times an instance in one entity can be associated with instances in another entity.
Modality is the minimum number of times an instance in one entity can be associated with instances from another entity.
<!--ID: 1716728594707-->


What are [Surrogate Keys](surrogate_keys), why are they used and why are they COOKED? #excalidraw 
Surrogate keys are auto-generated, system-maintained keys used to uniquely identify an entity where actual data is insufficient to use in a database system as a key. However, they have **no business meaning.**

How does markup work, and what does it allow us to do? #flashcard 
Markup adds additional information on top of our content, without changing the content itself.
It allows us to style/structure content (smells like CSS-talk) and also allows machines to parse the content.
<!--ID: 1716728594715-->


What is the difference between intension and extension? #flashcard 
Intension refers to semantics and definition ([Relation Schemas](relation_schemas))
Extension refers to examples, physical instances of schemas.
<!--ID: 1716728594724-->


What are the 5 properties of relations? #flashcard 
Each row in a relation represents a $k$-tuple of a relation
The ordering of rows does not matter (relations are [sets](sets) after all)
All rows are distinct (RELATIONS ARE SETS OK?)
The ordering of attributes doesn't matter
The significance of each column is based on the name its given
<!--ID: 1716728594743-->


In database systems, what is a key? #flashcard 
A key is a unique identifier that is distinct between any 2 tuples in a relation.
<!--ID: 1716728594751-->


In database systems, what are functional dependencies? #flashcard 
Functional dependencies are expressions in the form $A \rightarrow B$ where any two tuples have the same value for the set of attributes $A$, they also have the same values for the set of attributes $B$.
<!--ID: 1716728594761-->


How does [the closure algorithm](the_closure_algorithm) work? #flashcard 
You start with a set of attributes, and a set of [functional dependencies](functional_dependencies), and try to see how big the set of attributes can get by adding attributes using the set of functional dependencies.
<!--ID: 1716728594769-->


What are [Superkeys and Candidate Keys](superkeys_and_candidate_keys)? #flashcard 
Superkeys are keys that can determine all attributes in a relation.
Candidate keys are superkeys that cannot be reduced any further: remove an attribute from a candidate key, and it is no longer a superkey.
<!--ID: 1716728594779-->

What are the [BASE](base) Properties? #flashcard 
**Basic Availability**: The system will always remain operational, even if failures occur.
**Soft State**: The system should allow for eventual consistency, and does not always need to be consistent.
**Eventual Consistency**: This is the idea that the system can be temporarily inconsistent, but will always converge to a consistent state.
<!--ID: 1716729889035-->

What are [key-value databases](key-value_databases)? What are 2 advantages? What are some examples? When would you use one? #flashcard 
Key-value databases store data in key-value pairs.
They have a simple data model, which allows for flexibility in storing data.
They are basically [Hash Tables](hash_tables), which means insane access speed.
Examples include: Redis, AmazonDynoDB
They are good for data that need to be accessed quickly.
<!--ID: 1716731995784-->


What are [Document Databases](document_databases)? What are 4 advantages? What are some examples? When would you use one? #flashcard 
Document databases store data in documents, which are self-contained data units key-value pairs, and other documents.
They do not need a fixed schema and so are flexible
They typically store data in JSON/BSON which is highly compatible with modern programming languages and human-readable
Data is stored hierarchically, allowing for complex relationships to be stored in a single document
Querying is also flexible
Examples: MongoDB, CouchDB
They are good for applications that require evolving data models
<!--ID: 1716731995810-->



What are [Column-family Databases](column-family_databases)? What are 3 advantages? What are some examples? When would you use one? #flashcard 
Column-family databases store data in groups known as columns, with a key being used to get each column.
They are very flexible and allow for dynamic addition of columns
They allow for each row to have its own set of columns with their own data type
They can be queried efficiently using specific column-family querying languages.
Examples: Cassandra, Apache HBase
They are good for applications that require scalability, performance, and flexibility
<!--ID: 1716731995815-->


What are [Graph Databases](graph_databases)? What are 4 advantages? What are some examples? When would you use one? #flashcard 
Graph Databases store data using [Graphs](graphs), where entities are nodes and relationships are edges.
Their flexible schema allows for nodes to be added dynamically
Nodes and edges can have properties associated with them, allowing for metadata.
Edges aren't represented using [Adjacency Lists](adjacency_lists) or [Adjacency Matrices](adjacency_matrices), removing the cost of index lookups
Querying is efficient in graph databases.
Examples: Neo4j, Amazon Neptune
They are good for deeply interconnected data.
<!--ID: 1716731995824-->


What is a key advantages of all [NoSQL Databases](nosql_databases)? #flashcard 
They all scale horizontally well
<!--ID: 1716731995831-->

Oh no! You want to create a schema for a pizza, based on a pizza XML with 2 toppings. Most things have been defined correctly, but the parser keeps saying `Element 'Topping': This element is not expected.` when passed the original pizza XML. What could be wrong with your schema? #flashcard 
`maxOccurs` needs to be defined if allowing for more than one occurrence of an element.
<!--ID: 1716826610906-->


What are two important things to consider when choosing a primary key for your table? #flashcard 
**Consider the domain:** Don't just choose keys based off of the current [functional dependencies](functional_dependencies), but also based off of how those can change in the future
**Ideally avoid text:** Text is more subject to change than cold, hard, numbers, but we want our key to be as resilient to change as possible.
<!--ID: 1716826610925-->


#### Why:
#### How:








