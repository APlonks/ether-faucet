import axios from 'axios'

const BASE_URL = 'http://192.168.3.124:9999'

const walletService = {
    sendTransaction(wallet_to_send:string){
        return axios.post(BASE_URL+'/SendEthersToSpecificAddress',{
            wallet: wallet_to_send
        }).catch(error => {
            console.log(error);
        });
    }
}

export default walletService