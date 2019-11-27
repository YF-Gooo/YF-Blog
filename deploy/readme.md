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

# https自签名
apt-get install openssl
apt-get install libssl-dev
# cd /etc/nginx/
创建服务器私钥，命令会让你输入一个口令：
# openssl genrsa -des3 -out server.key 1024
创建签名请求的证书（CSR）：
# openssl req -new -key server.key -out server.csr
在加载SSL支持的Nginx并使用上述私钥时除去必须的口令：
# cp server.key server.key.org
# openssl rsa -in server.key.org -out server.key
最后标记证书使用上述私钥和CSR：
# openssl x509 -req -days 3650 -in server.csr -signkey server.key -out server.crt

# nginx
apt install nginx
scp nginx.conf root@ip:/etc/nginx/sites-available/
mv nginx.conf default
service nginx restart

