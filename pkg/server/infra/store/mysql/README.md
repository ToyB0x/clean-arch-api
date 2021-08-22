# CLI README
https://github.com/golang-migrate/migrate/tree/master/cmd/migrate

マイグレーション用ファイルの作成コマンド
```shell script
# migrate create -ext sql [操作内容] 
migrate create -ext sql create_user_table
```

# SQL BOILERによるコード生成

```shell
# install
go get github.com/volatiletech/sqlboiler/v4
go get github.com/volatiletech/null/v8
go get github.com/go-sql-driver/mysql
go get github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql@latest

# generate codes
sqlboiler mysql
```
