# General Notes- Active Directory
> General background information on Active Directory and the enumeration of it. Note, the enumeration is based on already gaining access to a machine domain-joined to the target network.

## General Information
- Domain Controller -> the server on the network with the AD Domain Services role set up for it. 
	- stores information on how AD is configured
	- stores all of the password hashes for the users in the domain
	- stores what tools and apps are available to users and how those objects are able to interact with one another
	- able to execute processes on domain joined computers

- Organizational Units -> basically a folder of objects
	- used to group items together based on commonalities
	- these objects are comprised of: 
		- computer objects -> the actual servers (domain-joined)
		- user objects

- AD relies extremely heavily on DNS, it is very common to see an authoritative DNS server hosted on the same server as the domain controller to provide lookups for domain-joined devices


### Checklist
- [ ] enumerate users
	- [ ] get information about the users and memberships to groups



