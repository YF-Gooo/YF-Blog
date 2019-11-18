# 安装docker以及docker-compose


# 镜像加速
sudo vim /etc/docker/daemon.json 
{
    "registry-mirrors": ["https://1bbsr4jc.mirror.aliyuncs.com","https://registry.docker-cn.com"]
}
sudo service docker restart

# 一键启动redis和mysql
    docker-compose up -d


