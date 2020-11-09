package main

import (
  "fmt"
  "net/http"
  "net/http/httputil"
  //"io/ioutil"
)

func main() {
  url := "http://google.co.jp"
  getByHttpClient(url)
  /*
  resp, _ := http.Get(url)
  defer resp.Body.Close()

  byteArray, _ := ioutil.ReadAll(resp.Body)
  fmt.Println(string(byteArray)) // htmlをstringで取得
  */
}
func getByHttpClient(url string ){
   req, _ := http.NewRequest("GET", url, nil)
   req.Header.Set("Authorization", "Bearer access-token")
   
   dump, _ := httputil.DumpRequestOut(req, true)
   fmt.Printf("%s", dump)
   
   client := new(http.Client)
   resp, _ := client.Do(req)
   dumpResp, _ := httputil.DumpResponse(resp, true)
   fmt.Printf("%s", dumpResp)
}
