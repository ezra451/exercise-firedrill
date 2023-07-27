all: ransomware discovery uac_bypass

ransomware:
	go build ./cmd/ransomware

discovery:
	go build github.com/FourCoreLabs/firedrill/cmd/discovery

uac_bypass:
	go build github.com/FourCoreLabs/firedrill/cmd/uac_bypass

registry_run:
	go build github.com/FourCoreLabs/firedrill/cmd/runkeyregistry
	
gorelease:
	goreleaser release --rm-dist --snapshot 
