package main
import (
"fmt"
//"os"
//"time"
"io/ioutil"
"net/http"
"log"
//"bytes"
"regexp"
//"strings"
"reflect"
)


func get_url(url string){
    log.SetFlags(log.Lshortfile | log.LstdFlags)

    fmt.Println("----------",url,"----------------")
    resp, err := http.Get(url)
    if err != nil {
        log.Fatalln(err)
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatalln(err)
    }
    //src := string(body)
    //log.Println(src)

    //r := bytes.NewReader(body)
    re := regexp.MustCompile(`<title>(.+?)</title>`)

fmt.Println(reflect.ValueOf(re).Type())

    var matchResult [][]byte
    matchResult = re.FindSubmatch(body)

fmt.Println(reflect.TypeOf(matchResult))

    fmt.Println("len:", len(matchResult))
    if len(matchResult) > 1 {
        fmt.Println(string(matchResult[1]))
    }

}

func main() {
    get_url("http://www.baidu.com")
}