# Install

```
ip addr add 129.254.169.254 dev lo
cp .env.example .env
```
Edit .env file to use ```LOCATION```, ``` RG_NAME``` & ```ID``` as the ones you want to map your creds to.
Run binary as service or manually.

# TODO
- Add service file for systemd service.
- Add pipeline to build binaries

# Build locally
```
go build -o bin/aadlogin-metadata main.go instanceCompute.go 
```