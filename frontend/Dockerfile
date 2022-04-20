FROM node:16-alpine
WORKDIR /src
RUN apk update && \
    yarn global add firebase-tools && \
    yarn --dev

ENV HOST 0.0.0.0
