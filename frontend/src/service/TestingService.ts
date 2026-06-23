import axios from 'axios'

const testingService = {
    TestingAPI(){
        return axios.get('/api/testing',{
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
              ws.close(); // Close connection once established
              resolve(true);
            };
        
            ws.onerror = (error) => {
              console.log("Connection failed", error);
              resolve(false);
            };
          });
    },
}

export default testingService