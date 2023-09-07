import axios from "axios"

class Xp {
    constructor(data) {
        this.check_id = parseInt(data.check_id, 10);
        this.xp_amount = parseInt(data.xp_amount, 10);
        if (data.hasOwnProperty('id')) {
            this.id = data.id
        }
    }
}

export const XpService = {
    async Get() {
        const response = await axios.get("http://localhost:8080/api/v1/xp")

        return response.data
    },

    async Add(data) {
        const xp = new Xp(data)
        return await axios.post("http://localhost:8080/api/v1/xp", xp)
    },

    async Delete(data) {
        return await axios.delete(`http://localhost:8080/api/v1/xp/${data}`)
    },

    async Update(data) {
        const xp = new Xp(data)
        return await axios.patch("http://localhost:8080/api/v1/xp", xp)
    },
}