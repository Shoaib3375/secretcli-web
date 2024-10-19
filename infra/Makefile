.PHONY: init-dev plan-dev apply-dev destroy-dev init-prod plan-prod apply-prod destroy-prod debug-state

# Development Environment
init-dev:
	cd environments/dev && terraform init -backend-config=../../dev.backend.hcl

plan-dev:
	cd environments/dev && terraform plan -out=dev.tfplan

apply-dev:
	cd environments/dev && terraform apply dev.tfplan

destroy-dev:
	cd environments/dev && terraform destroy -auto-approve

# Production Environment
init-prod:
	cd environments/prod && terraform init -backend-config=../../prod.backend.hcl

plan-prod:
	cd environments/prod && terraform plan -out=prod.tfplan

apply-prod:
	cd environments/prod && terraform apply prod.tfplan

destroy-prod:
	cd environments/prod && terraform destroy -auto-approve

# Debug command
debug-state:
	@echo "State files in root state directory:"
	@ls -l state/