package main

import (
	"backend/cake/config"
	"backend/cake/model"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)
	


func authenticator(role int8) func(next http.Handler) http.Handler{
	return func(next http.Handler) http.Handler{
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
			cookie,err := r.Cookie("jwt")		
			if err != nil{
				http.Error(w,"no authentication cookie attached to the request",400)
				return
			}
			jwtString := cookie.Value	
			err = model.Decode(jwtString,role)
			if err != nil{
				http.Error(w,err.Error(),401)
				return
			}
			config.Sugar.Info("Authenticated")
			next.ServeHTTP(w,r)
	})
	}
}

func main(){
	config.Sugar.Info("yooo app started")	
	mainRouter := chi.NewRouter()
	mainRouter.Use(middleware.Logger)
	//public
	mainRouter.Group(func(r chi.Router) {
		r.Post("/login",model.Login)
		r.Post("/register",model.Register)
	})

	//private
	mainRouter.Group(func(r chi.Router) {
		r.Use(authenticator(1))
		r.Put("/changePassword",model.ChangePassword)
	})

	http.ListenAndServe(":3000",mainRouter)	
}
