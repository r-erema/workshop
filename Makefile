GOLANGCI_IMAGE=golangci/golangci-lint:v2.12.2

CSR_CONF_PATH=./configs/csr.conf
CA_CERT_PATH=./assets/rootCA.crt
CA_KEY_PATH=./assets/rootCA.key
CSR_PATH=./assets/csr_for_common_cert_for_all.csr
CERT_KEY_PATH=./assets/common_cert_key_for_all.key
CERT_PATH=./assets/common_cert_for_all.crt

before-commit:
	sudo -S $(MAKE) test
	$(MAKE) lint
	sudo .bin/github/act -P ubuntu-latest=-self-hosted --var GITHUB_ACTIONS=true

test:
	go test -exec sudo -race -cover -coverprofile=coverage.txt -covermode=atomic ./...

lint:
	docker run --rm -v ${PWD}:/app -w /app ${GOLANGCI_IMAGE} sh -c "golangci-lint fmt && golangci-lint run --fix"

generate-certs:
	openssl genrsa -out ${CA_KEY_PATH} 2048
	openssl req -x509 -new -nodes -key ${CA_KEY_PATH} -days 10000 -out ${CA_CERT_PATH} \
		-config ${CSR_CONF_PATH} -extensions v3_ca
	openssl genrsa -out ${CERT_KEY_PATH} 2048
	openssl req -new -key ${CERT_KEY_PATH} -out ${CSR_PATH} -config ${CSR_CONF_PATH}
	openssl x509 -req -in ${CSR_PATH} -CA ${CA_CERT_PATH} -CAkey ${CA_KEY_PATH} \
		-CAcreateserial -out ${CERT_PATH} -days 10000 \
		-extensions v3_req -extfile ${CSR_CONF_PATH} -sha256
	openssl req  -noout -text -in ${CSR_PATH}
	openssl x509  -noout -text -in ${CERT_PATH}

trust-certs:
	sudo -S mkdir -p /usr/local/share/ca-certificates/erema
	sudo cp ${CA_CERT_PATH} ${CERT_PATH} /usr/local/share/ca-certificates/erema
	sudo update-ca-certificates -f

# https://about.gitlab.com/blog/2018/06/07/keeping-git-commit-history-clean/
start-changing-git-commit:
	# 1. Go to the previous commit before target commit
	git rebase -i `git log --pretty=%P -n 1 ${TARGET_COMMIT_TO_CHANGE}`
	# 2. Change "pick -> edit" desired commit(first in the list), example:
	# pick 74748f9 CI adding                edit 74748f9 CI adding
	# pick 63f7877 Brunch Sums Problem  =>  pick 63f7877 Brunch Sums Problem
	# ...                                   ...
	# 3. Make needed changes and add to commit changed files, example : git add .github/workflows/lint.yml
	# 4. Run `make finish-changing-git-commit`
finish-changing-git-commit:
	git rebase --continue
	git push --force-with-lease origin master
