const axios = require('axios')
const port = process.env.BACKEND_PORT
const baseUrl = `http://localhost:${port}/`

const mockNormalUser = {
  email: 'asdasd@mail.com',
  firstName: 'vasya',
  lastName: 'pupkin',
  password: 'lolkek',
  DateOfBirth: '2000-01-02T15:04:05Z',
}

const mockLoginUser = {
  email: 'asdasd@mail.com',
  password: 'lolkek',
}

const mockBadLoginUser = {
  email: 'asdasd@mail.com',
  password: 'lol',
}

const mockBadLoginUser2 = {
  email: 'asd@mail.com',
  password: 'lolkek',
}

const createUserWithoutField = (data) => {
  const res = []
  for (const key in data) {
    let { [key]: deletedKey, ...withoutKey } = data
    res.push(withoutKey)
  }
  return res
}

describe('user calls', () => {
  describe('user/register', () => {
    it('try register without one field', async () => {
      const data = createUserWithoutField(mockNormalUser)
      const requests = data.map((el) =>
        axios.post(`${baseUrl}user/register`, el)
      )
      const promises = await Promise.allSettled(requests)

      mapStatuses = promises.map((el) => el.reason.response.status)

      mapStatuses.forEach((el) => {
        expect(el).toBe(400)
      })
    })

    it('normal register', async () => {
      const res = await axios.post(`${baseUrl}user/register`, mockNormalUser)

      expect(res.status).toBe(201)
    })

    it('already registered', async () => {
      expect.assertions(1)
      return axios.post(`${baseUrl}user/register`, mockNormalUser).catch((err) => {
        expect(err.response.status).toBe(400)
      })
    })
  })

  describe('user/login', () => {
    it('normal login', async () => {
      const res = await axios.post(`${baseUrl}user/login`, mockLoginUser)

      expect(res.status).toBe(200)
    })

    it('bad login', async () => {
      expect.assertions(1)
      return axios.post(`${baseUrl}user/login`, mockBadLoginUser).catch((err) => {
        expect(err.response.status).toBe(400)
      })
    })

    it('bad password', async () => {
      expect.assertions(1)
      return axios.post(`${baseUrl}user/login`, mockBadLoginUser2).catch((err) => {
        expect(err.response.status).toBe(400)
      })
    })

    it('just fail', async () => {
      expect(10).toBe(1)
    })
  })
})
