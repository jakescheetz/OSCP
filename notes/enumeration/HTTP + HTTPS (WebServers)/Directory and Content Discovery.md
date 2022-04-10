# HTTP/HTTPS Enumeration
> Tips and tricks for proeoperly enumerating a web server. The attack surface for this topic is most likely the largest for any protocol so it is well worth your time to slow down and make sure all of your bases are covered when a web server is identified. 


## Directory Busting - Hidden Content
##### Gobuster
- general content discovery
	- ```gobuster dir -u [ip] -w /usr/share/seclists/Discovery/Web-Content/big.txt```
- common content
	- ```gobuster dir -u [ip] -w /usr/share/seclists/Discovery/Web-Content/Top1000-RobotsDisallowed.txt```
- once the backend server language has been discovered, look for content with corresponding file extensions: 
	- ```gobuster dir -u http://[ip] -w /usr/share/wordlists/dirbuster/directory-list-lowercase-2.3-medium.txt -x php```

##### Cewl
- generate a wordlist from words on webpages
	- ```cewl -d 2 -m 5 -w customlist.txt http://[ip]```

#### Wfuzz
- useful if a CMS is indentified
	- Find the appropriate wordlist for the identified CMS: 
		- /usr/share/SecLists/Discovery/Web-Content/CMS/*
- run on top of CMS specific scanner
	- ```wfuzz -c -z file,/path/to/wordlist.txt --sc 200,301,302 http://[ip]/FUZZ```




