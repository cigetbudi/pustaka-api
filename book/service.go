package book

type Service interface {
	FindAll() ([]Book, error)
	FindById(ID int) (Book, error)
	Create(bookRequest BookRequest) (Book, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Book, error) {
	books, err := s.repository.FindAll()
	return books, err
	//return s.Repository.FindAll()
}

func (s *service) FindById(ID int) (Book, error) {
	book, err := s.repository.FindById(ID)
	return book, err
}

func (s *service) Create(bookRequest BookRequest) (Book, error) {
	price, _ := bookRequest.Price.Int64()
	discount, _ := bookRequest.Discount.Int64()
	rating, _ := bookRequest.Rating.Int64()

	book := Book{
		Title:       bookRequest.Title,
		Price:       int(price),
		Decsription: bookRequest.Description,
		Discount:    int(discount),
		Rating:      int(rating),
	}

	newBook, err := s.repository.Create(book)
	return newBook, err

}
