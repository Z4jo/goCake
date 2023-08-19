package model

import (
	"backend/cake/config"
	"backend/cake/entity"
	"errors"
	"net/http"

	"github.com/go-chi/render"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)


func (payload *RegisterPayload) Bind(r *http.Request) error {
	if payload.Email == "" || payload.Password == "" || payload.Username == ""{
		config.Sugar.Debug("User fuckedUP")
		return errors.New("missing one of the fields")
	}
	return nil
}
func (payload *LoginPayload) Bind(r *http.Request) error {
	if (payload.Email == "" && payload.Username == "") || payload.Password == "" {
		config.Sugar.Debug("User fuckedUP")
		return errors.New("missing one of the fields")
	}
	return nil
}
func (payload *RegisterPayload) Render (w http.ResponseWriter, r *http.Request) error{
	payload.Password = ""
	return nil
}

func (payload *LoginPayload) Render (w http.ResponseWriter, r *http.Request) error{
	payload.Password = ""
	return nil
}

func encryptPassword(payload *RegisterPayload) error {
	encryptedPassword,err := bcrypt.GenerateFromPassword([]byte(payload.Password),10)	
	if err != nil{
		config.Sugar.Error("encryption didnt work mate")
		return err
	}
	payload.Password = string(encryptedPassword)
	return nil
}

func Login(w http.ResponseWriter, r *http.Request){
	data := &LoginPayload{}
	if err := render.Bind(r,data); err!=nil{
		render.Render(w,r,ErrorInvalidRequest(err))	
		return
	}
	var user *entity.User
	var result *gorm.DB
	if data.Username != ""{
		result = config.DB.Where("username = ?",data.Username).Find(&user)
	}else{
		result = config.DB.Where("email = ?",data.Email).Find(&user)
	}
	
	if result.RowsAffected == 0{
		config.Sugar.Debug(result.RowsAffected)
		render.Render(w,r,ErrorInvalidRequest(errors.New("bad username or email")))	
		return
	}
	config.Sugar.Debug(user)
	err := bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(data.Password))
	if err != nil{
		render.Render(w,r,ErrorInvalidRequest(errors.New("bad password")))	
		return
	}
	jwtToken,err := Encode(user)	
	if err != nil{
		config.Sugar.Error("jwt creation failed: ",err)
		render.Render(w,r,ErrorInternalServer(err))
		return
	}
	data.Id = user.ID
	http.SetCookie(w,&http.Cookie{Name:"jwt",Value:jwtToken})
	render.Status(r,http.StatusOK)
	render.Render(w,r,data)
}

func Register(w http.ResponseWriter, r *http.Request){
	data := &RegisterPayload{}
	if err := render.Bind(r,data); err != nil {
		render.Render(w,r,ErrorInvalidRequest(err))			
		return
	}	
	if err := encryptPassword(data); err != nil{
		render.Render(w,r,ErrorInternalServer(err))	
		return
	}

	user := &entity.User{Username:data.Username, Password: data.Password, Email: data.Email, Role: 1}
	config.DB.Create(user)
	config.Sugar.Info("User registered")
	jwtToken,err := Encode(user)

	if err != nil{
		config.Sugar.Error("jwt creation failed")
		render.Render(w,r,ErrorInternalServer(err))	
		return
	}

	data.Id = user.ID
	http.SetCookie(w,&http.Cookie{Name:"jwt",Value:jwtToken})
	render.Status(r,http.StatusCreated)
	render.Render(w,r,data)
}
func ChangePassword(w http.ResponseWriter, r *http.Request){


}
