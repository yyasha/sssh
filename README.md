# SSSH
Use ssh to communicate with other people.
***
## Quick start
![Made with VHS](https://vhs.charm.sh/vhs-21IpNFfBo4FEAS3UA2zy8D.gif)
**Installation**
```
git clone https://github.com/yyasha/sssh.git
cd sssh
go build
./sssh --help
```
**Usage**
``` console
Usage:
  sssh [OPTIONS]

Application Options:
  -b, --bind=          Host and port to listen on. (default: 0.0.0.0:22)
  -i, --identity=      Private key to identify server with. (default: ~/.ssh/id_rsa)
  -p, --password_mode  Enable mandatory password mode
      --whitelist=     Optional file of public keys who are allowed to connect.

Help Options:
  -h, --help           Show this help message
```
## Client connection
![Made with VHS](https://vhs.charm.sh/vhs-2J5QJXuXXusRe4XfsZFSy1.gif)
Use default ssh command to connect to server:
``` console
ssh -p 22 -i private_key username@127.0.0.1
```
**Server commands list:**
Type `/help` on the server for help message:
``` console
server commands:
/help - show this message
/exit - disconnect from the server
/new_password - new password
```
# License
MIT
