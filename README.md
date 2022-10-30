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



## ecs sample
最小構成でecsを構築

```shell
docker build -t ecs-sample -f Dockerfile_ecssample .
```
でbuild

ちなみに
```shell
docker run --publish 127.0.0.1:9998:9998 ecs-sample 
```
で
```shell
curl http://localhost:9997/healthcheck
```

でレスポンスが帰る

これをecrにあげる

### build
```shell
aws --region ap-northeast-1 ecr get-login-password | docker login --username AWS --password-stdin 523363116161.dkr.ecr.ap-northeast-1.amazonaws.com/ecs-sample
docker build -f Dockerfile_ecssample . -t ecs-sample -t 523363116161.dkr.ecr.ap-northeast-1.amazonaws.com/ecs-sample:1
docker push 523363116161.dkr.ecr.ap-northeast-1.amazonaws.com/ecs-sample:1 
```
