FROM phelewski/config-lint:v0.0.2

ARG INPUT_PARSER=''
ENV INPUT_PARSER="${INPUT_PARSER}"

ARG INPUT_INPUT_PATH=''
ENV INPUT_INPUT_PATH="${INPUT_INPUT_PATH}"

RUN apt-get update \
    && apt-get install -y curl

RUN curl -L https://github.com/stelligent/config-lint/releases/latest/download/config-lint_Linux_x86_64.tar.gz | tar xz -C /usr/local/bin config-lint \
    && chmod +rx /usr/local/bin/config-lint

COPY . .
ENTRYPOINT [ "sh", "-c", "/usr/local/bin/config-lint $INPUT_PARSER $INPUT_INPUT_PATH" ]