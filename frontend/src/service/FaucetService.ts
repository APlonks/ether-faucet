import axios from 'axios'

const faucetService = {
    SendEthersToSpecificAddress(wallet_to_send:string){
        console.log("The backend URL :",import.meta.env.VITE_BACKEND_URL)
        console.log("HTTP ENDPOINT NODE :",import.meta.env.VITE_HTTP_ENDPOINT_NODE)
        return axios.post(import.meta.env.VITE_BACKEND_URL+'/faucet',{
            wallet: wallet_to_send
        }).catch(error => {
            console.log(error);
        });
    }
}

export default faucetService