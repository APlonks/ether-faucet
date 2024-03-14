import axios from 'axios'

const faucetService = {
    SendEthersToSpecificAddress(wallet_to_send:string){
        const api_addr = localStorage.getItem('api_addr');
        if (!api_addr) {
            throw new Error('API URL not found in localStorage');
        }

        if (api_addr !== localStorage.getItem('api_addr')) {
            throw new Error('The API URL does not match the expected value in localStorage');
        }

        return axios.post(api_addr+"/faucet",{
            wallet: wallet_to_send
        }).catch(error => {
            console.log(error);
        });
    }
}

export default faucetService