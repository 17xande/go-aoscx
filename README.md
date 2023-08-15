# go-aoscx
Go library and CLI app for interacting with the Aruba AOS-S API.  
AOS-CX API will hopefully be added in the future.
Currently, up to v8 of the API is supported.

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

## References
**API v2-v8 FULL** [16.09 REST API and JSON Schema](https://h10145.www1.hpe.com/downloads/DownloadSoftware.aspx?SoftwareReleaseUId=26402&ProductNumber=J9625A&lang=&cc=&prodSeriesId=&SaidNumber=)  
**API v1:** [HPE ArubaOS-Switch REST API andJSON Schema Reference Guide 16.03Part Number: 5200-2937Published: January 2017Edition: 1](https://support.hpe.com/hpesc/public/docDisplay?docLocale=en_US&docId=emr_na-c05373669)  
**API v2-v8** [REST API for AOS-S 16.11](https://www.arubanetworks.com/techdocs/AOS-Switch/16.11/Aruba%20REST%20API%20for%20AOS-S%2016.11.pdf)  
**API v10.04** [ArubaOS-CX 10.05 REST v10.04 API Guide](https://www.arubanetworks.com/techdocs/AOS-CX/10.05/HTML/5200-7320/index.html)  

## TODO:
- [ ] Add other v1 functions
- [ ] Allow option for HTTPS (and default to HTTPS)
- [ ] Add authentication/session support
- [ ] Add API v10.04 support (for AOS-CX)
  - [ ] Support both API versions
