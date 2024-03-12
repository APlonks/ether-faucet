import axios from 'axios'

const testingService = {
    TestingAPI(api_addr:string){
        return axios.post(api_addr+'/testing',{
        }).catch(error => {
            console.log(error);
        });
    },
    // TestingHTTPEndpoint(api_addr:string){
    //     return axios.post(api_addr+'/testing',{
    //     }).catch(error => {
    //         console.log(error);
    //     });
    // },
    // TestingWSEndpointEndpoint(api_addr:string){
    //     return axios.post(api_addr+'/testing',{
    //     }).catch(error => {
    //         console.log(error);
    //     });
    // },
}

export default testingService