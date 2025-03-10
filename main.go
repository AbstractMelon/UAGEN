package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

var browsers = []string{
	"Mozilla/5.0 (%s) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/%d.0.%d.%d Safari/537.36",
	"Mozilla/5.0 (%s) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/%d.0 Mobile/15E148 Safari/604.1",
	"Mozilla/5.0 (%s; rv:%d.0) Gecko/20100101 Firefox/%d.0",
	"Mozilla/5.0 (%s) AppleWebKit/537.36 (KHTML, like Gecko) Edge/%d.0.%d.%d",
	"Mozilla/5.0 (%s) AppleWebKit/537.36 (KHTML, like Gecko) Opera/%d.0.%d.%d",
	"Mozilla/5.0 (%s) AppleWebKit/537.36 (KHTML, like Gecko) Brave/%d.0.%d.%d Safari/537.36",
}

var osSystems = []string{
	"Windows NT 10.0; Win64; x64",
	"Windows NT 6.1; Win64; x64",
	"Windows NT 10.0; WOW64",
	"Macintosh; Intel Mac OS X 10_%d_%d",
	"X11; Linux x86_64",
	"X11; Ubuntu; Linux x86_64",
	"X11; Fedora; Linux x86_64",
	"iPhone; CPU iPhone OS %d_%d like Mac OS X",
	"iPad; CPU OS %d_%d like Mac OS X",
	"Android %d; Mobile",
	"Android %d; Tablet",
}

func randomInt(min, max int) int {
	return rand.Intn(max-min+1) + min
}

func generateUserAgent() string {
	os := osSystems[rand.Intn(len(osSystems))]
	var osParams []any
	if strings.Contains(os, "%d") {
		osParams = make([]interface{}, strings.Count(os, "%d"))
		for i := range osParams {
			osParams[i] = randomInt(0, 99)
		}
		os = fmt.Sprintf(os, osParams...)
	}
	
	browser := browsers[rand.Intn(len(browsers))]
	return fmt.Sprintf(browser, os, randomInt(60, 120), randomInt(0, 9999), randomInt(0, 999))
}

var amount = flag.Int("amount", 1, "amount of user agents to generate")
var filename = flag.String("file", "", "filename to write to")

func main() {
	rand.Seed(time.Now().UnixNano())
	flag.Parse()
	if *filename != "" {
		data := make([]byte, 0)
		for i := 0; i < *amount; i++ {
			data = append(data, []byte(generateUserAgent()+"\n")...)
		}
		ioutil.WriteFile(*filename, data, 0644)
	} else {
		for i := 0; i < *amount; i++ {
			fmt.Println(generateUserAgent())
		}
	}
}


