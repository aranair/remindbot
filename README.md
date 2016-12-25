Hazel - A Telegram Bot written in Golang

### What is this?

- Telegram bot
- Golang
- Docker
- Sqlite
- Nginx / Self-Signed SSL Cert
- Digital Ocean

### See It in Action

<iframe src="//giphy.com/embed/Uct1vWeS03WLe" width="480" height="125" frameBorder="0" class="giphy-embed" allowFullScreen></iframe><p><a href="https://giphy.com/gifs/Uct1vWeS03WLe">via GIPHY</a></p>

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
