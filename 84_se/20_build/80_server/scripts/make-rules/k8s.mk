
# ==============================================================================
# Options

KUBECTL := kubectl
NAMESPACE ?= default
DEPLOYS=apiserver


# ==============================================================================
# Targets

.PHONY: k8s.deploy.all
k8s.deploy.all:
	@echo "===========> K8s Deploying all"
	@$(MAKE) k8s.deploy

.PHONY: k8s.deploy
k8s.deploy: $(addprefix k8s.deploy., $(DEPLOYS))

.PHONY: k8s.deploy.%
k8s.deploy.%:
	@echo "===========> K8s Deploying $* $(VERSION)-$(ARCH)"
	$(eval ARCH := $(word 2,$(subst _, ,$(PLATFORM))))
	$(shell ./scripts/conf2yaml.sh)
	$(KUBECTL) -n $(NAMESPACE) apply -f manifests/config.yaml
	$(KUBECTL) -n $(NAMESPACE) apply -f manifests/cert.yaml
	# echo @$(KUBECTL) -n $(NAMESPACE) set image deployment/$* $*=$(REGISTRY_PREFIX)/$*-$(ARCH):$(VERSION)
	$(KUBECTL) -n $(NAMESPACE) apply -f manifests/$*.yaml

.PHONY: k8s.undeploy
k8s.undeploy: $(addprefix k8s.undeploy., $(DEPLOYS))

.PHONY: k8s.undeploy.%
k8s.undeploy.%:
	$(eval ARCH := $(word 2,$(subst _, ,$(PLATFORM))))
	@echo "===========> K8s Undeploying $* $(VERSION)-$(ARCH)"
	@$(KUBECTL) -n $(NAMESPACE) delete -f manifests/$*.yaml
	$(KUBECTL) -n $(NAMESPACE) delete -f manifests/cert.yaml
	$(KUBECTL) -n $(NAMESPACE) delete -f manifests/config.yaml


