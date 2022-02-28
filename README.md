# Economicus

> submodule 사용법을 설명하기 위해 임시 작성했습니다. 추후 수정이 필요합니다.

## Requirements

- `.env` `.env.mysql` `nginx.conf`
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
git commit -m "chore: economicus-fe update" # 커밋 메세지는 마음대로 작성해주세요.
```

❗ `git submodule update --remote`는 필요할 때에만 해주세요! 일반적으로는 프론트엔드 작업자가 production 레포지토리에 업데이트하고 commit 했을 것이므로 `git pull`만 해도 되는 듯...합니다.

### frontend 작업 방법

- 작업(코드 수정, 커밋, 푸쉬, PR...)은 production 레포지토리가 아닌, economicus-fe 레포지토리에서 합니다.
- economicus-fe 레포지토리에 수정사항(새로운 commit)이 있다면, production 레포지토리에서 `git submodule update --remote`로 업데이트 후, `docker-compose up -d --build` 명령어 등을 통해 테스트해볼 수 있습니다.
