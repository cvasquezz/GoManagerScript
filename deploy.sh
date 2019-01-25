#!/bin/bash
echo "Deteniendo el servicio"
sudo systemctl stop pullService
echo "Borrando el binario"
rm -rf pullService
echo "creando nuevo binario"
GOOS=linux GOARCH=amd64 go build
mv pullShell pullService
echo "iniciando el servicio"
sudo systemctl start pullService