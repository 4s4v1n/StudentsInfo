import axios from "axios"

class P2p {
    constructor(data) {
        this.time = data.time
        this.checking_peer = data.checking_peer
        this.state = data.state
        this.check_id = parseInt(data.check_id, 10);
        if (data.hasOwnProperty('id')) {
            this.id = data.id
        }
    }
}

export const P2pService = {
    async Get() {
        const response = await axios.get("http://localhost:8080/api/v1/p2p")

        return response.data
    },

    async Add(data) {
        const p2p = new P2p(data)
        return await axios.post("http://localhost:8080/api/v1/p2p", p2p)
    },

    async Delete(data) {
        return await axios.delete(`http://localhost:8080/api/v1/p2p/${data}`)
    },

    async Update(data) {
        const p2p = new P2p(data)
        return await axios.patch("http://localhost:8080/api/v1/p2p", p2p)
    },
}