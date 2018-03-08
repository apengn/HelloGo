package main

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	//"fmt"
	//"encoding/json"
	_ "encoding/json"
	"fmt"
	//"encoding/json"
	//"os/user"
	"encoding/json"
)

func init() {
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/go_g?charset=utf8")
	orm.RegisterModel(new(User), new(Post), new(Profile), new(Tag))
	orm.RunSyncdb("default", false, true)
}

type User struct {
	Id      int      `json:"id"`
	Name    string   `orm:"size(100)",json:"name"`
	Profile *Profile `orm:"rel(one)"`      //one  to one
	Post    []*Post  `orm:"reverse(many)"` //设置一对多的反向关系
}
type Profile struct {
	Id   int
	Age  int16
	User *User `orm:"reverse(one)"` //一对一的反向关系（可选）
}

type Post struct {
	Id    int    `orm:"auto",json:"id"`
	Title string `orm:"size(100)",json:"title"`
	User  *User  `orm:"rel(fk)",json:"user"` //设置一对多的关系
	Tags  []*Tag `orm:"rel(m2m)"`            //多对多的关系
}
type Tag struct {
	Id    int
	Name  string
	Posts []*Post `orm:"reverse(many)"` //多对多的反向关系
}

func main() {
	orm.Debug = true

	o := orm.NewOrm()
	//
	//o.Using("default") //// 默认使用 default，你可以指定为其他数据库
	//
	//profile := new(Profile)
	//
	//profile.Age = 30
	//
	//user:=new(User)
	//
	//user.Profile=profile
	//
	//user.Name="beego"
	//
	//fmt.Println(o.Insert(profile))
	//fmt.Println(o.Insert(user))
	defer func() {

		if err := recover(); err != nil {
			fmt.Println(err)
		}

	}()
	//user := new(User)
	//user.Name = "beego"
	//通过指定的字段进行查询
	//o.Read(user,"Name")

	//user := User{Name: "beego222"}
	//user.Profile=&Profile{Age:99}
	//if created, id, erro := o.ReadOrCreate(&user, "Name"); erro == nil {
	//
	//	if created {
	//		fmt.Println("New Insert an object.Id:",id)
	//	} else {
	//		fmt.Println("Get an object.Id:",id)
	//	}
	//}else {
	//	fmt.Println(erro)
	//}

	//o.Read(user.Profile)

	//err:=o.Read(user.Profile.User)
	//if err ==orm.ErrNoRows{
	//	fmt.Println("查询不到")
	//}else if (err==orm.ErrMissPK) {
	//	fmt.Println("找不到主键")
	//} else {
	//
	//}
	//json, _ := json.Marshal(user)
	//fmt.Println(string(json))
	//users := []User{{Name: "xxxxxx",Profile: &Profile{Age: 32,Id:12}}, {Name: "WWWWWWWWWWWWW",Profile:&Profile{Age:89,Id:13}}}
	//id,err:=o.InsertMulti(100, users)
	//fmt.Println(id,"========",err)

	user := new(User)
	var profile []*User
	//o.QueryTable(user).Filter("profile__age__in", 20,30,32).Exclude("profile__lt",2).All(&profile)
	//
	//json, _ := json.Marshal(profile)
	//fmt.Println(string(json))

	cond := orm.NewCondition()
	cond.And("profile__isnull", false).AndNot("profile__age__in", 30)//.Or("profile__id__gt", 0)
	//Limit(5,10)  5条数据，从索引值为10的地方开始查询5条数据

	o.QueryTable(user).SetCond(cond).All(&profile)
	for _,p := range (profile) {
		o.Read(p.Profile)
	}
	json, _ := json.Marshal(profile)
	fmt.Println(string(json))
}
