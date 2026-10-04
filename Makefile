.PHONY: preview ## Run the preview server and open it in the browser.
preview:
	hugo server --buildDrafts F --buildFuture --openBrowser

.PHONY: mod-update ## Update all modules.
mod-update: hugo.toml
	hugo mod get ./...
	hugo mod tidy
