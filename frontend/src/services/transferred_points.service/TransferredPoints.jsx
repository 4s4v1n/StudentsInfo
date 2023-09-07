import axios from "axios"

class TransferredPoints {
    constructor(data) {
        this.checking_peer = data.checking_peer
        this.checked_peer = data.checked_peer
        this.points_amount = parseInt(data.points_amount, 10);
        if (data.hasOwnProperty('id')) {
            this.id = data.id
        }
    }
}

export const TransferredPointsService = {
    async Get() {
        const response = await axios.get("http://localhost:8080/api/v1/transferred_points")

        return response.data
    },

    async Add(data) {
        const points = new TransferredPoints(data)
        return await axios.post("http://localhost:8080/api/v1/transferred_points", points)
    },

    async Delete(data) {
        return await axios.delete(`http://localhost:8080/api/v1/transferred_points/${data}`)
    },

    async Update(data) {
        const points = new TransferredPoints(data)
        return await axios.patch("http://localhost:8080/api/v1/transferred_points", points)
    },
}