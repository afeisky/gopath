{{template "../public/header.tpl"}}
<script type="text/javascript">
    var userid = {{.userid}};
    var URL="/app1/user"
$(function(){
    //组用户列表
    $("#datagrid2").datagrid({
        url:URL+'/UserSelectCompanyList?Id='+userid,
        method:'get',
        fitColumns:false,
        striped:true,
        rownumbers:true,
        singleSelect: true,
        idField:'Id',
        columns:[[
            {field:'Id',title:'ID',width:50,align:'center'},
            {field:'Name',title:'公司名字',width:140,align:'center'},
            {field:'Longname',title:'全称',width:140,align:'center'}
        ]],
        onLoadSuccess:function(data){
            $("#datagrid2").datagrid('unselectAll');
            //默认选中已存在的对应关系
            for(var i=0;i<data.rows.length;i++){
                if(data.rows[i].checked == 1){
                    $(this).datagrid('selectRecord',data.rows[i].Id);
                }
            }
        }
    });
});
    //保存选择
    function saveselect(){
        var rows = $("#datagrid2").datagrid('getSelections');
        if(rows == null){
            vac.alert("最少要选中一行");
        }
        var ids = [];
        for(var i=0; i<rows.length; i++){
            ids.push(rows[i].Id);
        }
        vac.ajax(URL+'/UserUpdateCompany', {Id:userid,ids:ids.join(',')}, 'POST', function(r){
            $.messager.alert('提示',r.info,'info');
        })
    }
</script>
<body>
<table id="datagrid2" toolbar="#tb2"></table>
<div id="tb2" style="padding:5px;height:auto">
    <div style="margin-bottom:5px">
        <a href="#"  class="easyui-linkbutton" iconCls="icon-save" plain="true" onclick="saveselect()">保存</a>
    </div>
</div>
</body>
</html>