package src

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/aayushjoshi2709/mypic/src/comment"
	"github.com/aayushjoshi2709/mypic/src/group"
	"github.com/aayushjoshi2709/mypic/src/image"
	"github.com/aayushjoshi2709/mypic/src/presign"
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

func SetUpRepositories(ctx context.Context) {
	userRepository := user.Repository{}
	userRepository.Init()
	SetRepository("userRepository", &userRepository)

	imageRepository := image.Repository{}
	imageRepository.Init()
	SetRepository("imageRepository", &imageRepository)

	commentRepository := comment.Repository{}
	commentRepository.Init()
	SetRepository("commentRepository", &commentRepository)

	presignRepository := presign.Repository{}
	presignRepository.Init(ctx)
	SetRepository("presignRepository", &presignRepository)

	groupRepository := group.Repository{}
	groupRepository.Init()
	SetRepository("groupRepository", &groupRepository)

	PrintRepositories()
}
