# GoManagerScript

Ejecución y seguimiento a scripts desde api-rest con golang.

## Requerimientos

- Golang
- Postgres

## Instalación

Clone el proyecto en el local path de golang.

Ejemplo.

```
cd ~/go/src/github/ && git clone https://github.com/cvasquezz/GoManagerScript.git
```

En la ruta ```~/go/src/github/GoManagerScript/dataBase``` se encuentra los dump de la BDD en postgres

En el directorio del proyecto ejecute lo siguiete:
(no se si es necesario).

```
cd GoManagerScript
go get github.com/gorilla/mux
go get github.com/lib/pq
go get github.com/spf13/viper
go get -u github.com/golang/dep/cmd/dep
```
Antes de inciar la ejecución del api-rest, valide que las credenciales esten correctas según su configuración de clientes de BDD y puertos

el archivo de properties se encuenta en : ```GoManagerScript/config/properties.toml```

```
[dbProperties]
username = "postgres"
password = "123"
database = "postgres"

[serverRun]
port = "8080"

[mailProperties]
from = "xxxxxx@xxxx.com"
password = "xxxxxxx"
hostname = "smtp.gmail.com:587"
servername = "smtp.gmail.com"

```
el mailProperties es para configurar el envío de mails


Una vez instaladas las librerias hay que ejecutar el script firstExec.sh
este script creara un servicio llamado GoManagerScript y ejecutara un script que "compilara" y levantara la aplicación

```
sudo sh firstExec.sh
```
Cuando se realicen cambios a nivel de código en la aplicación solo se debe ejecutar el siguiete scritp para compilar y levantar los cambios.

```
sh deploy.sh
```

## Uso

Como es un api-rest se deben consumir los servicios a travez de postman, insomnia, etc. Por el momento existen 4 servicios:

### Ingreso de usuario:
URL Request
```
http://localhost:8080/api/v1/user/insert
```
Body
```
{
	"UserName": "myUser",
	"Password": "123456",
	"Nombre": "myName",
	"Apellidos": "Mock",
	"Fono":"+56966666666",
	"Permiso": 1
}
```
### Login
URL Request
```
http://localhost:8080/api/v1/user/login
```
Body
```
{
	"UserName": "myUser",
	"Password": "123456"
}
```
### Listar todos los scripts configurados
URL Request
```
http://localhost:8080/api/v1/script/getAll
```

### Ejecutar script
URL Request
```
http://locahost:8080/api/v1/execShell
```
Body
```
{
	"Mallid": "sh",
	"Kind": "ejemplo"
}
```
## Licencia
[MIT](https://choosealicense.com/licenses/mit/)