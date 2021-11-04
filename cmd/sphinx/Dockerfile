FROM scratch

COPY ServiceSample.viper.yaml /etc/ServiceSample/ServiceSample.viper.yaml
COPY sphinx-service /usr/local/bin/sphinx-service

CMD ["sphinx-service", "run"]
