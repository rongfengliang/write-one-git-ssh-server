# wriet one simple git ssh server

> use golang write one simple git ssh server

## how to running

* starting service

```code
docker-compose up -d
```

* add authorized_keys

> inside config/.ssh/authorized_keys

```code
command="/opt/git-shell",no-port-forwarding,no-X11-forwarding,no-agent-forwarding,no-pty <ssh public key>
```

* create git repo (--bare)

> inside docker container /opt/gitrepo

```code
cd /opt/gitrepo
git init --bare demoapp.git
```


* create git repo (local)

```code
git clone  ssh://git@localhost:2222/demoapp.git
```

* do push 

```code
git init 
echo "test" > demo.txt
git add --all
git commit -m "demo"
git push -u origin master
```