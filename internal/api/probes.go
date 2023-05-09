package api

import "net/http"

func (env *Env) Alive(_ http.ResponseWriter, _ *http.Request) {}

func (env *Env) Ready(_ http.ResponseWriter, _ *http.Request) {}
