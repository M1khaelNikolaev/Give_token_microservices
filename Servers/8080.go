package mainimport (
    "fmt"
    "net/http"
    "github.com/dgrijalva/jwt-go"
    "time"
)

var mySigningKey = []byte("supersecret")

func generateToken(w http.ResponseWriter, r *http.Request) {
    token := jwt.New(jwt.SigningMethodHS256)
    
    claims := token.Claims.(jwt.MapClaims)
    claims["authorized"] = true
    claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

    tokenString, err := token.SignedString(mySigningKey)
    if err != nil {
        fmt.Fprintln(w, err.Error())
    }

    fmt.Fprintln(w, tokenString)
}

func main() {
    http.HandleFunc("/generateToken", generateToken)
    http.ListenAndServe(":8080", nil)
}
