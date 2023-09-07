import axios from "axios"

export const PassOneTwoService = {
    async Get(value) {
        const response = await axios.get(`http://localhost:8080/api/v1/fnc/pass_one_two?task_1=${value.task_1}&task_2=${value.task_2}&task_3=${value.task_3}`)
        return response.data
    },
}