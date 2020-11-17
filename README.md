# Safebuddy

## Purpose
Create a safe environment for using a supervised PC

## What it does
- Takes a screenshot every 30-60 seconds
- Allows someone in a supervisor role to view
- Intended to allow parents to ensure safety of children's' online activities

## Requirements

 - Must be configurable remotely
	 - Local config only used to do initial connection
 - UI must be written in a web fmk

## Architecture
### Server
- Language: go
- Possibility to use Lambda?
- How would files upload to image storage
### Image Storage
 - Use Google Photos or Drive?
 - Compression
 - images only, not movies 
- Detect "locked"
-- Do I need to?
- Detect stable screen with zero changes
 ### Client
 - Each client uniquely identified with a token of sorts

## Features
- Auto Update
- SSO with Google, Facebook
- Liveness detection
- Circumvention detection
- Proxy support?
- Support Multiple screens
- Supervisor dashboard accessible from anywhere via web-based console
- 

## Retention
- Phase 0 - none
- Phase 1 - simple (60 days)
- Phase 2 - configurable

## Future
- Detect stable screen with near-zero changes



- 
<!--stackedit_data:
eyJoaXN0b3J5IjpbLTk4NTQwMDgyLDE1NjM1ODU2MjcsMTkwNz
E0MDM2NywtMTYwODk3MjMzLC0xMjg2ODI0OTAzXX0=
-->
