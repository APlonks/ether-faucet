import axios from 'axios'

const faucetService = {
    SendEthersToSpecificAddress(wallet_to_send:string){
        const api_addr = localStorage.getItem('api_addr');
        return axios.post(api_addr+"/faucet",{
            wallet: wallet_to_send
        });
        // }).catch(error => {
        //     console.log(error);
        // });
    }
}

export default faucetService