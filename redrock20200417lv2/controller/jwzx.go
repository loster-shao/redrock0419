package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"io/ioutil"
	"net/http"
	"redrock20200417lv2/mysql"
	"regexp"
	"strconv"
	"strings"
	"time"

	)

func Jwzx(c *gin.Context) {
	t1 := time.Now() // get current time
	db := mysql.DbConn()
	for i := 2019210001; i <= 2019215203; i++ {
		go Parse(i, db)
		time.Sleep(10 *time.Millisecond)
	}
	elapsed := time.Since(t1)
	time.Sleep(15 *time.Second)
	fmt.Println("爬虫结束,总共耗时: ", elapsed)
}

func (keyword Spiders) Get_html_header() string {
	client := &http.Client{}
	req, err := http.NewRequest("GET", keyword.url, nil)
	if err != nil {
		fmt.Println("err1:", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("err2:", err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("err3:", err)
	}
	return string(body)
}

func Parse(i int, db *gorm.DB) {
	//for i := s; i <= s + 1000; i++ {
	xh := strconv.Itoa(i)
	url := "http://jwc.cqupt.edu.cn/kebiao/kb_stu.php?xh=" + xh
	spider := &Spiders{url}
	html := spider.Get_html_header()
	reBody := strings.ReplaceAll(html, "\r\n", "") //去除所有换行 用""替代
	s := strings.ReplaceAll(reBody, " ", "")       //去除所有空格 用""替代
	body := strings.ReplaceAll(s, "\t", "")

	partern := `2019-2020学年2学期学生课表>>` + xh + `(.*?)</li>`
	liReg := regexp.MustCompile(partern)
	find := liReg.FindAllStringSubmatch(body, -1)
	student := find[0][1]

	//爬取课表
	partern1 := `<tdrowspan='\d+'>(.*?)<tdrowspan='\d+'></td></tr>`
	liReg1 := regexp.MustCompile(partern1)
	find1 := liReg1.FindAllString(body, -1)

	person := []Person{}
	for _, v := range find1 {
		liReg2 := regexp.MustCompile(`<tdrowspan='\d+'>(.*?)</td><tdrowspan='\d+'>(.*?)</td><tdrowspan='\d+'>(.*?)</td><tdrowspan='\d+'align='center'>(.*?)</td><td>(.*?)</td><td>(.*?)</td><td>(.*?)</td>`)
		find2 := liReg2.FindStringSubmatch(v)

		person = append(person, Person{
			Name:    find2[1],
			Class:   find2[2],
			Bx:      find2[3],
			Status:  find2[4],
			Teacher: find2[5],
			Time:    find2[6],
			Where:   find2[7],
		})
	}
	students := []Student{}
	students = append(students, Student{
		Stu:    student,
		Xh:     i,
		Person: person,
	})
	fmt.Println("students", students)
	//}
	db.AutoMigrate(&students)
	db.Create(Student{
		Stu:    student,
		Xh:     i,
		Person: person,
	})//创建表，鬼知道创建出来长啥样
}




