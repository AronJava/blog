# hugo的体验  
[toc]  

## 安装  
略过，最后生成的二进制会在go/bin目录下，和go在同一个目录  

## 建站  
首先验证hugo是否安装成功  
```bash
hugo version
```  

建站  
```
hugo new site mysite
```  
上面的命令就可以自动生成一个文件夹，这样就生成了网站  
目录：
```
archetypes: 储存.md的模板文件
content: 储存网站的所有内容
data: 储存数据文件供模板调用
layouts: 储存.html模板
static: 储存图片,css,js等静态文件，该目录下的文件会直接拷贝到/public
themes: 储存主题
```  

