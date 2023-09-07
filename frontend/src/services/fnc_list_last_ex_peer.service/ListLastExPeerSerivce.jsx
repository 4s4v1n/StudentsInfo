import axios from "axios"

export const ListLastExPeerSerivce = {
    async Get(value) {
        const response = await axios.get(`http://localhost:8080/api/v1/fnc/list_last_ex_peer?ex=${value.ex}`)
        return response.data
    },
}