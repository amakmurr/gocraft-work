package webui

import "github.com/gocraft/web"

func (c *context) BasicAuthMiddleware(rw web.ResponseWriter, r *web.Request, next web.NextMiddlewareFunc) {
	if c.basicAuth {
		username, password, ok := r.BasicAuth()
		if !ok || !(username == c.username && password == c.password) {
			rw.Header().Set("WWW-Authenticate", "Basic")
			rw.WriteHeader(401)
			return
		}
	}
	next(rw, r)
}
