package main

import (
	"bufio"
	"fmt"
	"os"
	/*   "strings"  */
	"flag"
	"regexp"
)

func main() {

/* command line options: */

/* input files: */
	file1Ptr := flag.String("f1", "", "name of first fasta file.")
	file2Ptr := flag.String("f2", "", "name of 2nd fasta file.")


	flag.Parse()

	fmt.Printf("%s\n", *file1Ptr)
	fmt.Printf("%s\n", *file2Ptr)

	fh1, err := os.Open(*file1Ptr)
	if err != nil {
		os.Exit(1)
	}

	
var sequences1 []string
index_id := make(map[int]string)
id_index := make(map[string]int)
index := 0
	scanner1 := bufio.NewScanner(fh1)
	for scanner1.Scan() {
		line := scanner1.Text()
		r, _ := regexp.Compile("^>([0-9]+).*")
		if r.MatchString(line) { /* >id line */
	match_strings := r.FindStringSubmatch(line)
			index_id[index] = match_strings[1]
			id_index[match_strings[1]] = index	
			fmt.Println(match_strings[1])
		} else { /* sequence line */
			sequences1 = append(sequences1, line)
			fmt.Printf("   [%s]\n", line)
			index++
		}
	}

/*
	fh2, err := os.Open(*file2Ptr)
	if err != nil {
	os.Exit(1)
}
*/


}



