import axios from 'axios'

const VITE_BACKEND_URL = 'http://192.168.5.10:8080'

const faucetService = {
    SendEthersToSpecificAddress(wallet_to_send:string){
        return axios.post(VITE_BACKEND_URL+'/SendEthersToSpecificAddress',{
            wallet: wallet_to_send
        }).catch(error => {
            console.log(error);
        });
    }
}

export default faucetService