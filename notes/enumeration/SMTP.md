# SMTP Enumeration
> tips and tricks for properly enumerating SMTP

- SMTP servers almost always support EXPN and VRFY: 
	- EXPN -> ask for membership of a mailing list
	- VRFY -> ask to verify an email address
		- will return a name if true

```
kali@kali:~$ nc -nv [IP] 25
(UNKNOWN) [IP] 25 (smtp) open
220 hotline.localdomain ESMTP Postfix
VRFY root
252 2.0.0 root
VRFY idontexist
550 5.1.1 <idontexist>: Recipient address rejected: User unknown in local recipient table
```

- see smtp-vrfy.py script in the #scripts folder

