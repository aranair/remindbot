Hazel - A Telegram Bot written in Golang

### What is this?

- Telegram bot
- Golang
- Docker
- Sqlite
- Nginx / Self-Signed SSL Cert
- Digital Ocean

### Walkthrough of Code

- Part 1: [https://aranair.github.io/posts/2016/12/25/how-to-set-up-golang-telegram-bot-with-webhooks/][1]
- Part 2: [https://aranair.github.io/posts/2017/01/21/how-i-deployed-golang-bot-on-digital-ocean/][2]

### See It in Action

![Commands!](https://github.com/aranair/remindbot/blob/master/commands.png?raw=true "Commands")

### Commands

- Hazel, anything: `replies with korean Hello :P`
- remind Do this and this
- remind me to Do this and this
- clear 2
- clearall
- list

### How to Run?

- Create a `configs.toml` file
- `docker-compose up`

### How to Deploy?

- Set up git hooks in production
- `git push production master`

Sample post-receive hook

```bash
#!/bin/sh

git --work-tree=/var/app/remindbot --git-dir=/var/repo/site.git checkout -f
cd /var/app/remindbot
docker-compose build
docker-compose down
docker-compose -d
```

### External Package Dependencies

- github.com/mattn/go-sqlite3
- github.com/BurntSushi/toml
- github.com/justinas/alice
- github.com/gorilla/context
- github.com/julienschmidt/httprouter

### License

MIT

[1]: https://aranair.github.io/posts/2016/12/25/how-to-set-up-golang-telegram-bot-with-webhooks/
[2]: https://aranair.github.io/posts/2017/01/21/how-i-deployed-golang-bot-on-digital-ocean/
