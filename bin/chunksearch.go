package main

import (
       "bufio"
       "fmt"
       "os"
    /*   "strings"  */
       "regexp"
       "flag"
       )


func main() {
 

file1Ptr := flag.String("f1", "", "name of first fasta file.")
file2Ptr := flag.String("f2", "", "name of 2nd fasta file.")

flag.Parse()

fmt.Printf("%s\n", *file1Ptr)
fmt.Printf("%s\n", *file2Ptr)

fh1, err := os.Open(*file1Ptr)
if(err != nil){
os.Exit(1)
}
scanner1 := bufio.NewScanner(fh1)

for scanner1.Scan() {
    line := scanner1.Text()
    r, _ := regexp.Compile("^>([0-9]+).*")
    if(r.MatchString(line)){
    match_strings := r.FindStringSubmatch(line)
    fmt.Println(match_strings[1])
    }else{
    fmt.Printf("   %s\n", line)
}
}



}