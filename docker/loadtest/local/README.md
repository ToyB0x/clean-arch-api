# locust(docker)によるローカル負荷テスト
```shell
# increase worker when each cpu core usage over 60%
cd local
docker-compose up --scale worker=8
```
