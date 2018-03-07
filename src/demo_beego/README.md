这些是beego创造的

----------------------------------------------------
cd \beego\src\demo_beego\  & bee new wwww1   #一个普通的beego例子
或在src/目录下 bee new demo_beego\www1 建一个项目
----------------------------------------------------
cd \beego\src\demo_beego\  
bee new hello_admin  #这个使用了admin 库
cd \beego\src\demo_beego\hello_admin 
bee run demo_beego\www1 
或 bee run  #要进这个目录里才能运行!
其它按 \beego\src\github.com\beego\admin\README.mk文档操作
D:\rd\beego\src\demo_beego\hello_admin>go build  & hello_admin  -syncdb & hello_admin run
ProductCompany....TableName
ProductUser....TableName
2018/03/06 14:24:53 Database  admin  created
insert user ...
insert user end
insert group ...
insert group end
insert role ...
insert role end
insert node ...
insert node end
database init is complete.
Please restart the application
ProductCompany....TableName
ProductUser....TableName
Starting....
Start ok
2018/03/06 14:25:02.856 [I] [asm_amd64.s:2337] http server Running on http://:82
----------------------------------------------------