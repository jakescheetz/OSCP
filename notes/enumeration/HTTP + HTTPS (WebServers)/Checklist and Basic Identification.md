# HTTP/HTTPS Enumeration
> Tips and tricks for proeoperly enumerating a web server. The attack surface for this topic is most likely the largest for any protocol so it is well worth your time to slow down and make sure all of your bases are covered when a web server is identified. 

### General Web Sever Information Gathering
- intial web service identification: 
	- ```nmap -sT -vv -p80,443,8080,8081,8082 --open --script-http-enum [ip range] -oG webserver-sweep.txt``` 
	- smart to rerun a full port scan on the IP's found to have one of the above ports open as its not uncommon for web management interfaces to be running on a wacky port for 'obscurity'
		- scan http-server vulnerabilties: 
			- ```grep "Status: Up" webserver-sweep.txt | awk  '{print $2}' > http-vuln-scan-list.txt```
			- ```nmap --script=http-vuln* -vv -iL http-vuln-scan-list.txt -oN http-vulns.nmap```

- run in a new terminal tab and then let it go (this one takes a long time)
	- ```nikto -h [ip]```

- get a general idea of backend languages and services: 
	- ```whatweb -u [ip]```

- pick up on potential areas of LFI/RFI: 
	- ```uniscan -u http://[ip]/ -qweds```

### Checklist
- [ ] robots.txt or sitemap.xml exist?
- [ ] whatweb to indentify backend languages if not glaringly obvious
- [ ] if on 443, SSL/TLS cert disclose any info about users? 
- [ ] rerun nmap full port scan to try to find management interface
- [ ] run initial dir bust / content discovery scans
- [ ] check source on pages and look at href's for links to internal pages and find hidden fields or comments in the source
	- [ ] spidering with burp works too
- [ ] look for CMS that the website is built on
	- [ ] if not common CMS google to see if vulnerable
- [ ] Injection points? SQLI, RFI/LFI, XXE, file upload? 
- [ ]  scan for heartbleed if 443 enabled
- [ ] login/admin pages? 


