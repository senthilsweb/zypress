import axios from 'axios'
    
    const apiClient = axios.create({
      baseURL: `http://localhost:9000`,
      withCredentials: false, 
      headers: {
        Accept: 'application/json',
        'Content-Type': 'application/json'
      }
    })
    
    export default {
      getDocs() {
        return apiClient.get('/content-api/docs')
      }
    }