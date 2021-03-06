package urlshort

import (
	"encoding/json"
	"net/http"

	"gopkg.in/yaml.v2"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path

		if path, ok := pathsToUrls[path]; ok {
			http.Redirect(w, r, path, http.StatusFound)
		}

		fallback.ServeHTTP(w, r)
	}
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	parsedYAML, err := parseYAML(yml)
	if err != nil {
		return nil, err
	}

	pathMap := buildMap(parsedYAML)
	return MapHandler(pathMap, fallback), nil
}

// JSONHandler will parse the provided JSON and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the JSON, then the
// fallback http.Handler will be called instead.
//
// JSON is expected to be in the format:
//	[
// 		{"path": "/json", "url": "https://gobyexample.com/json"},
//		{"path": "/hello", "url": "https://gobyexample.com/byebye"}
//	]
func JSONHandler(json []byte, fallback http.Handler) (http.HandlerFunc, error) {
	parsedJSON, err := parseJSON(json)
	if err != nil {
		return nil, err
	}

	pathMap := buildMap(parsedJSON)

	return MapHandler(pathMap, fallback), nil
}

func parseJSON(jsonData []byte) ([]pathURL, error) {
	var slice []pathURL

	err := json.Unmarshal(jsonData, &slice)
	if err != nil {
		return nil, err
	}

	return slice, nil
}

func parseYAML(yml []byte) ([]pathURL, error) {
	var slice []pathURL

	err := yaml.Unmarshal(yml, &slice)
	if err != nil {
		return nil, err
	}

	return slice, nil
}

func buildMap(parsedYml []pathURL) map[string]string {
	pathMap := make(map[string]string)

	for _, path := range parsedYml {
		pathMap[path.Path] = path.URL
	}

	return pathMap
}

type pathURL struct {
	Path string `yaml:"path" json:"path"`
	URL  string `yaml:"url" json:"url"`
}
