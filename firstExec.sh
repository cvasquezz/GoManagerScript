#!/bin/bash
################# NO MODIFICAR  #################
RUTA=`pwd`
APP="${RUTA}/pullService"
SERVICE=${RUTA}/service.inc
PKWORK=RUTA
#################################################
# VALIDO QUE SE EJECUTE CON SUDO
if [ "$(id -u)" != "0" ]; then
    echo "Lo siento, debe ejecutar la shell con sudo."
    echo "sudo ./firstExec.sh"
    exit 1
fi
# INSTALO LAS DEPENDENCIAS DEL PROYECTO
echo "Instalo las dependencias"
#go get -u github.com/golang/dep/cmd/dep
#export GOPATH=$(go env GOPATH)
#${GOPATH}/bin/dep ensure
go install .
#GENERO EL ARCHIVO PARA EL SERVICIO
echo "generando el servicio pullService"
cat .service.inc | sed "s~$PKWORK~$APP~" > pullService.service
sudo mv pullService.service /etc/systemd/system/pullService.service  
sudo systemctl daemon-reload                
#DEPLOYO EL PROYECTO
echo "deployando el proyecto"
sudo sh deploy.sh
echo "Para inciar el demonio ejecute"
echo "sudo systemctl start pullService"