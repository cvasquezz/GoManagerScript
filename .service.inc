[Unit]
Description=Servicio de ejecucion de script
After=network.target

[Service]
User=root
Group=www-data

ExecStart=RUTA

[Install]
WantedBy=multi-user.target