# General Notes- Active Directory
> General background information on Active Directory and the enumeration of it. Note, the enumeration is based on already gaining access to a machine domain-joined to the target network.

## General Information
- Domain Controller -> the server on the network with the AD Domain Services role set up for it. 
	- stores information on how AD is configured
		- stores all of the password hashes for the users in the domain
		- stores what tools and apps are available to users and how those objects are able to interact with one another
		- able to execute processes on domain joined computers
	- Domain Controller with the PdcRoleOwner property will always have the most up to date information regarding authentication statuses

- Organizational Units -> basically a folder of objects
	- used to group items together based on commonalities
	- these objects are comprised of: 
		- computer objects -> the actual servers (domain-joined)
		- user objects

- LDAP -> Lightweight Directory Access Protocol
	- used for authentication with DC's and 3rd party apps
	- LDAP query structure: 
		- ```LDAP://hostName:[port][/distinguishedName]```
	- the distinguished name of a domain is split into two parts: 
		- example.com -> ```DC=example```, ```DC=com```

- AD relies extremely heavily on DNS, it is very common to see an authoritative DNS server hosted on the same server as the domain controller to provide lookups for domain-joined devices

- Moving across machines is critical for privilege escalation since credentials are typically able to be dumped from chached memory
	- if a user account is compromised, a malicious user can grab creds for another user that could log into another machine and repeat the process until domain admin creds are dumped from memory 
		- "derivative local admin attack"

### Enumeration Checklist
- [ ] enumerate users
	- [ ] get information about the users and memberships to groups
- [ ] Do any users have an interesting title or belong to privileged groups?
- [ ] Identify if privileged users are logged on so creds can be grabbed from cached memory
- [ ] Identify if null sessions are possible







