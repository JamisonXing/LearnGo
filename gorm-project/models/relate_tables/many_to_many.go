package relate_tables

type Article2 struct {
	Id      int
	Title   string `gorm:"type:varchar(204)"`
	Content string `gorm:"type:varchar(204)"`
	Desc    string `gorm:"type:varchar(204)"`
	UId     uint
	Tag     []Tag `gorm:"many2many:Article2_Tags"`
}

type Tag struct {
	Id   int
	Name string `gorm:"type:varchar(204)"`
	Desc string `gorm:"type:varchar(204)"`
}
