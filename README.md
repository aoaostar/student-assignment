<h1 align="center">学生作业管理系统</h1>
<p align="center">
  <a href="https://github.com/aoaostar/student-assignment/releases"><img src="https://img.shields.io/github/release/aoaostar/student-assignment?style=flat-square" alt="Release version"></a>
  <a href="https://github.com/aoaostar/student-assignment/actions?query=workflow%3ABuild"><img src="https://img.shields.io/github/workflow/status/aoaostar/student-assignment/build?style=flat-square" alt="Build status"></a>
  <a href="https://github.com/aoaostar/student-assignment/releases"><img src="https://img.shields.io/github/downloads/aoaostar/student-assignment/total?style=flat-square" alt="Downloads"></a>
  <a href="https://github.com/aoaostar/student-assignment/blob/master/LICENSE"><img src="https://img.shields.io/github/license/aoaostar/student-assignment?style=flat-square" alt="License"></a>
</p>

### What's this？
这是一款`学生作业管理系统`，后端基于`go`。

### 演示地址

* <https://job.v8net.online>

### 演示图


### 部署

#### 使用说明
    + `releases`下载对应的包
    + 修改`config.yaml`内的数据库配置信息
    + 然后运行即可
```yaml
app_name: 学生作业提交系统 # 应用名称
debug: true # 是否启用debug
listen: 0.0.0.0:8080 # 监听地址
static: dist # 静态文件目录
upload_size: 30
database:
  host: 127.0.0.1 # 数据库地址
  port: 3306 # 数据库端口
  name: StudentJobs # 数据库库名
  table_prefix: aoaostar_
  username: root # 数据库用户名
  password: root # 数据库密码
```
```shell
chmod +x StudentAssignment
nohup ./StudentAssignment > log.log 2>&1 &
```

* 进程守护
    + vi /usr/lib/systemd/system/StudentAssignment.service
    + 添加以下内容，其中`app_path`为`StudentAssignment`所在的路径

```shell
[Unit]
Description=StudentAssignment
After=network.target
 
[Service]
Type=simple
WorkingDirectory=/www/wwwroot/job.v8net.online
ExecStart=/www/wwwroot/job.v8net.online/StudentAssignment
Restart=on-failure
 
[Install]
WantedBy=multi-user.target
```
/www/wwwroot/job.v8net.online

* 然后执行`systemctl daemon-reload`重载配置

```
启动: systemctl start StudentAssignment
关闭: systemctl stop StudentAssignment
自启: systemctl enable StudentAssignment
状态: systemctl status StudentAssignment
```

### 反向代理

* 由于项目运行在`8080`端口，所以需要`nginx`反向代理
```
location / {
    proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    proxy_set_header Host $http_host;
    proxy_set_header X-Real-IP $remote_addr;
    proxy_redirect off;
    proxy_pass http://127.0.0.1:8080;
}
```

* 代理静态资源需要注释以下内容，否则静态文件无法加载

```
location ~ .*\.(gif|jpg|jpeg|png|bmp|swf)$
{
    expires      30d;
    error_log /dev/null;
    access_log /dev/null;
}
location ~ .*\.(js|css)?$
{
    expires      12h;
    error_log /dev/null;
    access_log /dev/null;
}
```


### 交叉编译

* <https://golang.google.cn/doc/install/source#introduction>
* <https://golang.google.cn/doc/install/source#environment>