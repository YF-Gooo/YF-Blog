from fabric import Connection,task
from invoke import run

@task
def deploy(c):
    with Connection('root@47.240.45.192') as c:
        c.run("mkdir yfblog")
        c.run("cd yfblog && docker-compose rm -fsv",pty=True)
        c.run("rm -rf docker-compose.yml")
        c.put("docker-compose.yml","docker-compose.yml")
        c.run("docker-compose up -d")

@task
def init(c)
    run("echo hello")