FROM node:24-alpine AS deps
WORKDIR /app
COPY package.json pnpm-lock.yaml ./
RUN npm i -g pnpm && pnpm install --frozen-lockfile

FROM node:24-alpine AS builder
WORKDIR /app
COPY --from=deps /app/node_modules ./node_modules
COPY . .
RUN pnpm build

FROM node:24-alpine
RUN npm i -g pnpm
WORKDIR /app

COPY --from=deps /app/node_modules ./node_modules
COPY --from=builder /app/.output ./.output
COPY . .

CMD if [ "$APP_ENV" = "development" ]; then \
      pnpm dev; \
    else \
      node ./.output/server/index.mjs; \
    fi