# ======================================================================
# Entry Points

build: bin/hyparview-example
.PHONEY: build

root: cert/$(DOMAIN)-agent-ca.pem
.PHONEY: root

ssh-keys: terraform/demo-key.aws
.PHONEY: ssh-keys

# ======================================================================
# Implementation

# build
sources = $(wildcard *.go)
protobufs = proto/hyparview.pb.go proto/gossip.pb.go proto/stat.pb.go
bin/hyparview-example: $(sources) $(protobufs)
	go build -o $@
	chmod +x $@

proto/%.pb.go: proto/%.proto
	protoc -I . --go_out=plugins=grpc:. $<

# certs & keys
cert/$(DOMAIN)-agent-ca.pem:
	mkdir -p cert
	cd cert && consul tls ca create -domain=$(DOMAIN)

cert/$(DOMAIN)-server-%.pem: cert/$(DOMAIN)-agent-ca.pem
	cd cert && consul tls cert create -server -domain $(DOMAIN)
	mv cert/dc1-server-$(DOMAIN)-0.pem $@
	mv cert/dc1-server-$(DOMAIN)-0-key.pem $(basename $@)-key.pem

cert/$(DOMAIN)-client-%.pem: cert/$(DOMAIN)-agent-ca.pem
	cd cert && consul tls cert create -client -domain $(DOMAIN)
	mv cert/dc1-client-$(DOMAIN)-0.pem $@
	mv cert/dc1-client-$(DOMAIN)-0-key.pem $(basename $@)-key.pem

# ======================================================================
# Demo

terraform/hosts: terraform/apply
	(cd terraform; terraform show -json) \
	| jq -M '.values.root_module.resources[].values.public_ip' \
	| grep -v null \
	> $@

terraform/apply: terraform/demo-key.pub
	cd terraform; terraform apply
	touch $@

terraform/demo-key.pub:
	ssh-keygen -t rsa -f $(basename $@)
	chmod 600 demo-key*
