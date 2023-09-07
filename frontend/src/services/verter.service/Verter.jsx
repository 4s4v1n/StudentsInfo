import axios from "axios"

class Verter {
    constructor(data) {
        this.time = data.time
        this.state = data.state
        this.check_id = parseInt(data.check_id, 10);
        if (data.hasOwnProperty('id')) {
            this.id = data.id
        }
    }
}

export const VerterService = {
    async Get() {
        const response = await axios.get("http://localhost:8080/api/v1/verter")

        return response.data
    },

    async Add(data) {
        const verter = new Verter(data)
        return await axios.post("http://localhost:8080/api/v1/verter", verter)
    },

    async Delete(data) {
        return await axios.delete(`http://localhost:8080/api/v1/verter/${data}`)
    },

    async Update(data) {
        const verter = new Verter(data)
        return await axios.patch("http://localhost:8080/api/v1/verter", verter)
    },
}