package main

import (
	"fmt"
	"strings"
	"os"
	"net/http"
	"errors"
	"log"
	"bufio"
	"strconv"
)

func check(p_e error) {
	if p_e != nil {
		panic(p_e)
		log.Fatal(p_e)
	}
}

func keepCalm(p_e error) bool {
	if p_e != nil {
		return true
	}
	return false
}

func catchData(p_line string) ([5]string, error) {
	var l_doc [5]string
	l_fields := strings.Split(p_line, ";")
	if len(l_fields) != 5 {
		return l_doc, errors.New("strange line !")
	}

	l_doc[0] = l_fields[0]			// firstname
	l_doc[1] = l_fields[1]			// name
	l_doc[2] = l_fields[2]			// id
	l_doc[3] = l_fields[3]			// birthday
	l_doc[4] = l_fields[4]			// value

	return l_doc, nil
}

func creatIndex() {
	l_f, err := os.Open("/home/user/workspace/open/client_ES/go/index_clone.json")
	check(err)
	defer l_f.Close()
	l_req, err := http.NewRequest("PUT", "http://localhost:9200/clones", l_f)
	check(err)
	l_req.Header.Set("Content-Type", "application/x-ndjson")

	l_resp, err := http.DefaultClient.Do(l_req)
	check(err)
	defer l_resp.Body.Close()
}

func deleteIndex() {
	// Generated by curl-to-Go: https://mholt.github.io/curl-to-go
	l_req, err := http.NewRequest("DELETE", "http://localhost:9200/clones", nil)
	check(err)
	l_resp, err := http.DefaultClient.Do(l_req)
	check(err)
	defer l_resp.Body.Close()
}

func sendJson(p_file string) {
	// Generated by curl-to-Go: https://mholt.github.io/curl-to-go
	l_f, err := os.Open(p_file)
	check(err)
	defer l_f.Close()
	l_req, err := http.NewRequest("POST", "http://localhost:9200/_bulk", l_f)
	check(err)
	l_req.Header.Set("Content-Type", "application/x-ndjson")

	l_resp, err := http.DefaultClient.Do(l_req)
	check(err)
	defer l_resp.Body.Close()
}

func writeJson(p_file *os.File, p_doc [5]string) {
	fmt.Fprintf(p_file, "{ \"index\" : { \"_index\" : \"clones\", \"_type\" : \"clone\" } }\n")
	fmt.Fprintf(p_file, "{\"firstname\":\"%s\",\"name\":\"%s\",\"id\":\"%s\",\"birthday\":\"%s\",\"value\":\"%s\"}\n",
		p_doc[0], p_doc[1], p_doc[2], p_doc[3], p_doc[4])
}

func worker(p_c <-chan string, p_end chan string, p_n int) {
	l_bulk,_ := strconv.Atoi(os.Args[1])
	l_i := 1
	l_fileName := "bulk" + strconv.Itoa(p_n) + ".json"
	l_json, err := os.Create(l_fileName)
	check(err)
	for {
		l_line := <-p_c
		if l_line == "end" {
			p_end <- ""
			break
		}
		if l_i%l_bulk == 0 {
			sendJson(l_fileName)
			err := os.Remove(l_fileName)
			check(err)
			l_json, err = os.Create(l_fileName)
			check(err)
			//fmt.Println("bulk n°", strconv.Itoa(n))
		}

		l_res, err := catchData(string(l_line))
		if keepCalm(err) {
			fmt.Println("Problem at line", l_i)
		} else {
			writeJson(l_json, l_res)
		}
		l_i++
	}
}

func progress(p_c <-chan int, p_size int) {
	l_accSize := 0
	for {
		l_size := <-p_c
		if l_size < 0 {
			l_pc := l_accSize * 100 / p_size
			fmt.Println(strconv.Itoa(l_pc) + "% inject")
		}
		l_accSize += l_size
	}
}

func listenStdin(p_c chan int) {
	l_reader := bufio.NewReader(os.Stdin)
	for {
		_, err := l_reader.ReadString('\n')
		check(err)
		p_c <- -1
	}
}

func main() {
	deleteIndex()
	creatIndex()

	file, err := os.Open("/home/user/workspace/open/client_ES/tmp/data_clone.csv")
	check(err)
	stat, err := file.Stat()
	check(err)
	defer file.Close()
	reader := bufio.NewReader(file)

	c_progress := make(chan int)
	c_doc := make (chan string)
	c_end := make (chan string)

	nbWorker,_ := strconv.Atoi(os.Args[2])								// <====== number of goroutine
	for i := 1; i <= nbWorker; i++ {
		go worker(c_doc, c_end, i)
	}

	go progress(c_progress, int(stat.Size()))
	go listenStdin(c_progress)
	fmt.Println("Injection with " + strconv.Itoa(nbWorker) + " goroutines")
	fmt.Println("Press *enter* for know advancement")

	i := 1
	for {
		line, _, err := reader.ReadLine()
		if line == nil {
			break
		}
		check(err)
		c_doc <- string(line)
		c_progress <- len(line)
		i++
	}



	for i := 1; i <= nbWorker; i++ {
		c_doc <- "end"
	}
	i=0
	for {
		<-c_end
		i++
		if i == nbWorker {
			break
		}
	}

	for i := 1; i <= nbWorker; i++ {
		sendJson("bulk"+strconv.Itoa(i)+".json")
		err = os.Remove("bulk"+strconv.Itoa(i)+".json")
		check(err)
	}
}