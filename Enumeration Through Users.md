# Enumerating AD via Users 
- Need to understand what users are currently authenticated so that their credentials can be dumped from memory

### PowerShell Empire
- [GitHub Repo](https://github.com/EmpireProject/Empire)

- PowerView can be used to view the currently logged in users, but only a handful of functions are needed
	- import the module: 
		- ```PS C:\> Import-Module .\PowerView.ps1 ```
		- Get-Netloggedon -> view logged in users
			- ```PS C:\> Get-NetLoggedon -ComputerName [computer name]```
		- Get-NetSession -> get active sessions on the domain controller
			- ```PS C:\> Get-NetSession -ComputerName [DC name]```
-----
### Service Accounts
- another way to gain access is to identify a service account 
