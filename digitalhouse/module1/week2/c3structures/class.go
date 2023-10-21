package c3structures

// Usuario Estructura que describe objeto Usuario
type Usuario struct{
	Nombre 		string 	`json:"nombre"`
	Apellido 	string	`json:"apellido"`
	Edad 		int 	`json:"edad"`
	Correo		string 	`json:"correo"`
	Contrasena 	string 	`json:"-"`
}
type RequestPersona struct{
	Nombre 		string 	`json:"nombre"`
	Apellido 	string	`json:"apellido"`
	Edad 		int 	`json:"edad"`
	Correo		string 	`json:"correo"`
}
type ResponsePersona struct{
	Nombre 		string 	`json:"nombre"`
	Apellido 	string	`json:"apellido"`
	Edad 		int 	`json:"edad"`
	Correo		string 	`json:"correo"`
	Contrasena 	string 	`json:"password"`
	Username	string 	`json:"username"`
}

func Class3() {
 
}

//!return fmt.Sprintf("hola %s", p.Nombre)