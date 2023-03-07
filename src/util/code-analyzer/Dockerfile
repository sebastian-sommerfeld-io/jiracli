FROM jetbrains/qodana-go:2023.1-eap
LABEL maintainer="sebastian@sommerfeld.io"

RUN apt-get update \
    && apt-get install -y --no-install-recommends unzip=6.0-26+deb11u1 \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

# Download AsciiDoc plugin for AsciiDoc checks
RUN curl -L -o asciidoctor-intellij-plugin.zip https://github.com/asciidoctor/asciidoctor-intellij-plugin/releases/download/0.38.2/asciidoctor-intellij-plugin-0.38.2.zip \
    && unzip asciidoctor-intellij-plugin.zip \
    && cp -a asciidoctor-intellij-plugin /opt/idea/plugins/asciidoctor-intellij-plugin

# Download Grazie plugin for grammar checks
RUN curl -L -o grazie.zip 'https://plugins.jetbrains.com/plugin/download?rel=true&updateId=257026' \
    && unzip grazie.zip \
    && cp -a grazie /opt/idea/plugins/grazie
