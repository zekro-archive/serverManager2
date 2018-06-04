 <div align="center">
     <h1>~ ServerManager2 ~</h1>
     <strong>faster, harder, stronger</strong><br><br>
     <a href="https://travis-ci.org/zekroTJA/serverManager2"><img src="https://travis-ci.org/zekroTJA/serverManager2.svg?branch=master"/></a>
     <a href="https://github.com/zekroTJA/serverManager2/releases"><img src="https://img.shields.io/github/release/zekroTJA/serverManager2/all.svg"/></a>
 </div>

---

# Description

With this tool, you can easily hanlde multiple servers running with [screen](https://linux.die.net/man/1/screen):  
List, start, stop, restart, resume and packup them.

---

# Installation

Download the release binary of the tool for your system [**here**](https://github.com/zekroTJA/serverManager2/releases) or compile it by yourself:


> Go installation is required for this!  
> You can download it [here](https://golang.org/dl/).
```
$ git clone https://github.com/zekroTJA/serverManager2.git
$ cd serverManager2
$ go install github.com/logrusorgru/aurora
$ go build -o servermanager src/github.com/zekrotja/serverManager2/*.go
```

---

# Commands

Command | Description | Implemented
--------|---------|-------
help | Display help message about all commands | Yes
start `<index/name>` `[e]` | Start a server by name or index<br>Use `e` as argument if you want to run server in endless* mode | Yes
stop `<index/name>` | Stop a currently running server by name or index | Yes
resume `<index/name>` | Resume a started screen session by name or index | Yes
restart `<index/name>` `[e]` | Restart a server by name or index | WIP
backup `<index/name>` | Open backup manager for server by name or index | Not yet
config | edit the config of the program | Yes
exit | Exit the tool | Yes

---

# To do

- [ ] Exclude folders as servers beginnign with `_` or `.`
- [ ] Started Before Timestamp
- [ ] Restart command
- [ ] Backup System
- [x] Static Config + Config Editor / Setup menu setting up Config on Startup
- [x] Start Command
- [x] Stop Command
- [x] Resume Command
- [x] Help Command

---

# Used 3rd-Party-Packages

- [Aurora](https://github.com/logrusorgru/aurora) by logrusorgru

---

© 2018 - present Ringo Hoffmann (zekro Development)  
contact[at]zekro.de | https://zekro.de
