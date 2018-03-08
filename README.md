
Beego使用方法:   2018-01-01

----------------------------------------------
Beego环境搭建和bee工具安装使用
1.
创建D:\rd\beego目录，并建3个空的子目录 bin,pkg,src
在环境变量中设置系统变量 GOROOT=E:\IDEA\go1.9.2.windows-amd64\go   #go的压缩包解压后
在环境变量中设置系统变量 GOPATH=D:\rd\beego
在环境变量中设置系统变量 PATH=E:\IDEA\TDM_GCC_64\bin;C:\Program Files\Git\cmd;D:\rd\mysql-5.7.19;%GOROOT%\bin;%GOPATH%\bin;%GOPATH%;
在PATH中加上 %GOROOT%\bin;%GOPATH%\bin;%GOPATH%;

2. 下载并安装好git bash工具。
3. 打开gitcmd，输入 go get github.com/astaxie/beego  。稍等片刻即可在GOPATH的src目录下看到有\github.com\astaxie\beego目录。
4. 在gitcmd中输入 go get github.com/beego/bee  
5. 测试bee是否安装成功，可在命令行中输入bee，得到如下结果就成功了。

-------------------
beego admin请看admin目录下的Readme.md文件操作

如果是go项目:
在 GOPATH/src目录下， go new project1; go run project1 就跑一个项目了。

如果是beego项目，则:
在 GOPATH/src/demo_beego目录下， bee new project1; bee run project1 就跑一个项目了。

注意，在windows跑这些命令时，使用cmd.exe或使用git的Git CMD.ext也是OK的。

----------------------------------
有两个例子:
   \src\demo_beego\admin_copy  #这是从仿src/github.com/beego/admin的例子(copy from github.com/beego/admin)
   \src\demo_beego\hello_admin  #这是从github.com/beego/admin的例子