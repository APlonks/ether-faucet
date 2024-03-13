import axios from 'axios'
import Web3 from 'web3';

const testingService = {
    TestingAPI(api_addr:string){
        return axios.get(api_addr+'/testing',{
        }).catch(error => {
            console.log(error);
        });
    },

    TestingHTTPEndpoint(http_endpoint:string){
        console.log("http_endpoint:",http_endpoint)
        return axios.get(http_endpoint,{
        }).catch(error => {
            console.log(error);
        });
    },

    TestingWSEndpoint(ws_endpoint:string): Promise<boolean>{
        return new Promise((resolve) => {
            
            const ws = new WebSocket(ws_endpoint);
        
            ws.onopen = () => {
              console.log("Connection established");
              ws.close(); // Ferme la connexion une fois établie
              resolve(true);
            };
        
            ws.onerror = (error) => {
              console.log("Connection failed", error);
              resolve(false);
            };
        
            // Le gestionnaire onclose pourrait être utilisé pour des cas plus spécifiques
            // où la connexion est établie puis fermée pour d'autres raisons.
          });
    },
}

export default testingService