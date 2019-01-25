package libs

/*
	auth : Chrstian Vásquez
	karibu SPA
*/
import (
	"github.com/GoManagerScript/model"
)

/*User ...*/
type User struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

/*UserAll ...*/
type UserAll struct {
	UserName  string
	Password  string
	Nombre    string
	Apellidos string
	Fono      string
	Permiso   int
}

/*ActionLogin ...*/
type ActionLogin interface {
	ObtenerTokenUser() model.TokenResult
}

/*ActionSetUser ...*/
type ActionSetUser interface {
	SetNewUser() model.RespSetUser
}

/*SetNewUser ...*/
func (u UserAll) SetNewUser() model.RespSetUser {
	var resp model.RespSetUser
	if u.UserName != "" && u.Password != "" && u.Nombre != "" && u.Apellidos != "" && u.Fono != "" && u.Permiso > 0 {
		pass := model.GetMD5Hash(u.Password)
		resp = model.InserUser(u.UserName, pass, u.Nombre, u.Apellidos, u.Fono, u.Permiso)
	} else {
		resp = model.RespSetUser{
			Data:    false,
			Message: "Todos los campos son obligatorios",
		}
	}

	return resp
}

/*ObtenerTokenUser ...*/
func (u User) ObtenerTokenUser() model.TokenResult {
	pass := model.GetMD5Hash(u.Password)
	userData := model.GetToken(u.UserName, pass)
	return userData
}
