package swaggering

import (
	"fmt"
	"log"
)

func ResolveService(context *Context) {
	log.Print("Resolving types")

	log.Print(context.swaggers)
	for _, swagger := range context.swaggers {
		for adx := range swagger.Apis {
			context.apis = append(context.apis, &swagger.Apis[adx])
		}
		for _, model := range swagger.Models {
			context.models = append(context.models, model)
		}
	}

	log.Printf("  Found %d apis", len(context.apis))
	log.Printf("  Found %d models", len(context.models))

	context.openModels = make([]*Model, 0, len(context.models))

	resolveApis(context)
	resolveModels(context)
}

func resolveModels(self *Context) {
	var cur *Model

	for len(self.openModels) > 0 {
		cur, self.openModels = self.openModels[0], self.openModels[1:]
		if cur.GoUses {
			continue
		}
		cur.GoUses = true
		cur.GoName = cur.Id
		for name, prop := range cur.Properties {
			prop.GoName = capitalize(name)
			prop.findGoType(self)
		}
	}
}

func logErr(err error, format string, args ...interface{}) {
	if err != nil {
		args = append(args, err)
		log.Output(2, fmt.Sprintf(format, args...))
	}
}

func resolveApis(context *Context) {
	var err error

	for _, api := range context.apis {
		for _, op := range api.Operations {
			op.Path = api.Path
			err = op.findGoType(context)
			logErr(err, "Operation %s invalid: %v", op.Nickname)

			for _, parm := range op.Parameters {
				err = parm.findGoType(context)
				logErr(err, "Operation %s invalid: parameter %s: %v", op.Nickname, parm.Name)

				if parm.GoTypeInvalid {
					op.GoTypeInvalid = true
				}

				if parm.Name == "body" {
					op.HasBody = true
				}
			}
		}
	}

}

func (context *Context) modelFor(typeName string, to *DataType) (err error) {
	err = context.modelUsed(typeName)
	to.GoModel = true
	to.setGoType(typeName, err)
	return
}

func (self *Context) modelUsed(name string) (err error) {
	for _, model := range self.models {
		if model.Id == name {
			if !model.GoUses {
				self.openModels = append(self.openModels, model)
			}
			return
		}

	}
	err = fmt.Errorf("Model %q doesn't appear in known models.", name)
	return
}
