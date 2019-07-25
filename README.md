[![Build Status](https://drone-github.cdeyun.com/api/badges/ohko/pm/status.svg)](https://drone-github.cdeyun.com/ohko/pm)

# PM

项目管理协助工具

- 项目管理
- 任务管理
- 成员管理
- 项目任务纬度甘特图
- 成员任务纬度甘特图

## 启动服务
```shell
docker pull ohko/pm
docker rm -fv pm; \
	docker run -d --name pm --restart=always \
		-p 8082:8080 \
		-v /data/docker-pm:/db \
		-v /usr/share/zoneinfo:/usr/share/zoneinfo:ro \
		-e LOG_LEVEL=1 \
		-e TZ=Asia/Shanghai \
		-v /etc/ssl/certs:/etc/ssl/certs:ro \
		ohko/pm
```

## 修改admin密码
```shell
docker exec -it pm /pm_linux64 -resetAdmin newpassword
```