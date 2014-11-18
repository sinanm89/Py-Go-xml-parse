package main

import (
    "fmt"
    "io/ioutil"
    "github.com/moovweb/gokogiri"
    "net/http"
)

// VIEWS
func some_request(w http.ResponseWriter, r *http.Request) {
    uri := "http://ad4.liverail.com/?LR_PUBLISHER_ID=65617&LR_SCHEMA=vast2-vpaid"
    response, err := http.Get(uri)
    if err != nil {
        return
    }

    defer response.Body.Close()

    data, _ := ioutil.ReadAll(response.Body)
    doc, err := gokogiri.ParseXml(data)
    if err != nil {
        return
    }

    defer doc.Free()

    doc.Root().SetAttr("version", "3.0")

    ccc, err := doc.Root().Search("Ad/InLine/Creatives/Creative/Linear")
    if err != nil {
        return
    }
    ccc[0].SetAttr("offset", "00:00:08")

//  USE TO SEE METHODS OF INTERFACES
//  fooType := reflect.TypeOf(ccc)
//  fmt.Println(fooType)
//  for i := 0; i < fooType.NumMethod(); i++ {
//      method := fooType.Method(i)
//      fmt.Println(method.Name)
//  }
}
//  END VIEWS
func main() {
    fmt.Println("Server Running at 8000...")
//  URLS
    http.HandleFunc("/", some_request)
//  END URLS
    http.ListenAndServe(":8000", nil)
}

