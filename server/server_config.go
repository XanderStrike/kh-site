package server

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func BuildServer() *gin.Engine {
	assetProvider := &AssetProvider{getAssetsDirectory()}

	pageSpecs := []PageSpec{
		PageSpec{AssetName: "home", Title: "", Route: ""},
	}

	pageFactory := NewPageFactory(assetProvider, pageSpecs)

	assetNames := assetProvider.ListAllNonHTML()

	serverConfig := ServerConfig{
		AssetNames:    assetNames,
		PageFactory:   pageFactory,
		AssetProvider: assetProvider,
	}

	router := serverConfig.BuildRouter()
	return router
}

func getAssetsDirectory() string {
	assetsDirectory := os.Getenv("ASSETS_DIR")
	if assetsDirectory == "" {
		assetsDirectory = "assets"
	}
	path, err := filepath.Abs(assetsDirectory)
	if err != nil {
		panic("could not expand assets directory path: " + err.Error())
	}

	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			panic("Didn't find an assets directory at " + path)
		} else {
			panic("Could not stat assets directory: " + path)
		}
	}

	fmt.Printf("\n Looking for assets in: %s\n", path)
	return path

}
