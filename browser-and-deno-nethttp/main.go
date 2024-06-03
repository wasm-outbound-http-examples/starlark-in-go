package main

import (
	nethttp "github.com/pcj/starlark-go-nethttp"
	"go.starlark.net/starlark"
	"go.starlark.net/syntax"
)

func main() {
	const code = `
resp = http.get('https://httpbin.org/anything')
print(resp.body)
`

	predeclared := starlark.StringDict{
		"http": nethttp.NewModule(),
	}

	thread := &starlark.Thread{}
	_, err := starlark.ExecFileOptions(syntax.LegacyFileOptions(), thread, "", code, predeclared)
	if err != nil {
		panic(err)
	}
}
