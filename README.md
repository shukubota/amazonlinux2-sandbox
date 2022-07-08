# amazonlinux2-sandbox

## 手順

```sh
docker-compose up
docker-compose exec backend bash
```

## s3
```
aws s3 create-bucket --bucket test  --endpoint=http://localhost:14566
```

でbucket作成

```
aws s3 ls --endpoint=http://localhost:14566
```
で確認