package main

import (
	"github.com/qri-io/starlib"
	"go.starlark.net/starlark"
)

func main() {
	const code = `
load('http.star', 'http')

res = http.get('https://httpbin.org/anything')
print(res.body())
`

	thread := &starlark.Thread{
		Load: starlib.Loader,
	}
	_, err := starlark.ExecFile(thread, "", code, nil)
	if err != nil {
		panic(err)
	}
}
