# centos的docker初体验  

[toc]  

## 安装  
[参考文章](https://yeasy.gitbooks.io/docker_practice/install/centos.html)  

### 1、安装依赖包  
```bash
sudo yum install -y yum-utils \
           device-mapper-persistent-data \
           lvm2
```  

### 2、添加yum源  
```bash 
sudo yum-config-manager \
    --add-repo \
    https://mirrors.ustc.edu.cn/docker-ce/linux/centos/docker-ce.repo
```  
在添加yum源的时候报一个错误  
![image](6701C26CBD494878B8721C34065DFA38)  
这是因为python升级导致的  
解决方法:  
[参考文章](https://1028826685.iteye.com/blog/2404268)  
修改几个文件的python版本  
1、 /usr/libexec/urlgrabber-ext-down  
2、 /usr/bin/yum-config-manager  
将他们开头的#!/usr/bin/python换成#!/usr/bin/python2.7  


### 3、更新yum软件源缓存  
```bash
sudo yum makecache fast
```  
以上会同样报一个python版本问题的错误，只需要把对应文件的python版本换了就好  

### 4、安装docker-ce(ce是社区版不收费，ee是企业版)   
![image](CFA4CB7D307F42A389EC4CB2A66CEA55)  
```bash
sudo yum install docker-ce
```  

### 5、启动Docker CE
```bash
 sudo systemctl enable docker  
 sudo systemctl start docker
```  

### 6、测试docker是否正确安装  
```bash
docker run hello-world
```  

### 7、镜像加速器  
在/etc/docker目录下打开或者创建文件daemon.json  
在其内添加内容  
```json
{
  "registry-mirrors": [
    "https://registry.docker-cn.com"
  ]
}
```
其中可以使用的镜像有:  
[Docker官方提供的中国registry mirror              **https://registry.docker-cn.com**](https://docs.docker.com/registry/recipes/mirror/#use-case-the-china-registry-mirror)  
[阿里云加速器（需要登陆账号获取）](https://cr.console.aliyun.com/cn-hangzhou/new)  
[七牛云加速器 **https://reg-mirror.qiniu.com**](https://kirk-enterprise.github.io/hub-docs/#/user-guide/mirror)  

配置好配置文件后，重启docker  
```bash  
sudo systemctl daemon-reload
sudo systemctl restart docker
```  
检测镜像是否生效，用下面的命令  
```bash
docker info
```  
如果有如下内容，那么说明配置成功  
![image](B836D71EC8E642969A786B1FE045B8EE)  


## 使用（搭建一个WordPress网站）  
[阮一峰的网络日志——Docker微服务教程](http://www.ruanyifeng.com/blog/2018/02/docker-wordpress-tutorial.html)  

1、首先创建一个工作目录  
```
mkdir wordpress && cd wordpress
```

2、创建容器
```bash
docker container run \
  --rm \
  --name wordpress \
  --volume "$PWD/":/var/www/html \
  php:5.6-apache
```  
这条命令基于php 5.6的image文件新建一个容器，并且运行该容器  
```
--rm:  运行后停止，自动删除容器文件  
--name wordpress: 容器名字叫做wordpress  
--volume "$PWD/" :/var/www/html:将当前目录($PWD)映射到容器的/var/www/html(apache 对外访问的默认目录)。所以，当前目录的任何修改，都会反映到容器里面，进而被外部访问到。
```  
这样会有一个ip给出来，但是只能当前机器访问，其余机器是访问不了的，所以我们可以使用参数-p将容器中的服务端口映射出来  
```bash
docker container run   --rm -p 80:80   --name wordpress   --volume "$PWD/":/var/www/html   php:5.6-apache
```  
其中 -p 80:80的意思是将容器内的80端口映射到当前机器的80端口，那么只要能访问我这台服务器的都能访问到容器内的服务  

-p 外网端口:容器端口  

接下来启动一个后台运行的mysql容器，命名为wordpressdb
```bash
docker container run \
  -d \
  --rm \
  --name wordpressdb \
  --env MYSQL_ROOT_PASSWORD=123456 \
  --env MYSQL_DATABASE=wordpress \
  mysql:5.7
```  
然后再编写一个Dockerfile文件  
```
FROM php:5.6-apache
RUN docker-php-ext-install mysqli
CMD apache2-foreground
```  
用docker build来新建一个image文件
```
docker build -t phpwithmysql .
```
（最后有一个.，意思是依据当前目录下的Dockerfile来创建image文件）  


## 使用docker部署go服务  
有一个go项目weapons  
![image](D61905FD1A0D4474892CB75A1CDDC1D3)  
下面我们编写一个dockerfile  
这个文件的目录可以随意定  
![image](EDC106757A764B75B9E8C0233E1CF5B3)  
我们就定在这里  
![image](B0502995CA104B94A29D8D21A9FD16E1)  
dockerfile的内容如下  
其中workdir是在镜像中的路径（由于我们使用go的镜像，所以这个路径是相对/go的路径，也就是我们这个全路径是/go/code/weapons）   
参数ADD 是将本地的./weapons目录复制到镜像中的/go/code/weapons路径中  
然后RUN 是执行bash命令  
ENTRYPOINT 是程序入口

```bash
docker build -t weapons .
```
使用docker build来build出一个Image文件，image的名字是weapons，根据当前目录下的dockerfile文件  
然后再使用  
```bash
docker run -p 80:80 --rm weapons
```  
就可以运行这个image了  
这样我们就把我们的go服务部署到docker上了，然后可以把image放到docker hub上就可以共享了


