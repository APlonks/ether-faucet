# ether-faucet


## Dev Environment

OS :
- Ubuntu : 22.04

Install :
- nodejs 20.11.1 : https://nodejs.org/en
- Go 1.22.0 : https://go.dev/doc/install

Vscode extensions :
- Vue - Official

Browser extensions : 
- Vue Dev Tools : https://devtools.vuejs.org/guide/installation.html

## Installation

**Clone the project**
```bash
git clone https://github.com/APlonks/ether-faucet.git
cd ether-faucet
```
### Source code

**Frontend**
```bash
cd frontend
npm install
npm run dev # Dev mode
# or 
npm run build   # Prod mode with a nginx for example
```

**Backend**
```bash
cd backend
make dev # Dev mode
# or 
make build   # Build for prod mode
make run    # Run for prod mode
```

**Redis Stack**
```bash
# Using docker : 
docker run -d --name redis-stack -p 6379:6379 -p 8001:8001 redis/redis-stack:7.2.0-v9
```

### Docker Compose

#### Install docker

Docker engine for Ubuntu : https://docs.docker.com/engine/install/ubuntu/

#### Start

```bash
# Create the network if it doesn't exist
docker network create --driver bridge bcnetwork

# Modify the environment variables for the Backend
# Front variables will be modified in the web interface
vim .env_compose 

# Start the compose
docker compose up
```

#### Stop
```bash
# Stop the compose
docker compose down
```

## TODO :
- Implement transaction post EIP 1559
- Ethereum logo using ThreeJS
- Statistics page about Simulation
- Deployment using Helm for Kubernetes
- Add a Redis to manage the synchronization of multiple backends (Kubernetes deployment) and features such as simulation, which you can only run once at any given time. 
