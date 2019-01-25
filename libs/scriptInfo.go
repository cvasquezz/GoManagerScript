package libs

import "github.com/GoManagerScript/model"

/*GetAllScript ...*/
func GetAllScript() model.ScriptAll {
	data := model.GetScriptAll()
	return data
}
