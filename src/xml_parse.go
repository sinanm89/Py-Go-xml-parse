package main

import (
    "fmt"
    "io/ioutil"
    "strings"
    "github.com/moovweb/gokogiri"
    "runtime"
    "net/url"
    "net/http"
    "reflect"
)

func dir (obj interface{}) {
    //  USE TO SEE METHODS OF INTERFACES
    fooType := reflect.TypeOf(obj)
    fmt.Println("------DIR-------")
    for i := 0; i < fooType.NumMethod(); i++ {
        method := fooType.Method(i)
        fmt.Println(method.Name)
    }
    fmt.Println("------/DIR-------")
    fmt.Println(reflect.TypeOf(obj))
}

// VIEWS
func some_request(w http.ResponseWriter, r *http.Request) {

    u, _ := url.Parse(r.URL.String())
    queryParams := u.Query()
    pub_id := queryParams.Get("publisher_id")
    if pub_id == "" {
        url_path := strings.TrimLeft(r.URL.Path, "/")
        url_path_list := strings.Split(url_path, "/")
        pub_id = url_path_list[0]
    }

    uri := fmt.Sprintf("http://ad4.liverail.com/?LR_PUBLISHER_ID=%s&LR_SCHEMA", pub_id)

    response, err := http.Get(uri)
    if err != nil {
        return
    }
    defer response.Body.Close()

    data, err := ioutil.ReadAll(response.Body)
    if err != nil {
        http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusBadRequest)
        return
    }

    doc, err := gokogiri.ParseXml(data)
    if err != nil {
        http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusBadRequest)
        return
    }
    defer doc.Free()

    xml_linear, err := doc.Root().Search("Ad/InLine/Creatives/Creative/Linear")
    if err != nil {
        http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusBadRequest)
        return
    }
    if len(xml_linear) == 0 {
        http.Error(w, fmt.Sprintf("Error: Bad XML data Received check your ID"), http.StatusBadRequest)
        return
    }

    doc.Root().SetAttr("version", "3.0")
    xml_linear[0].SetAttr("skipoffset", "00:00:08")

//  xml_response, output_buffer := doc.Root().ToXml(nil, nil)
//  if output_buffer == 0 {
//      http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusBadRequest)
//      return
//  }
    string_doc := doc.String()
    w.Header().Set("Content-Type", "text/xml")
    fmt.Fprintf(w, string_doc, r.Body)
//  w.Write(string_doc)
}
//  END VIEWS

func main() {
    runtime.GOMAXPROCS(-1)   //  Sets the concurrency to all the cores available

    fmt.Println("Server Running at 8000...")
//  URLS
    http.HandleFunc("/", some_request)
//  END URLS
    http.ListenAndServe(":8000", nil)
}

