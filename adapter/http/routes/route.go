package route

type Route interface {
	Setup()
}

type Routes []Route

func (r *Routes) Setup() {
	for _, route := range *r {
		route.Setup()
	}
}

func NewRoute(bookRoute *BookRoute) *Routes {
	return &Routes{
		bookRoute,
	}
}
