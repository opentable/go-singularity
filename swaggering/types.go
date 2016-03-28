package swaggering

type (
	DataType struct {
		GoType                 string
		GoTypeInvalid, GoModel bool
		Type, Format           string
		Ref                    string `json:"$ref"`
		Enum                   []string
	}

	Collection struct {
		DataType
		Items       DataType
		UniqueItems bool
	}

	Swagger struct {
		BasePath, ResourcePath string
		Apis                   []Api
		Models                 map[string]*Model
	}

	Api struct {
		Path, Description string
		Operations        []*Operation
	}

	Operation struct {
		Nickname, Method, Path, Deprecated string
		HasBody                            bool
		Parameters                         []*Parameter
		ResponseMessages                   []*ResponseMessage
		Collection
	}

	Parameter struct {
		ParamType, Name         string
		Required, AllowMultiple bool
		Collection
	}

	ResponseMessage struct {
		Code                   int
		Message, ResponseModel string
		model                  *Model
	}

	Model struct {
		Id, Description, Discriminator string
		GoName                         string
		GoUses                         bool
		Required, SubTypes             []string
		Properties                     map[string]*Property
	}

	Property struct {
		GoName string
		Collection
	}
)
