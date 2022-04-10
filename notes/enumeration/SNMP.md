# SNMP Enumeration
> tips and tricks to properly enumerate SNMP. Note, this protocol is commonly misunderstood and not properly enumerated- make sure to take note of weaknesses to look for that are highlighted here as they could be potential entry points to a system.

- common points of weakness: 
	- unencrypted protocol
	- weak authentication schemes
	- VERY commonly left with defaul creds
	- community strings private/public left as defaults

[IBM SNMP MIB Tree Strings Database](https://www.ibm.com/support/knowledgecenter/ssw_aix_71/commprogramming/mib.html)

#### Nmap for SNMP 
- Note SNMP is UDP based so it requires both the -sU switch and sudo
	-  ```sudo nmap -sU --open -p161 [ip range] -oG snmp-sweep.txt```

### OneSixtyOne
- use seclists:
	- ```onesixtyone -c /usr/share/seclists/Discovery/SNMP/common-snmp-community-strings.txt [ip]```

#### snmpWalk 
- get a community string's corresponding values
	- ```snmpwalk -c [string name i.e. public] -v1 -t 10 [ip]
- get windows users: 
	- ```snmpwalk -c public -v1 -t 10 [ip] 1.3.6.1.4.1.77.1.2.25 ```
- get running processes: 
	- ```snmpwalk -c public -v1 10.11.1.73 1.3.6.1.2.1.25.4.2.1.2```
- get open ports: 
	- ```snmpwalk -c public -v1 10.11.1.14 1.3.6.1.2.1.6.13.1.3```
- get installed software on the system: 
	- ```snmpwalk -c public -v1 10.11.1.50 1.3.6.1.2.1.25.6.3.1.2```

- snmpwalk + snmpenum:
```for community in public private manager; do snmpwalk -c $community -v1 $ip; done```
```snmpwalk -c public -v1 $ip```
```snmpenum $ip public windows.txt```


