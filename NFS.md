# NFS Enumeration
> tips and tricks to properly enumerate NFS

- similar to SMB, network file share service but specific to linux
- RPCBind -> port 111, used to manage port space and directs sysems to correct port (often redirected to port 2049)

#### Nmap for NFS
- nmap sweep for rpcbind: 
	- ``` nmap -vv -p111 --script=rpcinfo [ip range]```
- nmap scripts available for nfs: 
	- /usr/share/nmap/nfs*
	- ```nmap -vv -p111 --script nfs* [ip]```

### Commands
- Show mounts: 
	- ```showmount -e [ip]```
- Mount an open share: 
	- ```mount -o nolock [ip]:/open/share /path/on/your/machine```

NFSspy -> commonly used to get around mounting permission errors: 
- ```nfsspysh -o server=[ip]:/open/share,hide,allow_other,ro,initr /path/on/your/system```
- [GitHub Repo](https://github.com/bonsaiviking/NfSpy)

- interesting file on share but not correct permissions to view file? 
	- list the permissions ofthe interesting file: 
		- ```ls -l [dir of file]
	- add a new user to the system
		- ```sudo adduser pwn3d```
	- edit the permissions of the new user to the permissions of the user that is able to edit/access the intersting file on the share: 
		- ```sudo sed -i -e 's/[current uid]/[target uid]/g' /etc/passwd```
		- check that the new user has the correct uid: 
			- ```cat /etc/passwd | grep pwn3d```
	- switch to new user: 
		- ```su pwn3d```


