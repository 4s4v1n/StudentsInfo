import axios from "axios"

export const StatisticBlockService = {
    async Get(value) {
        const response = await axios.get(`http://localhost:8080/api/v1/fnc/statistic_block?block_1=${value.block1}&block_2=${value.block2}`)
        return response.data
    },
}