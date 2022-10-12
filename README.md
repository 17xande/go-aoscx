# go-aoscx
Go library and CLI app for interacting with the Aruba AOS-CX API.

## CLI Usage:
```bash
go-aoscx --ip 192.168.0.1 --command vlans
# Lists all the VLANs configured on the switch with the IP specified.
```

## Requirements
- Rest interface to be enabled on the switch.
  - Run `show rest-interface` from the 'configure' context to enable rest interface.
- No authentication.
  - Switch authentication will be added in future.

## Warning!
This project is in Alpha.  
Contributions are welcome.

### References
[HPE ArubaOS-Switch REST API andJSON Schema Reference Guide 16.03Part Number: 5200-2937Published: January 2017Edition: 1](https://support.hpe.com/hpesc/public/docDisplay?docLocale=en_US&docId=emr_na-c05373669)
