package models

// Definición del modelo de usuario
type Usuario struct {
    Id		int		`form:"-"`
    Nombre      string	`form:"nombre"`
    App		string	`form:"app"`
    Apm		string	`form:"apm"`
    Correo	string	`form:"correo"`
    Edad		int		`form:"edad"`
}

