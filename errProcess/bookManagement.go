package bookManagement

type BookDatabase struct{}

func NewBookDatabase() *BookDatabase {
	var db BookDatabase
	return *db
}

func (d *BookDatabase) GetBook(id string) *Book {}

func (d *BookDatabase) CreateBook(data interface{}) {}
