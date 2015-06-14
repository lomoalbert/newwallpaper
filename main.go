package main

import (
    "net/http"
    "fmt"
    "io/ioutil"
    "encoding/xml"
    "os"
    "flag"
)

type bingimage struct {
    XMLName xml.Name `xml:"images"`
    Images []image `xml:"image"`
}

type image struct{
    Url string  `xml:"url"`
    UrlBase string  `xml:"urlBase"`
}

func getimg(url,finame string){
    response,err:=http.Get(url)
    if err!=nil{
        fmt.Println(err.Error())
    }
    defer response.Body.Close()
    fd,err :=ioutil.ReadAll(response.Body)

    fi, err := os.OpenFile(finame, os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        panic(err)
    }
    defer fi.Close()
    fi.Write(fd)
}

func main(){
    defer func(){
        if info:=recover();info!=nil{
            fmt.Println(info)
        }
    }()
    finame := flag.String("path", "", "path to backgroud picture.")
    flag.Parse()
    if *finame==""{
        fmt.Println("请输入路径: --path=/home/yourname/Pictures/wallpaper")
        return
    }
    url:="http://www.bing.com/HPImageArchive.aspx?format=xml&idx=0&n=8"
    response,err:=http.Get(url)
    if err!=nil{
        fmt.Println(err.Error())
    }
    defer response.Body.Close()
    fd,err :=ioutil.ReadAll(response.Body)
    var todayimg bingimage
    err =xml.Unmarshal(fd, &todayimg)
    imgurl:="https://bing.com"+todayimg.Images[0].Url
    getimg(imgurl,*finame)
}
