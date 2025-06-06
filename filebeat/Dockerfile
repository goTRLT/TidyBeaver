FROM docker.elastic.co/beats/filebeat:9.0.0
COPY --chown=root:filebeat filebeat.yml T:/Repo/filebeat/filebeat.yml
USER root
RUN chown -R root T:/Repo/filebeat/filebeat.yml