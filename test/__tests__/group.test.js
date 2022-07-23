const axios = require('axios')
const { baseUrl } = require('../config')
const { normalGroup } = require('../__mock__/group.mock')
const { mockLoginUser } = require('../__mock__/user.mock')

const instance = axios.create({
  withCredentials: true,
  baseURL: baseUrl,
})

describe('group calls', () => {
  it('create group not authed', async () => {
    expect.assertions(1)
    try {
      await instance.post('group/create', normalGroup)
    } catch (err) {
      expect(err.response.status).toBe(401)
    }
  })

  it('create group authed', async () => {
    const user2 = await instance.post(`user/login`, mockLoginUser)
    const cookie2 = user2.headers['set-cookie'][0].split(';')[0]

    const res = await instance.post('group/create', normalGroup, {
      headers: { Cookie: cookie2 },
      credentials: 'include',
    })

    expect(res.status).toBe(201)
  })
})
