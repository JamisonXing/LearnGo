package relate_tables

//一对多,外键放在多的那边

type User2 struct {
	Id       int
	Name     string `gorm:"type:varchar(204)"`
	Age      int
	Addr     string    `gorm:"type:varchar(204)"`
	Articles []Article `gorm:"ForeignKey:UId;AssociationForeignKey:Id"`
}

type Article struct {
	Id      int
	Title   string `gorm:"type:varchar(204)"`
	Content string `gorm:"type:varchar(204)"`
	Desc    string `gorm:"type:varchar(204)"`
	UId     uint
}
