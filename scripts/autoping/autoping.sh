#!/bin/bash

# Script to ping sweep a network for alive hosts
# Author: Jake Scheetz, for OSCP 2022

#color vars for output
RED=`tput setaf 1`
GREEN=`tput setaf 2`
YELLOW=`tput setaf 3`
WHITE=`tput setaf 7`
BG=`tput setab 7`
BLUE=`tput setaf 4`
NC=`tput sgr0` # color termination sequence
UL=`tput smul`
RMUL=`tput rmul`

#Spinner Progress item
spinner() {
    local i sp n
    sp='/-\|'
    n=${#sp}
    printf ' '
    while sleep 0.1; do
        printf "%s\b" "${sp:i++%n:1}"
    done
}

#Starting Prompt
echo "${YELLOW}+===================================================+"
echo "|${BG}${BLUE}  Tool to ping sweep a CIDR range for alive hosts  ${NC}${YELLOW}|"
echo "|${BG}${BLUE}            Written By: Jake Scheetz               ${NC}${YELLOW}|"
echo "|${BG}${BLUE}                @findingUrPasswd                   ${NC}${YELLOW}|"
echo "+===================================================+${NC}"

#Ask if a list of IP's should be created from output
read -p "${YELLOW} >> ${NC}If any hosts are identified alive during the pingsweep would you like to write them to a file? (y/n): " -n 1 -r
#Ask for file name if yes
if [[ $REPLY =~ ^[Yy]$ ]]
then
	echo ""
	read -p "Enter file name to write to: " filename
	touch $filename
fi

#Ask for first three octets of IP
echo ""
read -p "${YELLOW}>>${NC} Enter the first three octet's of the target IP (${UL}1.1.1${RMUL}.x): " ip
echo "${GREEN}[${BLUE}+${GREEN}]${WHITE} using ${UL}$ip${RMUL}.x as target IP identifier${NC}"
echo "" #line break

#Ask for CIDR range
read -p "${YELLOW}>>${NC} Enter the CIDR range to scan (/24, /16, etc.): " cidr
echo "${GREEN}[${BLUE}+${GREEN}]${WHITE} using ${UL}$ip.1$cidr${RMUL} as target network${NC}"
echo "${GREEN}[${BLUE}+${GREEN}]${WHITE} The pingsweep is underway, you'll be notified of all alive hosts - assume the rest are unreachable ${YELLOW}:)${NC}"
echo "${YELLOW}+${GREEN}=============${BLUE}SCAN${GREEN}=============${YELLOW}+${NC}"
# Logic to sort out CIDR address
case $cidr in
        '/24')
                range=255
                ;;
        '/23')
                range=511
                ;;
        '/22')
                range=1024
                ;;
        '/27')
                range=31
                ;;
        '/28')
                range=15
                ;;
        '/29')
                range=7
                ;;
        *)
                echo -n "currently unsupported CIDR range, modify the code to suit needs"
                ;;
esac


echo "${RED}**Note** Skipping $ip.0 --> Network Gateway${NC}"
echo "${GREEN}[${BLUE}+${GREEN}] ${YELLOW}Pingsweepin' away...${NC}"; spinner &
#begin logic for iterations of ping sweeping
for host in $(seq 1 $range)
do
        if ping -c1 $ip.$host &> /dev/null
        then
                echo "${GREEN}[${BLUE}!!${GREEN}]${WHITE} ${UL}$ip.$host${RMUL} is alive on the network.${NC}"
		if [[ $REPLY =~ ^[Yy]$ ]]
		then
			echo "$ip.$host" >> $filename #write to specified file if necessary
		fi
        fi
done

#kill spinner
kill "$!"

# notify end of ops
if [[ $REPLY =~ ^[Yy]$ ]]
then
        echo ""
        echo "${GREEN}[${BLUE}+${GREEN}] ${WHITE}Writing IP list of alive hosts to --> ${RED}${UL}$filename${RMUL}${NC}"
	echo "${GREEN} --> Operations complete. <-- ${NC}"
else
	echo "${GREEN} --> Operations complete. <-- ${NC}"
fi
