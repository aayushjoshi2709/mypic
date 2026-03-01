package src

import (
	"fmt"
	"log/slog"

	"github.com/aayushjoshi2709/mypic/src/comment"
	"github.com/aayushjoshi2709/mypic/src/image"
	"github.com/aayushjoshi2709/mypic/src/user"
)

var repositories = make(map[string]any)

func GetRepository[T any](name string) *T {
	return repositories[name].(*T)
}

func SetRepository[T any](name string, repository *T) {
	repositories[name] = repository
}

func PrintRepositories() {
	slog.Info("Registered repositories:")
	for name := range repositories {
		slog.Info(name, "type ", fmt.Sprintf("%T", repositories[name]))
	}
}

func SetUpRepositories() {
	userRepository := user.Repository{}
	userRepository.Init()
	SetRepository("user_repository", &userRepository)

	imageRepository := image.Repository{}
	imageRepository.Init()
	SetRepository("image_repository", &imageRepository)

	commentRepository := comment.Repository{}
	commentRepository.Init()
	SetRepository("comment_repository", &commentRepository)

	PrintRepositories()
}
