package app

import (
	"gonote/internal/handler"
	"gonote/internal/repository"
	"gonote/internal/routes"
	"gonote/internal/service"
)

type PostModule struct {
	postRoutes *routes.PostRoutes
}

func NewPostModule(ctx *AppContext) *PostModule {
	postRepo := repository.NewPostRepository(ctx.DB)
	postService := service.NewPostService(postRepo)
	postHandler := handler.NewPostHandler(postService)
	postRoutes := routes.NewPostRoutes(postHandler)
	return &PostModule{
		postRoutes: postRoutes,
	}
}

func (um *PostModule) Routes() routes.Routes {
	return um.postRoutes
}
