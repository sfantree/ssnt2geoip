package main

import (
	"bytes"
	"errors"
	"fmt"
	htmlquery "github.com/antchfx/xquery/html"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"net/http"
	"os/exec"
	"reflect"
	"regexp"
	"runtime"
	"strings"
	"sync"
	"time"
)

// https://www.cnblogs.com/Kingram/p/12627606.html
func GbkToUtf8(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

func getRawOutput() (string){
	var out []byte
	var err error
	if runtime.GOOS == "windows" {
		out, err = exec.Command("netstat", "-n").Output()
	} else {
		out, err = exec.Command("ss","-nt").Output()
	}
	if err != nil {
		panic(err)
	}
	utf8, err := GbkToUtf8(out)
	if err != nil {
		panic(err)
	}
	s := string(utf8[:])
	return (s)
}

func removeDuplicates(elements []string) []string {   // change string to int here if required
	// Use map to record duplicates as we find them.
	encountered := map[string]bool{}                   // change string to int here if required
	result := []string{}                               // change string to int here if required

	for v := range elements {
		if encountered[elements[v]] == true {
			// Do not add duplicate.
		} else {
			// Record this element as an encountered element.
			encountered[elements[v]] = true
			// Append to result slice.
			result = append(result, elements[v])
		}
	}
	// Return the new slice.
	return result
}

func getAllIpaddr(out string) ([]string){
	//out := getRawOutput()

	// net.ParseIP(out)

	ipReg := `((0|[1-9]\d?|1\d\d|2[0-4]\d|25[0-5])\.){3}[0-9]+`

	//ipReg := `(([1-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.)(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){2}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])`

	r := regexp.MustCompile(ipReg)
	//matches := r.FindStringSubmatch(out)
	matches := r.FindAllString(out, -1)
	var ans []string
	for _, v := range matches {
		ans = append(ans, v)
	}
	ret := removeDuplicates(ans)
	return ret
}

func getHtml(url_ string) string {
	var retries int = 10
	req, _ := http.NewRequest("GET", url_, nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3776.0 Safari/537.36")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,en-US;q=0.7")

	var resp *http.Response
	var err error
	var data []byte
	for retries > 0 {
		client := &http.Client{Timeout: time.Second * 10}
		resp, err = client.Do(req)
		if err != nil {
			//log.Println(err)
			retries -= 1
		} else {
			if resp != nil {
				defer resp.Body.Close()
				data, err = ioutil.ReadAll(resp.Body)
				if err != nil || len(data) < 2000 {
					// retries -= 1
					// log.Printf("%s response less than 2000", url_)
					continue
				} else {
					break
				}
			}
		}
	}
	// log.Printf("%s response data: %d", url_, len(data))
	return fmt.Sprintf("%s", data)

	//var (
	//	err      error
	//	response *http.Response
	//	retries  int = 3
	//)
	//for retries > 0 {
	//	response, err = http.Get("https://non-existent")
	//	// response, err = http.Get("https://google.com/robots.txt")
	//	if err != nil {
	//		log.Println(err)
	//		retries -= 1
	//	} else {
	//		break
	//	}
	//}
	//if response != nil {
	//	defer response.Body.Close()
	//	data, err := ioutil.ReadAll(response.Body)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	fmt.Printf("data = %s\n", data)
	//}
}

//func getHtml1(url_ string) string {
//
//	resp, err := retryablehttp.Get(url_)
//	if err != nil {
//		panic(err)
//	}
//
//	data, err := ioutil.ReadAll(resp.Body)
//
//	log.Printf("%s response data: %d", url_, len(data))
//	return fmt.Sprintf("data = %s\n", data)
//
//	//var (
//	//	err      error
//	//	response *http.Response
//	//	retries  int = 3
//	//)
//	//for retries > 0 {
//	//	response, err = http.Get("https://non-existent")
//	//	// response, err = http.Get("https://google.com/robots.txt")
//	//	if err != nil {
//	//		log.Println(err)
//	//		retries -= 1
//	//	} else {
//	//		break
//	//	}
//	//}
//	//if response != nil {
//	//	defer response.Body.Close()
//	//	data, err := ioutil.ReadAll(response.Body)
//	//	if err != nil {
//	//		log.Fatal(err)
//	//	}
//	//	fmt.Printf("data = %s\n", data)
//	//}
//}
//

func ReverseSlice(data interface{}) {
	value := reflect.ValueOf(data)
	if value.Kind() != reflect.Slice {
		panic(errors.New("data must be a slice type"))
	}
	valueLen := value.Len()
	for i := 0; i <= int((valueLen-1)/2); i++ {
		reverseIndex := valueLen - 1 - i
		tmp := value.Index(reverseIndex).Interface()
		value.Index(reverseIndex).Set(value.Index(i))
		value.Index(i).Set(reflect.ValueOf(tmp))
	}
}

func handleGeoString(raw string) string {
	var ret string
	replacer := strings.NewReplacer(",", "", "\n", "")
	ret = replacer.Replace(raw)

	var1 := removeDuplicates(strings.Split(ret, " "))
	ReverseSlice(var1)

	return strings.Join(var1[:], "")
}

func getIpinfo(ip string) [2]string {
	urlTemplate := "https://ip.sb/ip/%s"
	html := getHtml(fmt.Sprintf(urlTemplate, ip))
	root, _ := htmlquery.Parse(strings.NewReader(html))
	td := htmlquery.Find(root, "//*[@class='proto_location']")

	var protoLocation string
	var protoOrganization string

	if len(td) > 0 {
		protoLocation = htmlquery.InnerText(td[0])
		protoLocation = handleGeoString(protoLocation)
	}
	td = htmlquery.Find(root, "//*[@class='proto_organization']")
	if len(td) > 0 {
		protoOrganization = htmlquery.InnerText(td[0])
	}

	var last string
	if len(protoOrganization) > 0 {
		last = protoLocation + " " + protoOrganization
	} else {
		last = protoLocation
	}
	return [2]string{ip, last}
}

const xthreads = 5

func ConcurrentWork(jobList []string) [][2]string {
	var ch = make(chan string, 50) // This number 50 can be anything as long as it's larger than xthreads
	var wg sync.WaitGroup
	var mutex sync.Mutex

	// This starts xthreads number of goroutines that wait for something to do
	wg.Add(xthreads)

	var ResultList [][2]string

	for i := 0; i < xthreads; i++ {
		go func() {
			for {
				a, ok := <-ch
				if !ok { // if there is nothing to do and the channel has been closed then end the goroutine
					wg.Done()
					return
				}
				ret := getIpinfo(a) // do the thing

				mutex.Lock()
				ResultList = append(ResultList, ret)
				mutex.Unlock()

			}
		}()
	}

	// Now the jobs can be added to the channel, which is used as a queue

	for _, JobArgv := range jobList {
		ch <- JobArgv
	}

	//for i := 0; i < 50; i++ {
	//	ch <- i // add i to the queue
	//}

	close(ch) // This tells the goroutines there's nothing else to do
	wg.Wait() // Wait for the threads to finish

	return ResultList
}

func replaceOutput(raw string, tab [][2]string ) string {

	//fmt.Printf("%#v\n", tab)

	all := raw
	for _, v := range tab {
		regStr := v[0] + ":[0-9]+"
		regPattern := regexp.MustCompile(regStr)

		targets := regPattern.FindStringSubmatch(all)
		var target string
		if len(targets) > 0 {
			target = targets[0]
			all = regPattern.ReplaceAllString(all, target + " " + v[1])
		}
		// fmt.Printf("%#v\n", regPattern.FindStringSubmatch(all))
		// all = regPattern.ReplaceAllString(all, target + " " + v[1])
	}
	return all
}

func main() {
	out := getRawOutput()
	//var out string
	//out = "192.168.0.141 80.82.17.10 114.226.54.54 113.78.163.218 179.151.28.133 121.133.208.234 115.159.192.185 149.129.219.140 192.168.1.3 73.243.191.4 207.210.95.74 104.131.122.90 217.24.230.181 172.65.0.13 223.16.25.134 14.152.102.103 111.229.122.186 159.65.79.117 207.148.19.196 180.116.178.87 187.116.252.235 104.196.0.102 83.227.94.231 177.147.111.146 117.80.194.140 183.60.119.114 192.168.10.165 111.229.81.22 138.201.67.220 182.87.221.39 122.51.219.27 63.142.249.110 54.39.248.243 192.168.122.86 134.175.242.146 222.185.38.61 60.246.91.192 189.96.248.88 218.154.10.96 58.152.206.68 128.95.160.156 45.63.74.150 3.121.100.130 182.85.87.254 14.198.209.120 134.175.177.159 122.51.71.63 159.89.140.62 122.51.89.104 165.22.131.209 118.184.212.254 118.190.104.85 192.168.0.21 110.245.46.125 192.168.9.10 88.99.191.53 192.168.122.38 104.248.44.204 24.189.37.153 192.168.122.192 73.145.179.61 191.14.62.227 90.91.163.211 116.202.229.43 129.204.103.68 75.158.205.36 115.159.192.205 114.67.224.206 118.113.194.255 192.168.0.151 138.68.247.215 24.218.116.28 39.130.157.40 104.131.131.82 66.42.111.225 24.140.249.211 58.152.223.142 110.254.34.9 89.159.14.37 147.75.195.153 182.138.123.26 109.194.47.83 1.43.109.204 94.176.233.122 18.18.248.83 147.75.70.221 58.153.75.222 159.89.149.203 85.224.135.36 179.246.12.84 121.231.11.14"

	ips := getAllIpaddr(out)
	// log.Printf("ips: %v", ips)

	ipsgeo := ConcurrentWork(ips)
	// log.Printf("ips: %v", ipsgeo)

	//tab := getIpinfos(ips)
	ret := replaceOutput(out, ipsgeo)
	fmt.Printf("%s", ret)

}