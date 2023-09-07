import axios from "axios"

export const FriendsService = {
    async Get() {
        const response = await axios.get("http://localhost:8080/api/v1/friends")

        return response.data
    },

    async Add(data) {
        return await axios.post("http://localhost:8080/api/v1/friends", data)
    },

    async Delete(data) {
        return await axios.delete(`http://localhost:8080/api/v1/friends/${data}`)
    },

    async Update(data) {
        return await axios.patch("http://localhost:8080/api/v1/friends", data)
    },
}