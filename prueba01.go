package main
import (
  "net/http"
)

func main(){
  //println("hola peru")
  //routs
  http.HandleFunc("/",homeHandler)
  //inicia tu servior
  http.ListenAndServe(":3000",nil)
}
func homeHandler(w http.ResponseWriter,r *http.Request){
  w.Write([]byte("hola prueba de Go peru-hola compañeros de A&S"))
}
