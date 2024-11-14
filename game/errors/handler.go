package errors

import (
	"strings"

	"github.com/nika-gromova/o-architecture-patterns/game/base"
	"github.com/nika-gromova/o-architecture-patterns/game/commands"
)

type HandleFunc func(base.Command, error) base.Command

type Handler struct {
	storage        map[string]map[string]HandleFunc
	defaultHandler HandleFunc
}

func NewHandler() *Handler {
	return &Handler{
		storage: make(map[string]map[string]HandleFunc),
		defaultHandler: func(cmd base.Command, err error) base.Command {
			return &commands.LogCommand{
				Err: err,
				Cmd: cmd,
			}
		},
	}
}

func (h *Handler) Register(cmdType, errType string, handler HandleFunc) {
	if _, ok := h.storage[cmdType]; !ok {
		h.storage[cmdType] = make(map[string]HandleFunc)
	}
	h.storage[cmdType][errType] = handler
}

func (h *Handler) Handle(cmd base.Command, err error) base.Command {
	cmdType := base.GetVarType(cmd)
	cmdHandlers, ok := h.storage[cmdType]
	if ok {
		for key := range cmdHandlers {
			if strings.Contains(err.Error(), key) {
				return h.storage[cmdType][key](cmd, err)
			}
		}
	}
	return h.defaultHandler(cmd, err)
}
