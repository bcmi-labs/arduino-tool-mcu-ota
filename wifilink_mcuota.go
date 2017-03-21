package main

import (
    "os"
    "fmt"
    "strconv"
    "io"
    "bufio"
    "math"
    "gopkg.in/resty.v0"
    "flag"
)

var filename string
var address string
var port int
var nlines int

const endpoint = "otafile"


func upload_files(list []string){
    url_base := "http://" + address + ":" + strconv.Itoa(port) + "/" + endpoint
    totchunk := len(list)
    
    fmt.Println("Sending " + filename + " to host " + address)
    
    for index, item := range list{        
    
        url := url_base + "?totchunk=" + strconv.Itoa(totchunk) + "&numchunk=" + strconv.Itoa(index+1)
        resp, err := resty.R().
        SetFile("files", item).
        Post(url)
        
        if resp.StatusCode() == 200 || err != nil {
            fmt.Println("["+strconv.Itoa(index+1)+" / "+strconv.Itoa(totchunk)+"] Done")
                
            if index+1==totchunk {
                fmt.Println("Upload all done")
            }
        } else
        if resp.StatusCode() > 399 || err == nil {
            fmt.Println("["+strconv.Itoa(index+1)+" / "+strconv.Itoa(totchunk)+"]")
            fmt.Print("Error: ")
            fmt.Println(err)
        }
    }
    clean(list)
}

func split(){
    file_list := []string {}

    lines := nlines
     file, err := os.Open(filename)
     if err != nil {
             panic(err)
     }

     defer file.Close()
     
     index := 1
     counter := 0
     var output string
        reader := bufio.NewReader(file)
        for {
            line, _, err := reader.ReadLine()
            
            if err == io.EOF {
                 break
            }

          output += string(line) + "\n"
         
            if math.Mod(float64(index) , float64(lines)) == 0 || index == getFileLines(filename) {
                     output_file, err := os.Create(filename + "." + strconv.Itoa(counter+1) + ".out")
                     output += string(line) + "\n"
                     if err != nil {
                            fmt.Println(err)
                        }
                     output_file.WriteString(output)
                     output_file.Close();
                     file_list = append (file_list, filename + "." + strconv.Itoa(counter+1) + ".out")
                     output = ""
                     counter ++   
            }
         index ++ 
     }    
    upload_files(file_list)
}

func clean(list []string){
    counter:=1
    for _, item := range list{
        os.Remove(item)
        counter++
    }
}

func get_args(){
    flag.StringVar(&filename, "f", "", "firmware file to upload (required)")
    flag.StringVar(&address, "i", "", "ip of the esp8266 based board to upload to (required)")
    flag.IntVar(&port, "p", 80, "network port")
    flag.IntVar(&nlines, "l", 20, "max lines for each splitted files")

    flag.Parse()
    
    if len(filename)>0 && len(address)>0 {
        split()
    } else {
        fmt.Print("usage : ")
        fmt.Println("arduino_mcuota [-h] -f FILE -i IP [-p PORT] [-l LINES]")
        fmt.Print("error: ")
        
        if len(filename) == 0 {
            fmt.Println("-f is required")
        }
        if len(address) == 0 {
            fmt.Println("-i is required")
        }
    }
}

func getFileLines(filename string)(int){
        tot:=0
        file, err := os.Open(filename)
        if err != nil {
             panic(err)
        }
        defer file.Close()

        reader := bufio.NewReader(file)
        for {
                 _, _, err := reader.ReadLine()
                if err == io.EOF {
                         break
                 }
                tot++
        }
        return tot
}


func main(){
    get_args()
}