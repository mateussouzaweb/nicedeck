FROM ubuntu:latest
LABEL org.opencontainers.image.source=https://github.com/mateussouzaweb/nicedeck
LABEL maintainer="Mateus Souza <mateussouzaweb@gmail.com>"
ENV DEBIAN_FRONTEND=noninteractive

# Add languages
RUN apt-get update && apt-get install -y locales \
    && rm -rf /var/lib/apt/lists/* \
	&& localedef -i en_US -c -f UTF-8 -A \
    /usr/share/locale/locale.alias en_US.UTF-8
ENV LANG=en_US.utf8

# Install system dependencies
RUN apt update && apt install -y \
    build-essential g++ pkg-config git curl wget

# Install GitHub CLI
RUN mkdir -p -m 755 /etc/apt/keyrings \
    && wget -qO- https://cli.github.com/packages/githubcli-archive-keyring.gpg | tee /etc/apt/keyrings/githubcli-archive-keyring.gpg > /dev/null \
    && chmod go+r /etc/apt/keyrings/githubcli-archive-keyring.gpg \
    && echo "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/githubcli-archive-keyring.gpg] https://cli.github.com/packages stable main" | tee /etc/apt/sources.list.d/github-cli.list > /dev/null \
    && apt update \
    && apt install -y gh

# Install Golang
COPY --from=golang:latest /usr/local/go/ /usr/local/go/
ENV PATH="/usr/local/go/bin:${PATH}"

# Set command
CMD ["/bin/bash"]