# Go TODO API チュートリアル

## 取り組んだ記事

- [Go初心者のためのTODO APIハンズオン 前編](https://qiita.com/aaaasahi_17/items/673bc8bd5b5579aed28a)
- [Go初心者のためのTODO APIハンズオン 後編](https://qiita.com/aaaasahi_17/items/c9eee71eaad323e021e4)

## 起動方法

DBの起動

```bash
docker compose up
```

APIの起動

```bash
go run main.go
```

どうせ`docker-compose.yml`にもがっつり書いてあるので、`.env`もプッシュしておきます。

## TODOのデータ構造

```go
type Todo struct {
    ID        int       `json:"id"`
    Title     string    `json:"title"`
    Content   string    `json:"content"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
```

## エンドポイント

| パス | メソッド | 機能 |
| ---- | ---- | --- |
| /todos | GET | 全件取得 |
| /todos/:id | GET | 一件取得 |
| /todos | POST | 作成 |
| /todos/:id | PUT | 更新 |
| /todos/:id | DELETE | 削除 |
