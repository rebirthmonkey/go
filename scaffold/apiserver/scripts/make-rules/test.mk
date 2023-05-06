# ==============================================================================
# Targets

.PHONY: test.api.%
test.api.%:
	@./tests/api/rest.sh insecure::$*

.PHONY: test.api
test.api:
	@./tests/api/rest.sh insecure::user
	@./tests/api/rest.sh insecure::secret
	@./tests/api/rest.sh insecure::policy

.PHONY: test.api-basic
test.api-basic:
	@./tests/api/auth.sh insecure::basic

.PHONY: test.api-jwt.%
test.api-jwt.%:
	@./tests/api/auth.sh insecure::jwt::$*

.PHONY: test.api-authz
test.api-authz:
	@./tests/api/authz.sh insecure::authz

.PHONY: test.api-authz-jwt
test.api-authz-jwt:
	@./tests/api/authz.sh insecure::authz-jwt