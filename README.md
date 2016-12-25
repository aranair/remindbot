Hazel - A Telegram Bot written in Golang

### What is this?

- Telegram bot
- Golang
- Docker
- Sqlite
- Nginx / Self-Signed SSL Cert
- Digital Ocean

### Commands

- Hazel, anything: `replies with korean Hello :P`
- remind Do this and this
- remind me to Do this and this
- clear 2
- clearall
- list

![Commands!](https://github.com/aranair/remindbot/blob/master/commands.png?raw=true "Commands")

### How to Run?

- Create a `configs.toml` file
- `docker-compose up`

### How to Deploy?

- Set up git hooks in production

```bash
#!/bin/sh

# /var/repo/site.git/hooks/post-receive
git --work-tree=/var/app/remindbot --git-dir=/var/repo/site.git checkout -f
cd /var/app/remindbot
docker-compose build
docker-compose down
docker-compose -d
```

- `git push production master`
