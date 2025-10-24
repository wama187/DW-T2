package book

type BookRepository[T any, BID comparable, UID comparable] interface {
	Create(b *T) error
	Update(b *T) error
	Delete(bookID BID, ownerID UID) error
	FindByID(b *T) (*T, error)
	FindAllByOwner(ownerID UID) ([]*T, error)
}


type BookPostgre struct {
	ID    	uint
	Title   string
	Author  string
	OwnerID uint
}

type BookMongo struct {
	ID      string
	Title   string
	Author  string
	OwnerID uint
}
