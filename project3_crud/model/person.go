package model

//Community estructura de una comunidad
type Community struct {
	//Name: nombre de una comunidad, ejemplo: Globant
	Name string
}

//Communities slice de comunidades
type Communities []Community

//Person estructura de una persona
type Person struct {
	//Name nombre de la persona, ej: Andy
	Name string
	//Age edad de la persona, ej:23
	Age uint8
	//Communities comunidades a la que pertenece una persona
	Communities Communities
}

//Persons slice de personas
type Persons []Person
