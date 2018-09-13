package models

import (
	jwt "github.com/dgrijalva/jwt-go"
) //tocken

type Claim struct {
	User `json:"user"`
	jwt.StandardClaims
}
