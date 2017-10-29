jazel
------

### What is this?

- A Telegram Bot written in Golang - Parses messages and records reminders.
- A Cron-like Bot - Checks for overdue reminders periodically.
- Docker / Docker-Compose
- Sqlite3
- Nginx / Self-Signed SSL Cert
- Digital Ocean
- Github hooks for deployment

### Walkthrough of Code

- Part 1: [https://aranair.github.io/posts/2016/12/25/how-to-set-up-golang-telegram-bot-with-webhooks/][1]
- Part 2: [https://aranair.github.io/posts/2017/01/21/how-i-deployed-golang-bot-on-digital-ocean/][2]
- Part 3: [https://aranair.github.io/posts/2017/08/20/golang-telegram-bot-migrations-cronjobs-and-refactors/][3]

### See It in Action

![Commands!](https://github.com/aranair/remindbot/blob/master/commands.png?raw=true "Commands")

### Commands

- Hazel, anything: `replies with korean Hello :P`
- remind buy dinner
- remind me to do this and this
- remind me to sleep :9jul 10pm
- remind me to buy chocolate :today 10pm
- remind me to buy a gift :tomorrow 10pm
- clear 2
- clearall
- list

### Resetting Item Number for Resetting

- Unfortunately will deprecate support for multi-tenant bots but I've enabled this for personal use.

### How to Run?

- Create a `configs.toml` file
- `docker-compose up` or `go run cmd/webapp/main.go`

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

- No Telegram-"api" packages were used, just some regex and http.
- github.com/mattn/go-sqlite3
- github.com/BurntSushi/toml
- github.com/justinas/alice
- github.com/gorilla/context
- github.com/julienschmidt/httprouter

### License

MIT

[1]: https://aranair.github.io/posts/2016/12/25/how-to-set-up-golang-telegram-bot-with-webhooks/
[2]: https://aranair.github.io/posts/2017/01/21/how-i-deployed-golang-bot-on-digital-ocean/
[3]: https://aranair.github.io/posts/2017/08/20/golang-telegram-bot-migrations-cronjobs-and-refactors/
