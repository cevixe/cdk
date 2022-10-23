package handler

import (
	"github.com/cevixe/cdk/module"
	"github.com/cevixe/cdk/module/function"
)

/*
Propiedades de configuración de una función handler.
*/
type HandlerProps struct {
	/*
		Tipo de la función handler (HandlerType). Si no se especifica un valor
		se utilizará: HandlerType_Basic

		Ejemplos:
			-> HandlerType_Basic
			-> HandlerType_Standard
			-> HandlerType_Advanced
	*/
	Type HandlerType `field:"required" json:"type"`

	/*
		Tipos de eventos que deben ser escuchados y procesados por la función handler.
		Si se envía un array vacío, todos los eventos serán escuchados y procesados.
		Si no se envía un array, ningún evento será escuchado.

		Ejemplos:
			-> nil
			-> &[]string{}
			-> &[]string{"AccountCreated", "AccountUpdated", "AccountDeleted"}

	*/
	Events *[]string `field:"optional" json:"events"`

	/*
		Tipos de comandos que deben ser escuchados y procesados por la función handler.
		Si se envía un array vacío, todos los comandos serán escuchados y procesados.
		Si no se envía un array, ningún comando será escuchado.
		Ejemplos:
			-> nil
			-> &[]string{}
			-> &[]string{"CreateAccount", "UpdateAccount", "DeleteAccount"}

	*/
	Commands *[]string `field:"optional" json:"commands"`
}

func NewHandler(mod module.Module, alias string, props *HandlerProps) Handler {

	fn := function.NewFunction(mod, alias)

	return &handlerImpl{
		Function: fn,
		typ:      props.Type,
		events:   props.Events,
		commands: props.Commands,
	}
}
