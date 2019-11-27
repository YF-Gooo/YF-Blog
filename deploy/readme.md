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
# cd /etc/nginx/sites-available/
openssl genrsa -des3 -out server.key 1024
openssl rsa -in server.key -out server_nopass.key
openssl req -new -key server.key -out server.csr
    Enter pass phrase for server.key: # 之前设置的密码

    -----

    Country Name (2 letter code) [XX]:CN # 国家

    State or Province Name (full name) []:Jilin # 地区或省份

    Locality Name (eg, city) [Default City]:Changchun # 地区局部名

    Organization Name (eg, company) [Default Company Ltd]:Python # 机构名称

    Organizational Unit Name (eg, section) []:Python # 组织单位名称

    Common Name (eg, your name or your server's hostname) []:www.yfgooo.xyz # 网站域名

    Email Address []:123@server.com # 邮箱

    A challenge password []: # 私钥保护密码,可直接回车

    An optional company name []: # 一个可选公司名称,可直接回车
openssl x509 -req -days 365 -in server.csr -signkey server.key -out server.crt
下面一步极为重要，一定要用没有密码的server.key,否则报错
mv server_nopass.key server.key 
# nginx
apt install nginx
scp nginx.conf root@ip:/etc/nginx/sites-available/
mv nginx.conf default
service nginx restart

