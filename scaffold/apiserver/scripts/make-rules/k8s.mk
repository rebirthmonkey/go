
# ==============================================================================
# Options

KUBECTL := kubectl
NAMESPACE ?= default

.PHONY: k8s.deploy
k8s.deploy: $(addprefix k8s.deploy., $(BINS))

.PHONY: k8s.deploy.%
k8s.deploy.%:
	$(eval ARCH := $(word 2,$(subst _, ,$(PLATFORM))))
	@echo "===========> K8s Deploying $*-$(ARCH):$(VERSION)"
	$(shell ./scripts/conf2yaml.sh)
	$(KUBECTL) -n $(NAMESPACE) apply -f manifests/config.yaml
	$(KUBECTL) -n $(NAMESPACE) apply -f manifests/cert.yaml
	# echo @$(KUBECTL) -n $(NAMESPACE) set image deployment/$* $*=$(REGISTRY_PREFIX)/$*-$(ARCH):$(VERSION)
	$(KUBECTL) -n $(NAMESPACE) apply -f manifests/$*.yaml

.PHONY: k8s.undeploy
k8s.undeploy: $(addprefix k8s.undeploy., $(BINS))

.PHONY: k8s.undeploy.%
k8s.undeploy.%:
	$(eval ARCH := $(word 2,$(subst _, ,$(PLATFORM))))
	@echo "===========> K8s Undeploying $*-$(ARCH):$(VERSION)"
	@$(KUBECTL) -n $(NAMESPACE) delete -f manifests/$*.yaml
	$(KUBECTL) -n $(NAMESPACE) delete -f manifests/cert.yaml
	$(KUBECTL) -n $(NAMESPACE) delete -f manifests/config.yaml


