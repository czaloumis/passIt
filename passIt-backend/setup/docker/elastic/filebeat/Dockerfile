FROM docker.elastic.co/beats/filebeat:8.13.4

# Switch to root to change permissions
USER root

# Copy config
COPY filebeat.yml /usr/share/filebeat/filebeat.yml

# Fix permissions
RUN chmod go-w /usr/share/filebeat/filebeat.yml

# (Optional) Return to default user
USER filebeat