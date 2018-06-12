 <div align="center">
     <h1>~ ServerManager2 ~</h1>
     <strong>faster, harder, stronger</strong><br><br>
     <img src="https://forthebadge.com/images/badges/made-with-go.svg" height="25"/>&nbsp;
     <img src="https://forthebadge.com/images/badges/powered-by-jeffs-keyboard.svg" height="25" />
     <br>
     <br>
     <a href="https://travis-ci.org/zekroTJA/serverManager2"><img src="https://travis-ci.org/zekroTJA/serverManager2.svg?branch=master"/></a>&nbsp;
     <a class="badge-align" href="https://www.codacy.com/app/zekroTJA/serverManager2?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=zekroTJA/serverManager2&amp;utm_campaign=Badge_Grade"><img src="https://api.codacy.com/project/badge/Grade/a0d09c2e78f748e2ab81a236baeb5b44"/></a>&nbsp;
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

You can use the automatic script for building and installing:
```
$ wget -s https://raw.githubusercontent.com/zekroTJA/serverManager2/master/install.bash | bash
```

Or do it manually:
```
$ git clone https://github.com/zekroTJA/serverManager2.git src/github.com/zekroTJA/serverManager2
$ export GOAPTH=$PWD
$ go get github.com/logrusorgru/aurora
$ go build -o /usr/bin/servermanager src/github.com/zekroTJA/serverManager2/main.go
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

- [ ] Backup System
- [x] Exclude folders as servers beginnign with `_` or `.`
- [x] Started Before Timestamp
- [x] Restart command
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
