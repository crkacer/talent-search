import Api from '@/services/Api'

export default {
    register (credentials) {
        return Api().post('register', credentials)
    }
    
    getAllTalents() {
        return Api().get('/getall')
    }
} 
/*
AuthenticationService.register ({
    "email": 'testing@gmail.com'
    "password": '123456'
})
*/