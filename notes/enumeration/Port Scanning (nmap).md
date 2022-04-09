# Nmap
> useful nmap commands

- for exam use -> [AutoRecon](https://github.com/Tib3rius/AutoRecon)
	- I highly discourage the use of this tool for course exercises and the lab env, unless you're EXTREMELY pressed for time
	- it will fully automate the nmap process and even find exploits for you but you will not learn hardly anything about the enumeration process... which is a foundational skill to why you're studying for the OSCP in the first place. 
	- after getting a list of hosts from pingsweep: 
		- ```autorecon -t alivehosts.txt``` (run in a new directory so results are separated out, this gets messy really quick)

- all nmap scripts are stored in /usr/share/nmap/scripts
	- use ```--script-help [script name] ``` to see usage info
#### Commands
- pingsweep
	- ``` nmap -sn {target}
	- if runnning on a cidr network, grep for al of the hosts that are up and send to a file 
	- ``` nmap -sn {target} -oG alivehosts.txt | grep "Status: Up" | awk '{print $2}'
	- sweep for all hosts running a web service?? 
		-  ``` nmap -sn -p80,443 --open {target} -oG | grep open | cut -d" " -f2
		- this can be done for any service techinically

- initial scan to get some idea of open ports
	- like to run this one quickly to start manually looking through things while I run the longer scans in the backgorund
	- ```nmap -sT -A -T4 --open -vv --top-ports=50 -Pn {target} -O
