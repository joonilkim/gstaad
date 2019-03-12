FROM envoyproxy/envoy:latest

ARG AWS_REGION=ap-northeast-1
ARG COGNITO_USERPOOL_ID
ARG UPSTREAM=localhost
RUN test -n "$COGNITO_USERPOOL_ID"

COPY envoy /etc/envoy/
COPY pb /etc/envoy/pb/

RUN sed -i "s/\${UPSTREAM}/$UPSTREAM/" /etc/envoy/envoy.yaml
RUN sed -i "s/\${AWS_REGION}/$AWS_REGION/" /etc/envoy/envoy.yaml
RUN sed -i "s/\${COGNITO_USERPOOL_ID}/$COGNITO_USERPOOL_ID/" /etc/envoy/envoy.yaml

ENTRYPOINT ["/docker-entrypoint.sh"]
CMD ["envoy", "-c", "/etc/envoy/envoy.yaml"]
