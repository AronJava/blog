# hexo体验  
[toc]  

## 安装  
[搭建hexo个人博客](http://code.skyheng.com/post/4103.html)  

### 在centos下搭建hexo  
1、安装git（忽略）  
使用  
```bash
git --version
```  
使用上面的命令来确认是否安装了git  

2、安装Node.js  
[Node.js的安装、卸载与升级](https://blog.imzhengfei.com/node-js-an-zhuang-xie-zai-yu-sheng-ji/)  

下载node@6 LTS:  
```bash
curl --silent --location https://rpm.nodesource.com/setup_6.x | bash -
```  

使用yum安装node.js  
```bash
yum -y install nodejs
```  

安装编译工具  
```bash
yum -y install gcc-c++ make
```  
检测node.js是否安装成功  
```bash
node -v
```  
```bash
npm -v
```  
3、设置node.js的源  
```bash
alias cnpm="npm --registry=https://registry.npm.taobao.org \
--cache=$HOME/.npm/.cache/cnpm \
--disturl=https://npm.taobao.org/dist \
--userconfig=$HOME/.cnpmrc"
```  
检验源是否更换成功
```bash
cnpm info express
```  
安装hexo  
```bash
cnpm install -g hexo-cli
```  

### hexo项目搭建  
创建一个工作目录，我创建为/data/blog  
在/data/blog下  
创建一个hexo项目  
```bash
hexo init
```  
启动服务  
```bash
hexo server
```  
启动服务后就可以根据给的url来访问Blog了  

### hexo的一些命令  
加载静态资源（启动服务前加载一次）
```bash
hexo generate
```  

修改Hexo的监听端口  
在/node_modules/hexo-server目录下的index.js文件中修改  


### hexo的一些优秀文章  

[next主题的个性化打造](https://zhuanlan.zhihu.com/p/28128674)  

[博客优化——next主题](http://wangwlj.com/2017/09/09/blog-opti/)  

## 个性化设置博客  
1、设置网站图标  
在**主题的配置文件**下，找到配置字段favicon  
![image](5BCA463D5F6C458A8FE365B83D22FF42)  
配置好图片即可  

2、设置动态标题  
在主题的source/js里面新建一个js文件  
**dytitle.js**  
```javascript  
<!--动态标题-->
var OriginTitile = document.title;
var titleTime;
document.addEventListener('visibilitychange', function () {
    if (document.hidden) {
        document.title = ' 你不理我了！';
        clearTimeout(titleTime);
    }
    else {
        document.title = ' 么么哒 ' + OriginTitile;
        titleTime = setTimeout(function () {
            document.title = OriginTitile;
        }, 2000);
    }
});
```  
然后再打主题的/layout/_layout.swig文件里面  
在</body>之前添加  
```html
<!--动态标题-->
<script type="text/javascript" src="/js/src/dytitle.js"></script>
```  

3、背景图片  
在主题的/source/css/_custom文件下的custom.styl文件中添加  
```html
body{
    background:url(/img/background.jpg);
    background-attachment: fixed;
}
```