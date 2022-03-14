# Economicus

> submodule 사용법을 설명하기 위해 임시 작성했습니다. 추후 수정이 필요합니다.

## Requirements

- `.env` `.env.mysql` `.env.mongo` `nginx.conf`
- 실행 방법

  ```shell
  git clone --recursive https://github.com/economicus/production

  docker-compose up -d --build
  ```

## Submodule 추가 설명

### submodule 업데이트

economicus-fe 서브모듈을 최신 버전으로 업데이트 하고 싶은 경우, 다음의 명령어를 사용해주세요.

```shell
git submodule update --remote
```

위와 같이 서브모듈 업데이트 후 commit하면, production 레포지토리에 반영됩니다.

```shell
git add .
git commit -m "chore: update" # 커밋 메세지는 마음대로 작성해주세요.
```
