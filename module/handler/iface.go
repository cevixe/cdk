package handler

import (
	"github.com/cevixe/cdk/module/function"
)

/*
Tipo de la función handler.

El tipo está basado en los requerimientos sobre el backpressure y sobre los mensajes
que escucha y procesa la función handler.
*/
type HandlerType uint8

const (
	/*
		La función no posee un requerimiento especial sobre el backpressure ni en el orden
		en el que recibe los mensajes, por lo que puede escuchar los mensajes directamente del bus.
	*/
	HandlerType_Basic HandlerType = 0
	/*
		La función no posee un requerimiento especial sobre el orden en el que recibe los mensajes,
		sin embargo requiere una cola de contención previa al procesamiento de los mensajes, ya sea
		por mayor resilencia o por control del throughput.
	*/
	HandlerType_Standard HandlerType = 1
	/*
		La función requiere recibir los mensajes en el orden en el que fueron generados y publicado,
		así mismo requiere una cola de contención previa al procesamiento de los mensajes, ya sea
		por mayor resilencia o por control del throughput.
	*/
	HandlerType_Advanced HandlerType = 2
)

type Handler interface {
	Type() HandlerType
	Events() *[]string
	Commands() *[]string
	function.Function
}
