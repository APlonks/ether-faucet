import axios from 'axios'

const faucetService = {
    SendEthersToSpecificAddress(wallet_to_send:string){
        // const apiUrl = localStorage.getItem('api_addr') || 'URL_PAR_DÉFAUT_ICI';

        const apiUrl = localStorage.getItem('api_addr');
        if (!apiUrl) {
            // Si aucune URL n'est trouvée dans le localStorage, lancez une erreur
            throw new Error('API URL not found in localStorage');
        }

        if (apiUrl !== localStorage.getItem('api_addr')) {
            // Si l'URL obtenue est différente de celle du localStorage, lancez une erreur
            throw new Error('The API URL does not match the expected value in localStorage');
        }

        return axios.post(apiUrl+"/faucet",{
            wallet: wallet_to_send
        }).catch(error => {
            console.log(error);
        });
    }
}

export default faucetService