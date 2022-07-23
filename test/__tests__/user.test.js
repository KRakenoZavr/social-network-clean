const axios = require('axios')
const { baseUrl } = require('../config')
const {
  mockBadLoginUser,
  mockBadLoginUser2,
  mockLoginUser,
  mockNormalUser1,
  mockNormalUser2,
  mockNormalUserPrivate,
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
      // check user registration
      const res1 = await instance.post(`user/register`, mockNormalUser1)
      const res2 = await instance.post(`user/register`, mockNormalUser2)

      expect(res1.status).toBe(201)
      expect(res2.status).toBe(201)

      const cookie1 = res1.headers['set-cookie'][0].split(';')[0]
      const cookie2 = res2.headers['set-cookie'][0].split(';')[0]
      const userId1 = res1.data.userID
      const userId2 = res2.data.userID

      // check follow
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

      instance
        .post(
          `user/follow`,
          {
            userId2: userId2,
          },
          {
            headers: { Cookie: cookie2 },
            credentials: 'include',
          }
        )
        .catch((err) => {
          expect(err.response.status).toBe(400)
        })

      // check user follow list
      const follows = await instance.get('user/check-follow', {
        headers: { Cookie: cookie1 },
        credentials: 'include',
      })

      expect(follows.status).toBe(200)
      expect(follows.data).toBe(null)

      // check user friends
      const friends = await instance.get('user/friends', {
        headers: { Cookie: cookie1 },
        credentials: 'include',
      })

      expect(friends.status).toBe(200)
      expect(friends.data[0].Email).toBe(mockNormalUser2.email)
    })

    it('register new acc with private', async () => {
      const user1 = await instance.post('user/register', mockNormalUserPrivate)
      expect(user1.status).toBe(201)

      const cookie1 = user1.headers['set-cookie'][0].split(';')[0]
      const userId1 = user1.data.userID

      const user2 = await instance.post(`user/login`, mockLoginUser)
      expect(user2.status).toBe(200)
      const cookie2 = user2.headers['set-cookie'][0].split(';')[0]

      // check follow
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

      // check follow list
      const followList = await instance.get(`user/check-follow`, {
        headers: { Cookie: cookie1 },
        credentials: 'include',
      })
      expect(followList.status).toBe(200)
      expect(followList.data.length).toBe(1)
      expect(followList.data[0].Email).toBe(mockLoginUser.email)
      const userId2 = followList.data[0].UserID

      // pre-check friends list
      const friendsList = await instance.get(`user/friends`, {
        headers: { Cookie: cookie1 },
        credentials: 'include',
      })
      expect(friendsList.status).toBe(200)
      expect(friendsList.data).toBe(null)

      // accept invite
      const resolve = await instance.post(
        'user/resolve',
        {
          userId: userId2,
          accept: true,
        },
        {
          headers: { Cookie: cookie1 },
          credentials: 'include',
        }
      )
      expect(resolve.status).toBe(200)

      // after check follow list
      const followList2 = await instance.get(`user/check-follow`, {
        headers: { Cookie: cookie1 },
        credentials: 'include',
      })
      expect(followList2.status).toBe(200)
      expect(followList2.data).toBe(null)

      // after check friends list
      const friendsList2 = await instance.get(`user/friends`, {
        headers: { Cookie: cookie1 },
        credentials: 'include',
      })
      expect(friendsList2.status).toBe(200)
      expect(friendsList2.data.length).toBe(1)
      expect(friendsList2.data[0].Email).toBe(mockLoginUser.email)
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
