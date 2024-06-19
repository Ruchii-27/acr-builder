package graph

const globalDefaultYamlLinux string = `
# Native Task aliases, use as eg. $Commit
ID: "{{.Run.ID}}"
SharedVolume: "{{.Run.SharedVolume}}"
Registry: "{{.Run.Registry}}"
RegistryName: "{{.Run.RegistryName}}"
Date: "{{.Run.Date}}"
OS: "{{.Run.OS}}"
Architecture: "{{.Run.Architecture}}"
Commit: "{{.Run.Commit}}"
Branch: "{{.Run.Branch}}"

# Default image aliases, can be used without $ directive in cmd
acr: mcr.microsoft.com/acr/acr-cli:0.11
az: mcr.microsoft.com/acr/azure-cli:7045a1f
bash: mcr.microsoft.com/acr/bash:7045a1f
curl: mcr.microsoft.com/acr/curl:7045a1f
cssc: mcr.microsoft.com/acr/cssc:7045a1f
`
