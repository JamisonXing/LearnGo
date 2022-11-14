package relate_tables

//一对一，属于
//属于：关系和外键在同一方，有关系的哪一方属于另外一个模型
/*type User struct {
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
*/

//一对一,包含
//包含：关系和外键不在同一方，有关系的那一方包含另外一个模型

type User struct {
	Id   int
	Name string `gorm:"type:varchar(204)"`
	Age  int
	Addr string `gorm:"type:varchar(204)"`
	PId  int
}

type UserProfile struct {
	Id    int
	Pic   string `gorm:"type:varchar(204)"`
	CPic  string `gorm:"type:varchar(204)"`
	Phone string `gorm:"type:varchar(204)"`
	User  User   `gorm:"ForeignKey:PId;AssociationForeignKey:Id"` //关联关系
}
