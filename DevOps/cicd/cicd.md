### CI/CD Tool

### Docker+Jenkins+GitLab 搭建 CI/CD 系统
- 第一步：安装 Docker
- 第二步：安装 GitLab
  ```
    docker run -d \
    --hostname localhost \
    -p 8080:80 -p 2222:22 \
    --name gitlab \
    --restart always \
    --volume /tmp/gitlab/config:/etc/gitlab \
    --volume /tmp/gitlab/logs:/var/log/gitlab \
    --volume /tmp/gitlab/data:/var/opt/gitlab \
    gitlab/gitlab-ce:13.3.8-ce.0
  ```
- 第三步：安装 Jenkins
```
    docker run -d --name=jenkins \
    -p 8888:8080 \
    -u root \
    --restart always \
    -v /var/run/docker.sock:/var/run/docker.sock \
    -v /usr/bin/docker:/usr/bin/docker \
    -v /tmp/jenkins_home:/var/jenkins_home \
    jenkins/jenkins:lts
```
- 解锁jenkins
```
docker logs -f jenkins
```
- 配置 GitLab SSH 访问公钥
``` 
ssh-keygen -o -t rsa -b 2048 -C "email@example.com"
cat $HOME/.ssh/id_rsa.pub
```
- 上传服务代码到 GitLab
- 创建 Jenkins 任务
- 配置自动构建
- 配置自动部署

### 使用Gitlab Runner 部署
- 新建部署用户
```
#创建一个deployer用户
sudo useradd deployer

#修改用户权限，允许其新建目录等命令操作
ls -l /etc/sudoers
chmod -v u+w /etc/sudoers
#在最下面增加一个用户
vim /etc/sudoers
deployer        ALL=(ALL)       NOPASSWD: ALL
#写权限收回
chmod -v u-w /etc/sudoers
#设置用户密码
passwd deployer
#登录
ssh deployer@ip_address
#生成秘钥
ssh-keygen -t rsa
#查看私钥
cat /home/deployer/.ssh/id_rsa
```
- 安装 Gitlab Runner(docker)
- [Install GitLab Runner](https://docs.gitlab.com/runner/install/)
```
docker run -d --name gitlab-runner --restart always \
  -v /var/run/docker.sock:/var/run/docker.sock \
  -v /srv/gitlab-runner/config:/etc/gitlab-runner \
  gitlab/gitlab-runner:latest
```
- 新注册Gitlab Runner
```
gitlab-runner register
```
- 新建.gitlab.ci.yml配置文件(php demo)
```
variables:
  RELEASES_STORAGE_DIR: '/var/www/$CI_COMMIT_REF_NAME/$CI_PROJECT_PATH/storage'
  CREATE_RELEASES_STORAGE_DIR: '[ -d $RELEASES_STORAGE_DIR ] || sudo mkdir -p $RELEASES_STORAGE_DIR'
  RELEASES_DIR: '/var/www/$CI_COMMIT_REF_NAME/$CI_PROJECT_PATH/releases'
  CREATE_RELEASE_DIR: '[ -d $RELEASES_DIR ] || sudo mkdir -p $RELEASES_DIR'
  NEW_RELEASES_DIR: '$RELEASES_DIR/$CI_COMMIT_SHORT_SHA'
  CREATE_NEW_RELEASES_DIR: '[ -d $NEW_RELEASES_DIR ] || sudo mkdir -p $NEW_RELEASES_DIR'
  BEFORE_CHMOD: 'sudo chown -R deployer:deployer $NEW_RELEASES_DIR'
  BEFORE_CHMOD_VENDOR: 'sudo chown -R deployer:deployer $NEW_RELEASES_DIR/vendor'
  AFTER_CHMOD: 'sudo chown -R apache:apache /var/www/$CI_COMMIT_REF_NAME && sudo chown -R apache:apache $RELEASES_STORAGE_DIR && sudo chmod -R 777 $RELEASES_STORAGE_DIR'
  CD_NEW_RELEASES_DIR: 'cd $NEW_RELEASES_DIR'
  CD_RELEASES_DIR: 'cd $RELEASES_DIR'
  #Linux删除除了某个文件之外的所有文件/目录
  CLEAN_RELEASES_DIR: 'ls |grep -v $CI_COMMIT_SHORT_SHA |xargs sudo rm -rf'
  RM_RELEASE_STORAGE_DIR: 'sudo rm -rf $NEW_RELEASES_DIR/storage'
  LN_RELEASE_STORAGE_DIR: 'sudo ln -nfs $RELEASES_STORAGE_DIR $NEW_RELEASES_DIR/storage'
  LN_RELEASE_DIR: 'sudo ln -nfs $NEW_RELEASES_DIR /var/www/$CI_COMMIT_REF_NAME/$CI_PROJECT_PATH/current'
  MV_REPO: 'sudo mv -fv /home/deployer/$CI_PROJECT_DIR/* $NEW_RELEASES_DIR'
  CP_DEV_ENV: 'cp /home/deployer/$CI_PROJECT_DIR/.env.dev $NEW_RELEASES_DIR/.env'
  CREATE_FRAMEWORK_CACHE: '[ -d $RELEASES_STORAGE_DIR/framework/cache ] || sudo mkdir -p $RELEASES_STORAGE_DIR/framework/cache'
  CREATE_FRAMEWORK_SESSIONS: '[ -d $RELEASES_STORAGE_DIR/framework/sessions ] || sudo mkdir -p $RELEASES_STORAGE_DIR/framework/sessions'
  CREATE_FRAMEWORK_TESTING: '[ -d $RELEASES_STORAGE_DIR/framework/testing ] || sudo mkdir -p $RELEASES_STORAGE_DIR/framework/testing'
  CREATE_FRAMEWORK_VIEWS: '[ -d $RELEASES_STORAGE_DIR/framework/views ] || sudo mkdir -p $RELEASES_STORAGE_DIR/framework/views'


before_script:
  - echo "Before script"
  - echo $CI_COMMIT_REF_NAME
  - echo $CI_PROJECT_PATH
  - echo $CI_COMMIT_SHORT_SHA
  - echo $CI_REPOSITORY_URL
  - echo $CI_PROJECT_DIR
  - 'eval $CREATE_RELEASES_STORAGE_DIR'  # will execute
  - 'eval $CREATE_RELEASE_DIR'  # will execute
  - 'eval $CREATE_NEW_RELEASES_DIR'  # will execute
  - 'eval $CD_NEW_RELEASES_DIR'


stages:
  - build
  - test
  - deploy-dev

building:
  stage: build
  script:
    - echo "Move repo..."
    - echo $NEW_RELEASES_DIR
    - 'eval $BEFORE_CHMOD'
    - 'eval $MV_REPO'
    - composer install
    - 'eval $BEFORE_CHMOD_VENDOR'

testing:
  stage: test
  script:
    - echo "testing..."
    # - php ./vendor/bin/phpunit

deploying_dev:
  stage: deploy-dev
  script:
    - echo "deploying dev..."
    - 'eval $CP_DEV_ENV'
    - php artisan key:generate
    - 'eval $CREATE_FRAMEWORK_CACHE'
    - 'eval $CREATE_FRAMEWORK_SESSIONS'
    - 'eval $CREATE_FRAMEWORK_TESTING'
    - 'eval $CREATE_FRAMEWORK_VIEWS'
    - php artisan cache:clear
    - php artisan config:clear
    - php artisan storage:link
    - php artisan migrate --force
    - php artisan passport:keys
    - echo "Restarting supervisor"
    - sudo supervisorctl restart all
    - echo "Linking storage directory"
    - 'eval $RM_RELEASE_STORAGE_DIR'
    - 'eval $LN_RELEASE_STORAGE_DIR'
    - echo 'Linking current directory'
    - 'eval $AFTER_CHMOD'
    - 'eval $LN_RELEASE_DIR'
    - echo 'Removing earlier app'
    - 'eval $CD_RELEASES_DIR'
    - 'eval $CLEAN_RELEASES_DIR'
  only:
    - develop

```