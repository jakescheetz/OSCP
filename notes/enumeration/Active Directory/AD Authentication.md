# AD Authentication Notes
> General notes on AD Authentication- this is not necessarily a way to enumerate AD but it provides background information on how auth works in these environments

## Background - Authentication Types
### NTLM
- Authentication method used when a user authenticates to an IP address instead of a hostname of if the target hostname is not listed on the DNS Server on the DC 
- not the default authentication protocol, but is widely used for its flexibility and ease of implementation
- Steps of protocol: 
	1. NTLM hash is calculated -> the users password is hashed using NTLM algo
	2. client username is sent to the server and the server generates a nonce based on the username
	3. client encrypts the nonce with the NTLM hash and sends it back to the server -> "response"
	4. "response" is forwarded to the DC
	5. DC encrypts the "response" itself and and compares it with the stored credentials

### Kerberos
- "ticket" based authentication
- DC functions as a Key Distribution Center (KDC), that evaluates requests and if its successful then the client is granted a ticket
- default authentication protocol in AD env
- password hashes are stored in the Local Security Authority Subsystem Service (LSASS)

- **TGT** -> Ticket Granting Ticket
	- The TGT includes a ton of info like pass hash, session key, IP address, etc.
	- given to user after successful authentication evaluation by the KDC (usually DC)
	- TGT is encrypted by a secret key that is only known to the KDC, making it not decryptable by the client
	- TGT's are only valid for 10 hours, and then need to revalidate- revalidation does not require password
	- Separate process is used for accessing resources with a TGT like Mail, Network shares, etc.

- Steps of protocol: 
	- **#1 Getting a TGT:**
		1. Authentication request made from the client machine (logging into a machine), containing a time stamp of when the request was made and the time stamp is encrypyed with a hash of the username:password combination of the target user
		2. Server receives the authentication request and decrypts the information- if the time stamp is not a repeat and the password hash of the target user matches then auth is successful
		3. Server replies to the client with a TGT and a Session Key. The Session Key is encrypted with the user's password hash and is easily decrypted and reused by the user for the duration of the session. 
	- **#2 Accessing Resources:**
		1. Client generates a ticket granting request that has the current user,TGT encrypted with secret key, time encrypted with the password hash and SPN of desired resource
		2. Server receives the requests and validates if the SPN exists in the domain. If it does, the TGT is decrypted with the secret key and the session key is extracted so that the time stamp and username can be decrypted. 
		3. If the username, time stamp and IP address that are decrypted are determined to be valid by the server (KDC), a valid reply is sent back to the client
		4. The reply from the server contains: SPN of granted access, a session key specifically for use between SPN and client, **a service ticket with the username, group memberships and service session key**
		5. Authentication is complete when client is able to present a session key and a service ticket 
