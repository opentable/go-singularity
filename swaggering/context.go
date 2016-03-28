package swaggering

type Context struct {
	swaggers   []*Swagger
	apis       []*Api
	models     []*Model
	openModels []*Model
}

func NewContext() (context *Context) {
	context = &Context{}
	context.swaggers = make([]*Swagger, 0)

	return
}
