package get_port_token

//package get_port_token

import (
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"czw_lol_query_tools/lcu"
)

var (
	cli = http.Client{
		Transport: &http.Transport{
			ForceAttemptHTTP2: true,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
)

const (
	apiUrlFmt = "https://riot:%s@127.0.0.1:%d"
)

var (
	portFlag = flag.Int("port", 8098, "lcu 代理端口")
)

func Return_port_token() string {

	flag.Parse()
	//port := *portFlag
	lcuPort, lcuToken, err := lcu.GetLolClientApiInfo()

	if err != nil {

		log.Fatal(err)
	}

	proxyURL := fmt.Sprintf(apiUrlFmt, lcuToken, lcuPort)
	go func() {
		ticker := time.NewTicker(time.Second * 3)
		for { //这是个会一直执行的循环
			<-ticker.C
			var lcuPort, lcuToken, err = lcu.GetLolClientApiInfo()

			if err != nil {
				continue
			}
			updateProxyURL := fmt.Sprintf(apiUrlFmt, lcuToken, lcuPort)

			if updateProxyURL == proxyURL {
				continue
			}
			proxyURL = updateProxyURL
			//log.Println("update lcu:", proxyURL)
		}
	}()
	//log.Printf("listen on :%d, lcu api:%s\n", port, proxyURL)

	return proxyURL
	/*
		err = http.ListenAndServe(fmt.Sprintf(":%d", port), http.HandlerFunc(func(w http.ResponseWriter,
			r *http.Request) {
			req, _ := http.NewRequest(r.Method, proxyURL+r.URL.Path+"?"+r.URL.RawQuery, nil)
			req.Body = r.Body
			req.Header = r.Header
			resp, err := cli.Do(req)
			if err != nil {
				_, _ = fmt.Fprintf(w, "err : %v", err)
				return
			}
			w.Header().Set("Content-Length", resp.Header.Get("Content-Length"))
			w.Header().Set("Content-Type", resp.Header.Get("Content-Type"))
			_, _ = io.Copy(w, resp.Body)
			defer func() {
				_ = resp.Body.Close()
			}()
		}))
		if err != nil {
			log.Fatal(err)
		}
	*/

}
