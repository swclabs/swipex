<h3 align="center">
    <img src="./docs/logo/hcmut.png" align=top height="50px">
    <img src="./docs/logo/logo.svg" alt="logo" height="50px" align=top>
</h3>
<p align="center">
APIs Server for Swipe written in Go 1.2+ <br>
Designed for the final thesis at HCMUT-VNU
</p>

# Swipe

Swipe (code: `swipecore`) is a api server for [Swipe](https://github.com/swclabs/swipe). `swipecore` provides functions, and services through API and microservices. Designed for server, `swipecore` provides cli commands to run api services and redis-based distributed systems cluster.

The project is designed for the final thesis at the University of Technology, Vietnam National University, Ho Chi Minh City.

## Install

Before installing, you must install make (Makefile) if you use windows operating system

- Go v1.22.5 or higher
- PostgreSQL
- Redis

Update your environment variables. see [.env.example](./.env.example)

### Start

If you want to use makefile, see Other Command Below

Build applications

```bash
make
```

Run api server

```bash
./bin/swipe s
```

Run worker server

```bash
./bin/swipe w
```
