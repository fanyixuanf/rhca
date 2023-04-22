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