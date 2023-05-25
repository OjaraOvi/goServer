package main

import (
	"fmt"
	"net/http"
)

func CheckAuth() Middleware {
	return func (f http.HandlerFunc) http.HandlerFunc{
		return func (w http.ResponseWriter, r *http.Request){
			flag := false
			fmt.Println("checking authentication")
			if flag {
				f(w,r)
			}else {
				return
			}
		}
	}
}