FROM node:alpine

WORKDIR '/app'

COPY package.json .

RUN npm install

COPY . .

# Vite reads environment variables at build time, not at runtime
ENV VITE_BASE_URL=

RUN npm run build

EXPOSE 4173

CMD ["npm", "run", "preview", "--", "--host"]