// t03 project main.go
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

/*响应信息结构体*/
type fmRoot struct {
	R    int
	Song []fmSong
}

type fmSong struct {
	Album       string
	Picture     string
	Ssid        string
	Artist      string
	Url         string
	Company     string
	Title       string
	Rating_avg  int
	Length      int64
	Subtype     string
	Public_time string
	Sid         string
	Aid         string
	Sha256      string
	Kbps        string
	Albumtitle  string
	Like        int
}

/*统计信息*/
type Sinfo struct {
	etime     float32
	totalCot  int
	totalMer  int
	succeeCot int
}

const (
	dir = "H:\\life story\\music\\"
)

//抓取音乐列表
func smain() []byte {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://douban.fm/j/mine/playlist?type=n&sid=347955&pt=2.9&channel=-3&pb=64&from=mainsite&r=1d815d7ebf", nil)
	req.Header.Add("Accept:", "*/*")
	//req.Header.Add("Accept-Encoding", "gzip,deflate,sdch")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.8")
	req.Header.Add("Cookie", "openExpPan=Y; bid=\"8YicOUp+Kx0\"; dbcl2=\"56345753:wqSu+cIqVy8\"; fmNlogin=\"y\"; __utma=58778424.1673168777.1384438525.1385773681.1385773681.15; __utmb=58778424.1.9.1385790386347; __utmc=58778424; __utmz=58778424.1384438525.1.1.utmcsr=(direct)|utmccn=(direct)|utmcmd=(none)")
	req.Header.Add("Host", "douban.fm")
	req.Header.Add("Referer", "http://douban.fm/")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.2; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/30.0.1599.101 Safari/537.36")
	resp, err := client.Do(req)

	if resp.StatusCode == 200 {
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		return body
	} else {
		log.Fatal(err)
		os.Exit(1)
	}
	return nil
}

func main() {
	fmt.Println("Spilder start................!")
	startTime := time.Now().UnixNano()
	sinfo := Sinfo{}
	sum := 1

	for sum < 10 {
		data := smain()
		if data != nil {
			var root fmRoot
			err := json.Unmarshal(data, &root)
			if err != nil {
				fmt.Println("Start Down File....")
				crawMusic(root.Song, &sinfo)
			} else {
				log.Fatal(err)
			}
		}
	}

	sinfo.etime = float32(time.Now().UnixNano()-startTime) / 1e9
	fmt.Printf("success cot %d  memory %d(M) total cot %d \n", sinfo.totalCot, sinfo.totalMer, sinfo.totalCot)
	fmt.Printf("exceute time %.3fs\n", sinfo.etime)
}

/*
*抓取
 */
func crawMusic(song []fmSong, sinfo *Sinfo) {
	sinfo.totalCot = sinfo.totalCot + len(song)
	for _, v := range song {
		if !FileExist(dir+v.Title+".mp3") && v.Url != "" {
			log.Println(v.Url + " " + v.Title + "  " + strconv.FormatInt(v.Length, 10))
			downMusic(v, sinfo)
		}
	}
}

func downFile(url string) []byte {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Accept-Encoding", "gzip,deflate,sdch")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.8")
	req.Header.Add("Cookie", "openExpPan=Y; bid=\"8YicOUp+Kx0\"; dbcl2=\"56345753:wqSu+cIqVy8\"; fmNlogin=\"y\"; __utma=58778424.1673168777.1384438525.1385773681.1385773681.15; __utmb=58778424.1.9.1385790386347; __utmc=58778424; __utmz=58778424.1384438525.1.1.utmcsr=(direct)|utmccn=(direct)|utmcmd=(none)")
	req.Header.Add("Host", "mr3.douban.com")
	req.Header.Add("Referer", "http://douban.fm/")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.2; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/29.0.1547.76 Safari/537.36")
	resp, _ := client.Do(req)

	if resp.StatusCode == 200 {
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		return body
	} else {
		log.Println("error -->" + url)
	}
	return nil
}

/*
*下载文件
 */
func downMusic(song fmSong, sinfo *Sinfo) {

	file := dir + song.Title + ".mp3"
	dts := downFile(song.Url)
	if dts != nil {
		defer func() {
			f, _ := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
			defer f.Close()
			ioutil.WriteFile(file, dts, os.ModePerm)
			sinfo.succeeCot++
		}()
	} else {
		log.Println("save file error")
	}
}

/*
*保存歌词详细信息
 */
func FileInfo() {
	//TODO
}

/*
*文件是否存在
 */
func FileExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}
