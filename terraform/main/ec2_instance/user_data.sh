#!/bin/bash

exec > >(tee /var/log/user-data.log|logger -t user-data -s 2>/dev/console) 2>&1

yum update -y
yum install -y amazon-cloudwatch-agent docker amazon-ssm-agent

# Start Docker service
service docker start
usermod -a -G docker ec2-user




# Start SSM Agent service
service amazon-ssm-agent start
chkconfig amazon-ssm-agent on


adduser -m ssm-user
tee /etc/sudoers.d/ssm-agent-users <<'EOF'
# User rules for ssm-user
ssm-user ALL=(ALL) NOPASSWD:ALL
EOF
chmod 440 /etc/sudoers.d/ssm-agent-users 
# Now adding the ssm-user works!
usermod -a -G docker ssm-user


# Install Docker Compose
curl -L "https://github.com/docker/compose/releases/latest/download/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
chmod +x /usr/local/bin/docker-compose

# Create and populate the docker-compose.yml file
mkdir -p /apps
cat <<EOL > /apps/docker-compose.yml
networks:
  app-network:
    name: app-network
    driver: "bridge"

services:
  react-app:
    image: 
    ports:
      - "3000:4173"
    container_name: react-app
    environment:
      - VITE_BASE_URL=
    depends_on:
      - golang
    networks:
      - app-network

  golang:
    image: 
    container_name: golang
    restart: always
    environment:
      - PORT=8080
      - DB_USER=root
      - DB_PASSWORD=12345678
      - DB_HOST=mysql
      - DB_PORT=3306
      - DB_NAME=golang
      - JWT_SECRET_KEY=secretkey
      - REDIS_DB=0
      - REDIS_ADDR="redis:6379"
      - REDIS_PASS=12345678
    ports:
      - "4000:8080"
    depends_on:
      mysql:
        condition: service_healthy
      redis:
        condition: service_healthy
    healthcheck:
      test: ["CMD", "curl", "-f", "http://localhost:8080/v1"]
      interval: 10s
      retries: 5
      start_period: 30s
      timeout: 5s
    networks:
      - app-network
    deploy:
      resources:
        limits:
          cpus: "0.5"
          memory: "512M"

  mysql:
    image: mysql:latest
    environment:
      - MYSQL_ROOT_PASSWORD=12345678
      - MYSQL_DATABASE=golang
    volumes:
      - mysql-data:/var/lib/mysql
    ports:
      - "33066:3306"
    container_name: mysqlDb
    healthcheck:
      test:
        [
          "CMD",
          "mysqladmin",
          "ping",
          "-h",
          "localhost",
          "-u",
          "root",
          "-p12345678"
        ]
      interval: 10s
      retries: 5
      start_period: 30s
      timeout: 5s

    networks:
      - app-network

  redis:
    image: redis/redis-stack-server:latest
    command: >
      redis-server --maxmemory 1gb
                   --maxmemory-policy allkeys-lru
                   --protected-mode no
                   --requirepass 12345678
    ports:
      - "6379:6379"
    container_name: redis
    networks:
      - app-network
    healthcheck:
      test: ["CMD", "redis-cli", "-h", "localhost", "-p", "6379", "ping"]
      interval: 10s
      retries: 5
      start_period: 30s
      timeout: 5s

volumes:
  mysql-data:
EOL

# Change ownership of the file to ec2-user
# chown ec2-user:ec2-user /home/ec2-user/docker-compose.yml