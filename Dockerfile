FROM ubuntu:latest
LABEL org.opencontainers.image.source https://github.com/mateussouzaweb/nicedeck
LABEL maintainer="Mateus Souza <mateussouzaweb@gmail.com>"
ENV DEBIAN_FRONTEND=noninteractive

# Add languages
RUN apt-get update && apt-get install -y locales && rm -rf /var/lib/apt/lists/* \
	&& localedef -i en_US -c -f UTF-8 -A /usr/share/locale/locale.alias en_US.UTF-8
ENV LANG en_US.utf8

# Install system dependencies
RUN apt update && apt install -y \
    build-essential g++ pkg-config git curl wget \
    libgtk-4-dev libwebkitgtk-6.0-dev \
    qt6-base-dev qt6-webengine-dev

# Install golang
COPY --from=golang:latest /usr/local/go/ /usr/local/go/
ENV PATH="/usr/local/go/bin:${PATH}"

# Set command
CMD ["/bin/bash"]