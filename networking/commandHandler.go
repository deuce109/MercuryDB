package networking

import "fmt"

type CommandHandler struct {
	handlers map[string]func(args ...any) (any, error)
}

func DefaultCommandHandler() *CommandHandler {
	return &CommandHandler{
		handlers: map[string]func(args ...any) (any, error){
			"ping": func(args ...any) (any, error) {
				return "pong", nil
			},
		},
	}
}

func (ch *CommandHandler) RegisterCommand(command string, handler func(args ...any) (any, error)) {
	ch.handlers[command] = handler
}

func (ch *CommandHandler) HandleCommand(command string, args ...any) (any, error) {
	if handler, exists := ch.handlers[command]; exists {
		return handler(args...)
	}
	return nil, fmt.Errorf("command not found: %s", command)
}
