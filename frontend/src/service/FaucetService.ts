import axios from 'axios'

const faucetService = {
    SendEthersToSpecificAddress(wallet_to_send:string){
        return axios.post(import.meta.env.VITE_BACKEND_URL+'/SendEthersToSpecificAddress',{
            wallet: wallet_to_send
        }).catch(error => {
            console.log(error);
        });
    }
}

export default faucetService