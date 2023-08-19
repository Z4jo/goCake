package model

import (
	"backend/cake/config"
	"backend/cake/entity"
	"errors"
	"time"
	"github.com/go-chi/jwtauth/v5"
)

var tokenAuth *jwtauth.JWTAuth

func init(){
	tokenAuth = jwtauth.New("HS256",[]byte("sokerez"),nil)
}

func checkExpiration(incomingTime string)bool{
	parsedTime, err := time.Parse(time.RFC3339, incomingTime)	
	if err != nil{
		config.Sugar.Error("Unable to parse time in jwt token")	
		return false
	}
	now := time.Now()
	duration := now.Sub(parsedTime)
	if duration.Minutes() >= 0{
		config.Sugar.Info("token Expired")	
		return false 
	}
	return true
}

func Encode(user *entity.User)(string,error) {
	expirationTime := time.Now().Add(time.Minute * 15)
	_,token,err := tokenAuth.Encode(map[string]interface{}{"userId":user.ID,"expiration":expirationTime,"role":user.Role})
	if err != nil{
		return  "",err
	}
	return  token,nil
}

func Decode(incomingToken string, role int8) error {
	token,err := tokenAuth.Decode(incomingToken)		
	if err != nil{
		return err
	}
	expiration,valueExists := token.Get("expiration")
	if valueExists == false{
		config.Sugar.Debug("bad jwt structure")
		return errors.New("token was not issued by this server")
	}

	incomingRole,valueExists:= token.Get("role")
	if valueExists == false{
		config.Sugar.Debug("bad jwt structure")
		return errors.New("token was not issued by this server")
	}
	if int8(incomingRole.(float64)) != role{
		return  errors.New("not authorized")
	}
	if	!checkExpiration(expiration.(string)){
		return  errors.New("jwt expired")
		
	}
	return nil
}

