package swaggering

type (
	DataType struct {
		GoType        string
		GoTypeInvalid bool
		Type, Format  string
		Ref           string `json:"$ref"`
		Enum          []string
	}

	Collection struct {
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
		Method, Nickname, Deprecated string
		Parameters                   []*Parameter
		ResponseMessages             []*ResponseMessage
		DataType
		Collection
	}

	Parameter struct {
		ParamType, Name         string
		Required, AllowMultiple bool
		DataType
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
		DataType
		Collection
	}
)
