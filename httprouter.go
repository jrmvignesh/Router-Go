package main
import (
    "fmt"
    "github.com/julienschmidt/httprouter"
    "net/http"
    "encoding/json"
)


type User struct{
Name string `json:"name"`

}

type UserResp struct{
Greeting string `json:"greetings"`
}

func hello(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
    fmt.Fprintf(rw, "Hello, %s!\n", p.ByName("name"))
      }
    
  func post (rw http.ResponseWriter, req *http.Request, _ httprouter.Params) {  
    
   u := User{}
   var v UserResp
    json.NewDecoder(req.Body).Decode(&u)
    v.Greeting = "Hello,"+u.Name+"!"
    uj, _ := json.Marshal(v)
      
    fmt.Fprintf(rw, "%s", uj)
}

   
    
func main() {
    mux := httprouter.New()
    mux.GET("/hello/:name", hello)
    mux.POST("/user",post)
    server := http.Server{
            Addr:        "0.0.0.0:8080",
            Handler: mux,
    }
    server.ListenAndServe()
}