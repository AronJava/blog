# ssh免密登陆远程机器  
[contos7配置ssh基于秘钥对验证登陆](http://pcvc.net/blog/2015/08/10/centos-7-configuring-ssh-key-based-authentication-login/)

SSH为Secure Shell的缩写，SSH是目前较为可靠，专为远程登陆会话和其他网络服务提供安全性的协议。Linux一般自带的是OpenSSH，可用`ssh -v`查看当前版本

现在有两台机器A与B，要让A用ssh连接B的时候免密码连接，可以利用秘钥对来进行验证登陆从而实现免密码  
## 目标机器的配置
我们需要对B机器进行一些配置，用ssh登陆到远程主机，有两种认证方式，一种是基于口令的安全验证，另外一种是基于秘钥的安全验证。可用通过修改`/etc/ssh/sshd_config`这个文件对ssh服务的一些行为进行配置  

1、首先检查`~/`目录下是否有`.ssh`文件夹，如果没有，可用利用命令`ssh-keygen`生成秘钥对
```bash
ssh-keygen -t rsa
```

2、修改ssh配置文件`/etc/ssh/sshd_config`
```
PubkeyAuthentication yes  # 启用基于密匙的安全验证
PasswordAuthentication yes  # 启用基于口令的安全验证
PermitRootLogin no  # 禁止 root 登录 ssh（设置了这个就不能用root来登陆这台机器了）
```

3、获取客户机A的公钥，复制放入B机器的`~/.ssh/authorized_keys`文件中（一行代表一个公钥）如果没有该文件，那么可以自己创建

4、修改相关目录、文件的权限(假如B机器的用户是dev)
```bash
chmod 700 /home/dev
chmod 700 /home/dev/.ssh
chmod 400 /home/dev/.ssh/authorized_keys
```

5、重启SSH服务
```bash
systemctl restart sshd.service
```

## 客户机A
1、获取公钥

在`~/.ssh/id_rsa.pub`中就是公钥（全部复制给B机器的文件中）

2、利用`ssh user@ip`来连接B机器