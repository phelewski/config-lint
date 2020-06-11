FROM phelewski/config-lint:v0.0.1

ARG PARSER=''
ENV PARSER="${PARSER}"

ARG INPUT_PATH=''
ENV INPUT_PATH="${INPUT_PATH}"

COPY . .
ENTRYPOINT [ "sh", "-c", "./config-lint $PARSER $INPUT_PATH" ]
