# DNS Enumeration
> tricks and items to cover for DNS enum

#### Record types: 
- NS -> nameserver
	- contains record of the authoritative server hosting the DNS record
- A -> host 
	- hostname to IP entry
- MX -> mail exchange
	- names of servers responsible for handling email
- PTR -> pointer
	- (reverse A) IP to hostname
- CNAME -> canonical name
	- alias for other host records
- TXT -> text
	- used to store txt records like info on ownership
- AXFR -> zone transfer


### Commands
- host 
	- ```host -t [record type] www.example.com```
	- zone transfer: 
		- ```host -l [domain] [NS server for domain]```

- DNSRecon
	- zone transfer: 
		- ``` dnsrecon -d [domain] -t axfr
	- bruteforce domains: 
		- ``` densrecon -d [domain] -D [list.txt] -t brt```



