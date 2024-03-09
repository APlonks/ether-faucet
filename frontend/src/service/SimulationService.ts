import axios from 'axios'

const SimulationService = {
    StartSimulation(accounts_per_wallets:number, ethers_per_transactions:number, transactions_per_blocks:number){
        return axios.post(import.meta.env.VITE_BACKEND_URL+'/startsimu',{
            accounts_per_wallets: accounts_per_wallets,
            ethers_per_transactions: ethers_per_transactions,
            transactions_per_blocks: transactions_per_blocks
        }).catch(error => {
            console.log(error);
        });
    },

    StopSimulation(){

    }
}

export default SimulationService