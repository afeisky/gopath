
xcopy /e/y ..\..\github.com\beego\admin\views views
#xcopy /e/y ..\..\github.com\beego\admin\static static
go build
hello_admin  -syncdb
hello_admin run

