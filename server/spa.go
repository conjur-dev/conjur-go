package server

import (
	"conjur/logger"
	"net/http"
	"os"
	"path/filepath"
)

type SpaHandler struct {
	StaticPath string
	IndexPath  string
}

func (spaHandler SpaHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	workingDir, _ := os.Getwd()
	buildDir := filepath.Join(workingDir, "client/build")
	requestPath := filepath.Join(buildDir, request.URL.Path)
	info, err := os.Stat(requestPath)
	if os.IsNotExist(err) || info.IsDir() {
		http.ServeFile(writer, request, filepath.Join(buildDir, "index.html"))
		return
	}
	if err != nil {
		logger.Error(err.Error())
		http.Error(writer, "", http.StatusInternalServerError)
		return
	}
	http.FileServer(http.Dir(buildDir)).ServeHTTP(writer, request)
}
