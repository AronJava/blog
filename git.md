# 搭建git项目

  自己在GitHub上面搭建一个仓库，并且在本地建立仓库连接到远程仓库，能提交代码
  
## 搭建过程
1、在github上面创建一个仓库

2、在本地创建一个目录

3、使用命令`git init`初始化一个本地git

4、使用`git remote add origin`来添加远程仓库

5、使用`git config`设置全局变量
```
git config --global user.name "Aron"
git config --global user.email "1217924725@qq.com"
```

5、然后就可以提交代码了，使用`git push 远程主机名 分支名`来push代码
```
git push origin master
```

6、每次都要输入密码很麻烦，使用`git config`来记住密码
```
git config --global credential.helper store
```
上面的命令运行完之后，只需要输入一次密码之后就不需要再输入密码了