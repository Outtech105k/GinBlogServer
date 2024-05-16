package main

import (
	"fmt"
	"strings"

	"github.com/gin-contrib/multitemplate"
)

func createPageRender() (multitemplate.Renderer, error) {
	directries := []string{"templates", "content", "contentBody"}
	r := multitemplate.NewRenderer()

	templateNames := map[string][]string{
		"index":           {"base", "topPage"},
		"blogArticleList": {"base", "customContent", "blogArticleList"},
		"404":             {"base", "customContent", "404"},
	}

	for pageName, files := range templateNames {
		filePaths, err := genTemplateFilePaths(directries, files)
		if err != nil {
			return nil, err
		}
		r.AddFromFiles(pageName, filePaths...)
	}

	return r, nil
}

func genTemplateFilePaths(directries, fileNames []string) ([]string, error) {
	if len(directries) < len(fileNames) {
		return nil, fmt.Errorf("argument number is `directries` < `fileNames`")
	}
	paths := make([]string, 0, len(fileNames))

	for i := 0; i < len(fileNames); i++ {
		paths = append(paths, fmt.Sprintf(
			"%s/%s.html",
			strings.Join(directries[:i+1], "/"),
			fileNames[i],
		))
	}

	return paths, nil
}
