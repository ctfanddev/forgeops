FROM gcr.io/forgerock-io/am/pit1:7.0.0-007878d135fb859c709c1c7c03ddccb582c9b2f5

# BEGIN - temporary work around for testing if the ds-idrepo is configured before AM starts. Remove when FBC is integrated
# Installs ldap utils for querying the repo.
USER root
RUN apt-get update && apt-get install -y ldap-utils && apt-get clean all
COPY --chown=forgerock:root openam /home/forgerock/openam
# END

COPY logback.xml /usr/local/tomcat/webapps/am/WEB-INF/classes
RUN chown -R forgerock:root /usr/local/tomcat

#  JDK11 with GC trace options
# ENV CATALINA_OPTS "-server -XX:MaxRAMPercentage=75 -Dcom.sun.identity.configuration.directory=/home/forgerock/openam -Dcom.iplanet.services.stats.state=off -Xlog:gc*,gc+ref=debug,gc+heap=debug,gc+age=trace:file=/tmp/gc-%p-%t.log:tags,uptime,time,level:filecount=10,filesize=10m -Djavax.net.ssl.trustStore=/var/run/secrets/truststore/cacerts -Djavax.net.ssl.trustStorePassword=changeit"
# JDK 11
ENV CATALINA_OPTS "-server -XX:MaxRAMPercentage=75 -Dcom.sun.identity.configuration.directory=/home/forgerock/openam -Dcom.iplanet.services.stats.state=off -Djavax.net.ssl.trustStore=/var/run/secrets/truststore/cacerts -Djavax.net.ssl.trustStorePassword=changeit"

#ENV CATALINA_OPTS "-server -Dcom.sun.identity.configuration.directory=/home/forgerock/openam -Dcom.iplanet.services.stats.state=off -Djavax.net.ssl.trustStore=/var/run/secrets/truststore/cacerts -Djavax.net.ssl.trustStorePassword=changeit"

# This copies in the boot.json and other files...
COPY --chown=forgerock:root openam /home/forgerock/openam

# Override the entrypoint and just start tomcat
#CMD ["catalina.sh", "run"]

RUN chown -R forgerock:root /home/forgerock
#ENTRYPOINT ["/bin/sh"]
# Use sh - because windows has issues with +x on the script
USER forgerock
CMD ["bash", "-c", "/home/forgerock/openam/boot.sh"]
