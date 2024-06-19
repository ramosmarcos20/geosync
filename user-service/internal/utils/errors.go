package utils

import "errors"

var (
	UserNotFound 	= errors.New("Usuario no encontrado")
	InvalidInput 	= errors.New("Input Invalido")
	ErrorDataBase	= errors.New("Error de Base de datos")
	EmailInUse		= errors.New("Este Email ya esta registrado")
	UserNameInUse	= errors.New("UserName ya esta registrado")
)