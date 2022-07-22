const axios = require('axios')
const { baseUrl } = require('../config')
const {
  mockBadLoginUser,
  mockBadLoginUser2,
  mockLoginUser,
  mockNormalUser1,
  mockNormalUser2,
} = require('../__mock__/user.mock')

const instance = axios.create({
  withCredentials: true,
  baseURL: baseUrl,
})

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
    it('normal register and follow', async () => {
      const res1 = await instance.post(`user/register`, mockNormalUser1)
      const res2 = await instance.post(`user/register`, mockNormalUser2)

      expect(res1.status).toBe(201)
      expect(res2.status).toBe(201)

      const cookie1 = res1.headers['set-cookie'][0].split(';')[0]
      const cookie2 = res2.headers['set-cookie'][0].split(';')[0]
      const userId1 = res1.data.userID
      const userId2 = res2.data.userID

      const res = await instance.post(
        `user/follow`,
        {
          userId2: userId1,
        },
        {
          headers: { Cookie: cookie2 },
          credentials: 'include',
        }
      )

      expect(res.status).toBe(201)

      const follows = await instance.get('user/check-follow', {
        headers: { Cookie: cookie1 },
        credentials: 'include',
      })

      console.log(follows.data)
    })

    it('try register without one field', async () => {
      const data = createUserWithoutField(mockNormalUser1)
      const requests = data.map((el) => instance.post(`user/register`, el))
      const promises = await Promise.allSettled(requests)

      mapStatuses = promises.map((el) => el.reason.response.status)

      mapStatuses.forEach((el) => {
        expect(el).toBe(400)
      })
    })

    it('already registered', async () => {
      expect.assertions(1)
      return instance.post(`user/register`, mockNormalUser1).catch((err) => {
        expect(err.response.status).toBe(400)
      })
    })
  })

  describe('user/login', () => {
    it('normal login', async () => {
      const res = await instance.post(`user/login`, mockLoginUser)

      expect(res.status).toBe(200)
    })

    it('bad login', async () => {
      expect.assertions(1)
      return instance.post(`user/login`, mockBadLoginUser).catch((err) => {
        expect(err.response.status).toBe(400)
      })
    })

    it('bad password', async () => {
      expect.assertions(1)
      return instance.post(`user/login`, mockBadLoginUser2).catch((err) => {
        expect(err.response.status).toBe(400)
      })
    })
  })
})
