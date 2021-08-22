FROM locustio/locust:1.5.3

COPY ./locustfile.py /home/locust
COPY run.sh /run.sh

ENTRYPOINT ["/run.sh"]
