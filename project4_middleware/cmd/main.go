package main

import "github.com/andymartinezot/api_go_project/project4_middleware/funciones"

func execute(name string, f funciones.MyFunction) {
	f(name)
}

func main() {
	name := "Andy Martinez (NubeScript!)"
	execute(name, funciones.MiddlewareLog(funciones.Saludar))
	execute(name, funciones.MiddlewareLog(funciones.Despedirse))
}
