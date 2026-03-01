package src

import (
	"fmt"
	"log/slog"

	"github.com/aayushjoshi2709/mypic/src/comment"
	"github.com/aayushjoshi2709/mypic/src/image"
	"github.com/aayushjoshi2709/mypic/src/user"
)

var handlers = make(map[string]any)

func GetHandler[T any](name string) *T {
	return handlers[name].(*T)
}

func SetHandler[T any](name string, handler T) {
	handlers[name] = handler
}

func PrintHandlers() {
	slog.Info("Registered handlers:")
	for name := range handlers {
		slog.Info(name, "type ", fmt.Sprintf("%T", handlers[name]))
	}
}

func SetUpHandlers() {
	userHandler := user.Handler{}
	userHandler.New(GetRepository[user.Repository]("user_repository"))
	SetHandler("user_handler", &userHandler)

	imageHandler := image.Handler{}
	imageHandler.New(GetRepository[image.Repository]("image_repository"))
	SetHandler("image_handler", &imageHandler)

	commentHandler := comment.Handler{}
	commentHandler.New(GetRepository[comment.Repository]("comment_repository"))
	SetHandler("comment_handler", &commentHandler)

	PrintHandlers()
}
