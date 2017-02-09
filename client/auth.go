package client

import "net/http"
import "errors"

//Auth is a decorator to check if the request is Authorized
func Auth(token string) Decorator {
	return func(c Client) Client {
		return ClientFunc(func(r *http.Request) (res *http.Response, err error) {
			if token != "verysecuretoken" {
				return res, errors.New("Not Authorized")
			}
			return c.Do(r)
		})
	}
}
