package app

import "gonote/internal/routes"

type Module interface {
	Routes() routes.Routes
}

func GetRoutesFromModules(modules []Module) []routes.Routes {
	routesList := make([]routes.Routes, len(modules))
	for i, module := range modules {
		routesList[i] = module.Routes()
	}
	return routesList
}
