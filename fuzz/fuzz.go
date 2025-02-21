package fuzz

import (
	"strings"

	"github.com/evanw/esbuild/pkg/api"
)

func FuzzTransformJS(data []byte) int {
	result := api.Transform(string(data), api.TransformOptions{
		Loader: api.LoaderJS,
	})
	for _, err := range result.Errors {
		if strings.Contains(err.Text, "panic:") {
			panic(err.Text)
		}
	}
	return 0
}
