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

### Docker Compose

#### Install docker

Docker engine for Ubuntu : https://docs.docker.com/engine/install/ubuntu/

#### Start

```bash
# Modify the environment variables for the Backend
# Front variables will be modified in the web interface
vim .env_compose 

# Start the compose
docker compose up
```

#### Stop
```bash
docker compose down
```

## TODO :
- Error management, if blockchain not running and request are sent from the front : dont write it's ok
- Modify nginx configuration to allow reload on SPA without returning 404
- Implement transaction post EIP 1559
- Ethereum logo using ThreeJS
- Statistics page about Simulation
- Deployment using Helm for Kubernetes