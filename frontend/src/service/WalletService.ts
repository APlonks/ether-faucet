import axios from 'axios'

const BASE_URL = 'http://golangserver:8080'

const walletService = {
    sendTransaction(wallet_to_send:string){
        return axios.post(BASE_URL+'/sendTransaction',{
            wallet: wallet_to_send
        }).catch(error => {
            console.log(error);
        });
    }
}

export default walletService