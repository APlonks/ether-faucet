import axios from 'axios'

const SimulationService = {
    StartSimulation(accounts_per_wallet:number, ethers_per_wallet:number, ethers_per_transaction:number, transactions_per_block:number){
        return axios.post(import.meta.env.VITE_BACKEND_URL+'/start-simulation',{
            accounts_per_wallet: accounts_per_wallet,
            ethers_per_wallet: ethers_per_wallet,
            ethers_per_transaction: ethers_per_transaction,
            transactions_per_block: transactions_per_block
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