const axios = require("axios")
const port = process.env.BACKEND_PORT
const baseUrl = `http://localhost:${port}/`

const mockUser = {
    "email": "asdasd@mail.com",
    "firstName": "armna",
    "lastName": "asd",
    "password": "asdasd",
    "DateOfBirth": "2000-01-02T15:04:05Z"
}

describe("user calls", () => {
    it("user/register", async () => {
        const res = await axios.post(`${baseUrl}user/register`, mockUser)

        expect(res.status).toBe(201)
    })
})
