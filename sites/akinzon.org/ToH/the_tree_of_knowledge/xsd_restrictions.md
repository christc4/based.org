# XSD Restrictions

Up: [XSD](xsd)
Brother(s):
TARGET DECK

**Maximum and minimum occurrences**
	`minOccurs` -  minimum required amount of element 
		- Use a value of 0 to make the element optional
	`maxOccurs` - maximum required amount of element
		- Use a value of `"unbounded"` to make the element possibly infinite


**Defining sub-elements**
	 - `<xs:sequence>` - This denotes that a sequence of elements will follow
	 - `<xs:all>` - Denotes that elements can appear in any order


**Defining types**
	 - `<xs:element ... type="type">` - Gives a type to an element
	 - `<xs:attribute ... type="type">` - Gives a type to an attribute 


**Defining attributes**
	- Attributes come at the end of an element's definition
	- `<xs:attribute name="[name]" type="type">`
		- `required` keyword can denote that the attribute must be given

**Mixed elements**
	- Allows for elements to contain both text and other elements
	- `<xs:complexType ... mixed="true">`




































#### Why:
#### How:









