#!/bin/bash
echo "Deteniendo el servicio"
sudo systemctl stop GoManagerScript
echo "Borrando el binario"
rm -rf GoManagerScript
echo "creando nuevo binario"
GOOS=linux GOARCH=amd64 go build
#mv GoManagerScript GoManagerScript
echo "iniciando el servicio"
sudo systemctl start GoManagerScript