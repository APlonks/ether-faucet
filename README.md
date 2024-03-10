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

### Docker

#### Build Dockerfile

**Frontend**
```bash
cd frontend
docker build -t ether-faucet-frontend:0.1 .
docker run -p 8888:80 --rm --name ether-faucet-frontend ether-faucet-frontend:0.1
```

**Backend**
```bash
cd backend
docker build -t ether-faucet-backend:0.1 .
docker run
```

## TODO : 
- Implement transaction post EIP 1559