import axios from 'axios'

const SimulationService = {
    StartSimulation(accounts_per_wallets:number, ethers_per_wallets:number, ethers_per_transactions:number, transactions_per_blocks:number){
        return axios.post(import.meta.env.VITE_BACKEND_URL+'/start-simulation',{
            accounts_per_wallets: accounts_per_wallets,
            ethers_per_wallets: ethers_per_wallets,
            ethers_per_transactions: ethers_per_transactions,
            transactions_per_blocks: transactions_per_blocks
        }).catch(error => {
            console.log(error);
        });
    },

    StopSimulation(){
        return axios.post(import.meta.env.VITE_BACKEND_URL+'/stop-simulation',{
        }).catch(error => {
            console.log(error);
        });
    }
}

export default SimulationService