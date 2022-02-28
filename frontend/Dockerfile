FROM node:16-alpine
WORKDIR /src
RUN apk update && \
    yarn global add create-nuxt-app && \
    yarn global add firebase-tools

ENV HOST 0.0.0.0
