package relate_tables

//一对一

type User struct {
	Id   int
	Name string `gorm:"type:varchar(204)"`
	Age  int
	Addr string
}

type UserProfile struct {
	Id    int
	Pic   string
	CPic  string
	Phone string

	User User `gorm:"ForeignKey:UId;AssociationForeignKey:Id"` //关联关系
	//UserID int  //默认关联字段
	UId int //自定义外键名
}
