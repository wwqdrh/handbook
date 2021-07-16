package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/antchfx/htmlquery"
)

var client = &http.Client{Timeout: time.Second * 10}

var startYear = flag.Int("start", 0, "搜索开始年份")
var endYear = flag.Int("end", 0, "搜索结束年份")

type event struct {
	MovieID string
	Name    string
	Time    string   // 2020-04-16
	Keyword []string // 定档、预告、海报
}

// 根据url发送get，返回数据内容
// 设置猫眼cookie
// __mta=108886717.1618662419949.1618671165063.1618671428226.4; uuid_n_v=v1;
// uuid=31075C709F7811EB895E79679BE70E79C4DA301379A249619B1E25A3F3F0DC8E;
// _csrf=6bf238dae15de4aa06f1292b65f2a4b99caf2e2ec287d4677d9b2bfc98dc486c;
// Hm_lvt_703e94591e87be68cc8da0da7cbd0be2=1618662418;
// _lx_utm=utm_source%3Dgoogle%26utm_medium%3Dorganic;
// _lxsdk_cuid=178dfcc6660c8-02250de790d554-6a15217c-1fa400-178dfcc6660c8;
// _lxsdk=31075C709F7811EB895E79679BE70E79C4DA301379A249619B1E25A3F3F0DC8E; __mta=108886717.1618662419949.1618671165063.1618671236874.4; Hm_lpvt_703e94591e87be68cc8da0da7cbd0be2=1618671428; _lxsdk_s=178e04809d9-e01-693-11%7C%7C11
func addCookies(req *http.Request, cookies string) {
	for _, item := range strings.Split(cookies, ";") {
		keyValue := strings.Split(item, "=")
		if len(keyValue) != 2 {
			continue
		}
		req.AddCookie(&http.Cookie{
			Name:  keyValue[0],
			Value: keyValue[1],
		})
	}
}

func doGet(url string) string {
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3776.0 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		log.Println(fmt.Sprintf("[发送请求失败] doget - %s", url))
		return ""
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil && data == nil {
		log.Println(fmt.Sprintf("[解析数据失败] doget - %s", url))
		return ""
	}
	return string(data)
}

/**
 * 根据start, end获取对应的yearid 全部 2021(16) 2020(15) 2019 2018 2017 2016 2015 2014 2013 2012 2011 2000-2010 90年代 80年代 70年代 更早
 * 1到17，最后一个不要参数
 * 2010-2021
 */
func getYearIDByRange(start, end int) (res []string) {
	curYear := time.Now().Year()
	maxID := curYear - 2010 + 5

	var left, right int
	// start确定左边界
	if start < 1970 {
		left = 1 // 更早
	} else if start < 1980 {
		left = 2 // 70年代
	} else if start < 1990 {
		left = 3 // 80年代
	} else if start < 2000 {
		left = 4
	} else if start < 2010 {
		left = 5 // 2000-2010
	}
	// end确定右边界
	if end < 1970 {
		right = 1
	} else if end < 1980 {
		right = 2
	} else if end <= 1990 {
		right = 3
	} else if end < 2000 {
		right = 4
	} else if end <= 2010 {
		right = 5
	} else if end <= curYear {
		right = end - 2010 + 5 // 比如2011，就是6，2021就是16
	} else {
		right = maxID
	}

	if left == 1 && right == maxID {
		res = append(res, "0") // 所有
		return res
	}

	for ; left <= right; left++ {
		res = append(res, strconv.Itoa(left))
	}
	return res
}

/**
 * 给定开始以及结束时间来搜索电影的id，返回[]string, 如果是0那就是所有的内容
 * https://maoyan.com/films?yearId=12&showType=3&offset=30
 */
func searchMovieIDByYearID(yearID string) (res []string) {
	// 根据url，获取最大的页数，以及offset规则，拼接url，并解析出内容来
	data := doGet(fmt.Sprintf("https://maoyan.com/films?yearId=%s&showType=3", yearID))
	if data == "" {
		return nil
	}

	root, _ := htmlquery.Parse(strings.NewReader(data))
	// 解析最大页数
	node := htmlquery.Find(root, "//ul[@class='list-pager']/li[last()-1]/a/text()")
	if node == nil {
		log.Println("[searchmovieidbyyearid] - 数据解析失败可能需要验证")
		return nil
	}
	maxPage, err := strconv.Atoi(htmlquery.InnerText(node[0]))
	if err != nil {
		log.Println(fmt.Sprintf("[解析最大页数失败] searchid - %s", yearID))
		return nil
	}

	// 解析第一页的内容里面的movieid
	movieNodes := htmlquery.Find(root, "//dl[@class='movie-list']/dd/div[1]/a")
	if movieNodes == nil {
		log.Println("[searchmovieidbyyearid] - 数据解析失败可能需要验证")
	} else {
		for _, node := range movieNodes {
			res = append(res, strings.TrimLeft(htmlquery.SelectAttr(node, "href"), "/films/")) // /films/36
		}
	}
	// 解析页数*30 offset的内容movieid
	for idx := 2; idx <= maxPage; idx++ {
		time.Sleep(time.Second * 1 / 2)
		curUrl := fmt.Sprintf("https://maoyan.com/films?yearId=%s&showType=3&offset=%d", yearID, idx*30)
		data := doGet(curUrl)
		if data == "" {
			continue
		}

		curRoot, _ := htmlquery.Parse(strings.NewReader(data))
		movieNodes := htmlquery.Find(curRoot, "//dl[@class='movie-list']/dd/div[1]/a")
		if movieNodes == nil {
			log.Println("[searchmovieidbyyearid] - 数据解析失败可能需要验证")
			continue
		}
		for _, node := range movieNodes {
			res = append(res, strings.TrimLeft(htmlquery.SelectAttr(node, "href"), "/films/")) // /films/36
		}
	}
	return res
}

/**
 * 根据电影id获取对应的营销事件
 */
func getEventByID(movieID string) (res []event) {
	return res
}

func main() {
	// 命令行，指定开始年份，结束年份
	flag.Parse()
	if *startYear == 0 || *endYear == 0 {
		flag.Usage()
		os.Exit(0)
	}
}
