package main

import(
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"log"
	"time"
	"sync"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/driver/mysql"
	//"github.com/jinzhu/gorm"
	//_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/tidwall/gjson"
)

// 存入数据库的结构体
type Video struct{
	Id		int64
	Channel_id	int64
	Name		string
	View_count	int64
	Like_count	int64
	Author_name	string
	Author_id	int64
	Bvid		string
}

var wg sync.WaitGroup

// 将点赞数、播放量转换成整数
func SwitchNum(str string) int64 {
	slice := []rune(str)
	lastword := string(slice[len(slice)-1:])
	num, _ := strconv.ParseFloat(string(slice[:len(slice)-1]), 64)

	if lastword == "万" {
		return int64(num * 10000)
	}

	return int64(num)
}

// 连接数据库
func InitDB()*gorm.DB {
	dsn := "root:hu20010326@tcp(127.0.0.1:3306)/bilibili?charset=utf8mb4&parseTime=True&loc=Local"
	//DB, err := gorm.Open("mysql", dns)
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	fmt.Println("连接数据库成功")
	return DB
}

// 连接单个频道的视频数据，判断频道是否存在
func ConnectChannel(channel_id int, ch chan string, chId chan int64) {
	defer wg.Done()

	fmt.Println("连接频道", channel_id)

	requestUrl := "https://api.bilibili.com/x/web-interface/web/channel/multiple/list?channel_id="+strconv.Itoa(channel_id)+"&sort_type=hot&page_size=30"
	req, err := http.NewRequest("GET", requestUrl, nil)
	if err != nil {
		log.Println(err.Error())
		return
	}

	response, _ := http.DefaultClient.Do(req)

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer response.Body.Close()
	content := string(body)

	videoDatas := gjson.Get(content, "data.list").String()
	if videoDatas == "" {
		fmt.Println("频道", channel_id, "无内容")
		return
	}

	fmt.Println("频道", channel_id, "连接成功")

	ch <- content
	chId <- int64(channel_id)
}

// 得到单个视频信息
func GetAVideo(video_No int, content string, channel_id int64, videoCh chan Video) {

	var video Video

	videoInfo := gjson.Get(content, "data.list."+strconv.Itoa(video_No)).String()

	video.Id = SwitchNum(gjson.Get(videoInfo, "id").String())
	video.Channel_id = channel_id
	video.Name = gjson.Get(videoInfo, "name").String()
	video.View_count = SwitchNum(gjson.Get(videoInfo, "view_count").String())
	video.Like_count = SwitchNum(gjson.Get(videoInfo, "like_count").String())
	video.Author_name = gjson.Get(videoInfo, "author_name").String()
	video.Author_id = SwitchNum(gjson.Get(videoInfo, "author_id").String())
	video.Bvid = gjson.Get(videoInfo, "bvid").String()

	fmt.Println("获取第", video_No, "视频")

	videoCh <- video
}

// 获取一个频道下的所有视频
func GetAllVideo(ch chan string, chId chan int64, videoCh chan Video) {
	defer wg.Done()

	content := <-ch
	channel_id := <-chId
	for i := 1; i <= 30; i ++ {
		GetAVideo(i, content, channel_id, videoCh)
	}
}

// 将视频信息存入数据库
func SaveInDB(DB *gorm.DB, videoCh chan Video) {
	defer wg.Done()

	video := <-videoCh

	DB.Clauses(clause.OnConflict{
		UpdateAll: true,
		}).Create(&video)

	fmt.Println("存入数据库")
}

func Spider(channel_id int, ch chan string, chId chan int64) {
	ConnectChannel(channel_id, ch, chId)

}

func main(){
	t1:=time.Now()
	fmt.Println("开始爬虫")

	DB := InitDB()
	//defer DB.Close()

	// 视频信息相关的通道
	videoContent := make(chan string)
	defer close(videoContent)
	chId := make(chan int64)
	defer close(chId)
	videoCh := make(chan Video)
	defer close(videoCh)

	for i := 100; i <= 200; i++ {
		wg.Add(1)
		go ConnectChannel(i, videoContent, chId)
	}

	wg.Add(1)
	go GetAllVideo(videoContent, chId, videoCh)

	for i:=0; i < 30; i++ {
		wg.Add(1)
		go SaveInDB(DB, videoCh)
	}

	wg.Wait()
	fmt.Println("结束爬虫")

	elapsed:=time.Since(t1)
	fmt.Println(elapsed)
}
