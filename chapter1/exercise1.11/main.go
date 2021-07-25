// Fetchall fetches URLs in parallel and reports their names and sizes.
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine; calls fetch asynchronously
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.2f elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body) // discard the response body by writing to ioutil.Discard output stream
	resp.Body.Close()                                 // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}

/*
\chapter1> .\exercise1.11 https://www.alexa.com/siteinfo/google.com https://www.alexa.com/siteinfo/youtube.com https://www.alexa.com/siteinfo/tmall.com https://www.alexa.com/siteinfo/qq.com https://www.alexa.com/siteinfo/baidu.com https://www.alexa.com/siteinfo/sohu.com https://www.alexa.com/siteinfo/zoom.us https://www.alexa.com/siteinfo/live.com https://www.alexa.com/siteinfo/netflix.com https://www.alexa.com/siteinfo/yy.com
0.76s   294626  https://www.alexa.com/siteinfo/yy.com
0.78s   299158  https://www.alexa.com/siteinfo/tmall.com
0.79s   299099  https://www.alexa.com/siteinfo/sohu.com
0.88s   300110  https://www.alexa.com/siteinfo/live.com
0.96s   298305  https://www.alexa.com/siteinfo/netflix.com
1.02s   299519  https://www.alexa.com/siteinfo/zoom.us
2.20s   298079  https://www.alexa.com/siteinfo/baidu.com
2.55s   299327  https://www.alexa.com/siteinfo/qq.com
3.21s   251025  https://www.alexa.com/siteinfo/youtube.com
4.45s   298271  https://www.alexa.com/siteinfo/google.com
4.45 elapsed
*/
