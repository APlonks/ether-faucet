import axios from 'axios'

const SimulationService = {
    StartSimulation(accounts_per_wallet:number, ethers_per_wallet:number, ethers_per_transaction:number, transactions_per_block:number){
        const api_addr = localStorage.getItem('api_addr');
        return axios.post(api_addr+'/start-simulation',{
            accounts_per_wallet: accounts_per_wallet,
            ethers_per_wallet: ethers_per_wallet,
            ethers_per_transaction: ethers_per_transaction,
            transactions_per_block: transactions_per_block
        }).catch(error => {
            console.log(error);
        });
    },

    StopSimulation(){
        const api_addr = localStorage.getItem('api_addr')
        return axios.post(api_addr+'/stop-simulation',{
        }).catch(error => {
            console.log(error);
        });
    }
}

export default SimulationService