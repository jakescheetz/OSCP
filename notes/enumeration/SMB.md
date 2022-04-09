# SMB Enumeration
> tips and tricks to get good coverage and identify weaknesses in SMB

>🚧
>Note: SMB's track record with security is extremely poor. Take the time to ensure this is properly enumerated- you'll thank yourself later. The course materially is also SEVERALLY lacking on SMB. 

- NetBIOS? 
	- listens on 139 and is technically different than SMB but it is needed for backwards compatibility issues so it is very often coupled with SMB

- Enum4Linux (Updated)
	- provides an overall enumeration of SMB
	- be careful to only use this on hosts you know have SMB open because it can cause weird errors
	- [GitHub Repo](https://github.com/cddmp/enum4linux-ng)

- nmap SMB/NetBIOS sweep
	- ``` nmap -sn -T4 -vv -p139,445 -oG smbsweep.txt | grep open | cut -d" " -f2 ```

- Get the hostname of target with NetBIOS/SMB 
	- ```nmblookup -A [ip]```

#### Scanning NetBIOS
- nbtscan
	- ```sudo nbtscan -r [cidr range]```

#### Nmap for SMB 
- get all of the smb scripts nmap can run: 
	- ```ls -1 /usr/share/nmap/scripts/smb*```
- get shares on a host with SMB open: 
	- ```nmap --script smb-enum-shares -p139,445 --open [ip]```
- check for well known SMB vulns: 
	- ```nmap --script smb-vuln* -p139,445 [ip]``` 

#### SMBmap
- extremely simplified syntax- HIGHLY RECOMMEND, don't make things more complicated than they'll already be :)
- list shares + show your access to the shares, check null sessions
	- ``` smbmap -H [ip or hostname]```
- if you have creds, list all shares
	- ```smbmap -H [ip] -d [domain] -u [user] -p [password]



