package main

import (
	"io"
	"net/http"

	"go.starlark.net/starlark"
)

func main() {
	const code = `
# custom function
print(httpget('https://httpbin.org/anything'))
`

	httpGetFunc := func(thread *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
		var url string
		err := starlark.UnpackPositionalArgs(fn.Name(), args, kwargs, 1, &url)
		if err != nil {
			panic(err)
		}

		resp, err := http.Get(url)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		return starlark.String(string(body)), err
	}
	predeclared := starlark.StringDict{
		"httpget": starlark.NewBuiltin("httpget", httpGetFunc),
	}

	thread := &starlark.Thread{}
	_, err := starlark.ExecFile(thread, "", code, predeclared)
	if err != nil {
		panic(err)
	}
}
