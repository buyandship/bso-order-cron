.PHONY: build_all
build_all: build docker

.PHONY: build
build:
	@echo "Compiling..."
	@GOOS=linux GOARCH=amd64 go build -o output/bso-order-cron ./main.go

.PHONY: docker
docker:
	@echo "Building docker image..."
	@docker build -f Dockerfile -t ghcr.io/bravo-lu/bso-order-cron:latest .
	@echo "Pushing docker image..."
	@docker push ghcr.io/bravo-lu/bso-order-cron:latest

.PHONY: t1
t1:
	@echo "Updating service on t1"
	@ssh -i ~/.ssh/bns_mkp_dev.pem ec2-user@mkp-ssh1.buynship.com 'bash /home/ec2-user/build_bso_order_cron.sh'


.PHONY: t2
t2:
	@echo "Updating service on t2"
	@ssh -i ~/.ssh/bns_mkp_dev.pem ec2-user@mkp-ssh2.buynship.com 'bash /home/ec2-user/build_bso_order_cron.sh'
