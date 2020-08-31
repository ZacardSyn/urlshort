package urlshort

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"net/http"
)


func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		path := r.URL.Path
		if dest, ok := pathsToUrls[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}

func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var pathUrls []pathUrl
	err := yaml.Unmarshal(yml, &pathUrls)
	if err != nil{
		return nil, err
	}

	pathsToUrls := make(map[string]string)
	//for i := 0; i < len(pathUrls); i++{
	//	pathsToUrls[pathUrls[i].path] = pathUrls[i].url
	//}
	for _, pu := range pathUrls {
		pathsToUrls[pu.Path] = pu.Url
	}
	fmt.Println(pathsToUrls)
	return MapHandler(pathsToUrls, fallback), nil
}

type pathUrl struct {
	Path string
	Url string
}