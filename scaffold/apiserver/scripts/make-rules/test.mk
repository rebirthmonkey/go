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
