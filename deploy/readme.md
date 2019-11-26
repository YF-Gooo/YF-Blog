# step 1: 安装必要的一些系统工具
apt-get update
apt-get -y install apt-transport-https ca-certificates curl software-properties-common
# step 2: 安装GPG证书
curl -fsSL http://mirrors.aliyun.com/docker-ce/linux/ubuntu/gpg | sudo apt-key add -
# Step 3: 写入软件源信息
add-apt-repository "deb [arch=amd64] http://mirrors.aliyun.com/docker-ce/linux/ubuntu $(lsb_release -cs) stable"
# Step 4: 更新并安装Docker-CE
apt-get -y update
apt-get -y install docker-ce
# 配置镜像加速器
mkdir -p /etc/docker
tee /etc/docker/daemon.json <<-'EOF'
{
  "registry-mirrors": ["https://bi49zw24.mirror.aliyuncs.com"]
}
EOF
systemctl daemon-reload
systemctl restart docker

# portainer
docker volume create portainer_data
docker run -d -p 9000:9000 -p 8000:8000 --name portainer --restart always -v /var/run/docker.sock:/var/run/docker.sock -v portainer_data:/data portainer/portainer
Registries->添加镜像仓库
stacks->创建docker-compose

# nginx
apt install nginx
scp nginx.conf root@ip:/etc/nginx/sites-available/
mv nginx.conf default
service nginx restart
