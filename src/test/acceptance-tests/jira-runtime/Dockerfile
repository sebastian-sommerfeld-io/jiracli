FROM atlassian/jira-software:9.6.0-ubuntu-jdk11
LABEL maintainer="sebastian@sommerfeld.io"

COPY --chown=jira:jira assets/jira/dbconfig.xml /var/atlassian/application-data/jira/dbconfig.xml
COPY --chown=jira:jira assets/jira/database /var/atlassian/application-data/jira/database

ENV JVM_MINIMUM_MEMORY=1024m
ENV JVM_MAXIMUM_MEMORY=4096m

# See "Shared directory and user IDs" (https://hub.docker.com/r/atlassian/jira-software)
# By default the Jira application runs as the user jira, with a UID and GID of `2001`.
# Consequently this UID must have write access to the shared filesystem.
RUN chmod -R o-rw /var/atlassian/application-data

USER 2001
