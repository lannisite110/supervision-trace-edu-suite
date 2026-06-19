.PHONY: validate compliance register smoke dev-frontend

CORE ?= ../web3-edu-platform-core
PLUGINS_DIR ?= ..

validate:
	@for m in plugins/*/plugin.manifest.yaml; do \
		echo "==> $$m"; \
		cd "$(CORE)" && MANIFEST="../supervision-trace-edu-suite/$$m" make validate-plugin; \
	done

compliance:
	cd "$(CORE)" && bash ci/compliance/compliance-check.sh ..

register:
	cd "$(CORE)" && make register-plugins PLUGINS_DIR="$(PLUGINS_DIR)"

smoke:
	bash scripts/joint-debug-smoke.sh

dev-frontend:
	cd dev-frontend && npm install && npm run dev
