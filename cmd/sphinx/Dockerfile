FROM jrei/systemd-centos:7

COPY SphinxService.viper.yaml /etc/SphinxService/SphinxService.viper.yaml
COPY sphinx-service /usr/local/bin/sphinx-service

CMD ["sphinx-service", "run"]
