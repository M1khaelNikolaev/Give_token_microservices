package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	 "github.com/dgrijalva/jwt-go"
)

var MySecretKey = [] byte (os.Getenv("My_Secret_Key"))

func homePage (w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w,"Secret")
}

func isAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header ["Token"] !=nil {
jwt.Parse (r.Header["Token"][0], func (token *jwt.Token)(interface{},error){
	if_, ok:= token.Method.(*jwt.SigningMethodHMAC); !=ok{
		return nil, fmt.Errorf("Invalid signin method")
	}
aud:= "billing.jwtgo.io"
	checkAudience:=token.Claims.(jwt.MapClaims).VerifyAudience(aud,false)

	if !checkAudience{
		return nil, fmt.Errorf("Invalid aud")
	}
	iss:= "jwtgo.io"
	checkIss:= token.Claims.(jwt.MapClaims).VerifyIssuer(iss,false)
	if !checkIss{
		return nil, fmt.Errorf("Invalid iss")
	}
	return MySigningKey, nil
})
if err !=nil {
	fmt.Printf(w,err.Error())
}
if token.Valid{
	endpoint(w,r)
}

}
		else{
	fmt.Printf(w,"No authorization token provided")
}
	})
}



func handleReqests(){
	http.Handle("/", isAuthorized(homePage))
	log.Fatal(http.ListenAndServe(":9001",nil))
}

func main (){

	fmt.Printf("server")
	handleReqests()
}
