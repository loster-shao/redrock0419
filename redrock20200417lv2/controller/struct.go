package controller

type Spider struct {
	url    string
	header map[string]string
}

type Spiders struct {
	url    string
	//header map[string]string
}

//豆瓣电影
type Movie struct {
	Id    uint     `gorm:"primary_key"`
	Name  string   //名字
	P     string   //评语
	Num   string   //评论数量
	Score float64  //评分
	Jpg   string   //图片地址
	D     string   //导演
	Ok    int
}

//学生课程
type Person struct {
	Stu     string
	Xh      int
	Name    string  //课程名
	Class   string  //编号
	Bx      string  //必修/选修/重修
	Status  string  //课程状态
	Time    string  //时间
	Where   string  //地点
	Teacher string  //老师
}

//课程
type Class struct {
	Name    string  //课程名
	Class   string  //编号
	Bx      string  //必修/选修/重修
	Status  string  //课程状态
	Time    string  //时间
	Where   string  //地点
	Teacher string  //老师
}

//学生课程
type Students struct {
	Stu    string
	Xh     int
	Class  []Class
}

//学生课程
type Student struct {
	Stu    string
	Xh     int
	Person  []Person
}
