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
### Storage
 - Use Google Photos or Drive?
 - Compression
 - Should we create a movie?
 - Retention / cleanup
 ### Client
 - Each client uniquely identified

## Features
- SSO with Google, Facebook
- Liveness detection
- Circumvention detection

## - Proxy support

<!--stackedit_data:
eyJoaXN0b3J5IjpbMTAwMTA4NDUyOSwxOTA3MTQwMzY3LC0xNj
A4OTcyMzMsLTEyODY4MjQ5MDNdfQ==
-->