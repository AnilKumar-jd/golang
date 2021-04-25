package main

import (
  "fmt"
  "net/http"
  "log"
  "encoding/json"
  )
  type Erticle struct{
  Name string `json:"name"`
  Title string `json:"title"`
  Job string `json:"job"`
  }

type Erticles []Erticle

  func homepage(w http.ResponseWriter, r *http.Request){
  fmt.Fprintf(w,"homepge endpoint.....")
  }

  func erticlehh(w http.ResponseWriter, r *http.Request){
  erticle:= Erticles{
          Erticle{Name:"anil",Title:"JD",Job:"Softwre Developer"},}
          json.NewEncoder(w).Encode(erticle)
  }

  func homehandlerRequest(){
        http.HandleFunc("/",homepage)
        http.HandleFunc("/erticle",erticlehh)
    log.Fatal(http.ListenAndServe(":8081",nil))
  }

  func main(){
  homehandlerRequest()
  }
