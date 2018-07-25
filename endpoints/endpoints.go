package endpoints

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/minhkhiemm/example-go/service"

	"github.com/minhkhiemm/example-go/endpoints/book"
	"github.com/minhkhiemm/example-go/endpoints/category"
	"github.com/minhkhiemm/example-go/endpoints/user"
)

// Endpoints .
type Endpoints struct {
	//struct user endpoints
	FindUser    endpoint.Endpoint
	FindAllUser endpoint.Endpoint
	CreateUser  endpoint.Endpoint
	UpdateUser  endpoint.Endpoint
	DeleteUser  endpoint.Endpoint
	//struct category endpoint
	FindCategory    endpoint.Endpoint
	FindAllCategory endpoint.Endpoint
	CreateCategory  endpoint.Endpoint
	UpdateCategory  endpoint.Endpoint
	DeleteCategory  endpoint.Endpoint
	//struct book endpoint
	FindBook    endpoint.Endpoint
	FindAllBook endpoint.Endpoint
	CreateBook  endpoint.Endpoint
	UpdateBook  endpoint.Endpoint
	DeleteBook  endpoint.Endpoint
}

// MakeServerEndpoints returns an Endpoints struct
func MakeServerEndpoints(s service.Service) Endpoints {
	return Endpoints{
		//return user endpoint
		FindUser:    user.MakeFindEndPoint(s),
		FindAllUser: user.MakeFindAllEndpoint(s),
		CreateUser:  user.MakeCreateEndpoint(s),
		UpdateUser:  user.MakeUpdateEndpoint(s),
		DeleteUser:  user.MakeDeleteEndpoint(s),
		//return category endpoint
		FindCategory:    category.MakeFindEndPoint(s),
		FindAllCategory: category.MakeFindAllEndpoint(s),
		CreateCategory:  category.MakeCreateEndpoint(s),
		UpdateCategory:  category.MakeUpdateEndpoint(s),
		DeleteCategory:  category.MakeDeleteEndpoint(s),

		//return book endpoint
		FindBook:    book.MakeFindEndPoint(s),
		FindAllBook: book.MakeFindAllEndpoint(s),
		CreateBook:  book.MakeCreateEndpoint(s),
		UpdateBook:  book.MakeUpdateEndpoint(s),
		DeleteBook:  book.MakeDeleteEndpoint(s),
	}
}
