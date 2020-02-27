## 公有云安装gitlab-runner和docker方法.md
`以下内容以腾讯云服务器(centos 7.5+)操作为例（root用户），阿里云、aws、华为云对照操作即可`

### install docker-ce 
 yum remove docker docker-client docker-client-latest docker-common docker-latest docker-latest-logrotate  docker-logrotate docker-engine -y

 yum install -y yum-utils device-mapper-persistent-data lvm2

 yum-config-manager --add-repo https://download.docker.com/linux/centos/docker-ce.repo

 yum install docker-ce docker-ce-cli  containerd.io -y

##### 增加腾讯云加速器
cat /etc/docker/daemon.json
{
   "registry-mirrors": [
       "https://mirror.ccs.tencentyun.com"
  ]
}

##### 启动docker服务
 systemctl start docker 

##### 拉起一个容器，测试验证docker
 docker run hello-world


### install gitlab-runner
wget https://mirrors.cloud.tencent.com/gitlab-runner/yum/el7/gitlab-runner-12.7.1-1.x86_64.rpm

yum install -y gitlab-runner-12.7.1-1.x86_64.rpm

>> 补充：如果是二进制安装gitlan-runner,这需要以下添加：

```
### creat gitlab-runner user and set docker user-power
useradd --comment 'GitLab Runner' --create-home gitlab-runner --shell /bin/bash

groupadd docker    
gpasswd -a gitlab-runner docker 
newgrp docker  
su gitlab-runner 
docker ps   

gitlab-runner install --user=gitlab-runner --working-directory=/home/gitlab-runner

gitlab-runner start
```

##### gitlab-runner向gitlab注册，并install和start gitlab-runner
gitlab-runner register

gitlab-runner install

gitlab-runner start

##### git push code and check gitlab-cicd jobs

