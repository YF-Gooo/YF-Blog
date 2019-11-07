# 镜像加速
sudo vim /etc/docker/daemon.json 
{
    "registry-mirrors": ["https://1bbsr4jc.mirror.aliyuncs.com","https://registry.docker-cn.com"]
}
sudo service docker restart