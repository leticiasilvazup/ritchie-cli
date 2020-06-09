package tpl_shell

const (
	Main = `#!/bin/sh

. ./{{bin-name}}/{{bin-name}}.sh --source-only

run $SAMPLE_TEXT $SAMPLE_LIST $SAMPLE_BOOL`

	Makefile = `# SH
BINARY_NAME={{bin-name}}.sh
DIST=../dist
DIST_DIR=$(DIST)/commons/bin
build:
	mkdir -p $(DIST_DIR)
	cp main.sh $(DIST_DIR)/$(BINARY_NAME) && cp -r {{bin-name}} Dockerfile set_umask.sh $(DIST_DIR)
	chmod +x $(DIST_DIR)/$(BINARY_NAME)`

	Dockerfile = `
FROM alpine:latest

WORKDIR /app
COPY . /app

RUN chmod +x set_umask.sh
RUN chmod +x {{bin-name}}.sh

WORKDIR /app

ENTRYPOINT ["./set_umask.sh"]
CMD ["./{{bin-name}}.sh"]`

	Umask = `#!/bin/sh
umask 0011
$1`

	File = `#!/bin/sh
run() {
  echo "Hello World! "
  echo "You receive $SAMPLE_TEXT in text. "
  echo "You receive $SAMPLE_LIST in list. " 
  echo "You receive $SAMPLE_BOOL in boolean. "  
}
`
)
