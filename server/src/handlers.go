package src

import (
	"fmt"
	"log/slog"

	"github.com/aayushjoshi2709/mypic/src/comment"
	"github.com/aayushjoshi2709/mypic/src/group"
	"github.com/aayushjoshi2709/mypic/src/image"
	"github.com/aayushjoshi2709/mypic/src/presign"
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
	userHandler.New(map[string]any{
		"userRepository": GetRepository[user.Repository]("userRepository"),
	})
	SetHandler("userHandler", &userHandler)

	imageHandler := image.Handler{}
	imageHandler.New(map[string]any{
		"imageRepository": GetRepository[image.Repository]("imageRepository"),
	})
	SetHandler("imageHandler", &imageHandler)

	commentHandler := comment.Handler{}
	commentHandler.New(map[string]any{
		"commentRepository": GetRepository[comment.Repository]("commentRepository"),
	})
	SetHandler("commentHandler", &commentHandler)

	presignHandler := presign.Handler{}
	presignHandler.New(map[string]any{
		"presignRepository": GetRepository[presign.Repository]("presignRepository"),
	})
	SetHandler("presignHandler", &presignHandler)

	groupHandler := group.Handler{}
	groupHandler.New(map[string]any{
		"groupRepository": GetRepository[group.Repository]("groupRepository"),
		"imageRepository": GetRepository[image.Repository]("imageRepository"),
		"userRepository": GetRepository[user.Repository]("userRepository"),
	})
	SetHandler("groupHandler", &groupHandler)

	PrintHandlers()
}
