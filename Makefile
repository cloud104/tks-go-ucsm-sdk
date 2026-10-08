EXTRACT_DIR := tools/mogen/extract
VENV        := $(EXTRACT_DIR)/.venv
PYTHON      := $(VENV)/bin/python3

.PHONY: schema generate check-generated

# Cria o venv do extrator na primeira vez (ou quando o requirements mudar).
$(PYTHON): $(EXTRACT_DIR)/requirements.txt
	python3 -m venv $(VENV)
	$(VENV)/bin/pip install -q -r $(EXTRACT_DIR)/requirements.txt
	touch $(PYTHON)

# Reextrai o schema do ucsmsdk. Só é preciso ao mudar classes.txt ou a versão do SDK.
schema: $(PYTHON)
	$(PYTHON) $(EXTRACT_DIR)/extract.py --classes tools/mogen/classes.txt --out schema

# Regera as structs a partir do schema versionado (não precisa de Python).
generate:
	go generate ./...

# Falha se o código gerado commitado estiver desatualizado em relação ao schema.
check-generated: generate
	git diff --exit-code -- mo/generated schema
 