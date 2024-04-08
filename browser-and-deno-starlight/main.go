package main

import (
	"io"
	"net/http"

	"github.com/starlight-go/starlight"
)

func main() {
	httpGetFunc := func(url string) (string, error) {
		resp, err := http.Get(url)
		if err != nil {
			return "", err
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		return (string(body)), nil
	}

	globals := map[string]interface{}{
		"httpget": httpGetFunc,
	}

	code := []byte(`
# custom function
print(httpget('https://httpbin.org/anything'))
`)

	_, err := starlight.Eval(code, globals, nil)
	if err != nil {
		panic(err)
	}
}
