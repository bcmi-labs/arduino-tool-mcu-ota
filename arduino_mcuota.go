package main

import (
    "os"
    "fmt"
    "strconv"
    "io"
    "bufio"
    "math"
    "gopkg.in/resty.v0"
)

var filename string
var nlines string
var port string 
var address string

const endpoint = "otafile"


func upload_files(list []string){
    url_base := "http://" + address + ":" + port + "/" + endpoint
    totchunk := len(list)
    
    fmt.Println("Sending " + filename + " to host " + address)
    
    for index, item := range list{        
    
        url := url_base + "?totchunk=" + strconv.Itoa(totchunk) + "&numchunk=" + strconv.Itoa(index+1)
        resp, err := resty.R().
        SetFile("files", item).
        Post(url)
        
        if resp.StatusCode() == 200 || err != nil {
            fmt.Println("["+strconv.Itoa(index+1)+" / "+strconv.Itoa(totchunk)+"] Done")
            if index==totchunk {
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

func parse_args(){
    port = "80";
    nlines = "20";
    args := os.Args[1:]

    for i:=0; i < len(args); i+=2 {
        switch args[i]{
            case "-f" , "--file":
                filename = args[i+1]
            case "-l" , "--lines":
                nlines = args[i+1] 
            case "-p" , "--port":
                port = args[i+1]
            case "-i" , "--ip":
                address = args[i+1]
        }
    }
    split()
}

func split(){
    file_list := []string {}

    lines, err := strconv.Atoi(nlines)
    if err != nil {
        fmt.Println(err)
    }
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

//            if err == nil {
//                 fmt.Print("Error: ")
//                 fmt.Println(err)
//                 break
//            }
            
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
    parse_args()
}