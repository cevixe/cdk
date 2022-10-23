package handler

import "github.com/cevixe/cdk/module/function"

type handlerImpl struct {
	function.Function
	typ      HandlerType
	events   *[]string
	commands *[]string
}

func (h *handlerImpl) Type() HandlerType {
	return h.typ
}

func (h *handlerImpl) Events() *[]string {
	return h.events
}

func (h *handlerImpl) Commands() *[]string {
	return h.commands
}
