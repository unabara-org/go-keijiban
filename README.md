## マイグレーションを実行する

- [golang\-migrate/migrate](https://github.com/golang-migrate/migrate) を使う
- 使い方として Go プログラムから呼び出すか、CLI から呼び出す
- 本プロジェクトでは CLI から呼び出す
  - CLI についてのインストール方法などはこちらを参考にする
  - https://github.com/golang-migrate/migrate/tree/master/cmd/migrate
- 以下のコマンド例は fish shell で確認した

マイグレーションコマンド

```
migrate --path /db/migrations -database <database_url> up
```

MySQL のデータベース URL は以下

```
mysql://user:password@tcp\(host:port\)/dbname?query
```

例

```
$ migrate --path ./db/migrations -database mysql://user:password@tcp\(localhost:3306\)/go_keijiban up
```
