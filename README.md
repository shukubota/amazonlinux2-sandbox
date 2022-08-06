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

## sqs
```
aws sqs list-queues --endpoint-url=http://localhost:4566 --profile=localstack
```

### circleci api
例
```sh
curl -u token: -X POST --header "Content-Type: application/json" -d '{
  "parameters": {},
  "branch": "main"
}' https://circleci.com/api/v2/project/gh/shukubota/amazonlinux2-sandbox/pipeline
```

branchを指定できる
tokenはhttps://app.circleci.com/settings/user/tokensから指定できる.