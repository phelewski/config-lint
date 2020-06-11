FROM phelewski/config-lint:v0.0.2

ARG PARSER=''
ENV PARSER="${PARSER}"

ARG INPUT_PATH=''
ENV INPUT_PATH="${INPUT_PATH}"

COPY . .
ENTRYPOINT [ "sh", "-c", ".github/actions/entrypoint.sh $PARSER $INPUT_PATH" ]