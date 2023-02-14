FROM atlassian/jira-software:8.22.6-ubuntu-jdk11
LABEL maintainer="sebastian@sommerfeld.io"

COPY --chown=jira:jira assets/jira/dbconfig.xml /var/atlassian/application-data/jira/dbconfig.xml
COPY --chown=jira:jira assets/jira/database /var/atlassian/application-data/jira/database

ENV JVM_MINIMUM_MEMORY=1024m
ENV JVM_MAXIMUM_MEMORY=4096m
