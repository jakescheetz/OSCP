# Nmap
> useful nmap commands

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
