# =============== #
# Stage 1: Build  #
# =============== #
FROm node:25 as builder

WORKDIR /app

copy package*.json pnpm-lock.yaml pnpm-workspace.yaml ./

RUN npm install -g pnpm
RUN pnpm install

copy . .

run pnpm build

# =================== #
# Stage 2: Production #
# =================== #

from node:25

WORKDIR /app

copy  package*.json ./

COPY --from=builder /app/dist ./dist

COPY package.json ./dist/package.json

copy . .

CMD ["pm2-runtime", "start", "ecosystem.config.js"]
