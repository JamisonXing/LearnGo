package store

type Book struct {
	Id      string   `json:"id"`
	Name    string   `json:"name"`
	Authors []string `json:"authors"`
	Press   string   `json:"press"` //出版社
}

type Store interface {
	Create(*Book) error       //创建一个新的图书目录
	Update(*Book) error       //更新某个图书目录
	Get(string) (Book, error) //获取某个图书信息
	GetAll() ([]Book, error)  //获取所有图书的信息
	Delete(string) error      //删除某图书条目
}
