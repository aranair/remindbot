Endpoint for @hn_remind_bot (A Telegram Bot)

### What is this?

- Telegram bot
- Golang
- Docker
- SQLITE

### Commands

- remind __
- remind me to __
- clear _ID_
- clearall
- list

### How to Run?

- Create a `configs.toml` file
- `docker-compose up`

### How to Deploy?

- Set up git hooks in production
- `git push production master`
- `docker-compose build`
- `docker-compose up -d`
